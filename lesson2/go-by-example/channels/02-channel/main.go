package main

import "fmt"

func main() {
	messages := make(chan string)

	// 从goroutine发送数据
	go func() { messages <- "ping" }()
	// 从channel接收数据
	msg := <-messages
	fmt.Println(msg)
}
