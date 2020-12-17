package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRule(ruleStr string) []int {
	rule := make([]int, 0)
	for _, orPart := range strings.Split(ruleStr, " or ") {
		for _, partStr := range strings.Split(orPart, "-") {
			partInt, _ := strconv.Atoi(partStr)
			rule = append(rule, partInt)
		}
	}
	return rule
}

func parseTicket(line string) []int {
	ticket := make([]int, 0)
	for _, numStr := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(numStr)
		ticket = append(ticket, num)
	}
	return ticket
}

func matchRule(rule []int, num int) bool {
	return (num >= rule[0] && num <= rule[1]) || (num >= rule[2] && num <= rule[3])
}

func matchRuleAll(rule []int, nums []int) bool {
	for _, num := range nums {
		if !matchRule(rule, num) {
			return false
		}
	}
	return true
}

func matchRules(rules map[string][]int, nums []int) []string {
	rulesMatched := make([]string, 0)
	for ruleName, rule := range rules {
		if matchRuleAll(rule, nums) {
			rulesMatched = append(rulesMatched, ruleName)
		}
	}
	return rulesMatched
}

func transpose(tickets [][]int) map[int][]int {
	rowCount := len(tickets)
	colCount := len(tickets[0])
	transposed := make(map[int][]int, colCount)
	for i := range tickets[0] {
		transposed[i] = make([]int, rowCount)
	}
	for i, ticket := range tickets {
		for j, num := range ticket {
			transposed[j][i] = num
		}
	}
	return transposed
}

func ticketValid(rules map[string][]int, ticket []int) (bool, int) {
	valid := true
	errorRate := 0
	for _, num := range ticket {
		matched := false
		for _, rule := range rules {
			// at least one rule matches
			if matchRule(rule, num) {
				matched = true
			}
		}

		// no number matched => invalid ticket
		if !matched {
			valid = false
			errorRate += num
		}
	}
	return valid, errorRate
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)

	// parse rules
	rules := make(map[string][]int)
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) == 1 {
			// last rule parsed
			break
		}

		rules[parts[0]] = parseRule(parts[1])
	}

	fscanner.Scan()
	fscanner.Scan()

	// parse my ticket
	myTicket := parseTicket(fscanner.Text())

	fscanner.Scan()
	fscanner.Scan()

	totalErrorRate := 0
	nearbyTickets := make([][]int, 0)
	for fscanner.Scan() {
		ticket := parseTicket(fscanner.Text())
		if valid, errorRate := ticketValid(rules, ticket); valid {
			nearbyTickets = append(nearbyTickets, ticket)
		} else {
			totalErrorRate += errorRate
		}
	}

	fmt.Println("Part 1", totalErrorRate)

	transposedTickets := transpose(nearbyTickets)

	myTicketMap := make(map[string]int)

	for {
		matchedRulesByColumn := make(map[int]string)

		// find columns that have one rule match only
		for colIdx, col := range transposedTickets {
			matchedRules := matchRules(rules, col)
			if len(matchedRules) == 1 {
				matchedRulesByColumn[colIdx] = matchedRules[0]
			}
		}

		for colIdx, ruleName := range matchedRulesByColumn {
			// update my ticket
			myTicketMap[ruleName] = myTicket[colIdx]
			// drop rules that are matched
			delete(rules, ruleName)
			// drop columns that are matched
			delete(transposedTickets, colIdx)
		}

		// all rules matched
		if len(rules) == 0 {
			break
		}
	}

	total := 1
	for ruleName, ticketNum := range myTicketMap {
		if strings.HasPrefix(ruleName, "departure") {
			total *= ticketNum
		}
	}
	fmt.Println("Part 2", total)
}
