package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(file *os.File) [][]map[string]int {
	fscanner := bufio.NewScanner(file)
	groups := make([][]map[string]int, 0)
	group := make([]map[string]int, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		if line == "" {
			// new group
			groups = append(groups, group)
			group = make([]map[string]int, 0)
			continue
		}
		person := make(map[string]int)
		for _, part := range line {
			person[string(part)]++
		}
		group = append(group, person)
	}
	// add last group
	groups = append(groups, group)
	return groups
}

func main() {
	file, _ := os.Open("data.txt")
	groups := parseFile(file)

	unions := 0
	intersections := 0
	for _, group := range groups {
		groupSize := len(group)
		counter := make(map[string]int)
		for _, person := range group {
			for k := range person {
				counter[k]++
			}
		}
		unions += len(counter)
		for _, v := range counter {
			if v == groupSize {
				intersections++
			}
		}
	}
	fmt.Println("Union: ", unions)
	fmt.Println("Intersection: ", intersections)
}
