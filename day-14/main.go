package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getMemoryLocations(mask string, value string) []string {
	if len(mask) != len(value) {
		panic("Lengths don't match")
	}
	newValue := ""
	for i := 0; i < len(mask); i++ {
		if mask[i] == '1' {
			newValue += "1"
		} else if mask[i] == '0' {
			newValue += string(value[i])
		} else if mask[i] == 'X' {
			newValue += "X"
		}
	}
	return convertMemory(newValue, []string{""})
}

func convertMemory(location string, acc []string) []string {
	if location == "" {
		// processed last character
		return acc
	}
	currentChar := location[0]
	if currentChar == '0' || currentChar == '1' {
		for i := range acc {
			acc[i] = acc[i] + string(currentChar)
		}
	} else if currentChar == 'X' {
		for i := range acc {
			acc = append(acc, acc[i]+"0")
			acc[i] = acc[i] + "1"
		}
	} else {
		panic("Impossible state")
	}
	return convertMemory(location[1:], acc)
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)

	// parse file
	lines := make([]string, 0)
	for fscanner.Scan() {
		lines = append(lines, fscanner.Text())
	}

	// part 1
	maskRegex := regexp.MustCompile(`^mask = ([01X]+)$`)
	memRegex := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	mask := int64(0)
	passMask := int64(0)
	mem := make(map[string]int64)
	for _, line := range lines {
		if submatches := maskRegex.FindStringSubmatch(line); len(submatches) > 0 {
			maskStr := submatches[1]
			mask, _ = strconv.ParseInt(strings.ReplaceAll(maskStr, "X", "0"), 2, 64)
			passMask, _ = strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(maskStr, "1", "0"), "X", "1"), 2, 64)
		} else if submatches := memRegex.FindStringSubmatch(line); len(submatches) > 0 {
			value, _ := strconv.ParseInt(submatches[2], 10, 64)
			maskedValue := mask | (value & passMask)
			mem[submatches[1]] = maskedValue
		} else {
			panic("Impossible state!")
		}
	}
	sum := int64(0)
	for _, val := range mem {
		sum += val
	}
	fmt.Println(sum)

	// part 2
	maskStr := ""
	mem = make(map[string]int64)
	for _, line := range lines {
		if submatches := maskRegex.FindStringSubmatch(line); len(submatches) > 0 {
			maskStr = submatches[1]
		} else if submatches := memRegex.FindStringSubmatch(line); len(submatches) > 0 {
			// get all memory locations to update
			memLocationInt, _ := strconv.ParseInt(submatches[1], 10, 64)
			memLocation := strconv.FormatInt(memLocationInt, 2)

			// pad with zeroes if necessary
			if len(memLocation) < 36 {
				memLocation = strings.Repeat("0", 36-len(memLocation)) + memLocation
			}

			// update memory locations with value
			value, _ := strconv.ParseInt(submatches[2], 10, 64)
			for _, loc := range getMemoryLocations(maskStr, memLocation) {
				mem[loc] = value
			}
		} else {
			panic("Impossible state!")
		}
	}
	sum = int64(0)
	for _, val := range mem {
		sum += val
	}
	fmt.Println(sum)
}
