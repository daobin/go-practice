package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

// World 定义服务类
type World struct{}

/*
Hello
1、方法只能有两个可序列化的参数
2、参数的类型不能是：channel（通道）、complex（复数类型）、func（函数）,因为不能进行序列化
3、方法要返回一个error类型，同时必须是公开的方法
*/
func (w *World) Hello(req string, res *string) error {
	*res = "Hello " + req
	return nil
}

func (w *World) Print(req string, res *string) error {
	*res = "Print " + req
	return nil
}

// abc 无法被调用，需要首字母大写导出
func (w *World) abc(req string, res *string) error {
	*res = "Abc " + req
	return nil
}

func main() {
	// 注册RPC服务
	err := rpc.Register(new(World))
	if err != nil {
		fmt.Println("rpc.Register Error: ", err.Error())
		return
	}
	// 使用HTTP作为RPC载体
	rpc.HandleHTTP()
	// 监听
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("net.Listen Error: ", err.Error())
		return
	}
	// 链接
	err = http.Serve(listen, nil)
	if err != nil {
		fmt.Println("http.Serve Error: ", err.Error())
		return
	}
}
