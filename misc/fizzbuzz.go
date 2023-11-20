package main

import (
	"fmt"
)

func fizzbuzz(n int32) {
	i := n % 3
	j := n % 5

	var str string

	if i == 0 {
		str = "fizz"
	}
	if j == 0 {
		str = str + "buzz"
	}

	if len(str) == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(str)
	}
}

func main() {
	for i := 1; i <= 50; i++ {
		fizzbuzz(int32(i))
	}
}
