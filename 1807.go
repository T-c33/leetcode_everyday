package main

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	replaceMap := make(map[string]string, 0)
	for _, kv := range knowledge {
		replaceMap[kv[0]] = kv[1]
	}
	fmt.Printf("%v\n", replaceMap)
	firstStep := strings.Split(s, "(")
	for _, s1 := range firstStep {
		if s1 == "" {
			continue
		}
		secondStep := strings.Split(s1, ")")
		var replace string
		toReplace := "(" + secondStep[0] + ")"
		fmt.Printf("to replace: %v\n", toReplace)

		if t, ok := replaceMap[secondStep[0]]; ok {
			replace = t
		} else {
			replace = "?"
		}
		fmt.Printf("replace: %v\n", replace)

		s = strings.Replace(s, toReplace, replace, -1)
		fmt.Printf("s: %v\n", 1)

	}
	return s
}

func evaluate1(s string, knowledge [][]string) string {
	var m map[string]string = make(map[string]string)
	for _, vec := range knowledge {
		m[vec[0]] = vec[1]
	}
	vec := strings.Split(s, "(")
	res := ""
	for _, sub := range vec {
		vecsub := strings.Split(sub, ")")
		if len(vecsub) == 1 {
			res += sub
		} else {
			if ele, ok := m[vecsub[0]]; !ok {
				res += "?"
			} else {
				res += ele
			}
			res += vecsub[1]
		}
	}
	return res

}
