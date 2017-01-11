package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

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

	return m
}

func main() {
	s := "I and I am learning Go!"
	fmt.Println(WordCount(s))
}
