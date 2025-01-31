package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if !strings.HasPrefix(str1, str2) && !strings.HasPrefix(str2, str1) {
		return ""
	}

	if str1 == "" || str2 == "" {
		return ""
	}

	var checkString string
	if len(str1) >= len(str2) {
		checkString = str2
	} else {
		checkString = str1
	}
	var result string

	for i := 1; i < len(checkString)+1; i++ {
		considerString := checkString[:i]
		if isDivisible(str1, considerString) && isDivisible(str2, considerString) {
			result = considerString
		}
	}

	return result
}

func isDivisible(s, t string) bool {
	if len(s)%len(t) != 0 {
		return false
	}
	expected := strings.Repeat(t, len(s)/len(t))
	return expected == s
}
