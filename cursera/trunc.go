package main

import "fmt"

func trunc() {
	var x float32
	fmt.Println("Enter a float number:")
	fmt.Scan(&x)
	fmt.Println("Truncated result: ", int32(x))
}

func main() {
	trunc()
}
