package main

import (
	"fmt"
)

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	result := make([]bool, len(candies))
	for indx, candy := range candies {
		if candy+extraCandies >= max {
			result[indx] = true
		}
	}
	return result
}
