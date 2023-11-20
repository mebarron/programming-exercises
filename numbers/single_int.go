/* Find the single instance of an integer */

package main

import (
	"fmt"
)

func single_int(a []int32) int32 {
	nums := make(map[int32]int)

	for _, num := range a {
		if _, twice := nums[num]; twice {
			nums[num] += 1
		} else {
			nums[num] = 1
		}
	}

	var n int32
    
    for num, count := range nums {
		if count == 1 {
			n = num
		}
	}

	return n
}

func main() {
	a := []int32{1, 2, 3, 7, 3, 2, 1, 8, 9, 10, 4, 2, 2, 6, 7, 11, 4, 6, 8, 9, 10}

	fmt.Println(single_int(a))

}