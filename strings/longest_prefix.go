/* Find the longest common prefix in an array of strings */

package main

import (
	"fmt"
	"sort"
)

func compare(s []string) string {
	var str string

	sort.Strings(s)

	for i := 0; i < len(s); i++ {
		w0 := s[i]
		w1 := s[len(s)-1]

		k := count(w0)

		for j := 0; j < k; j++ {
			if w0[j] != w1[j] {
				str = w0[:j]

				return str
			}
		}
	}

	return str
}

func prefix(sl []string) {
	var pref string

	pref = compare(sl)
	if pref == "" {
		fmt.Println("No common prefix.")
	} else {
		fmt.Println("Longest prefix: ", pref)
	}
}

func count(w string) int {
	c := len(w)

	return c
}

func main() {

	a := []string{"flower", "flow", "flight", "flumpy", "florida", "flex"}
	b := []string{"dog", "racecar", "macbook"}
	c := []string{"clumpy", "clasp", "cling", "clothesline"}
	d := []string{"dddan", "ddddann", "ddddef", "ddddxj"}

	fmt.Println("\n")

	prefix(a)
	prefix(b)
	prefix(c)
	prefix(d)
}
