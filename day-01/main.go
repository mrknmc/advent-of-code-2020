package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	fs := make([]int64, 0)
	for fscanner.Scan() {
		i, _ := strconv.ParseInt(fscanner.Text(), 10, 64)
		fs = append(fs, i)
	}
	for _, f1 := range fs {
		for _, f2 := range fs {
			if f1+f2 == 2020 && f1 != f2 {
				// found it
				fmt.Println(f1 * f2)
			}
		}
	}
	for _, f1 := range fs {
		for _, f2 := range fs {
			for _, f3 := range fs {
				if f1+f2+f3 == 2020 && f1 != f2 && f1 != f3 && f2 != f3 {
					// found it
					fmt.Println(f1 * f2 * f3)
				}
			}
		}
	}
}
