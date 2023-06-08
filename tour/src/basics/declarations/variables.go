package main

import "fmt"

// package level
// var c, python, java bool

// A var declaration can include initializers
var i, j int = 1, 2

func main() {
	// Function level
	// var i int
	// fmt.Println(i, c, python, java)

	// initializers
	c, python, java := true, false, "no!"
	// or
	// var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

}
