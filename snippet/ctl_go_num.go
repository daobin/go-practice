package main

import (
	"log"
	"math"
	"runtime"
	"sync"
)

// 1. 通过带buffer的channel控制协程数量，实际上限制的是循环速度

//func biz(ch chan bool, i int) {
//	log.Println("Go >>>>>>>>> ", i, "   Count = ", runtime.NumGoroutine())
//	<-ch
//}
//
//func main() {
//	taskCnt := math.MaxInt64
//
//
//	ch := make(chan bool, 3)
//	for i := 0; i < taskCnt; i++ {
//		ch <- true
//		go biz(ch, i)
//	}
//}

// 2. 利用无缓冲channel与任务发送、执行分离的方式控制协程数量

var wg = sync.WaitGroup{}

func biz(ch chan int) {
	for t := range ch {
		log.Println("Go >>>>>>>>> ", t, "   Count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	log.Println(">>>>>>>>>> [Start]")

	ch := make(chan int)

	goCnt := 3 // 启动协程数量
	for i := 0; i < goCnt; i++ {
		go biz(ch)
	}

	taskCnt := math.MaxInt64
	for i := 0; i < taskCnt; i++ {
		sendTask(i, ch)
	}

	wg.Wait()
	log.Println(">>>>>>>>>> [End]")
}
