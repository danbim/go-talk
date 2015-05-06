package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		cnt, ok := wordMap[word]
		if !ok {
			wordMap[word] = 1
		} else {
			wordMap[word] = cnt + 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
