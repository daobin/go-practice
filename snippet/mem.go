package main

import (
	"log"
	"runtime"
	"time"
)

func test() {
	container := make([]int, 8)

	log.Println("======> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println("======> loop end.")
}

func main() {
	log.Println("Start.")

	test()

	log.Println("force gc.")
	runtime.GC()

	time.Sleep(3600 * time.Second)
}
