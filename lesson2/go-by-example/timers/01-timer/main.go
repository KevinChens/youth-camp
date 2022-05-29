package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器等待2s
	timer1 := time.NewTimer(2 * time.Second)
	// 一直阻塞，直到C明确发送定时器失效的值
	<-timer1.C
	fmt.Println("Timer 1 fired")
	// 取消定时器
	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
