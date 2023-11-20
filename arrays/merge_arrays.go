/* Merges two arrays of integers into a sorted array */

package main

import (
	"fmt"
	"sort"
)

func merge(a []int, b []int) []int {
	var m []int

	m = append(m, a...)
	m = append(m, b...)

	sort.Ints(m)

	return m
}

func main() {
	a := []int{1, 1, 4, 5, 6, 9, 11}
	b := []int{2, 7, 8, 12, 13, 15, 16, 5}

	m := merge(a, b)

	fmt.Println(m)
}
