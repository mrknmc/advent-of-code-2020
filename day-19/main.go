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

func split(message string) [][]string {
	result := make([][]string, 0)
	// rule 8: make first part as large as possible since that's what's repeated
	for i := len(message) - 1; i > 0; i-- {
		result = append(result, []string{message[:i], message[i:]})
	}
	return result
}

func split2(message string) [][]string {
	result := make([][]string, 0)
	// rule 11: make middle part as large as possible since that's what's repeated
	for i := 1; i < len(message); i++ {
		for j := len(message) - 1; j > i; j-- {
			result = append(result, []string{message[:i], message[i:j], message[j:]})
		}
	}
	return result
}

func explore(rules map[string]string, rule string, message string, cache *map[string]bool) bool {
	if result, ok := (*cache)[rule+"."+message]; ok {
		return result
	}

	// raw string
	var value string
	if _, err := fmt.Sscanf(rule, "%q", &value); err == nil {
		return message == value
	}

	// or nodes
	if strings.Contains(rule, "|") {
		for _, orPart := range strings.Split(rule, " | ") {
			if explore(rules, orPart, message, cache) {
				(*cache)[rule+"."+message] = true
				return true
			}
		}
		(*cache)[rule+"."+message] = false
		return false
	}

	// chain nodes
	chainParts := strings.Split(rule, " ")
	if len(chainParts) == 1 {
		// single rule => has to match the remaining string
		return explore(rules, rules[chainParts[0]], message, cache)
	} else if len(chainParts) == 2 && len(message) > 1 {
		// two rules in one => both have to match any possible combination
		for _, split := range split(message) {
			// valid way to split these messages
			if explore(rules, chainParts[0], split[0], cache) && explore(rules, chainParts[1], split[1], cache) {
				(*cache)[rule+"."+message] = true
				return true
			}
		}
		(*cache)[rule+"."+message] = false
		return false
	} else if len(chainParts) == 3 && len(message) > 2 {
		// two rules in one => both have to match any possible combination
		for _, split := range split2(message) {
			// valid way to split these messages
			if explore(rules, chainParts[0], split[0], cache) && explore(rules, chainParts[1], split[1], cache) && explore(rules, chainParts[2], split[2], cache) {
				(*cache)[rule+"."+message] = true
				return true
			}
		}
		(*cache)[rule+"."+message] = false
		return false
	}

	(*cache)[rule+"."+message] = false
	return false
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
	counter := 0
	for _, message := range messages {
		if matchMessage(root, message) {
			counter++
		}
	}

	fmt.Println("Part 1", counter)

	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"

	counter = 0

	cache := make(map[string]bool)
	for _, message := range messages {
		if explore(rules, rules["0"], message, &cache) {
			counter++
		}
	}
	fmt.Println("Part 2", counter)
}
