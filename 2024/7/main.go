package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation int

const (
	add operation = iota
	mult
	concat
)

type runningTotal struct {
	numbers    []int
	operations []operation
	value      int
}

func performOperation(target, number int, o operation, runningTotals []runningTotal) []runningTotal {
	newRunningTotals := make([]runningTotal, 0)
	for _, rt := range runningTotals {
		var result int
		switch o {
		case add:
			result = rt.value + number
		case mult:
			result = rt.value * number
		case concat:
			result, _ = strconv.Atoi(strconv.Itoa(rt.value) + strconv.Itoa(number))
		}
		if result <= target {
			newRunningTotals = append(newRunningTotals, runningTotal{
				numbers:    append(rt.numbers, number),
				operations: append(rt.operations, o),
				value:      result,
			})
		}
	}
	return newRunningTotals
}

func calculateEquation(target int, numbers []int, operations []operation) int {
	runningTotals := make([]runningTotal, 0)
	if numbers[0]+numbers[1] <= target {
		runningTotals = append(runningTotals, runningTotal{
			numbers:    numbers[0:2],
			operations: []operation{add},
			value:      numbers[0] + numbers[1],
		})
	}
	if numbers[0]*numbers[1] <= target {
		runningTotals = append(runningTotals, runningTotal{
			numbers:    numbers[0:2],
			operations: []operation{mult},
			value:      numbers[0] * numbers[1],
		})
	}
	// concat case
	r, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))
	if r <= target {
		runningTotals = append(runningTotals, runningTotal{
			numbers:    numbers[0:2],
			operations: []operation{concat},
			value:      r,
		})
	}

	for i := 2; i < len(numbers); i++ {
		newRunningTotals := make([]runningTotal, 0)
		for _, o := range operations {
			newRunningTotals = append(newRunningTotals, performOperation(target, numbers[i], o, runningTotals)...)
		}
		runningTotals = newRunningTotals
	}

	for _, rt := range runningTotals {
		if rt.value == target {
			return target
		}
	}
	return 0
}

func parseLine(l string) (int, []int) {
	temp := strings.Split(l, ":")
	target, _ := strconv.Atoi(temp[0])
	temp = strings.Split(strings.Trim(temp[1], " "), " ")
	numbers := make([]int, len(temp))
	for i, n := range temp {
		numbers[i], _ = strconv.Atoi(n)
	}
	return target, numbers
}

func part1(lines []string) int {
	sum := 0
	for _, l := range lines {
		target, numbers := parseLine(l)
		sum += calculateEquation(target, numbers, []operation{add, mult})
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, l := range lines {
		target, numbers := parseLine(l)
		sum += calculateEquation(target, numbers, []operation{add, mult, concat})
	}
	return sum
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2024/7/input.txt")
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
