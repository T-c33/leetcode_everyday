package main

func digitCount(num string) bool {
	numMap := map[rune]int{}
	for _, r := range num {
		numMap[r-'0']++
	}
	for i, c := range num {
		if numMap[rune(i)] != int(c-'0') {
			return false
		}
	}
	return true
}
