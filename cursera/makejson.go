package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	m := make(map[string]string)

	go func() {
		<-c
		fmt.Println("")

		jsonObj, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(jsonObj))
		os.Exit(0)
	}()

	var name string
	var address string
	for {
		fmt.Println("Ctrl+c to cancel")
		fmt.Print("Enter name: ")
		fmt.Scan(&name)
		fmt.Print("Enter address: ")
		fmt.Scan(&address)
		m[name] = address
	}

}
