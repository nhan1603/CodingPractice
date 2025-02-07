package main

import (
	"fmt"
)

func main() {
	// fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	fmt.Println(increasingTriplet([]int{10, 12, 9, 11, 3, 13}))
}

func increasingTriplet(nums []int) bool {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	if len(nums) < 3 {
		return false
	}

	var first, second int
	first = MaxInt
	second = MaxInt
	for _, num := range nums {
		if num <= first { // before finding the correct second, anything smaller than second can be the first
			first = num
		} else if num <= second {
			second = num
		} else { // reach this means the index is greater and the seconds, forming an increasing triplet
			return true
		}
	}

	return false
}
