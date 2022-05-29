package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 无符号整型变量表示计数器
	var ops uint64
	// 等待所有goroutine完成工作
	var wg sync.WaitGroup
	// 启动50个goroutine，每个goroutine会将计数器递增1000次
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// 使用AddUint64让计数器自增
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// 等待所有goroutine结束
	wg.Wait()
	// 安全访问ops，50000
	fmt.Println("ops:", ops)
}
