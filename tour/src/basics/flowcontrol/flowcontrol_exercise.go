// https://go.dev/tour/flowcontrol/8

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < int(x); i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
