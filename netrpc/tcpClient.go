package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("rpc.Dial Error: ", err.Error())
		return
	}
	defer conn.Close()

	res := ""
	//err = conn.Call("Hello.Say", "ZS", &res)
	err = conn.Call("h.Say", "ZS", &res)
	if err != nil {
		fmt.Println("conn.Call Error: ", err.Error())
		return
	}
	fmt.Println(res)
}
