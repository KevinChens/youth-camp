package way

import (
	"fmt"
	"lesson1/simpleDict/v3/query"
	"sync"
	"time"
)

func CPrint(word string) {
	// 并行
	fmt.Println("********并行*************")
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		query.WithVolcano(word)
	}()
	go func() {
		defer wg.Done()
		query.WithCaiyun(word)
	}()
	go func() {
		defer wg.Done()
		query.WithYoudao(word)
	}()
	wg.Wait()
	elapsedTime := time.Since(startTime) / time.Millisecond
	fmt.Printf("并行时间：%d ms\n", elapsedTime)
}
