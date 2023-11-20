/* Shifts elements of an array left and right using recursion */

package main

import (
	"fmt"
)

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
		fmt.Println(p)
	}

	return copy_arr, lim
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
		fmt.Println(p)
	}

	return copy_arr, lim
}

func main() {
	arr0 := []int{-1, 0, 1, 2, -1, 4, 9, 7, 8, 11, 10, 4, 6, 17, 2, 11, 7, 7, 22, 24, 25, 71, 34, 19, 46, 33, 31, 8, 6, -5, -9, -11, -5, 67, 89, 44, 32, 86, 15, 16, 17, 100}

	_, shift0 := shiftleft(arr0, len(arr0))

	_, shift1 := shiftright(arr0, len(arr0))

	fmt.Println(shift0 + shift1)
}
