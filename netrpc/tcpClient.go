package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("rpc.Dial Error: ", err.Error())
		return
	}
	defer client.Close()

	res := ""
	//err = client.Call("Hello.Say", "ZS", &res)
	err = client.Call("h.Say", "ZS", &res)
	if err != nil {
		fmt.Println("client.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)
}
