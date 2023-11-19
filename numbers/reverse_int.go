/* Reverse the digits in an integer */

package main

import (
	"fmt"
)

func reverse(i int) int {
	r := 0
	for i != 0 {
		r = r*10 + i%10
		i = i / 10
	}
	return r
}

func main() {

	num := -321

	x := reverse(num)

	fmt.Println(x)

}
