package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State int

const (
	Missing State = iota
	Present
	Usable
)

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

func main() {
	file, _ := os.Open("data.txt")
	nums := parseFile(file)
	maxAdapterJolts := 0
	adapterMap := make(map[int]State)
	adapterMap[0] = Usable
	for _, num := range nums {
		adapterMap[num] = Present
		if num > maxAdapterJolts {
			maxAdapterJolts = num
		}
	}
	builtInAdapterJolts := maxAdapterJolts + 3
	adapterMap[builtInAdapterJolts] = Present
	diff1 := 0
	diff3 := 0
	longestChain := []int{0}
	for i := 0; i <= builtInAdapterJolts; i++ {
		if adapterMap[i] != Missing {
			// it's one of the adapters we have
			if adapterMap[i-1] == Usable {
				adapterMap[i] = Usable
				longestChain = append(longestChain, i)
				diff1++
			} else if adapterMap[i-3] == Usable {
				adapterMap[i] = Usable
				longestChain = append(longestChain, i)
				diff3++
			} else if adapterMap[i-2] == Usable {
				adapterMap[i] = Usable
				longestChain = append(longestChain, i)
			}
		}
	}
	fmt.Println(diff1 * diff3)

	count := 1.0
	inbetween := 0.0

	for i := 1; i < len(longestChain)-1; i++ {
		diff := longestChain[i+1] - longestChain[i-1]
		if diff >= 3 {
			// cannot drop
			combs := 1.0
			if inbetween == 0 {
				combs = 1
			} else if inbetween == 1 {
				combs = 2
			} else if inbetween == 2 {
				combs = 4
			} else if inbetween == 3 {
				combs = 7
			} else {
				panic("This cannot happen")
			}
			count *= combs
			inbetween = 0
		} else if diff == 2 {
			// can drop
			inbetween++
		} else {
			panic("This cannot happen")
		}
	}

	fmt.Println(count)
}
