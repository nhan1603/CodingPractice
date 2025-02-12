package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	total := len(nums)
	last := total - 1
	indx := 0
	for indx < total {
		num := nums[indx]
		if num == 0 {
			copy(nums[indx:], nums[indx+1:])
			nums[last] = 0
			indx--
			total--
		}
		indx++
	}
}
