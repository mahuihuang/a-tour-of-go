package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	splits := strings.Fields(s)
	for _, word := range splits {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
