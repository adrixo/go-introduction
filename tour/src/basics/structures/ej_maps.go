package main

/*
Implement WordCount. It should return a map of the counts of
each “word” in the string s. The wc.Test function runs a
test suite against the provided function and prints success
or failure.

You might find strings.Fields helpful.
*/

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	fmt.Println(words)

	var m map[string]int = make(map[string]int)

	for _, v := range words {
		m[v] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
