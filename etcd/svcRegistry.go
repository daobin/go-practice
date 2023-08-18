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

	// 为服务分配一个租约
	lease, err := cli.Grant(context.Background(), 10)
	if err != nil {
		log.Println("cli.Grant Err: ", err.Error())
		return
	}

	// 注册服务
	//svcAddr := "127.0.0.1:8080"
	svcAddr := "127.0.0.1:8081"
	svcKey := "/services/shop-user/" + svcAddr
	_, err = cli.Put(context.Background(), svcKey, svcAddr, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Println("cli.Put Err: ", err.Error())
		return
	}

	// 定期刷新租约，确保服务继续保持注册状态
	_, err = cli.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		log.Println("cli.KeepAlive Err: ", err.Error())
		return
	}

	log.Println("Service Registry Success !!!")
}
