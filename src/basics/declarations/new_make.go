package main

import "fmt"

type Register struct {
	service_type int
	oid          float64
	olt          []int
	pon          *int
	port         string
	value        string
}

func NewRegister() *Register {
	r := new(Register)
	return r
}

func main() {
	a := new(Register)
	fmt.Println("Hello", a)

	a = NewRegister()
	fmt.Println("new", a)

	a = &Register{}
	fmt.Println("composite literal", a)
}
