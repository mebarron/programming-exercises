/* Similar implementation as two-sum, except with three arrays */

package main

import (
	"fmt"
)

func three_sum(arry [][]int, t int) map[int]int {

	r := make(map[int]int)

	if len(arry) < 3 {
		return r
	}

	a := arry[0]
	b := arry[1]
	c := arry[2]

	for i := 0; i < len(a); i++ {
		for j := len(b) - 1; j >= 0; j-- {
			for k := 0; k < len(c); k++ {
				if a[i]+b[j]+c[k] == t {
					r[i] = a[i]
					r[j] = b[j]
					r[k] = c[k]

					return r
				}
			}
		}
	}

	return r

}

func main() {

	arr0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	arr2 := []int{19, 20, 21, 22, 23, 24, 25, 26, 27}

	arrs := make([][]int, 3)

	arrs[0] = arr0
	arrs[1] = arr1
	arrs[2] = arr2

	t := 45

	fmt.Println(three_sum(arrs, t))

}
