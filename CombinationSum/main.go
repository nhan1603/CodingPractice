package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var doBackTrack func(candidates []int, target int, current []int, currentIndex int, currentSum int)

	doBackTrack = func(candidates []int, target int, current []int, currentIndex int, currentSum int) {
		if currentSum == target {
			result = append(result, current)
			return
		} else if currentSum > target {
			return
		}

		for indx := currentIndex; indx < len(candidates); indx++ {
			newList := make([]int, len(current))
			copy(newList, current)
			newList = append(newList, candidates[indx])
			doBackTrack(candidates, target, newList, indx, currentSum+candidates[indx])
		}
	}

	doBackTrack(candidates, target, make([]int, 0), 0, 0)
	return result
}
