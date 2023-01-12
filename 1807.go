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
