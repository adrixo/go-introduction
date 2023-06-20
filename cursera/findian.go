package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	println("Enter a string > ")
	fmt.Scan(&s)
	s = strings.ToLower(s)

	m := make(map[rune]bool)

	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case 'i':
			if i == 0 {
				m['i'] = true
			}
		case 'a':
			m['a'] = true
		case 'n':
			if i == len(s)-1 {
				m['n'] = true
			}
		}
	}

	if len(m) == 3 {
		fmt.Println("Found!")
		return
	}
	fmt.Println("Not Found!")
}
