package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	newString := strings.TrimSpace(s)
	listWord := strings.Fields(newString)
	result := ""
	for _, word := range listWord {
		result = word + " " + result
	}
	return strings.TrimSpace(result)
}
