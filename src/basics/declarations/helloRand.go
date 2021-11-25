package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(100))
	//Exported names
	fmt.Println("Properly accesed name", math.Pi)
	// fmt.Println("Wrongly accesed name", math.pi)
}
