/* Go implementation of a stack data structure */

package main

import (
	"fmt"
)

type Stack []string

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() string {
	var item string

	if len(*s) > 0 {
		item = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return item
	} else {
		item = ""
		return item
	}
}

func main() {
	i0 := "First"
	i1 := "Second"
	i2 := "Third"

	var stack Stack

	stack.push(i0)
	stack.push(i1)
	stack.push(i2)

	for i, _ := range stack {
		fmt.Println(stack.pop())
		i++
	}
}
