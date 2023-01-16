package main

func minMaxGame(nums []int) int {

	newNums := nums
	for len(newNums) != 1 {
		var newNumsCopy []int
		l := len(newNums)
		for i := 0; i < l-1; i = i + 2 {
			if i/2%2 == 0 {
				newNumsCopy = append(newNumsCopy, minFunc(newNums[i], newNums[i+1]))
			}
			if i/2%2 == 1 {
				newNumsCopy = append(newNumsCopy, maxFunc(newNums[i], newNums[i+1]))
			}
		}
		newNums = newNumsCopy
	}
	return newNums[0]
}

func minFunc(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFunc(a, b int) int {
	if a < b {
		return b
	}
	return a
}
