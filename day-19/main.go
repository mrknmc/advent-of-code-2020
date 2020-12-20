package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	children []*Node
	value    string
	rule     string
}

func reverse(items *[]string) {
	for i, j := 0, len(*items)-1; i < j; i, j = i+1, j-1 {
		(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
	}
}

func buildGraph(rules map[string]string, rule string, appendix *Node) *Node {
	// or nodes
	if strings.Contains(rule, "|") {
		orParts := strings.Split(rule, " | ")
		children := make([]*Node, len(orParts))
		for i, orPart := range orParts {
			children[i] = buildGraph(rules, orPart, appendix)
		}
		return &Node{children, "", rule}
	}

	// raw string
	var value string
	if _, err := fmt.Sscanf(rule, "%q", &value); err == nil {
		if appendix != nil {
			// raw string => appendix attached at the end
			return &Node{[]*Node{appendix}, value, rule}
		}
		return &Node{nil, value, rule}
	}

	// chain nodes
	chainParts := strings.Split(rule, " ")
	// start from the back
	reverse(&chainParts)
	for _, chainPart := range chainParts {
		appendix = buildGraph(rules, rules[chainPart], appendix)
	}
	return &Node{[]*Node{appendix}, "", rule}
}

func matchMessage(node *Node, message string) bool {
	if node.value != "" && node.value != string(message[0]) {
		return false
	}
	trimmedMessage := strings.TrimPrefix(message, node.value)
	for _, child := range node.children {
		if matchMessage(child, trimmedMessage) {
			return true
		}
	}
	return trimmedMessage == ""
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	rules := make(map[string]string)
	for fscanner.Scan() {
		line := fscanner.Text()
		if line == "" {
			// last rule
			break
		}
		parts := strings.Split(line, ": ")
		rules[parts[0]] = parts[1]
	}

	messages := make([]string, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		messages = append(messages, line)
	}

	root := buildGraph(rules, rules["0"], nil)
	// fmt.Println(root)
	counter := 0
	for _, message := range messages {
		if matchMessage(root, message) {
			counter++
		}
	}
	fmt.Println(counter)
}
