package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	fmt.Println(n)
}

func main() {
	// 1. concurrent execution
	// We can not assume which one of the two goroutines is going to be processed first
	go printNumber(1)
	go printNumber(2)
	time.Sleep(1 * time.Second)

	// 2. Race condition
	// We can not make asumptions about the variable access
	var number int = 0
	go increaseNumber(&number)
	go increaseNumber(&number)
	time.Sleep(1 * time.Second)
}

func increaseNumber(n *int) {
	*n++
}
