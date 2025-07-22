package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01", "10"}))
}

func findDifferentBinaryString(nums []string) string {
	n := int64(len(nums))
	check := make(map[int64]bool)

	for _, str := range nums {
		val, _ := strconv.ParseInt(str, 2, 64)

		check[val] = true
	}
	var i int64
	var result string
	for i = 0; i <= n; i++ {

		ok := check[i]
		if !ok {
			result = strconv.FormatInt(i, 2)
			for int64(len(result)) < n {
				result = "0" + result
			}
			return result
		}
	}
	return result
}

func removeStars(s string) string {
	var result []rune
	for _, val := range s {
		if val != '*' {
			result = append(result, val)
		} else {
			result = result[:len(result)-1]
		}
	}
	return string(result)
}
