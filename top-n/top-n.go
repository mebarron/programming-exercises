/* Read and rank the top n values, such as count, from a file */
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

var proc_counter int
var elapsed time.Time

// Parse the values from the data we're expecting to receive
func parse(data string, delimeter string, index int) map[string]int {
	elapsed = time.Now()

	values := make(map[string]int)

	for _, line := range strings.Split(data, "\n") {
		proc_counter++

		contents := strings.Split(line, delimeter)

		if index > len(contents)-1 {
			index = len(contents) - 1
		}

		if _, exists := values[contents[index]]; exists {
			values[contents[index]] += 1
		} else {
			values[contents[index]] = 1
		}
	}
	return values
}

// Sort and rank top n values in descending order
func rank(values map[string]int, n int) {
	var keys []string

	for k, _ := range values {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return values[keys[i]] < values[keys[j]]
	})

	rank := 1
	index := len(values) - 1

	for i := n; i > 0; i-- {
		for j, _ := range keys {
			if index == len(values)-n-1 {
				break
			}
			j = index
			fmt.Println(rank, " ", keys[j], " ", values[keys[j]])
			index--
			rank++
		}
		if rank == len(keys) {
			break
		}
	}

	fmt.Println("Processed ", proc_counter, " lines in ", time.Now().Sub(elapsed).String())
}

func main() {
	data, err := ioutil.ReadFile("logfile.txt")
	if err != nil {
		panic(err)
	}

	values := parse(string(data), ",", 3)

	rank(values, 10)
}
