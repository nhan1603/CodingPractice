package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	sort.Ints(candidates)

	var doBackTrack func(candidates []int, target int, current []int, currentIndex int, currentSum int)

	doBackTrack = func(candidates []int, target int, current []int, currentIndex int, currentSum int) {
		if currentSum == target {
			result = append(result, current)
			return
		} else if currentSum > target {
			return
		}
		previous := 0
		for indx := currentIndex; indx < len(candidates); indx++ {
			if candidates[indx] == previous {
				continue
			}
			previous = candidates[indx]
			// copy of current list
			newList := make([]int, len(current))
			copy(newList, current)
			newList = append(newList, candidates[indx])
			doBackTrack(RemoveIndex(candidates, indx), target, newList, indx, currentSum+candidates[indx])
		}
	}

	doBackTrack(candidates, target, make([]int, 0), 0, 0)
	return result
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
