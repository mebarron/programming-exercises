/* Parses text from a file and counts/ranks the words in the file */

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Remove any special characters and get the distinct count of each word
func wordcount(c string) map[string]int {
	words := strings.Split(c, " ")

	sc := []string{",", ".", ":", ";", "|", "`", "\"", "!", "?", "'", "[", "]", "(", ")", "{", "}", "-", "_", "@", "#", "$", "%", "^", "&", "*", "<", ">", "/", "~", "\n"}

	wc := make(map[string]int)

	for _, word := range words {
		if word == "" {
			continue
		}

		if word == "\n" {
			continue
		}

		word = strings.ToLower(word)
		word = strings.ReplaceAll(word, " ", "")
		for _, c := range sc {
			word = strings.ReplaceAll(word, c, "")
		}
		if _, seen := wc[word]; seen {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}

	return wc
}

// Rank the top n words in the file in descending order
func rank(wc map[string]int, n int) {
	var keys []string

	for k, _ := range wc {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return wc[keys[i]] < wc[keys[j]]
	})

	rank := 1
	index := len(wc) - 1

	for i := n; i > 0; i-- {
		for j, _ := range keys {
			if index == len(wc)-n-1 {
				break
			}
			j = index
			fmt.Println(rank, " ", keys[j], " ", wc[keys[j]])
			index--
			rank++
		}
		if rank == len(keys) {
			break
		}
	}
}

func main() {
	c, err := os.ReadFile("comedy_of_errors.txt")
	if err != nil {
		panic(err)
	}

	wc := wordcount(string(c))

	rank(wc, 50)
}
