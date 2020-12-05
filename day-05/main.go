package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	maxSeatID := 0
	seats := make([]int, 1024)

	for fscanner.Scan() {
		line := fscanner.Text()
		rowHigh := 127
		rowLow := 0

		for _, i := range line[0:7] {
			pivot := rowLow + (rowHigh+1-rowLow)/2
			if i == 'F' {
				rowHigh = pivot - 1
			} else if i == 'B' {
				rowLow = pivot
			}
		}
		if rowLow != rowHigh {
			panic("high and low should be equal")
		}

		colHigh := 7
		colLow := 0
		for _, i := range line[7:10] {
			pivot := colLow + (colHigh+1-colLow)/2
			if i == 'L' {
				colHigh = pivot - 1
			} else if i == 'R' {
				colLow = pivot
			}
		}
		if colLow != colHigh {
			panic("high and low should be equal")
		}

		seatID := rowLow*8 + colLow
		seats[seatID] = 1
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Println("Max seat id: ", maxSeatID)

	mySeatID := maxSeatID - 1
	for {
		if seats[mySeatID] == 0 && seats[mySeatID-1] == 1 && seats[mySeatID+1] == 1 {
			break
		} else {
			mySeatID--
		}
	}
	fmt.Println("My seat: ", mySeatID)
}
