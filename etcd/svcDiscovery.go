package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	// 创建 ETCD 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.190.134:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("clientv3.New Err: ", err.Error())
		return
	}
	defer cli.Close()

	svcKey := "/services/shop-user"

	// 服务发现
	res, err := cli.Get(context.Background(), svcKey, clientv3.WithPrefix())
	if err != nil {
		log.Println("cli.Get: ", err.Error())
		return
	}

	for _, kv := range res.Kvs {
		log.Println("Service: ", string(kv.Value))
	}

	log.Println("Service Discovery Success !!!")
}
