package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go dostuff()
	go dostuff2()
	wg.Wait()
}

var on sync.Once

func setup() {
	fmt.Println("init")
}
func dostuff() {
	on.Do(setup)
	fmt.Println("hola")
	wg.Done()
}

func dostuff2() {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}
