package main

import (
	"fmt"
	"sort"
)

func main() {
	l := []int{9, 2, 6, 4, 22, 6, 70, 8, 9}
	l = channelSort(l)
	fmt.Println(l)
}

func channelSort(l []int) []int {
	// 1. divide slice in chunks
	cutter_c := make(chan []int)
	go getSubSlices(l, cutter_c)

	// 2. Sort divided slices
	sorter_c := make(chan []int)
	go sortSlice(cutter_c, sorter_c)

	// 3. Get all slices into var
	subSlices := make([][]int, 4)
	for value := range sorter_c {
		fmt.Println("Getting one", value)
		subSlices = append(subSlices, value)
	}
	// 4. Merge slices
	sortedSlice := []int{}
	for _, slice := range subSlices {
		sortedSlice = merge(sortedSlice, slice)
	}

	return sortedSlice
}

func insert(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func merge(host []int, minions []int) []int {
	for _, e := range minions {
		insertIndex := 0
		for _, e2 := range host {
			if e > e2 {
				insertIndex++
				continue
			}
		}
		host = insert(host, insertIndex, e)
	}
	return host
}

func sortSlice(consume chan []int, dispatch chan []int) {
	for slice := range consume {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		fmt.Println("sorted one", slice)
		dispatch <- slice
	}
	close(dispatch)
}

func getSubSlices(l []int, c chan []int) {
	chunkSize := len(l) / 4
	offset := len(l) % 4

	for chunkSize < len(l) {
		extra := 0
		if offset > 0 {
			extra = 1
			offset--
		}

		slice := l[0 : chunkSize+extra]
		fmt.Println("slicing one", slice)
		c <- slice
		l = l[chunkSize+extra:]
	}
	if len(l) > 0 {
		c <- l
	}

	close(c)
}
