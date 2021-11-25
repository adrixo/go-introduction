package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// when the param type is shared it can be ommited
func add_continued(x, y int) int {
	return x + y
}

func swap_multiple_results(x, y string) (string, string) {
	return y, x
}

func split_named_return_values(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(43, 13))
	fmt.Println(add_continued(43, 13))

	a, b := "hello", "world"
	fmt.Println(a, b)
	a, b = swap_multiple_results(a, b)
	fmt.Println(a, b)

	fmt.Println(split_named_return_values(17))

}
