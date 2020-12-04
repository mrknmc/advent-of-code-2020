package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseFile(file *os.File) []map[string]string {
	fscanner := bufio.NewScanner(file)
	passports := make([]map[string]string, 0)
	passport := make(map[string]string)
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, " ")
		if line == "" {
			// new passport
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}
		for _, part := range parts {
			keyval := strings.Split(part, ":")
			key := keyval[0]
			val := keyval[1]
			passport[key] = val
		}
	}
	// add last passport
	passports = append(passports, passport)
	return passports
}

func main() {
	file, _ := os.Open("data.txt")
	passports := parseFile(file)

	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}

	total := len(passports)
	invalidPart1 := 0
	invalidPart2 := 0
	byrRegex := regexp.MustCompile(`^(19[2-9][0-9]|200[0-2])$`)
	iyrRegex := regexp.MustCompile(`^(201[0-9]|2020)$`)
	eyrRegex := regexp.MustCompile(`^(202[0-9]|2030)$`)
	hgtRegex := regexp.MustCompile(`^(1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in$`)
	hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eclRegex := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegex := regexp.MustCompile(`^\d{9}$`)

	for _, passport := range passports {
		for _, reqKey := range required {
			if _, ok := passport[reqKey]; !ok {
				// missing required key
				invalidPart1++
				break
			}
		}
		invalid := false

		if val, ok := passport["byr"]; !ok || !byrRegex.MatchString(val) {
			invalid = true
		} else if val, ok := passport["iyr"]; !ok || !iyrRegex.MatchString(val) {
			invalid = true
		} else if val, ok := passport["eyr"]; !ok || !eyrRegex.MatchString(val) {
			invalid = true
		} else if val, ok := passport["hgt"]; !ok || !hgtRegex.MatchString(val) {
			invalid = true
		} else if val, ok := passport["hcl"]; !ok || !hclRegex.MatchString(val) {
			invalid = true
		} else if val, ok := passport["ecl"]; !ok || !eclRegex.MatchString(val) {
			invalid = true
		} else if val, ok := passport["pid"]; !ok || !pidRegex.MatchString(val) {
			invalid = true
		}
		if invalid {
			invalidPart2++
		}
	}
	fmt.Printf("Part 1 Total %d Valid %d\n", total, total-invalidPart1)
	fmt.Printf("Part 2 Total %d Valid %d\n", total, total-invalidPart2)
}
