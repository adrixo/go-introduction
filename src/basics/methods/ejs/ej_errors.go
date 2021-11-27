// package main

// import (
// 	"fmt"
// )

// type ErrNegativeSqrt float64

// func (f ErrNegativeSqrt) Error() string {
// 	return fmt.Sprintf("math: square root of negative number %g", float64(f))
// }

// func Sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, &ErrNegativeSqrt
// 	}
// 	return 2, nil
// }

// func main() {
// 	fmt.Println(Sqrt(-2))
// 	fmt.Println(Sqrt(-2))
// }
