package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSequentialSort(t *testing.T) {
	slice := []int{9, 2, 6, 4, 22, 6, 70, 8, 9}
	expected := []int{2, 4, 6, 6, 8, 9, 9, 22, 70}

	slice = channelSort(slice)
	if !Equal(slice, expected) {
		fmt.Println("Got ", slice)
		fmt.Println("Expected ", expected)
		t.Fatal("The array is not sorted")
	}
}
