package main

import (
    "fmt"

    "github.com/adrixo/introduction-to-kubernetes/go/simple/library"
)

func main() {
    // Get a greeting message and print it.
    message := library.Function()
    fmt.Println(message)
}
