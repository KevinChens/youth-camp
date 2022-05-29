package main

import "fmt"

func ping(pings chan<- string, msg string) {
	// pings只写
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// pings只读
	msg := <- pings
	// pongs只写
	pongs <- msg
}

func main() {
	// pings, pongs可读可写
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}