package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	var result float64
	var sum int
	for i := 0; i < k; i++ {
		sum = sum + nums[i]
	}
	result = float64(sum) / float64(k)
	if k == len(nums) {
		return result
	}

	start := 1
	end := k
	for end < len(nums) {
		sum = sum + nums[end] - nums[start-1]
		temp := float64(sum) / float64(k)
		if temp > result {
			result = temp
		}
		start++
		end++
	}
	return result
}
