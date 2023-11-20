/* Functions over arrays, along with min, max, sum, mean, median */

package main

import (
	"fmt"
	"sort"
)

// Combine and dedup two arrays
func dedup(a0 []int, a1 []int) []int {
	var a2 []int

	a2 = append(a2, a0...)
	a2 = append(a2, a1...)

	sort.Ints(a2)

	for i, v := range a2 {
		var n int
		if i+1 == len(a2) {
			break
		}
		n = a2[i+1]
		if n == v {
			a2 = remove_element(a2, i+1)
		}
	}

	var copy_arr []int

	for j, v := range a2 {
		if j == len(a2)-1 {
			if a2[j] != 0 {
				copy_arr = append(copy_arr, v)
			}
		}
		if j+1 < len(a2) {
			if v != a2[j+1] {
				if a2[j] != 0 {
					copy_arr = append(copy_arr, v)
				}
			}
		}
	}

	a2 = copy_arr

	return a2
}

func remove_element(a []int, index int) []int {
	var new_arr []int

	for i, v := range a {
		if i == index {
			continue
		} else {
			new_arr = append(new_arr, v)
		}
	}

	return new_arr
}

func add_element(a []int, e int) []int {
	a = append(a, e)
	sort.Ints(a)

	return a
}

func mean(a0 []int, a1 []int) int {
	var x int

	a2 := dedup(a0, a1)

	x = sum(a2)

	x = x / (len(a2))

	return x
}

func median(a0 []int, a1 []int) int {
	var x int

	a2 := dedup(a0, a1)

	index := len(a2) / 2

	x = a2[index]

	return x
}

func sum(a []int) int {
	var x int
	for _, v := range a {
		x = x + v
	}
	return x
}

func main() {
	a0 := []int{5, 4, 6, 8, 3, 6, 11, 6, 9, 17, 5, 21, 22, 10}
	a1 := []int{8, 4, 2, 7, 1, 13, 28, 36, 21, 22, 25, 26, 27, 19, 18, 16, 13}

	a2 := dedup(a0, a1)

	md := median(a0, a1)
	mn := mean(a0, a1)
	sm := sum(a2)

	min := a2[0]
	max := a2[len(a2)-1]

	fmt.Println("Array: ", a2)
	fmt.Println("Median: ", md)
	fmt.Println("Mean: ", mn)
	fmt.Println("Sum: ", sm)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}