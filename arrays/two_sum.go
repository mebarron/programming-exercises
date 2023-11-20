/* Find the first instance of target sum t of values x,y in two arrays */

package main

import (
	"fmt"
	"sort"
)

func two_sum(a0 []int, a1 []int, t int) []int {
	var result []int

	sort.Ints(a0)
	sort.Ints(a1)

	for i := 0; i < len(a0); i++ {
		for j := (len(a1) - 1); j > 0; j-- {
			sum := a0[i] + a1[j]
			if sum == t {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}

	return result
}

func main() {
	var t int

	a0 := []int{5, 4, 6, 8, 3, 16, 9, 12, 29, 24, 13, 16, 11}
	a1 := []int{5, 4, 2, 7, 1, 6, 11, 10}

	sort.Ints(a0)
	sort.Ints(a1)

	fmt.Println(a0, a1, "\n")

	fmt.Scanln(&t)

	result := two_sum(a0, a1, t)

	fmt.Println(result)
}
