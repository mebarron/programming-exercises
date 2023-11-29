package main

import (
	"fmt"
)

func check(s string) bool {
	lrunes := []rune{'(', '[', '{'}
	rrunes := []rune{')', ']', '}'}

	var runesets [][]rune
	result := false

	l := len(s)

	for i := 0; i < l; i += 2 {
		if i >= l {
			break
		}

		var runeset []rune
		runeset = append(runeset, rune(s[i]))
		runeset = append(runeset, rune(s[i+1]))

		runesets = append(runesets, runeset)
	}

	for _, runeset := range runesets {
		for i := 0; i < len(runeset); i++ {
			if runeset[i] == lrunes[0] {
				if runeset[i+1] == rrunes[0] {
					result = true
				} else {
					return false
				}
			}

			if runeset[i] == lrunes[1] {
				if runeset[i+1] == rrunes[1] {
					result = true
				} else {
					return false
				}
			}

			if runeset[i] == lrunes[2] {
				if runeset[i+1] == rrunes[2] {
					result = true
				} else {
					return false
				}
			}
		}
	}

	return result
}

func main() {
	s0 := "()"
	s1 := "()[]{}"
	s2 := "(]"
	s3 := "()[]{][)"
	s4 := "()[]{}()[]{}()()()[][][]{}{}{}"
	s5 := "()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}"

	fmt.Println(check(s0))
	fmt.Println(check(s1))
	fmt.Println(check(s2))
	fmt.Println(check(s3))
	fmt.Println(check(s4))
	fmt.Println(check(s5))
}
