package ovsdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"k8s.io/klog/v2"

	"github.com/ibm/ovsdb-etcd/pkg/common"
	"github.com/ibm/ovsdb-etcd/pkg/libovsdb"
	"github.com/ibm/ovsdb-etcd/pkg/ovsjson"
)

type JrpcServer interface {
	Wait() error
	Stop()
	Notify(ctx context.Context, method string, params interface{}) error
}

type Handler struct {
	db         Databaser
	etcdClient *clientv3.Client

	jrpcServer     JrpcServer
	handlerContext context.Context
	clientCon      net.Conn
	closed         bool // false by default
	mu             sync.Mutex

	// dbName->dbMonitor
	monitors map[string]*dbMonitor
	// json-value string to handler monitor related data
	handlerMonitorData map[string]handlerMonitorData

	databaseLocks map[string]Locker
}

func (ch *Handler) Transact(ctx context.Context, params []interface{}) (interface{}, error) {
	klog.V(5).Infof("Transact request from %v, params %v", ch.GetClientAddress(), params)
	if ch.closed {
		klog.V(5).Infof("Transact request from %v, the handler is closed", ch.GetClientAddress())
		// prevents old transactions
		return nil, nil
	}
	req, err := libovsdb.NewTransact(params)
	if err != nil {
		return nil, err
	}
	txn := NewTransaction(ch.etcdClient, req)
	txn.schemas = ch.db.GetSchemas()
	// temporary solution to provide consistency
	ch.db.DbLock(req.DBName)
	rev, err := txn.Commit()
	ch.db.DbUnlock(req.DBName)

	if err != nil {
		return nil, err
	}
	if monitor, ok := ch.monitors[txn.request.DBName]; ok {
		klog.V(5).Infof("Transact sending to monitor to %v: %s", ch.GetClientAddress(), txn.etcd.EventsDump())
		monitor.notify(txn.etcd.Events, rev)
	}

	klog.V(5).Infof("Transact response to %v: %s", ch.GetClientAddress(), txn.response)
	return txn.response.Result, nil
}

func (ch *Handler) Cancel(ctx context.Context, param interface{}) (interface{}, error) {
	klog.V(5).Infof("Cancel request from %v, param %v", ch.GetClientAddress(), param)

	return "{Cancel}", nil
}

func (ch *Handler) Monitor(ctx context.Context, params []interface{}) (interface{}, error) {
	klog.V(5).Infof("Monitor request from %v, params %v", ch.GetClientAddress(), params)
	updatersMap, err := ch.addMonitor(params, ovsjson.Update)
	if err != nil {
		klog.Errorf("Monitor from %v, params %v got an error: %s", ch.GetClientAddress(), params, err)
		return nil, err
	}
	data, err := ch.getMonitoredData(updatersMap)
	klog.V(5).Infof("Monitor response to %v, params %v, err %v", ch.GetClientAddress(), params, err)
	if err != nil {
		ch.removeMonitor(params[1], false)
		return nil, err
	}
	return data, nil
}

func (ch *Handler) MonitorCancel(ctx context.Context, param interface{}) (interface{}, error) {
	klog.V(5).Infof("MonitorCancel request from %v, param %v", ch.GetClientAddress(), param)
	err := ch.removeMonitor(param, true)
	if err != nil {
		return nil, err
	}
	return "{}", nil
}

func (ch *Handler) Lock(ctx context.Context, param interface{}) (interface{}, error) {
	klog.V(5).Infof("Lock request from %v, param %v", ch.GetClientAddress(), param)
	id, err := common.ParamsToString(param)
	if err != nil {
		return map[string]bool{"locked": false}, err
	}
	ch.mu.Lock()
	myLock, ok := ch.databaseLocks[id]
	ch.mu.Unlock()
	if !ok {
		myLock, err = ch.db.GetLock(ch.handlerContext, id)
		if err != nil {
			klog.Warningf("Lock returned error %v\n", err)
			return nil, err
		}
		ch.mu.Lock()
		// validate that no other locks
		otherLock, ok := ch.databaseLocks[id]
		if !ok {
			ch.databaseLocks[id] = myLock
		} else {
			// What should we do ?
			myLock.cancel()
			myLock = otherLock
		}
		ch.mu.Unlock()
	}
	err = myLock.tryLock()
	if err == nil {
		return map[string]bool{"locked": true}, nil
	} else if err != concurrency.ErrLocked {
		klog.Errorf("Locked %s got error %v", id, err)
		// TOD is it correct?
		return nil, err
	}
	go func() {
		err = myLock.lock()
		if err == nil {
			// Send notification
			klog.V(5).Infof("%v %s Locked", ch.GetClientAddress(), id)
			if err := ch.jrpcServer.Notify(ch.handlerContext, "locked", []string{id}); err != nil {
				klog.Errorf("notification %v\n", err)
				return
			}
		} else {
			klog.Errorf("Lock %s error %v\n", id, err)
		}
	}()
	return map[string]bool{"locked": false}, nil
}

