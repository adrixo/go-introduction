package main

import (
	"fmt"
	"sync"
	"time"
)

type Common struct {
	sync.Mutex
}

type Executor struct {
	c  Common
	id int
}

func (e Executor) run() {
	fmt.Printf("starting %d\n", e.id)
	e.c.Lock()
	fmt.Printf("executing task on %d\n", e.id)
	time.Sleep(2*time.Second)
	e.c.Unlock()
	fmt.Printf("finish %d\n", e.id)

}

func main() {
	var com Common

	com.

}
