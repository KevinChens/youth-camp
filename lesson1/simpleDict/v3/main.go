package main

import (
	"fmt"
	"lesson1/simpleDict/v3/way"
	"os"
)

func main() {
	// 从命令行读取参数
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD(example: simpleDict hello)`)
		os.Exit(1)
	}
	word := os.Args[1]
	// 并行
	//way.CPrint(word)
	// 串行
	way.SPrint(word)
}