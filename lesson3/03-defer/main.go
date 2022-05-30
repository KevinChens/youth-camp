package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_WRONLY|os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	n, err := file.WriteString("test")
	if err != nil {
		fmt.Printf("write in file failed, err:%v\n", err)
		return
	}
	fmt.Println(n)
}
