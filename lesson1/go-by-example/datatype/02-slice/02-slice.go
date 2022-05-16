package _2_slice

import "fmt"

func Simulation() {
	// 模拟内部不同长度的多维数组
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
	// stack
	stack := make([]int, 0)
	// push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	// read
	top := stack[len(stack)-1]
	// pop
	stack = stack[:len(stack)-1]
	fmt.Println("stack:", stack, "top:", top)
	// queue
	queue := make([]int, 0)
	// push
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	// read
	front := queue[0]
	// pop
	queue = queue[1:]
	fmt.Println("queue:", queue, "front:", front)
}
