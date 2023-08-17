package main

import (
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Println("os.Create Err: ", err.Error())
		return
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Println("trace.Start Err: ", err.Error())
		return
	}
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "Go 编程之乐"
	}()

	<-ch
}
