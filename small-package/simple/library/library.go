package library

import (
 "fmt"   
 "errors"
)

func Function() string {
    message := "Hello, World!"
    return message
}

// Hello returns a greeting for the named person.
func Hello(name string) (message string, e error) {
    if (name == "") {
        return name, errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    message = fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, e
}
