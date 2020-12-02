package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountChars(word string) map[string]int {
	countMap := make(map[string]int)
	for _, l := range word {
		stringL := string(l)
		if _, ok := countMap[stringL]; ok {
			// val already counted => increment
			countMap[stringL]++
		} else {
			countMap[stringL] = 1
		}
	}
	return countMap
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	validPart1 := 0
	validPart2 := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, " ")
		highLowPart := strings.Split(parts[0], "-")
		low, _ := strconv.Atoi(highLowPart[0])
		high, _ := strconv.Atoi(highLowPart[1])
		letter := strings.TrimRight(parts[1], ":")
		password := parts[2]
		countMap := CountChars(password)
		if countMap[letter] >= low && countMap[letter] <= high {
			validPart1++
		}
		if (string(password[low-1]) == letter && string(password[high-1]) != letter) ||
			(string(password[low-1]) != letter && string(password[high-1]) == letter) {
			validPart2++
		}
	}
	fmt.Println("valid part 1: ", validPart1)
	fmt.Println("valid part 2: ", validPart2)
}
