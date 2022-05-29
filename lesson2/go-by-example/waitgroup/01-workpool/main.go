package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// 启动3个worker, 初始阻塞，因为没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 发送5个jobs，任务完成后关闭channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	// 收集任务的返回值，确保所有worker协程都完成
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
