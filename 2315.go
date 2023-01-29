package main

func countAsterisks(s string) int {
	var lCount, starCount int
	for _, r := range s {
		if r == rune('|') {
			lCount++
		}
		if lCount%2 == 0 {
			if r == rune('*') {
				starCount++
			}
		}
	}
	return starCount
}
