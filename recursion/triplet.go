/* Finds all the triplets where i+j+k = t using examples of recursion and shifting the
elements of an array */

package main

import (
	"encoding/json"
	"fmt"
)

var m [][]int
var r []string

func chkarray(a []int, l int) int {
	if l/3 <= 0 {
		return l
	}

	if l/3 != 1 {
		chk := chkarray(a, l-3)

		return chk
	} else {
		return l
	}

	return l
}

func shiftleft(a []int, lim int) ([]int, int) {
	l := len(a) - 1

	copy_arr := make([]int, l+1)

	copy_arr[l] = a[0]

	for i := 1; i <= l; i++ {
		copy_arr[i-1] = a[i]
	}

	if lim == 0 {
		return copy_arr, lim

	} else {
		p, _ := shiftleft(copy_arr, lim-1)

		m = append(m, p)
	}

	return copy_arr, lim
}

func shiftright(a []int, lim int) ([]int, int) {
	l := len(a) - 1

	copy_arr := make([]int, l+1)

	copy_arr[0] = a[l]

	for i := l; i >= 1; i-- {
		copy_arr[i] = a[i-1]
	}

	if lim == 0 {
		return copy_arr, lim
	} else {
		p, _ := shiftright(copy_arr, lim-1)

		m = append(m, p)
	}

	return copy_arr, lim
}

func split(a []int) [][]int {
	chk := chkarray(a, len(a))

	var arrays [][]int
	var l int

	if (chk / 1) == 3 {
		_, l = shiftleft(a, len(a))
		_, l = shiftright(a, len(a))

		i := 0

		if l != 0 {
			for _, v := range m {
				for i != len(v) {
					n := make([]int, 3)

					n[0] = v[i]
					n[1] = v[i+1]
					n[2] = v[i+2]

					arrays = append(arrays, n)

					i = i + 3
				}

				i = 0
			}
		}
	}

	return arrays
}

func triplet(a []int, t int) map[int][]int {
	arrays := split(a)

	results := make(map[int][]int)

	for _, k := range arrays {
		if k[0]+k[1]+k[2] != t {
			continue
		}
		if k[0]+k[1]+k[2] == t {
			results[t] = k
		}
	}

	return results
}

func search(a []int, i int) {
	result := triplet(a, i)

	res, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	if len(result) != 0 {
		r = append(r, string(res))
	}
}

func main() {
	a := []int{-1, 0, 1, 2, -1, 4, 9, 7, 8, 11, 10, 4, 6, 17, 2, 11, 7, 7, 22, 24, 25, 71, 34, 19, 46, 33, 31, 8, 6, -5, -9, -11, -5, 67, 89, 44, 32, 86, 15, 16, 17, 100}

	for i := 0; i < 100; i++ {
		search(a, i)
	}

	fmt.Println(r)
}