func (ch *Handler) Unlock(ctx context.Context, param interface{}) (interface{}, error) {
	klog.V(5).Infof("Unlock request from %v, param %v", ch.GetClientAddress(), param)
	id, err := common.ParamsToString(param)
	if err != nil {
		return ovsjson.EmptyStruct{}, err
	}
	ch.mu.Lock()
	myLock, ok := ch.databaseLocks[id]
	delete(ch.databaseLocks, id)
	ch.mu.Unlock()
	if !ok {
		klog.V(4).Infof("Unlock non existing lock %s", id)
		return ovsjson.EmptyStruct{}, nil
	}
	myLock.cancel()
	return ovsjson.EmptyStruct{}, nil
}

func (ch *Handler) Steal(ctx context.Context, param interface{}) (interface{}, error) {
	klog.V(5).Infof("Steal request from %v, param %v", ch.GetClientAddress(), param)
	// TODO
	return "{Steal}", nil
}

func (ch *Handler) MonitorCond(ctx context.Context, params []interface{}) (interface{}, error) {
	klog.V(1).Infof("MonitorCond request from %v, param %v", ch.GetClientAddress(), params)
	updatersMap, err := ch.addMonitor(params, ovsjson.Update2)
	if err != nil {
		klog.Errorf("MonitorCond from remote %v got an error: %s", ch.GetClientAddress(), err)
		return nil, err
	}
	data, err := ch.getMonitoredData(updatersMap)
	klog.V(5).Infof("MonitorCond response to %v, params %v, err %v", ch.GetClientAddress(), params, err)
	if err != nil {
		ch.removeMonitor(params[1], false)
		return nil, err
	}
	return data, nil
}

func (ch *Handler) MonitorCondChange(ctx context.Context, params []interface{}) (interface{}, error) {
	klog.V(5).Infof("MonitorCondChange request from %v, params %v", ch.GetClientAddress(), params)
	// TODO implement
	return "{Monitor_cond_change}", nil
}

func (ch *Handler) MonitorCondSince(ctx context.Context, params []interface{}) (interface{}, error) {
	klog.V(5).Infof("MonitorCondSince request from %v, parameters %v", ch.GetClientAddress(), params)
	updatersMap, err := ch.addMonitor(params, ovsjson.Update3)
	if err != nil {
		klog.Errorf("MonitorCondSince from remote %v got an error: %s", ch.GetClientAddress(), err)
		return nil, err
	}
	data, err := ch.getMonitoredData(updatersMap)
	klog.V(5).Infof("MonitorCondSince response to %v, params %v, err %v", ch.GetClientAddress(), params, err)
	if err != nil {
		ch.removeMonitor(params[1], false)
		return nil, err
	}
	return []interface{}{false, ovsjson.ZERO_UUID, data}, nil
}

func (ch *Handler) SetDbChangeAware(ctx context.Context, param interface{}) interface{} {
	klog.V(5).Infof("SetDbChangeAware request from %v, param %v", ch.GetClientAddress(), param)
	return ovsjson.EmptyStruct{}
}

// RFC 7047 section 4.1.11
// Can be used by both clients and servers to verify the liveness of a database connection.
// "params": JSON array with any contents
// Returns : "result": same as "params"
func (ch *Handler) Echo(ctx context.Context, param interface{}) interface{} {
	klog.V(5).Infof("Echo request from %v, parameters %v", ch.GetClientAddress(), param)
	return param
}

func NewHandler(tctx context.Context, db Databaser, cli *clientv3.Client) *Handler {
	return &Handler{
		handlerContext: tctx, db: db, databaseLocks: map[string]Locker{}, handlerMonitorData: map[string]handlerMonitorData{},
		etcdClient: cli, monitors: map[string]*dbMonitor{},
	}
}

