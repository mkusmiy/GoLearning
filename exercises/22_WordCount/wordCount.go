package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		wordMap[word]++
	}

	/*
		var str []string
		str = strings.Fields(s)
		for i := range str {
			b := str[i]
			if _, ok := m[b]; ok {
				m[b]++
			} else {
				m[b] = 1
			}
		}
	*/
	return wordMap
}

func main() {
	s := "I and I am learning Go!"
	fmt.Println(WordCount(s))
}
