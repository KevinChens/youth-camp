package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d strarting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	// 启动goroutine，并递增WaitGroup的计数器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		// 将worker调用封装到一个闭包中，可以确保通知WaitGroup，此工作线程已完成
		// worker线程本身也就不需要知道执行涉及的并发原语
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	// 阻塞，直到WaitGroup计数器恢复为0
	wg.Wait()
}