func (ch *Handler) Cleanup() error {
	klog.Infof("CLEAN UP do something from %v", ch.GetClientAddress())
	ch.mu.Lock()
	defer ch.mu.Unlock()
	ch.closed = true
	for _, m := range ch.databaseLocks {
		m.unlock()
	}

	for _, monitor := range ch.monitors {
		monitor.cancelDbMonitor()
	}
	return nil
}

func (ch *Handler) SetConnection(jrpcSerer JrpcServer, clientCon net.Conn) {
	ch.jrpcServer = jrpcSerer
	ch.clientCon = clientCon
}

func (ch *Handler) notify(jsonValueString string, updates ovsjson.TableUpdates) {
	monitorData, ok := ch.handlerMonitorData[jsonValueString]
	if !ok {
		klog.Errorf("Unknown jsonValue %s", jsonValueString)
		return
	}
	if klog.V(5).Enabled() {
		klog.V(5).Infof("Monitor notification jsonValue %v to %v: %s", monitorData.jsonValue, ch.GetClientAddress(), updates)
	} else {
		klog.V(5).Infof("Monitor notification jsonValue %v to %v", monitorData.jsonValue, ch.GetClientAddress())
	}
	var err error
	switch monitorData.notificationType {
	case ovsjson.Update:
		err = ch.jrpcServer.Notify(ch.handlerContext, UPDATE, []interface{}{monitorData.jsonValue, updates})
	case ovsjson.Update2:
		err = ch.jrpcServer.Notify(ch.handlerContext, UPDATE2, []interface{}{monitorData.jsonValue, updates})
	case ovsjson.Update3:
		err = ch.jrpcServer.Notify(ch.handlerContext, UPDATE3, []interface{}{monitorData.jsonValue, ovsjson.ZERO_UUID, updates})
	}
	if err != nil {
		// TODO should we do something else
		klog.Errorf("Monitor notification jsonValue %v to %v returned error: %v", monitorData.jsonValue, ch.GetClientAddress(), err)
	}
}

func (ch *Handler) monitorCanceledNotification(jsonValue interface{}) {
	klog.V(5).Infof("monitorCanceledNotification %v to %v", jsonValue, ch.GetClientAddress())
	err := ch.jrpcServer.Notify(ch.handlerContext, MONITOR_CANCELED, jsonValue)
	if err != nil {
		// TODO should we do something else
		klog.Errorf(" error monitorCanceledNotification to %v : %v", ch.GetClientAddress(), err)
	}
}

func (ch *Handler) removeMonitor(jsonValue interface{}, notify bool) error {
	klog.V(5).Infof("removeMonitor %v", jsonValue)

	jsonValueString := jsonValueToString(jsonValue)
	ch.mu.Lock()
	defer ch.mu.Unlock()
	monitorData, ok := ch.handlerMonitorData[jsonValueString]
	if !ok {
		klog.Errorf("removing unexisting dbMonitor with jsonValue = %v", jsonValue)
		err := fmt.Errorf("unknown monitor")
		return err
	}
	monitor, ok := ch.monitors[monitorData.dataBaseName]
	if !ok {
		klog.Warningf("there is no monitor to %s", monitorData.dataBaseName)
	}

	monitor.removeUpdaters(monitorData.updatersKeys, jsonValueString)

	if !monitor.hasUpdaters() {
		monitor.cancel()
		delete(ch.monitors, monitorData.dataBaseName)
	}
	delete(ch.handlerMonitorData, jsonValueString)
	if notify {
		ch.monitorCanceledNotification(jsonValue)
	}
	return nil
}

