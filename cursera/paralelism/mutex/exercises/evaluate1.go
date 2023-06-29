package main

import (
	"fmt"
	"sync"
)

type Turn struct {
	mu   sync.Mutex
	phis []int
}

func RequestEat(turn *Turn, phis int, wg *sync.WaitGroup) {
	turn.mu.Lock()
	turn.phis = append(turn.phis, phis)
	if len(turn.phis)%2 == 0 || len(turn.phis) == 15 {
		t := (len(turn.phis) + 1) / 2
		fmt.Printf("Turn %d:\n", t)

		fmt.Printf("starting to eat %d\n", turn.phis[2*t-2])
		if len(turn.phis) != 15 {
			fmt.Printf("starting to eat %d\n", turn.phis[2*t-1])
		}

		fmt.Printf("finishing eating %d\n", turn.phis[2*t-2])
		if len(turn.phis) != 15 {
			fmt.Printf("finishing eating %d\n\n", turn.phis[2*t-1])
		}
	}
	turn.mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup              // Used for program finalization
	turn := Turn{phis: make([]int, 0)} // Mutex and list of philosofers

	for j := 0; j < 3; j++ {
		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go RequestEat(&turn, i, &wg)
		}
	}

	wg.Wait()
}
