package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(punishmentNumber(10))
}

func punishmentNumber(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		if (i%9 == 0 || i%9 == 1) && satisfiedCondition(i) {
			result = result + i*i
		}
	}
	return result
}

func satisfiedCondition(n int) bool {
	if n == 1 {
		return true
	}
	var doBackTrack func(origin, current int, part1 string) bool
	doBackTrack = func(origin, current int, part1 string) bool {
		value, _ := strconv.Atoi(part1)
		calculate := value + current
		if calculate == origin {
			return true
		}
		ind := 1
		for ind <= len(part1)-1 {
			add := part1[:ind]
			val, _ := strconv.Atoi(add)
			newPart := part1[ind:]
			check := current + val
			// fmt.Println("n:= " + strconv.Itoa(origin))
			// fmt.Println(check)
			// fmt.Println(newPart)
			if check <= origin && doBackTrack(origin, check, newPart) {
				return true
			}
			ind++
		}
		return false
	}

	return doBackTrack(n, 0, strconv.Itoa(n*n))
}
