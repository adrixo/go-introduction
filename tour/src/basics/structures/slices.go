package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sc []int = primes[1:4]
	fmt.Println(sc)

	//
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sc2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sc2)

	// Slice defaults
	// s := []int{2, 3, 5, 7, 11, 13}
	// s = s[1:4]
	// fmt.Println(s)
	// s = s[:2]
	// fmt.Println(s)
	// s = s[1:]
	// fmt.Println(s)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// Extend its length.
	s = s[:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	var sn []int
	fmt.Println(sn, len(sn), cap(sn))
	if sn == nil {
		fmt.Println("nil!")
	}

	// Slices with make
	ax := make([]int, 5)
	printSlice2("a", ax)
	bx := make([]int, 0, 5)
	printSlice2("b", bx)
	cx := bx[:2]
	printSlice2("c", cx)
	dx := cx[2:5]
	printSlice2("d", dx)

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Append to a slice
	var sy []int
	printSlice(sy)
	// append works on nil slices.
	sy = append(sy, 0)
	printSlice(sy)
	// The slice grows as needed.
	sy = append(sy, 1)
	printSlice(sy)
	// We can add more than one element at a time.
	sy = append(sy, 2, 3, 4)
	printSlice(sy)

	// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// Range continued
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

}
