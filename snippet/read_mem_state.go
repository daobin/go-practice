package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func readMemStats() {
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf("=========> 分配内存：%d 字节\n", ms.Alloc)
	log.Printf("=========> 空闲内存：%d 字节\n", ms.HeapIdle)
	log.Printf("=========> 归还内存：%d 字节\n", ms.HeapReleased)
}

func test2() {
	container := make([]int, 8)

	log.Println("======> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}
	log.Println("======> loop end.")
}

func main() {
	// 启动 pprof 监听，访问地址如下
	// http://127.0.0.1:10000/debug/pprof/heap?debug=1
	go func() {
		log.Println(http.ListenAndServe(":10000", nil))
	}()

	log.Println("======> [Start.]")

	test2()

	log.Println("======> [force gc.]")
	runtime.GC()

	log.Println("======> [Done.]")
	readMemStats()

	go func() {
		for {
			readMemStats()
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(3600 * time.Second)
}
