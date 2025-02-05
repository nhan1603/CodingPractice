package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 && n == 1 {
		return true
	}

	if n == 0 {
		return true
	}

	indx := 0
	for indx < len(flowerbed) {
		switch indx {
		case len(flowerbed) - 1:
			if flowerbed[indx] == 0 && flowerbed[indx-1] == 0 {
				n--
				indx = indx + 1
				break
			}
		case 0:
			if flowerbed[indx] == 0 && flowerbed[indx+1] == 0 {
				n--
				indx = indx + 1
				break
			}
		default:
			if flowerbed[indx] == 0 && flowerbed[indx+1] == 0 && flowerbed[indx-1] == 0 {
				n--
				indx = indx + 1
			}
		}

		if n == 0 {
			return true
		}
		indx++
	}
	return false
}
