package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("net.Dial Error: ", err.Error())
		return
	}
	defer conn.Close()

	// 基于JSON的RPC
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	defer client.Close()

	res := ""
	err = client.Call("dog.Run", "旺财", &res)
	if err != nil {
		fmt.Println("conn.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)
}
