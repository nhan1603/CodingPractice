package main

import (
	"fmt"
)

func main() {
	fmt.Println(areAlmostEqual("oh", "ho"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	listDiff := []int{}
	for indx := range s1 {
		if s1[indx] != s2[indx] {
			listDiff = append(listDiff, indx)
		}
	}
	if len(listDiff) == 0 {
		return true
	}

	if len(listDiff) == 2 {
		firstIndx := listDiff[0]
		secIndx := listDiff[1]
		return s1[firstIndx] == s2[secIndx] && s2[firstIndx] == s1[secIndx]
	}

	return false
}
