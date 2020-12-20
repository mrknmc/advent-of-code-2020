package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	children []*Node
	value    string
	rule     string
}

// func parseNode(text string) Node {
// 	if strings.Contains(text, "|") {
// 		// or node
// 		orParts := strings.Split(text, " | ")
// 		for _, orPart := range orParts {

// 		}
// 	} else if strings.Contains(text, " ") {
// 		andParts := strings.Split(text, " ")
// 		for _, andPart := range andParts {

// 		}
// 	} else if regexp.MatchString(text, `"\w+"`) {

// 	}
// }

func reverse(items *[]string) {
	for i, j := 0, len(*items)-1; i < j; i, j = i+1, j-1 {
		(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
	}
}

func buildGraph(rules map[string]string, rule string, appendix *Node) *Node {
	rawStringRegex := regexp.MustCompile(`"(\w+)"`)
	singleRuleRegex := regexp.MustCompile(`(\d)`)
	if strings.Contains(rule, "|") {
		// or node
		orParts := strings.Split(rule, " | ")
		children := make([]*Node, len(orParts))
		for _, orPart := range orParts {
			child := buildGraph(rules, orPart, nil)
			children = append(children, child)
		}
		return &Node{children, "", rule}
	} else if strings.Contains(rule, " ") {
		// chain nodes
		chainParts := strings.Split(rule, " ")
		// start from the back
		reverse(&chainParts)
		var appendix *Node = nil
		for _, chainPart := range chainParts {
			appendix = buildGraph(rules, chainPart, appendix)
		}
		return &Node{[]*Node{appendix}, "", rule}
	} else if matches := rawStringRegex.FindStringSubmatch(rule); len(matches) > 1 {
		// raw string => appendix attached at the end
		return &Node{[]*Node{appendix}, matches[1], rule}
	} else if singleRuleRegex.MatchString(rule) {
		// one rule Node
		child := buildGraph(rules, rules[rule], nil)
		return &Node{[]*Node{child}, "", rule}
	}

	panic("Impossible state")
}

func explore(rules map[string]string, rule string, message string) bool {
	rawStringRegex := regexp.MustCompile(`"(\w+)"`)
	singleRuleRegex := regexp.MustCompile(`(\d)`)
	if strings.Contains(rule, "|") {
		// or node
		orParts := strings.Split(rule, " | ")
		for _, orPart := range orParts {
			if explore(rules, orPart, rule) {
				return true
			}
		}
		// no child matched
		return false
	} else if strings.Contains(rule, " ") {
		// chain nodes
		// TODO: doesn't work
		chainParts := strings.Split(rule, " ")
		for _, chainPart := range chainParts {
			// tmpNode := returnNode
			if !explore(rules, chainPart, message) {
				return false
			}
		}
		// all children matched
		return true
	} else if matches := rawStringRegex.FindStringSubmatch(rule); len(matches) > 1 {
		// raw string
		return matches[1] == rule
	} else if singleRuleRegex.MatchString(rule) {
		// one rule Node
		return explore(rules, rule, message)
	}

	panic("Impossible state")
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

	// root := buildGraph(rules, rules["0"])
	match := buildGraph(rules, rules["0"], nil)
	fmt.Println(match)
}
