package main

import "fmt"

func main() {
	ans := evaluate("(name)is(age)yearsold", [][]string{
		{"name", "bob"}, {"age", "two"},
	})
	fmt.Printf("ans : %v", ans)

}