func (ch *Handler) addMonitor(params []interface{}, notificationType ovsjson.UpdateNotificationType) (Key2Updaters, error) {

	cmpr, err := parseCondMonitorParameters(params)
	if err != nil {
		return nil, err
	}
	if len(cmpr.DatabaseName) == 0 {
		return nil, fmt.Errorf("monitored dataBase name is empty")
	}

	jsonValueString := jsonValueToString(cmpr.JsonValue)
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if _, ok := ch.handlerMonitorData[jsonValueString]; ok {
		return nil, fmt.Errorf("duplicate monitor ID")
	}
	updatersMap := Key2Updaters{}
	var updatersKeys []common.Key
	for tableName, mcrs := range cmpr.MonitorCondRequests {
		var updaters []updater
		for _, mcr := range mcrs {
			updater := mcrToUpdater(mcr, jsonValueString, notificationType == ovsjson.Update)
			updaters = append(updaters, *updater)
		}
		key := common.NewTableKey(cmpr.DatabaseName, tableName)
		updatersMap[key] = updaters
		updatersKeys = append(updatersKeys, key)
	}
	monitor, ok := ch.monitors[cmpr.DatabaseName]
	if !ok {
		monitor = ch.db.CreateMonitor(cmpr.DatabaseName, ch)
		monitor.start()
		ch.monitors[cmpr.DatabaseName] = monitor
	}
	monitor.addUpdaters(updatersMap)
	ch.handlerMonitorData[jsonValueString] = handlerMonitorData{
		dataBaseName:     cmpr.DatabaseName,
		notificationType: notificationType,
		updatersKeys:     updatersKeys,
		jsonValue:        cmpr.JsonValue}

	return updatersMap, nil
}

func (ch *Handler) getMonitoredData(updatersMap Key2Updaters) (ovsjson.TableUpdates, error) {
	returnData := ovsjson.TableUpdates{}
	for tableKey, updaters := range updatersMap {
		if len(updaters) == 0 {
			// nothing to update
			continue
		}
		// validate that Initial is required
		reqInitial := false
		for _, updater := range updaters {
			reqInitial = reqInitial || libovsdb.MSIsTrue(updater.Select.Initial)
			if reqInitial {
				break
			}
		}
		resp, err := ch.db.GetData(tableKey, false)
		if err != nil {
			return nil, err
		}
		d1 := ovsjson.TableUpdate{}
		for _, kv := range resp.Kvs {
			for _, updater := range updaters {
				row, uuid, err := updater.prepareCreateRowInitial(&kv.Value)
				if err != nil {
					klog.Errorf("prepareCreateRowInitial returned %s", err)
					return nil, err
				}
				klog.V(8).Infof("processing getMonitoredData %v  row %v", d1, row)
				// TODO merge
				if row != nil {
					d1[uuid] = *row
				} else {
					klog.Info("row is nil")
				}
			}
		}
		if len(d1) > 0 {
			returnData[tableKey.TableName] = d1
		}
	}
	klog.V(6).Infof("getMonitoredData: %v", returnData)
	return returnData, nil
}

func (ch *Handler) GetClientAddress() string {
	if ch.clientCon != nil {
		return ch.clientCon.RemoteAddr().String()
	}
	return ""
}

func parseCondMonitorParameters(params []interface{}) (*ovsjson.CondMonitorParameters, error) {
	l := len(params)
	if l < 2 || l > 4 {
		err := fmt.Errorf("wrong length of condition dbMonitor parameters: %d", l)
		klog.Errorf("parseCondMonitorParameters %v params = %v", err, params)
		return nil, err
	}
	cmp := ovsjson.CondMonitorParameters{}
	var ok bool
	cmp.DatabaseName, ok = params[0].(string)
	if !ok {
		err := fmt.Errorf("parseCondMonitorParameters, cannot assert dbname interface (type %T, value %v) to string", params[0], params[0])
		klog.Errorf("%v", err)
		return nil, err
	}
	cmp.JsonValue = params[1]
	buf, err := json.Marshal(params[2])
	if err != nil {
		klog.Errorf("marshal dbMonitor conditional request returned %v", err)
		return nil, err
	}
	if err := json.Unmarshal(buf, &cmp.MonitorCondRequests); err != nil {
		obj := map[string]ovsjson.MonitorCondRequest{}
		if err := json.Unmarshal(buf, &obj); err != nil {
			return nil, fmt.Errorf("unmarshal dbMonitor condition requests returned: %v", err)
		}
		cmp.MonitorCondRequests = map[string][]ovsjson.MonitorCondRequest{}
		for k, v := range obj {
			cmp.MonitorCondRequests[k] = []ovsjson.MonitorCondRequest{v}
		}
	}
	if l == 4 {
		str, ok := params[3].(string)
		if !ok {
			err := fmt.Errorf("parseCondMonitorParameters, cannot assert last txn ID interface (type %T, value %v) to string", params[3], params[3])
			klog.Errorf("%v", err)
			return nil, err
		}
		cmp.LastTxnID = &str
	}
	return &cmp, nil
}

func jsonValueToString(jsonValue interface{}) string {
	return fmt.Sprintf("%v", jsonValue)
}
