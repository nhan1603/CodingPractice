package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func maxVowels(s string, k int) int {
	vowels := "aeiouAEIOU"
	var sum int
	var result int

	startString := s[:k]
	for _, char := range startString {
		if strings.Contains(vowels, string(char)) {
			sum++
		}
	}
	result = sum
	for i := k; i < len(s); i++ {
		if strings.Contains(vowels, string(s[i])) {
			sum++
		}
		if strings.Contains(vowels, string(s[i-k])) {
			sum--
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
