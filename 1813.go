package main

import (
	"strings"
)

func areSentencesSimilar(sentence1, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")
	i, n := 0, len(words1)
	j, m := 0, len(words2)
	for i < n && i < m && words1[i] == words2[i] {
		i++
	}
	for j < n-i && j < m-i && words1[n-j-1] == words2[m-j-1] {
		j++
	}
	return i+j == minf(n, m)
}

func minf(a, b int) int {
	if a > b {
		return b
	}
	return a
}
