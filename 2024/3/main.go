package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	multRegex    = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	disableRegex = regexp.MustCompile(`don't\(\).*?do\(\)`)
)

func part1(lines []string) int {
	sum := 0
	for _, l := range lines {
		matches := multRegex.FindAllStringSubmatch(l, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			sum += x * y
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	s := ""
	for _, l := range lines {
		s += l
	}
	s = disableRegex.ReplaceAllString(s, "")
	matches := multRegex.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}
	return sum
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2024/3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 sum is %d", part1(lines)))
	fmt.Println(fmt.Sprintf("part 2 sum is %d", part2(lines)))
}
