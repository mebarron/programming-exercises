/* Uses the folding method to generate a hash of some string */

package main

import (
	"fmt"
	"math/rand"
)

func fold(s string, km uint64) (uint64, uint64) {
	ht := make([]uint64, len(s)/8)

	var k uint64

	if km == 0 {
		k = rand.Uint64()
	} else {
		k = km
	}

	index := 0

	if len(s) < 8 {
		for _, v := range s {
			n := (uint64(v)) + k
			ht = append(ht, n)
		}
	} else {
		for i := 0; i < len(ht); i++ {
			n := (uint64(s[index]) + uint64(s[index+1]) +
				uint64(s[index+2]) + uint64(s[index+3]) +
				uint64(s[index+4]) + uint64(s[index+5]) +
				uint64(s[index+6]) + uint64(s[index+7])) + k
			ht[i] = n
			index = index + 8
		}
	}

	if len(s)-index < 8 {
		if len(s)-index > 0 {
			for j := index; j < len(s); j++ {
				n := (uint64(s[index]) + k)
				ht = append(ht, n)
				index = index + 1
			}
		}
	}

	var h uint64

	for _, v := range ht {
		h = h + uint64(v)
	}

	return h, k
}

func verify(s string, hash uint64, k uint64) (uint64, uint64, bool) {
	h, new_key := fold(s, k)

	var v bool

	if h == hash {
		if new_key == k {
			v = true
		}
	} else {
		v = false
	}

	return h, new_key, v
}

func main() {
	s0 := "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand."

	s1 := "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand."

	h, k := fold(s0, 0)

	fmt.Print(h, " ", k, "\n")

	fmt.Println(verify(s1, h, k))
}
