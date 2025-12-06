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

type FreshRange struct {
	min int
	max int
}

func part1(lines []string) int {
	count := 0

	ranges := make([]FreshRange, 0)
	var availableIngredientsPos int
	for i, line := range lines {
		if line == "" {
			availableIngredientsPos = i + 1
			break
		}
		nums := strings.Split(line, "-")
		low, _ := strconv.Atoi(nums[0])
		high, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, FreshRange{
			min: low,
			max: high,
		})
	}

	for _, ingredient := range lines[availableIngredientsPos:] {
		n, _ := strconv.Atoi(ingredient)
		for _, r := range ranges {
			if n >= r.min && n <= r.max {
				count++
				break
			}
		}
	}

	return count
}

func part2(lines []string) int {
	count := 0

	ranges := make([]FreshRange, 0)
	for _, line := range lines {
		if line == "" {
			break
		}
		nums := strings.Split(line, "-")
		low, _ := strconv.Atoi(nums[0])
		high, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, FreshRange{
			min: low,
			max: high,
		})
	}

	slices.SortFunc(ranges, func(a FreshRange, b FreshRange) int {
		if a.min < b.min {
			return -1
		}
		if a.min > b.min {
			return 1
		}
		return 0
	})

	finalRanges := make([]FreshRange, 0)

	workingRange := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].min > workingRange.max {
			finalRanges = append(finalRanges, workingRange)
			workingRange = ranges[i]
			continue
		}
		if ranges[i].min <= workingRange.max && ranges[i].max > workingRange.max {
			workingRange.max = ranges[i].max
			continue
		}
	}

	finalRanges = append(finalRanges, workingRange)

	for _, r := range finalRanges {
		count += r.max - r.min + 1
	}

	return count
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 count of available ingredients is %d", part1(lines)))
	fmt.Println(fmt.Sprintf("part 2 count of fresh ingredients is %d", part2(lines)))
}
