package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	// 工作协程
	go func() {
		for {
			// 循环从jobs接收数据, 接收完毕，more为false
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	// 使用jobs发送3个任务到工作协程，然后关闭jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all job")
	// main函数同步等待任务结束
	<-done
}
