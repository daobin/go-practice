package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct{}

func (h *Hello) Say(req string, res *string) error {
	fmt.Println("Req: ", req)
	*res = "你好 " + req

	fmt.Println("Res: ", *res)
	return nil
}

func main() {
	//err := rpc.Register(new(Hello))
	//if err != nil {
	//	fmt.Println("rpc.Register Error: ", err.Error())
	//	return
	//}
	err := rpc.RegisterName("h", new(Hello))
	if err != nil {
		fmt.Println("rpc.Register Error: ", err.Error())
		return
	}

	// 监听
	listen, err := net.Listen("tcp", ":8080")
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

		rpc.ServeConn(conn)
	}
}
