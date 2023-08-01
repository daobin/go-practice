package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("rpc.DialHTTP Error: ", err.Error())
		return
	}
	defer conn.Close()

	res := ""
	err = conn.Call("World.Hello", "ZS", &res)
	if err != nil {
		fmt.Println("conn.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)

	err = conn.Call("World.Print", "ZS", &res)
	if err != nil {
		fmt.Println("conn.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)

	err = conn.Call("World.Abc", "ZS", &res)
	if err != nil {
		fmt.Println("conn.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)
}
