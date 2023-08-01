package main

import (
	"context"
	"fmt"
	"github.com/daobin/go-practice/grpc/hello/helloService"
	"google.golang.org/grpc"
	"net"
)

// Hello 服务接口实现类
type Hello struct{}

func (h *Hello) SayHello(c context.Context, req *helloService.HelloReq) (*helloService.HelloRes, error) {
	fmt.Println("请求：", req)

	return &helloService.HelloRes{Msg: "你好：" + req.Name}, nil
}

func main() {
	// 创建 grpc 服务
	grpcServer := grpc.NewServer()

	// 注册服务
	helloService.RegisterHelloServer(grpcServer, new(Hello))

	// 设置监听
	listen, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println("net.Listen Error: ", err.Error())
		return
	}
	defer listen.Close()

	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("grpcServer.Serve Error: ", err.Error())
		return
	}
}
