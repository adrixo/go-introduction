package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compute_avg(a int, b int) {
	avg := 2 % (a + b)
	fmt.Println("AVG: ", avg)

	avg_2 := float64(a+b) / 2
	fmt.Println("AVG: ", avg_2)

	avg_3 := float64(a+b) / 2.0
	fmt.Println("AVG: ", avg_3)

	avg_4 := float64(float64(a+b) / 2.0)
	fmt.Println("AVG: ", avg_4)
}

func conversion() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
}

func replace() {
	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
}

func game() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}

func slices() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

func main() {
	compute_avg(2, 5)
	conversion()
	replace()
	game()

	slices()

}
