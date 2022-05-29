package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
在这个例子中，state将被一个单独的协程拥有。
这能保证数据在并行读取时不会混乱。
为了对state进行读取或者写入，其它的协程将发送一条数据到目前拥有数据的协程中，然后等待接收对应的回复。
结构体readOp和writeOp封装了这些请求，并提供了响应协程的方法。
 */

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64
	// 其他协程通过read，writes来发布读和写请求
	reads := make(chan readOp)
	writes := make(chan writeOp)
	// 拥有state的协程
	go func() {
		// state是这个协程私有的
		var state = make(map[int]int)
		for {
			select {
			case read := <- reads:
				read.resp <- state[read.key]
			case write := <- writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	// 启动100个goroutine，通过reads通道向拥有state的协程发起读取请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<- read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 用相同的方法启动10个写操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 让协程跑1s
	time.Sleep(time.Second)
	// 获取并报告ops
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
