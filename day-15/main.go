package main

import "fmt"

func loop(nums []int, iters int) int {
	maxSeen := make(map[int]int)
	prevSeen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		maxSeen[num] = i
	}

	// start with last number
	prevNum := nums[len(nums)-1]
	for i := len(nums); i < iters; i++ {
		if _, ok := prevSeen[prevNum]; ok {
			// seen more than once
			prevNum = maxSeen[prevNum] - prevSeen[prevNum]
		} else {
			// first time
			prevNum = 0
		}
		if maxIt, ok := maxSeen[prevNum]; ok {
			// if seen before move prev
			prevSeen[prevNum] = maxIt
		}
		maxSeen[prevNum] = i
	}

	return prevNum
}

func main() {
	nums := []int{0, 14, 6, 20, 1, 4}

	num := loop(nums, 2020)
	fmt.Println(num)

	num = loop(nums, 30000000)
	fmt.Println(num)
}
