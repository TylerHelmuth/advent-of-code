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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	pRules := make(map[string]map[string]bool)
	i := 0
	for lines[i] != "" {
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

	compareFunc := func(a, b string) int {
		if rules, ok := pRules[a]; ok {
			if _, ok := rules[b]; ok {
				return -1
			}
		}
		return 1
	}

	correctSum := 0
	correctedSum := 0
	for _, update := range updates {
		if slices.IsSortedFunc(update, compareFunc) {
			v, _ := strconv.Atoi(update[len(update)/2])
			correctSum += v
		} else {
			slices.SortFunc(update, compareFunc)
			v, _ := strconv.Atoi(update[len(update)/2])
			correctedSum += v
		}
	}

	fmt.Println(fmt.Sprintf("part 1 sum is %d", correctSum))
	fmt.Println(fmt.Sprintf("part 2 sum is %d", correctedSum))
}
