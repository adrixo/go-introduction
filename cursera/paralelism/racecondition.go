package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	fmt.Println(n)
}

func main() {
	// We can not assume which one of the two goroutines is going to be processed first
	go printNumber(1)
	go printNumber(2)
	time.Sleep(1 * time.Second)
}
