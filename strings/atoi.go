/* Converts a string to a 64 bit signed integer */

package main

import (
	"fmt"
)

func atoi(a string) int64 {
	max_digits := 19

	if a[:1] != "+" {
		if a[:1] != "-" {
			return 0
		}
	}

	if len(a)-1 > max_digits {
		return 0
	}

	var n uint64
	var b byte

	for i := 0; i < len(a); i++ {
		digit := a[i]

		if '0' <= digit && digit <= '9' {
			b = digit - '0'
		} else if 'a' <= digit && digit <= 'z' {
			b = digit - 'a' + 10
		} else if 'A' <= digit && digit <= 'Z' {
			b = digit - 'A' + 10
		}

		n *= uint64(10)
		n += uint64(b)
	}

	var number int64

	if a[:1] == "+" {
		number = int64(n)
	} else {
		number = int64(n)
		number = 0 - int64(n)
	}

	return number
}

func main() {
	a0 := "+123456789123456789"
	a1 := "-123456789123456789"

	str0 := atoi(a0)
	str1 := atoi(a1)

	fmt.Println(str0)
	fmt.Println(str1)
}
