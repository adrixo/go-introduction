package main

import (
	"fmt"
	"strconv"
)

func main() {
	l := make([]int, 0, 10)

	for i := 0; i < cap(l); i++ {
		input := userInput(i)

		if input == "x" || input == "X" {
			break
		}

		l = addToSlice(l, i, input)

		printSlice(l)
	}
}

func printSlice(l []int) {
	bubblesort(l)
	fmt.Print(l)
}

func addToSlice(l []int, i int, input string) []int {
	if len(l) <= i {
		n, _ := strconv.Atoi(input)
		l = append(l, n)
	}

	return l
}

func userInput(i int) string {
	var input string
	fmt.Printf("Enter number %d or type exit: ", i+1)
	fmt.Scan(&input)
	return input
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
