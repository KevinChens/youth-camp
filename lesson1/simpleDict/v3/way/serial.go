package way

import (
	"fmt"
	"lesson1/simpleDict/v3/query"
	"time"
)

func SPrint(word string) {
	// 串行
	fmt.Println("********串行*************")
	startTime := time.Now()
	query.WithVolcano(word)
	query.WithYoudao(word)
	query.WithCaiyun(word)
	elapsedTime := time.Since(startTime) / time.Millisecond
	fmt.Printf("串行时间：%d ms\n", elapsedTime)
}
