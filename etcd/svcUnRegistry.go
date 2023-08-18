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

	// 取消服务
	//svcAddr := "127.0.0.1:8080"
	svcAddr := "127.0.0.1:8081"
	svcKey := "/services/shop-user/" + svcAddr
	_, err = cli.Delete(context.Background(), svcKey)
	if err != nil {
		log.Println("cli.Put Err: ", err.Error())
		return
	}

	log.Println("Service UnRegistry Success !!!")
}
