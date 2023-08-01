package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

/*
安装 protoc：当前版为 v3.15.5
https://github.com/protocolbuffers/protobuf/releases/tag/v3.15.5

安装 protoc-gen-go：当前版本 v1.5.3
go get github.com/golang/protobuf@v1.5.3
*/

func main() {
	searchReq := &SearchReq{
		Query:     "name=fish",
		PageIndex: 1,
		PageSize:  20,
		State:     1,
	}
	fmt.Println("Query: ", searchReq.GetQuery())

	reqBytes, err := proto.Marshal(searchReq)
	if err != nil {
		fmt.Println("Marshal Error: ", err.Error())
		return
	}

	req := &SearchReq{}
	err = proto.Unmarshal(reqBytes, req)
	if err != nil {
		fmt.Println("Unmarshal Error: ", err.Error())
		return
	}

	fmt.Println("State: ", searchReq.GetState())
}
