/* Finds matches of a pattern in a string using regular expressions */

package main

import (
	"fmt"
	"log"
	"regexp"
)

func create_pattern(w string) string {
	var r []rune

	for _, v := range w {
		r = append(r, v)
	}

	c0 := r[0]
	c1 := r[len(w)-1]

	pattern := string(c0) + "([a-z]+)" + string(c1)

	return pattern
}

func match(pattern string, s string) bool {
	result, err := regexp.MatchString(pattern, s)

	if err != nil {
		log.Panic(err)
	}

	r, err := regexp.Compile(pattern)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Pattern:", r)

	if r.MatchString(s) == true {
		result = true
	}

	if len(r.FindString(s)) >= 1 {
		index := r.FindStringIndex(s)
		fmt.Println(index)
		result = true
	}

	if len(r.FindStringSubmatch(s)) >= 1 {
		index := r.FindStringSubmatchIndex(s)
		fmt.Println(index)
		result = true
	}

	return result
}

func main() {
	s := "This is a string with seven words"

	w := "words"

	pattern := create_pattern(w)

	r := match(pattern, s)

	fmt.Println("String:", s, "\nMatch?", r)
}
