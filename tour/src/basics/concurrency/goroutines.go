package main

import (
	"fmt"
	"time"
)

func say(s string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("worlds", 50)
	say("hello", 5)
}
