package practice

import (
	"fmt"
	"testing"
)

func TestGetChunkSize(t *testing.T) {
	var tests = []struct {
		slice                        []int
		expectedOffset, expectedSize int
	}{
		{[]int{}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1, 1}, 2, 0},
		{[]int{1, 1, 1}, 3, 0},
		{[]int{1, 1, 1, 1}, 0, 1},
		{[]int{1, 1, 1, 1, 1}, 1, 1},
		{[]int{1, 1, 1, 1, 1, 1}, 2, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 3, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 0, 2},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("%d, %d", i, test.slice)
		t.Run(testname, func(t *testing.T) {
			offset, size := getChunkSize(test.slice)
			if offset != test.expectedOffset || size != test.expectedSize {
				t.Fatalf(`offset %d != %d || size %d != %d`, offset, test.expectedOffset, size, test.expectedSize)
			}
		})
	}
}
