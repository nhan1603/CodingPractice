package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}

func maxOperations(nums []int, k int) int {
	var left = 0
	var right = len(nums) - 1
	if right == -1 {
		return 0
	}
	operationCount := 0
	slices.Sort(nums)
	for left < right {
		if nums[left]+nums[right] == k {
			operationCount++
			left++
			right--
		} else if nums[left]+nums[right] > k {
			right--
		} else if nums[left]+nums[right] < k {
			left++
		}
	}
	return operationCount
}
