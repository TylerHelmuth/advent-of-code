package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isCorrectOrdering(pRules map[string]map[string]bool, update []string) bool {
	for i, page := range update {
		if rules, ok := pRules[page]; ok {
			for _, pageBeforeP := range update[:i] {
				if _, ok2 := rules[pageBeforeP]; ok2 {
					return false
				}
			}
		}
	}
	return true
}

func part1(pRules map[string]map[string]bool, updates [][]string) int {
	sum := 0
	for _, update := range updates {
		if isCorrectOrdering(pRules, update) {
			v, _ := strconv.Atoi(update[len(update)/2])
			sum += v
		}
	}
	return sum
}

func getMiddleNumberForCorrectedUpdates(pRules map[string]map[string]bool, update []string) int {
	isCorrect := true
	for i := 0; i < len(update); i++ {
		page := update[i]
		if rules, ok := pRules[page]; ok {
			for j := 0; j < i; j++ {
				pageBeforeP := update[j]
				if _, ok2 := rules[pageBeforeP]; ok2 {
					isCorrect = false
					// then add pageBeforeP to be right after page
					update = slices.Insert(update, i+1, pageBeforeP)
					// first remove element from current position
					update = append(update[:j], update[j+1:]...)
					i--
				}
			}
		}
	}

	if isCorrect {
		return 0
	}
	v, _ := strconv.Atoi(update[len(update)/2])
	return v
}

func part2(pRules map[string]map[string]bool, updates [][]string) int {
	sum := 0
	for _, update := range updates {
		sum += getMiddleNumberForCorrectedUpdates(pRules, update)
	}
	return sum
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2024/5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	pRules := make(map[string]map[string]bool, 0)
	i := 0
	for {
		if lines[i] == "" {
			break
		}
		rules := strings.Split(lines[i], "|")

		if pRules[rules[0]] == nil {
			pRules[rules[0]] = map[string]bool{
				rules[1]: true,
			}
		} else {
			pRules[rules[0]][rules[1]] = true
		}
		i++
	}

	updates := make([][]string, 0)
	for j := i + 1; j < len(lines); j++ {
		updates = append(updates, strings.Split(lines[j], ","))
	}

	fmt.Println(fmt.Sprintf("part 1 count is %d", part1(pRules, updates)))
	fmt.Println(fmt.Sprintf("part 2 count is %d", part2(pRules, updates)))
}
