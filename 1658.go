package main

import "fmt"

func minOperations(nums []int, x int) int {
	sum := 0
	var ans []int
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Printf("sum : %v\n", sum)
	if sum < x {
		return -1
	}
	extra := sum - x
	if extra == 0 {
		return len(nums)
	}
	fmt.Printf("extra : %v\n", extra)

	left, right := -1, 0
	newSum := nums[0]
	for right < len(nums)  && left < len(nums)  {
		if newSum < extra {
			right++
			if right < len(nums){
				newSum = newSum + nums[right]
			}
		}
		if newSum > extra {
			newSum = newSum - nums[left+1]
			left++
		}
		if newSum == extra {
			fmt.Printf("left :%v, right : %v\n", left, right)
			ans = append(ans, right - left)
			right++
			if right < len(nums){
				newSum = newSum + nums[right]
			}
		}
	}
	fmt.Printf("%v\n", ans)
	if len(ans) == 0 {
		return -1
	}
	max := 0
	for _, s := range ans {
		if s > max {
			max = s
		}
	}
	return len(nums)-max
}
