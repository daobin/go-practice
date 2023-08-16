package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for i := 1; i < 20000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}

	return &buff
}

func test3() {
	log.Println("======> loop begin.")
	for i := 0; i < 1000; i++ {
		log.Println("======> 随机生成：", genSomeBytes())
	}
	log.Println("======> loop end.")
}

func main() {
	go func() {
		for {
			test3()
			time.Sleep(time.Second)
		}
	}()

	log.Println(http.ListenAndServe(":10000", nil))
}
