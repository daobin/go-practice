package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("rpc.DialHTTP Error: ", err.Error())
		return
	}
	defer client.Close()

	res := ""
	err = client.Call("World.Hello", "ZS", &res)
	if err != nil {
		fmt.Println("client.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)

	err = client.Call("World.Print", "ZS", &res)
	if err != nil {
		fmt.Println("client.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)

	err = client.Call("World.Abc", "ZS", &res)
	if err != nil {
		fmt.Println("client.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)
}
