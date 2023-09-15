package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var cnt int32

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//cnt += 1
			atomic.AddInt32(&cnt, 1)
		}()
	}

	wg.Wait()

	fmt.Println("cnt: ", cnt)
}
