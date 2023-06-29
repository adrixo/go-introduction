//
// This is assignment for the week 4 of the
// Coursera's course "Concurrency in Go"
//
// Implement the dining philosopher’s problem with the following constraints/modifications.
//
// 1. There should be 5 philosophers sharing chopsticks,
// with one chopstick between each adjacent pair of philosophers.
//
// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
//
// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered
// first (which we did in lecture).
//
// 4. In order to eat, a philosopher must get permission from a host which
// executes in its own goroutine.
//
// 5. The host allows no more than 2 philosophers to eat concurrently.
//
// 6. Each philosopher is numbered, 1 through 5.
//
// 7. When a philosopher starts eating (after it has obtained necessary locks)
// it prints “starting to eat <number>” on a line by itself,
// where <number> is the number of the philosopher.
//
// 8. When a philosopher finishes eating (before it has released its locks)
// it prints “finishing eating <number>” on a line by itself,
// where <number> is the number of the philosopher.
//

package main

// *********************
// Imports
// *********************

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// *********************
// Types
// *********************

type Chopstick struct {
	index int
	mut   sync.Mutex
}
type Philosopher struct {
	index   int
	numEats int
	lChop   *Chopstick
	rChop   *Chopstick
}

// *********************
// Const
// *********************

const NUM_PHILOSOPHERS int = 5

const MAX_EATERS int = 2

const NUM_FOODS int = 3

// *********************
// Vars
// *********************

// used to sync eating process with main goroutine finalization
var wg sync.WaitGroup

// Host: array with each eater state (true, if eating. false, otherwise)
var eatState [NUM_PHILOSOPHERS]bool

// mutex is used to prevent simulteneous access to
// host state array
var mutexHost sync.Mutex

// *********************
// Funcs
// *********************

// used to inform host about eating state from each eater
func hostInformEating(index int, isEat bool) {
	mutexHost.Lock()
	eatState[index] = isEat
	mutexHost.Unlock()
}

// permission function by host: is allowed to start to eat.
// return faklse if currently already 2 eaters we have on the table
func hostIsAllowedToEat() bool {
	mutexHost.Lock()
	numEaters := 0
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		if eatState[i] == true {
			numEaters = numEaters + 1
		}
	}
	mutexHost.Unlock()

	if numEaters >= MAX_EATERS {
		return false
	}
	return true
}

func eat(p *Philosopher) {

	// fmt.Println("Philo ", p.index, " start eating ...")

	// each philosopher should eat only 3 times
	for p.numEats < NUM_FOODS {
		if hostIsAllowedToEat() {
			hostInformEating(p.index, true)
			// select random order of takin chopsticks:
			// (left, right) or (right, left)
			if rand.Intn(2) == 0 {
				p.lChop.mut.Lock()
				p.rChop.mut.Lock()
			} else {
				p.rChop.mut.Lock()
				p.lChop.mut.Lock()
			}

			// nee to print eaters numbers 1,2,.. instead of 0,1,...
			fmt.Println("starting to eat", p.index+1)

			// give some non-zero time for philosopher to eat his food
			time.Sleep(110 * time.Millisecond)

			fmt.Println("finishing eating", p.index+1)

			// release chopsiticks for other eaters on table
			p.rChop.mut.Unlock()
			p.lChop.mut.Unlock()
			// mark that philosopher has finished current food
			p.numEats = p.numEats + 1
			hostInformEating(p.index, false)
		}
	}
	// each philosopher decrease global wait counter to
	// check when main should be finished
	wg.Done()
}

func main() {
	// create sticks array
	sticks := make([]*Chopstick, NUM_PHILOSOPHERS)
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		sticks[i] = new(Chopstick)
		sticks[i].index = i
	}

	// create philosophers array and init it
	philos := make([]*Philosopher, NUM_PHILOSOPHERS)
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		philos[i] = new(Philosopher)
		philos[i].lChop = sticks[i]
		philos[i].rChop = sticks[(i+1)%NUM_PHILOSOPHERS]
		philos[i].index = i
		philos[i].numEats = 0
	}

	// init eating state array
	// fmt.Println("Philosophers are starting eating...")
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		eatState[i] = false
	}

	// init wait group with number of eaters
	wg.Add(NUM_PHILOSOPHERS)

	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		go eat(philos[i])
	}
	// wait until all eaters finished their food several time
	wg.Wait()
}
