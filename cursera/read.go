package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
File example: names.json
{
  "names": [
    {
      "Name": "Mohamed",
      "Surname": "Ali"
    },
    {
      "Name": "Leonardo",
      "Surname": "Caprio"
    }
  ]
}
*/

type Person struct {
	Name    string
	Surname string
}

func main() {
	var fileName string

	fmt.Print("Filename: ")
	fmt.Scan(&fileName)

	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var result map[string][]Person

	json.Unmarshal(contents, &result)
	fmt.Println("Contents of file:", result)

	for i, person := range result["names"] {
		fmt.Println("Person number ", i)
		fmt.Println(person.Name)
		fmt.Println(person.Surname)
		fmt.Println()
	}
}
