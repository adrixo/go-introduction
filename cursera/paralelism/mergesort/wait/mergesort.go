package main

import (
	"sort"
	"sync"
)

func main() {
	l := []int{9, 2, 6, 4, 22, 6, 70, 8, 9}
	cSort(l)
}

func cSort(l []int) []int {
	// 1. Divide slices
	subSlices := getSubSlices(l)

	var wg sync.WaitGroup
	// 2. Sort divided slices
	for _, slice := range subSlices {
		go sortSlice(slice, &wg)
	}
	wg.Wait()

	// 3. Merge divided slices
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

func sortSlice(slice []int, wg *sync.WaitGroup) {
	wg.Add(1)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	wg.Done()
}

func getSubSlices(l []int) [][]int {
	chunkSize := len(l) / 4
	offset := len(l) % 4

	chunks := make([][]int, 0)
	for chunkSize < len(l) {
		extra := 0
		if offset > 0 {
			extra = 1
			offset--
		}
		chunks = append(chunks, l[0:chunkSize+extra])
		l = l[chunkSize+extra:]
	}
	if len(l) > 0 {
		chunks = append(chunks, l)
	}
	return chunks
}
