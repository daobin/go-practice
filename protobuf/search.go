package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

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
