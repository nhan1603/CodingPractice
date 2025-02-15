package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aaa"))
}

func partition(s string) [][]string {
	var result = make([][]string, 0)
	if len(s) == 1 {
		return [][]string{{s}}
	}
	var doBackTrack func(current []string, indx int, s string)
	doBackTrack = func(current []string, indx int, s string) {
		if len(s) == 0 {
			result = append(result, current)
			return
		}

		// copy of current list
		newList := make([]string, len(current))
		copy(newList, current)
		add := s[:indx]
		newPart := s[indx:]
		if isPalindrome(add) {
			fmt.Println(add)
			newList = append(newList, add)
			doBackTrack(newList, 1, newPart)
		}
		if indx < len(s) {
			doBackTrack(current, indx+1, s)
		}
	}

	doBackTrack([]string{}, 1, s)
	return result
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
