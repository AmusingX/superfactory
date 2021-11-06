package etcd

import (
	"go.etcd.io/etcd/client/v3"
	"log"
)

var Client *clientv3.Client

func InitEtcdClientV3(conf clientv3.Config) {
	var err error
	Client, err = clientv3.New(conf)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseEtcdClientV3(conf clientv3.Config) {
	if Client != nil {
		Client.Close()
		Client = nil
	}
}
