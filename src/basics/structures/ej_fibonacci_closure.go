package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f := 1
	anterior := 1
	return func() int {
		var aux = f
		f += anterior
		anterior = aux
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
