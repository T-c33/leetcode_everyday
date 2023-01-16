package main

func countDifferentSubsequenceGCDs(nums []int) (ans int) {
	maxVal := 0
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	occured := make([]bool, maxVal+1)
	for _, num := range nums {
		occured[num] = true
	}
	for i := 1; i <= maxVal; i++ {
		subGcd := 0
		for j := i; j <= maxVal; j += i {
			if occured[j] {
				if subGcd == 0 {
					subGcd = j
				} else {
					subGcd = gcd(subGcd, j)
				}
				if subGcd == i {
					ans++
					break
				}
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func gcd(num1, num2 int) int {
	for num1 != 0 {
		num1, num2 = num2%num1, num1
	}
	return num2
}
