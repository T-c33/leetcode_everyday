package main

func reinitializePermutation(n int) int {
	target := make([]int, n)
	for i := range target {
		target[i] = i
	}
	perm := append([]int(nil), target...)
	ans := 0
	for {
		ans++
		arr := make([]int, n)
		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+i/2]
			}
		}
		perm = arr
		if equal(perm, target) {
			return ans
		}
	}
}

func equal(l1, l2 []int) bool {
	for i, val := range l1 {
		if val != l2[i] {
			return false
		}
	}
	return true
}
