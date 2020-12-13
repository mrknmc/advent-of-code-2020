package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	fscanner.Scan()
	departureTime, _ := strconv.Atoi(fscanner.Text())
	fscanner.Scan()
	busIds := strings.Split(fscanner.Text(), ",")

	// parse running bus ids
	runningBusIds := make([]int, 0)
	for _, busID := range busIds {
		if busID != "x" {
			busIDInt, _ := strconv.Atoi(busID)
			runningBusIds = append(runningBusIds, busIDInt)
		}
	}
	foundBusID := -1
	time := -1
	for t := departureTime; foundBusID == -1; t++ {
		for _, busID := range runningBusIds {
			if t%busID == 0 {
				foundBusID = busID
				time = t - departureTime
				break
			}
		}
	}
	fmt.Println(foundBusID * time)

	initialBusID := runningBusIds[0]

	funcs := make(map[int]func(int) int)
	for offset, busID := range busIds {
		if busID != "x" {
			// skip xs
			busIDInt, _ := strconv.Atoi(busID)
			if busIDInt != initialBusID {
				// skip max bus id
				o := offset
				function := func(t int) int {
					return (t + o) % busIDInt
				}
				funcs[busIDInt] = function
			}
		}
	}

	mult := 1
	for time = initialBusID; true; time += initialBusID * mult {
		for busID, function := range funcs {
			if function(time) == 0 {
				// update multiplier with matched ID
				mult *= busID
				// no need to check that function anymore
				delete(funcs, busID)
			}
		}
		if len(funcs) == 0 {
			// all functions are matched
			break
		}
	}

	fmt.Println(time)
}
