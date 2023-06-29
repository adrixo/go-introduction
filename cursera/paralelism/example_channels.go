package main

import (
	"fmt"
	"time"
)

type prod struct {
	c chan int
}

func (p *prod) Work() {
	for i := 0; i < 10; i++ {
		p.c <- i
		time.Sleep(500 * time.Millisecond)
	}
}

type consumer struct {
	c  chan int
	id int
}

func (c *consumer) Work() {
	for i := range c.c {
		fmt.Println("Im ", c.id, " and i received: ", i)
	}
}

func main() {
	c := make(chan int)
	prod := &prod{
		c,
	}
	consumer1 := &consumer{
		c,
		1,
	}
	consumer2 := &consumer{
		c,
		2,
	}

	go prod.Work()
	go consumer1.Work()
	go consumer2.Work()
	time.Sleep(10 * time.Second)
}
