package main

import (
	"context"
	"fmt"
	"github.com/daobin/go-practice/grpc/hello/helloService"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务
	// insecure 包只有 grpc 1.30 版本及以上有，但是没有 naming 包，与 ETCD 冲突
	// grpcClient, err := grpc.Dial("127.0.0.1:8899", grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcClient, err := grpc.Dial("127.0.0.1:8899", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial Error: ", err.Error())
		return
	}

	// 注册客户端
	client := helloService.NewHelloClient(grpcClient)

	// 调用服务端方法
	res, err := client.SayHello(context.Background(), &helloService.HelloReq{Name: "Fish Lai"})
	if err != nil {
		fmt.Println("client.SayHello Error: ", err.Error())
		return
	}

	fmt.Println(res)
}
