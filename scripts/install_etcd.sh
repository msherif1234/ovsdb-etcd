#!/bin/bash

ETCD_VERSION=${ETCD_VERSION:-v3.4.16}

curl -L https://github.com/coreos/etcd/releases/download/$ETCD_VERSION/etcd-$ETCD_VERSION-linux-amd64.tar.gz -o etcd-$ETCD_VERSION-linux-amd64.tar.gz

tar xzvf etcd-$ETCD_VERSION-linux-amd64.tar.gz
rm etcd-$ETCD_VERSION-linux-amd64.tar.gz

cd etcd-$ETCD_VERSION-linux-amd64
sudo cp etcd /usr/local/bin/
sudo cp etcdctl /usr/local/bin/

cd ..
rm -r etcd-$ETCD_VERSION-linux-amd64

etcdctl version
