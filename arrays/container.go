/* Solves the container problem, where a := h x w, but you cannot tilt the
container, what is the maximum amount of x that can go in the container */

package main

import (
	"fmt"
	"sort"
)

// Search for the side that is a variable number of units shorter including repeating integers
func sides_diff() {
	h := []int{1, 8, 6, 36, 5, 4, 4, 29, 3, 7, 21, 18, 46, 36, 29, 11, 46}

	copy := h

	sort.Ints(copy)

	height := copy[len(copy)-1]

	w := 0

	for i := len(copy) - 1; i > 0; i-- {
		if copy[i-1] != copy[len(copy)-1] {
			height = copy[i-1]
			w = i
			break
		}
	}

	a := height * w

	fmt.Println("Side 0:", height)
	fmt.Println("Side 1:", copy[len(copy)-1])
	fmt.Println(a)
}

// Search for the side that is x units shorter 
func sides(x int) {
	h := []int{1, 2, 8, 2, 5, 4, 3, 3, 7, 10, 9}

	height := h[0]

	copy := h
	w := 0

	sort.Ints(copy)

	for i := 0; i < len(copy); i++ {
		for j := len(copy) - 1; j > 0; j-- {
			if copy[i] == copy[len(copy)-1]-x {
				height = copy[i]
				w = i
				break
			}
		}
	}

	a := height * w

	if a == 0 {
		fmt.Println("Won't work.")
	} else {
		fmt.Println("Side 0:", height)
		fmt.Println("Side 1:", copy[len(copy)-1])
		fmt.Println(a)
	}
}

func main() {
	sides(1)
	sides_diff()
}
