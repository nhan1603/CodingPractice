package main

import (
	"fmt"
)

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 || length == 1 {
		return nums
	}
	var before []int = make([]int, length)
	var after []int = make([]int, length)
	before[0] = 1
	after[length-1] = 1
	for indx := 1; indx < length; indx++ {
		before[indx] = before[indx-1] * nums[indx-1]
		after[length-1-indx] = after[length-indx] * nums[length-indx]
	}
	var result []int = make([]int, length)
	for indx := range before {
		result[indx] = before[indx] * after[indx]
	}
	return result
}
