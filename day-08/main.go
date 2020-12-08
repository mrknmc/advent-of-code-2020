package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Inst struct {
	id  string
	num int
}

func parseFile(file *os.File) []Inst {
	fscanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`^(\w+) ([+-]\d+)$`)
	insts := make([]Inst, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		match := regex.FindStringSubmatch(line)
		inst := match[1]
		num, _ := strconv.Atoi(match[2])
		insts = append(insts, Inst{inst, num})
	}
	return insts
}

func loops(insts []Inst) (bool, int) {
	executed := make(map[int]bool)
	acc := 0
	i := 0
	loops := false
	for true {
		if executed[i] {
			// loops
			loops = true
			break
		}
		if i >= len(insts) {
			// finished
			break
		}
		executed[i] = true
		if insts[i].id == "nop" {
			// noop
			i++
		} else if insts[i].id == "jmp" {
			i += insts[i].num
		} else {
			// assume acc
			acc += insts[i].num
			i++
		}
	}
	return loops, acc
}

func main() {
	file, _ := os.Open("data.txt")
	insts := parseFile(file)

	_, acc := loops(insts)
	fmt.Println("Part1 ", acc)

	for i := range insts {
		var tmp string
		if insts[i].id == "jmp" {
			tmp = "jmp"
			insts[i].id = "nop"
		} else if insts[i].id == "nop" {
			tmp = "nop"
			insts[i].id = "jmp"
		}
		loops, acc := loops(insts)
		if !loops {
			fmt.Println("Part2", acc)
		}
		// reverse change
		insts[i].id = tmp
	}
}
