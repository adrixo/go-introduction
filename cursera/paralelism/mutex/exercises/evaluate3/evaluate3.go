package main

import (
	"fmt"
	"sync"
)

var chopsticks []bool //true = unused, false = in use
var mut sync.Mutex
var wg sync.WaitGroup

func main() {
	fmt.Println("Hello World!")
	chopsticks = make([]bool, 5)
	for idx := range chopsticks {
		chopsticks[idx] = true
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosopher(i)
	}
	wg.Wait()
	println("Every philosopher has eaten their fill")
}
func philosopher(number int) {
	eaten := 0
	for eaten < 3 {
		if canIEat(number) {
			mut.Lock()
			fmt.Println("Starting to eat ", number)
			useChopsticks(number, false)
			eaten++
			useChopsticks(number, true)
			fmt.Println("Finishing eating ", number)
			mut.Unlock()
		}
	}
	fmt.Printf("Philosopher %v has eaten %d times\n", number, eaten)
	wg.Done()
}
func canIEat(number int) bool {
	if number == 4 {
		return (chopsticks[4] && chopsticks[0])
	}
	return (chopsticks[number] && chopsticks[number+1])
}
func useChopsticks(number int, use bool) {
	if number == 4 {
		chopsticks[4], chopsticks[0] = use, use
	} else {
		chopsticks[number], chopsticks[number+1] = use, use
	}
}
