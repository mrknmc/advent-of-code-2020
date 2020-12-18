package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	data *[]string
}

func newStack() *Stack {
	data := make([]string, 0)
	return &Stack{&data}
}

func (s Stack) empty() bool {
	return len(*s.data) == 0
}

func (s Stack) peek() string {
	if len(*s.data) == 0 {
		return ""
	}
	return (*s.data)[len(*s.data)-1]
}

func (s *Stack) pop() string {
	item := s.peek()
	newData := (*s.data)[:len(*s.data)-1]
	s.data = &newData
	return item
}

func (s *Stack) push(item string) {
	newData := append((*s.data), item)
	s.data = &newData
}

func reverse(items *[]string) {
	for i, j := 0, len(*items)-1; i < j; i, j = i+1, j-1 {
		(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
	}
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	precedence := make(map[string]int)
	precedence["+"] = 3
	precedence["*"] = 2
	precedence["-"] = 1
	precedence["("] = 0
	leftRegex := regexp.MustCompile(`\(`)
	rightRegex := regexp.MustCompile(`\)`)
	total := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		line = leftRegex.ReplaceAllLiteralString(rightRegex.ReplaceAllLiteralString(line, " )"), "( ")
		rpnStack := newStack()
		stack := newStack()
		for _, token := range strings.Split(line, " ") {
			if _, err := strconv.Atoi(token); err == nil {
				rpnStack.push(token)
			} else if token == "(" {
				stack.push(token)
			} else if token == ")" {
				for !stack.empty() && stack.peek() != "(" {
					rpnStack.push(stack.pop())
				}
				if stack.peek() == "(" {
					stack.pop()
				}
			} else {
				// operator
				for !stack.empty() && precedence[stack.peek()] >= precedence[token] {
					rpnStack.push(stack.pop())
				}
				stack.push(token)
			}
		}
		for !stack.empty() {
			rpnStack.push(stack.pop())
		}

		// to act as queue, we reverse the stack
		reverse(rpnStack.data)
		execStack := newStack()

		for !rpnStack.empty() {
			item := rpnStack.pop()
			if _, err := strconv.Atoi(item); err == nil {
				execStack.push(item)
			} else {
				right, _ := strconv.Atoi(execStack.pop())
				left, _ := strconv.Atoi(execStack.pop())
				if item == "+" {
					execStack.push(fmt.Sprintf("%d", left+right))
				} else if item == "-" {
					execStack.push(fmt.Sprintf("%d", left-right))
				} else if item == "*" {
					execStack.push(fmt.Sprintf("%d", left*right))
				} else {
					panic("Impossible state")
				}
			}
		}
		val, _ := strconv.Atoi(execStack.pop())
		fmt.Println(val)
		total += val
		if !execStack.empty() {
			panic("stack not empty")
		}
	}
	fmt.Println(total)
}
