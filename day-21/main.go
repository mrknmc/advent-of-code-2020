package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func intersection(map1 map[string]bool, map2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for key := range map1 {
		if map2[key] || map2 == nil {
			// both contain key
			result[key] = true
		}
	}
	return result
}

func difference(map1 map[string]bool, map2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for key := range map1 {
		if !map2[key] {
			// only map1 contains key
			result[key] = true
		}
	}
	return result
}

func makeSet(array []string) map[string]bool {
	result := make(map[string]bool)
	for _, val := range array {
		result[val] = true
	}
	return result
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for fscanner.Scan() {
		line := fscanner.Text()
		lines = append(lines, line)
	}

	knownAllergens := make(map[string]string)
	knownIngredients := make(map[string]bool)
	ingredientCounts := make(map[string]int)

	// count ingredients once
	for _, line := range lines {
		for _, ingredient := range strings.Split(strings.Split(line, " (contains ")[0], " ") {
			ingredientCounts[ingredient]++
		}
	}

	for {
		allergenMap := make(map[string]map[string]bool)
		for _, line := range lines {
			parts := strings.Split(line, " (contains ")
			ingredients := makeSet(strings.Split(parts[0], " "))
			for _, allergen := range strings.Split(strings.TrimSuffix(parts[1], ")"), ", ") {
				if knownAllergens[allergen] == "" {
					allergenMap[allergen] = intersection(difference(ingredients, knownIngredients), allergenMap[allergen])
				}
			}
		}

		if len(allergenMap) == 0 {
			// done
			break
		}

		// update known allergens
		for allergen, ingredientMap := range allergenMap {
			if len(ingredientMap) == 1 {
				for ingredient := range ingredientMap {
					knownAllergens[allergen] = ingredient
					knownIngredients[ingredient] = true
					delete(ingredientCounts, ingredient)
				}
			}
		}
	}

	total := 0
	for _, count := range ingredientCounts {
		total += count
	}

	fmt.Println("Part 1", total)

	allergens := make([]string, 0)
	for allergen := range knownAllergens {
		allergens = append(allergens, allergen)
	}

	sort.Strings(allergens)

	ingredients := make([]string, 0)
	for _, allergen := range allergens {
		ingredients = append(ingredients, knownAllergens[allergen])
	}

	fmt.Println("Part 2", strings.Join(ingredients, ","))
}
