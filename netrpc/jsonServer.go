package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
1、net/rpc标准库默认采用Go特有的gob编码，无法实现语言调用
2、net/rpc/jsonrpc采用JSON进行数据编码，支持跨语言调用
*/

type Dog struct{}

func (h *Dog) Run(req string, res *string) error {
	fmt.Println("请求参数：", req)
	*res = "快跑 " + req
	fmt.Println("响应结果：", *res)

	return nil
}

func main() {
	err := rpc.RegisterName("dog", new(Dog))
	if err != nil {
		fmt.Println("rpc.RegisterName Error: ", err.Error())
		return
	}

	// 监听
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("net.Listen Error: ", err.Error())
		return
	}
	defer listen.Close()

	for {
		fmt.Println("开始链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept Error: ", err.Error())
			return
		}

		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
