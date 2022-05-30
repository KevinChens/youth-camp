package main

import "os"

func main() {
	_, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
}
