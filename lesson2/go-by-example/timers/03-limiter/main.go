package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// 每200ms接收一个值，是任务速率限制的调度器
	limiter := time.Tick(200*time.Millisecond)

	for req := range requests {
		// 每次请求前阻塞limiter，实现了每200ms执行一次请求
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println()
	// 速率限制的同时，允许短暂并发，burstyLimiter允许最多3个爆发
	burstyLimiter := make(chan time.Time, 3)
	// 填充通道，表示允许的爆发
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 每200ms添加值到burstyLimiter，直到3个的限制
	go func() {
		for t := range time.Tick(200*time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	// 模拟另外5个传入请求，得益于burstyLimiter的爆发能力，前3个请求可以快速完成
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
