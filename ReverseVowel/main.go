package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	listChar := []rune{}

	for _, char := range s {
		if strings.Contains(vowels, string(char)) {
			listChar = append(listChar, char)
		}
	}

	runeList := []rune(s)
	var x rune
	for indx, runee := range runeList {
		if strings.Contains(vowels, string(runee)) {
			x, listChar = listChar[len(listChar)-1], listChar[:len(listChar)-1]
			runeList[indx] = x
		}
	}

	return string(runeList)
}
