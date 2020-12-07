package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	id       string
	children map[string]*Node
	counts   map[string]int
}

func parseFile(file *os.File) Node {
	fscanner := bufio.NewScanner(file)
	parentRegex := regexp.MustCompile(`^(?P<parentId>\w+ \w+) bags contain (.+)\.$`)
	childRegex := regexp.MustCompile(`^(?P<count>\d+) (?P<id>\w+ \w+) bags?`)
	root := Node{"", make(map[string]*Node, 0), nil}
	for fscanner.Scan() {
		line := fscanner.Text()
		parent := parentRegex.FindStringSubmatch(line)
		parentID := parent[1]
		children := make(map[string]*Node, 0)
		counts := make(map[string]int)
		if parent[2] != "no other bags" {
			for _, part := range strings.Split(parent[2], ", ") {
				child := childRegex.FindStringSubmatch(part)
				childID := child[2]
				count, _ := strconv.Atoi(child[1])
				counts[childID] = count
				children[childID] = &Node{childID, make(map[string]*Node), make(map[string]int)}
			}
		}
		root.children[parentID] = &Node{parentID, children, counts}
	}
	return root
}

func fill(node *Node, root *Node) {
	for _, child := range node.children {
		child.children = root.children[child.id].children
		child.counts = root.children[child.id].counts
		fill(child, root)
	}
}

func find(node *Node, id string) bool {
	if node.id == id {
		return true
	}

	for _, child := range node.children {
		if find(child, id) {
			return true
		}
	}
	return false
}

func count(node *Node) int {
	counter := 1
	for _, child := range node.children {
		counter += node.counts[child.id] * count(child)
	}
	return counter
}

func main() {
	file, _ := os.Open("data.txt")
	root := parseFile(file)

	fill(&root, &root)
	countPart1 := 0
	countPart2 := 0
	for _, child := range root.children {
		if child.id != "shiny gold" && find(child, "shiny gold") {
			countPart1++
		} else if child.id == "shiny gold" {
			countPart2 = count(child) - 1
		}
	}
	fmt.Println(countPart1)
	fmt.Println(countPart2)
}
