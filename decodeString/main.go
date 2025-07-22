package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

func decodeString(s string) string {
	var stack string
	number := "0123456789"
	for indx := 0; indx < len(s); indx++ {
		val := s[indx]
		if val == ']' {
			cut := len(stack) - 1
			for stack[cut] != '[' {
				cut--
			}
			stringPart := stack[cut+1:]
			stack = stack[:cut]
			for cut > 0 && strings.Contains(number, string(stack[cut-1])) {
				cut--
			}
			numPart := stack[cut:]
			stack = stack[:cut]
			count, _ := strconv.Atoi(numPart)
			for i := 0; i < count; i++ {
				stack = stack + stringPart
			}
		} else {
			stack = stack + string(val)
		}
	}

	return stack
}
