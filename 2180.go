package leetcode_everyday

func countEven(num int) int {
	if num < 10 {
		return num / 2
	}
	if num == 1000 {
		return 499
	}
	x := num / 10
	base := 5 * x - 1
	b := num /100
	g := num % 10
	s := x-b*10
	add := 0
	if (g+s+b)%2 == 0 {
		add = (g + 2)/2
	} else {
		add = (g + 1)/2
	}
	return base + add
}

func countEvenv1(num int) int {
	ans := 0
	for i := 2; i <= num; i++ {
		sum := 0
		n := i
		for n > 0 {
			sum = sum + n%10
			n=n/10
		}
		if sum % 2 == 0{
			ans++
		}
	}
	return ans
}
