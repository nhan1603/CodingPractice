package main

import "fmt"

func main() {
	counter := Constructor()
	fmt.Println(counter.Ping(1))
	fmt.Println(counter.Ping(100))
	fmt.Println(counter.Ping(3001))
	fmt.Println(counter.Ping(3002))
}

type RecentCounter struct {
	request []int
}

func Constructor() RecentCounter {
	result := RecentCounter{}
	result.request = make([]int, 0)
	return result
}

func (this *RecentCounter) Ping(t int) int {
	this.request = append(this.request, t)
	count := 0
	var i int
	for i = len(this.request) - 1; this.request[i] >= t-3000; i-- {
		count++
		if i == 0 {
			break
		}
	}
	fmt.Println(len(this.request) - i)
	return count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
