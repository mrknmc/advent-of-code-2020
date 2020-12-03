package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(file *os.File) map[int]map[int]rune {
	fscanner := bufio.NewScanner(file)
	data := map[int]map[int]rune{}
	lineNumber := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		data[lineNumber] = make(map[int]rune)
		for i, a := range line {
			data[lineNumber][i] = a
		}
		lineNumber++
	}
	return data
}

func traverse(data map[int]map[int]rune, moveX int, moveY int) int {
	positionX := 0
	positionY := 0

	rowCount := len(data)
	colCount := len(data[0])

	treeCount := 0

	for positionY < rowCount {
		entry := data[positionY][positionX]
		if entry == '#' {
			treeCount++
		}

		positionX += moveX
		positionY += moveY

		// reset to first column
		if positionX >= colCount {
			positionX = positionX % colCount
		}
	}

	return treeCount
}

func main() {
	file, _ := os.Open("data.txt")
	data := parseFile(file)

	pattern1 := traverse(data, 1, 1)
	pattern2 := traverse(data, 3, 1)
	pattern3 := traverse(data, 5, 1)
	pattern4 := traverse(data, 7, 1)
	pattern5 := traverse(data, 1, 2)

	fmt.Println("Pattern 1 ", pattern1)
	fmt.Println("Pattern 2 ", pattern2)
	fmt.Println("Pattern 3 ", pattern3)
	fmt.Println("Pattern 4 ", pattern4)
	fmt.Println("Pattern 5 ", pattern5)
	fmt.Println("Multiplied ", pattern1*pattern2*pattern3*pattern4*pattern5)
}
