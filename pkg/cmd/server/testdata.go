package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ibm/ovsdb-etcd/pkg/common"
	"github.com/ibm/ovsdb-etcd/pkg/libovsdb"
	"github.com/ibm/ovsdb-etcd/pkg/ovsdb"
	"github.com/ibm/ovsdb-etcd/pkg/types/OVN_Northbound"
)

func newUUID() libovsdb.UUID {
	return common.ToUUID(uuid.NewString())
}

func newSet(s interface{}) (*libovsdb.OvsSet, error) {
	set, err := libovsdb.NewOvsSet(s)
	if err != nil {
		return nil, err
	}
	return set, nil
}

func newMap(s interface{}) (*libovsdb.OvsMap, error) {
	set, err := libovsdb.NewOvsMap(s)
	if err != nil {
		return nil, err
	}
	return set, nil
}

func putNbGlobalOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, connections []libovsdb.UUID, options map[string]string) error {
	connectionsSet, err := newSet(connections)
	if err != nil {
		return err
	}
	optionsMap, err := newMap(options)
	if err != nil {
		return err
	}

	nbGlobal := OVN_Northbound.NB_Global{
		Connections:      *connectionsSet,
		External_ids:     libovsdb.OvsMap{},
		Hv_cfg:           0,
		Hv_cfg_timestamp: 0,
		Ipsec:            false,
		Name:             "",
		Nb_cfg:           0,
		Nb_cfg_timestamp: 0,
		Options:          *optionsMap,
		Sb_cfg:           0,
		Sb_cfg_timestamp: 0,
		Ssl:              libovsdb.OvsSet{},
		Version:          newUUID(),
		Uuid:             libovsdb.UUID{GoUUID: uuid},
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "NB_Global", uuid), nbGlobal)
	return nil
}

func putAclOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, action string, direction string, match string, priority int64) error {
	acl := OVN_Northbound.ACL{
		Action:       action,
		Direction:    direction,
		External_ids: libovsdb.OvsMap{},
		Log:          false,
		Match:        match,
		Meter:        libovsdb.OvsSet{},
		Name:         libovsdb.OvsSet{},
		Priority:     priority,
		Severity:     libovsdb.OvsSet{},
		Version:      newUUID(),
		Uuid:         libovsdb.UUID{GoUUID: uuid},
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "ACL", uuid), acl)
	return nil
}

func putAddressSetOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, addresses []string, externalIds map[string]string, name string) error {
	addressesSet, err := newSet(addresses)
	if err != nil {
		return err
	}
	externalIdsMap, err := newMap(externalIds)
	if err != nil {
		return err
	}
	addressSet := OVN_Northbound.Address_Set{
		Addresses:    *addressesSet,
		External_ids: *externalIdsMap,
		Name:         name,
		Version:      newUUID(),
		Uuid:         libovsdb.UUID{GoUUID: uuid},
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Address_Set", uuid), addressSet)
	return nil
}

func putConnectionOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, status map[string]string, target string) error {
	statusMap, err := newMap(status)
	if err != nil {
		return err
	}
	connection := OVN_Northbound.Connection{
		External_ids:     libovsdb.OvsMap{},
		Inactivity_probe: libovsdb.OvsSet{},
		Is_connected:     false,
		Max_backoff:      libovsdb.OvsSet{},
		Other_config:     libovsdb.OvsMap{},
		Status:           *statusMap,
		Target:           target,
		Version:          newUUID(),
		Uuid:             libovsdb.UUID{GoUUID: uuid},
	}

	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Connection", uuid), connection)
	return nil
}

func putForwardingGroupOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, childPorts []string, externalIds map[string]string, name string) error {

	childPortsSet, err := newSet(childPorts)
	if err != nil {
		return err
	}
	externalIdsMap, err := newMap(externalIds)
	if err != nil {
		return err
	}
	fowardingGroup := OVN_Northbound.Forwarding_Group{
		Child_port:   *childPortsSet,
		External_ids: *externalIdsMap,
		Liveness:     false,
		Name:         name,
		Vip:          "",
		Vmac:         "",
		Version:      newUUID(),
		Uuid:         libovsdb.UUID{GoUUID: uuid},
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Forwarding_Group", uuid), fowardingGroup)
	return nil
}

func putloadBalancerOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, externalIds map[string]string, name string, options map[string]string, protocol []string, vips map[string]string) error {
	externalIdsMap, err := newMap(externalIds)
	if err != nil {
		return err
	}
	optionsMap, err := newMap(options)
	if err != nil {
		return err
	}
	vipsMap, err := newMap(vips)
	if err != nil {
		return err
	}
	protocolSet, err := newSet(protocol)
	if err != nil {
		return err
	}
	loadBalancer := OVN_Northbound.Load_Balancer{
		External_ids:     *externalIdsMap,
		Health_check:     libovsdb.OvsSet{},
		Ip_port_mappings: libovsdb.OvsMap{},
		Name:             name,
		Options:          *optionsMap,
		Protocol:         *protocolSet,
		Selection_fields: libovsdb.OvsSet{},
		Vips:             *vipsMap,
		Version:          newUUID(),
		Uuid:             libovsdb.UUID{GoUUID: uuid},
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Load_Balancer", uuid), loadBalancer)
	return nil
}

func putlogicalRouterOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, externalIds map[string]string, name string, options map[string]string, ports []string, staticRoutes []string) error {
	externalIdsMap, err := newMap(externalIds)
	if err != nil {
		return err
	}
	optionsMap, err := newMap(options)
	if err != nil {
		return err
	}
	portsSet, err := newSet(common.ToUUIDSlice(ports))
	if err != nil {
		return err
	}
	staticRoutesSet, err := newSet(common.ToUUIDSlice(staticRoutes))
	if err != nil {
		return err
	}
	logicalRouter := OVN_Northbound.Logical_Router{
		Enabled:       libovsdb.OvsSet{},
		External_ids:  *externalIdsMap,
		Load_balancer: libovsdb.OvsSet{},
		Name:          name,
		Nat:           libovsdb.OvsSet{},
		Options:       *optionsMap,
		Policies:      libovsdb.OvsSet{},
		Ports:         *portsSet,
		Static_routes: *staticRoutesSet,
		Version:       newUUID(),
		Uuid:          common.ToUUID(uuid),
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Logical_Router", uuid), logicalRouter)
	return nil
}

func putGatewayChassisOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, chassis_name string, externalIds map[string]string, name string, priority int64) error {
	externalIdsMap, err := newMap(externalIds)
	if err != nil {
		return err
	}
	gatewayChassis := OVN_Northbound.Gateway_Chassis{
		Chassis_name: chassis_name,
		External_ids: *externalIdsMap,
		Name:         name,
		Options:      libovsdb.OvsMap{},
		Priority:     priority,
		Version:      newUUID(),
		Uuid:         libovsdb.UUID{GoUUID: uuid},
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Gateway_Chassis", uuid), gatewayChassis)
	return nil
}

func putlogicalSwitchOnEtcd(ctx context.Context, con *ovsdb.DatabaseEtcd, uuid string, externalIds map[string]string, name string, ports []string) error {
	externalIdsMap, err := newMap(externalIds)
	if err != nil {
		return err
	}
	portsSet, err := newSet(common.ToUUIDSlice(ports))
	if err != nil {
		return err
	}

	logicalSwitch := OVN_Northbound.Logical_Switch{
		Acls:              libovsdb.OvsSet{},
		Dns_records:       libovsdb.OvsSet{},
		External_ids:      *externalIdsMap,
		Forwarding_groups: libovsdb.OvsSet{},
		Load_balancer:     libovsdb.OvsSet{},
		Name:              name,
		Other_config:      libovsdb.OvsMap{},
		Ports:             *portsSet,
		Qos_rules:         libovsdb.OvsSet{},
		Version:           newUUID(),
		Uuid:              common.ToUUID(uuid),
	}
	con.PutData(ctx, common.NewDataKey("OVN_Northbound", "Logical_Switch", uuid), logicalSwitch)

	return nil
}

func loadServerData(con *ovsdb.DatabaseEtcd) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	// OVN_Northbound

	// NB_Global

	err := putNbGlobalOnEtcd(ctx, con, "a5088a51-7756-4dd4-909c-b7c59c9fcce7", []libovsdb.UUID{{GoUUID: "413afe3e-79ff-4583-88a6-f02b70b8e927"}},
		map[string]string{
			"e2e_timestamp":           "1612817071",
			"mac_prefix":              "86:a9:cb",
			"max_tunid":               "16711680",
			"northd_internal_version": "20.12.0-20.14.0-52.0",
			"northd_probe_interval":   "5000",
			"svc_monitor_mac":         "5a:d9:62:39:9f:87",
		})
	if err != nil {
		cancel = nil
		return err
	}

	//ACL

	err = putAclOnEtcd(ctx, con, "aa2bab19-9b31-4d01-b1ad-f5e49dd269f8", "allow-related", "to-lport", "ip4.src==10.244.0.2", 1001)
	if err != nil {
		cancel = nil
		return err
	}

	err = putAclOnEtcd(ctx, con, "3ed181f9-7c68-47ee-bcdc-6cf393a02772", "allow-related", "to-lport", "ip4.src==10.244.1.2", 1001)
	if err != nil {
		cancel = nil
		return err
	}

	err = putAclOnEtcd(ctx, con, "7071b927-cc6d-4145-8849-395e6226fdac", "allow-related", "to-lport", "ip4.src==10.244.1.2", 1001)
	if err != nil {
		cancel = nil
		return err
	}

	//Address_Set

	err = putAddressSetOnEtcd(ctx, con, "532757d0-bc2e-41b9-bafe-2542f995b011", []string{"10.244.0.5"}, map[string]string{"name": "local-path-storage_v4"}, "a10956707444534956691")
	if err != nil {
		cancel = nil
		return err
	}

	err = putAddressSetOnEtcd(ctx, con, "8e33c234-2da4-4e5f-858f-4bcd5bc3c68b", []string{}, map[string]string{"name": "local-path-storage_v4"}, "a5154718082306775057")
	if err != nil {
		cancel = nil
		return err
	}

	err = putAddressSetOnEtcd(ctx, con, "3581fd85-1428-45a8-9702-edec71dda0a1", []string{"10.244.0.3", "10.244.0.4"}, map[string]string{"name": "kube-system_v4"}, "a6937002112706621489")
	if err != nil {
		cancel = nil
		return err
	}

	err = putAddressSetOnEtcd(ctx, con, "99ad8ae1-bc86-4662-bca4-a88fd675ee3d", []string{}, map[string]string{"name": "ovn-kubernetes_v4"}, "a5675285926127865604")
	if err != nil {
		cancel = nil
		return err
	}

	err = putAddressSetOnEtcd(ctx, con, "fde500ad-eff5-47a3-be0b-02e7c23a1357", []string{}, map[string]string{"name": "kube-public_v4"}, "a18363165982804349389")
	if err != nil {
		cancel = nil
		return err
	}

	err = putAddressSetOnEtcd(ctx, con, "0af13342-2ea7-486d-825a-b57bd70a8cbc", []string{}, map[string]string{"name": "kube-node-lease_v4"}, "a16235039932615691331")
	if err != nil {
		cancel = nil
		return err
	}

	// Connection

	err = putConnectionOnEtcd(ctx, con, "413afe3e-79ff-4583-88a6-f02b70b8e927", map[string]string{"bound_port": "6641", "n_connections": "3", "sec_since_connect": "0", "sec_since_disconnect": "0"}, "ptcp:6641:172.18.0.4")
	if err != nil {
		cancel = nil
		return err
	}

	// Forwarding_Group

	err = putForwardingGroupOnEtcd(ctx, con, "6be9235a-b3b6-41d7-a5aa-356b5b3c96cc", []string{"25f2e69e-4bac-4529-9082-9f94da060cf1",
		"73000cf3-73d0-4283-8aad-bcf181626a40", "be25033c-27df-42a2-9765-52bc06acc71c"}, map[string]string{
		"name": "clusterPortGroup"}, "clusterPortGroup")
	if err != nil {
		cancel = nil
		return err
	}

	err = putForwardingGroupOnEtcd(ctx, con, "ee4d82d2-3a7d-4737-be8d-656374f5d56c", []string{"b4298483-cf17-46d4-9da1-034eab065ff1",
		"b6e1fc02-0306-4887-8e36-e8b0ec22b16c", "fcf06a69-16c2-4f34-b3a4-282a641862f8"},
		map[string]string{"name": "clusterRtrPortGroup"}, "clusterRtrPortGroup")
	if err != nil {
		cancel = nil
		return err
	}

	// Gateway_Chassis

	err = putGatewayChassisOnEtcd(ctx, con, "99c45e0b-3688-4992-900c-7d5a25930ba3", "1bd76edb-8626-4ecd-8185-788bd2121bda",
		map[string]string{"dgp_name": "rtos-node_local_switch"}, "rtos-node_local_switch_1bd76edb-8626-4ecd-8185-788bd2121bda", 100)
	if err != nil {
		cancel = nil
		return err
	}

	// Load_Balancer (only first one out of 12)

	err = putloadBalancerOnEtcd(ctx, con, "32cc16f0-cda7-4c63-87d7-d30349ce32d7",
		map[string]string{"k8s-cluster-lb-tcp": "yes"}, "", map[string]string{}, []string{"tcp"}, map[string]string{"10.96.0.10:53": "10.244.0.3:53,10.244.0.4:53",
			"10.96.0.10:9153": "10.244.0.3:9153,10.244.0.4:9153", "10.96.0.1:443": "172.18.0.4:6443"})
	if err != nil {
		cancel = nil
		return err
	}

	// Logical_Router (only first two out of 4)

	err = putlogicalRouterOnEtcd(ctx, con, "22c3143c-15c8-4018-91b6-8fe4b0ffab80",
		map[string]string{"physical_ip": "172.18.0.4", "physical_ips": "172.18.0.4"},
		"GR_ovn-control-plane",
		map[string]string{"always_learn_from_arp_request": "false", "chassi": "1bd76edb-8626-4ecd-8185-788bd2121bda", "dynamic_neigh_router": "true"},
		[]string{"af4e1844-8479-476f-a45a-6444475f0062", "d54fc12c-ecd5-44c6-a2b9-bda4e0534d6b"},
		[]string{"7115ddef-8cad-4fc9-8471-63715480e4fd", "e4c2100f-f7ba-4129-a11f-0b4d854a7c28"},
	)
	if err != nil {
		cancel = nil
		return err
	}

	err = putlogicalRouterOnEtcd(ctx, con, "70f7ff95-16c1-4832-9073-5c5ba807d205",
		map[string]string{"physical_ip": "172.18.0.2", "physical_ips": "172.18.0.2"},
		"GR_ovn-worker2",
		map[string]string{"always_learn_from_arp_request": "false", "chassis": "8ec06983-c3c3-4687-b2b0-6283ee76b252", "dynamic_neigh_routers": "true"},
		[]string{"a3be4bc9-9bd0-496f-9796-afc43f6cae12", "a5731950-de12-4d55-8687-5fa7d2571742"},
		[]string{"815544ad-3bd9-4224-8776-75a3effabe28", "93431fc9-c663-4ca9-9e7c-bbf964c39471"},
	)
	if err != nil {
		cancel = nil
		return err
	}

	// Logical_Switch (only 5'th and 6'th out of 12)

	err = putlogicalSwitchOnEtcd(ctx, con, "4cccc9ad-5ba6-42f9-a749-2d5f3c54ace4", map[string]string{}, "ext_ovn-control-plane",
		[]string{"006d780f-7169-44f6-8eb2-9df757feef61", "a8fe81ba-4746-4d6f-aebe-d53cfd4ae46f"})
	if err != nil {
		cancel = nil
		return err
	}

	err = putlogicalSwitchOnEtcd(ctx, con, "1822aa4e-9d58-4261-a22b-ecb6f864a3bc", map[string]string{}, "ext_ovn-worker",
		[]string{"2a77a299-b713-4461-b3f6-4e91485fae79", "e172ea88-8779-4fec-8fb4-553efb426880"})
	if err != nil {
		cancel = nil
		return err
	}

	cancel()
	return err
}
