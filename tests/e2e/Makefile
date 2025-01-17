ROOT_DIR := ../..

TCP_ADDRESS = :12345
UNIX_ADDRESS = /tmp/ovnnb_db.db

ETCD_NAME := ovsdb

DATABASE-PREFIX := ovsdb
SERVICE-NAME :=nb
SCHEMA-FILE := ovn-nb.ovsschema
LOAD-SERVER-DATA=TRUE := true
PID-FILE :=/tmp/nb-ovsdb.pid

.PHONY: etcd
etcd:
	etcd \
		--name $(ETCD_NAME) \
		--data-dir /tmp/$(ETCD_NAME).etcd

SERVER_FILES := \
	$(ROOT_DIR)/pkg/cmd/server/server.go \
	$(ROOT_DIR)/pkg/cmd/server/testdata.go

SERVER_ARGS := \
		-tcp-address $(TCP_ADDRESS) \
		-unix-address $(UNIX_ADDRESS) \
		-schema-basedir $(ROOT_DIR)/schemas \
		-database-prefix $(DATABASE-PREFIX) \
		-service-name $(SERVICE-NAME) \
		-schema-file $(SCHEMA-FILE) \
		-load-server-data $(LOAD-SERVER-DATA) \
		-pid-file $(PID-FILE)

.PHONY: dbg-server
dbg-server: 
	dlv debug $(SERVER_FILES) -- $(SERVER_ARGS)

.PHONY: server
server: 
	go run $(SERVER_FILES) $(SERVER_ARGS)

CLIENT_ARGS := \
		-server $(TCP_ADDRESS)

CLIENT_FILES := \
	$(ROOT_DIR)/pkg/cmd/client/client.go

.PHONY: dbg-client
dbg-client:
	dlv debug $(CLIENT_FILES) -- $(CLIENT_ARGS)

.PHONY: client
client:
	go run $(CLIENT_FILES) $(CLIENT_ARGS)

INSTALL_TOOLS += /usr/bin/pv
/usr/bin/pv:
	sudo apt-get install -y pv

INSTALL_TOOLS += /usr/bin/ovsdb-client
/usr/bin/ovsdb-client:
	sudo apt-get install -y ovn-common

INSTALL_TOOLS += /usr/local/sbin/demo-magic.sh
/usr/local/sbin/demo-magic.sh:
	sudo wget -O $@ https://raw.githubusercontent.com/paxtonhare/demo-magic/master/demo-magic.sh

.PHONY: demo
demo: $(INSTALL_TOOLS)
	./demo.sh
