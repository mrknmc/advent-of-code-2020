package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const maxUint = ^uint(0)

func parseFile(file *os.File) []int {
	fscanner := bufio.NewScanner(file)
	nums := make([]int, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums
}

func checkBatch(batch []int, num int) bool {
	for _, i := range batch {
		for _, j := range batch {
			if i != j && i+j == num {
				return true
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("data.txt")
	nums := parseFile(file)
	preambleSize := 25
	pivot := preambleSize
	resultPart1 := 0
	for {
		if !checkBatch(nums[pivot-preambleSize:pivot], nums[pivot]) {
			resultPart1 = nums[pivot]
			break
		}
		pivot++
	}

	fmt.Println(resultPart1)

	maxLength := 0
	maxI := 0
	maxJ := 0

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum == resultPart1 && j-i > maxLength {
				maxLength = j - i
				maxI = i
				maxJ = j
			}
		}
	}

	minNum := int(^uint(0) >> 1)
	maxNum := 0

	for _, num := range nums[maxI:maxJ] {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}

	fmt.Println(minNum + maxNum)
}
