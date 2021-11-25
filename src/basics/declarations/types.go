package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type conversions
	var x, y int = 3, 4
	var l float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(l)
	fmt.Println(x, y, z)

	// Type inference
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	// constants
	const World = "世界"
	fmt.Printf("Hello", World)
	fmt.Printf("\nType: %T Value: %v\n", needInt(Small), needInt(Small))
	fmt.Printf("Type: %T Value: %v\n", needFloat(Small), needFloat(Small))
	fmt.Printf("Type: %T Value: %v\n", needFloat(Big), needFloat(Big))

}
