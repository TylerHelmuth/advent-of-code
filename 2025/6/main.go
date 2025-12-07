package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(grid [][]string) int64 {
	grandTotal := int64(0)
	lastRow := len(grid) - 1
	for col := 0; col < len(grid[0]); col++ {
		operation := grid[lastRow][col]
		result, _ := strconv.ParseInt(grid[0][col], 10, 64)
		switch operation {
		case "+":
			for row := 1; row < lastRow; row++ {
				n, _ := strconv.ParseInt(grid[row][col], 10, 64)
				result += n
			}
		case "*":
			for row := 1; row < lastRow; row++ {
				n, _ := strconv.ParseInt(grid[row][col], 10, 64)
				result *= n
			}
		default:
			panic("invalid operation")
		}
		grandTotal += result
	}
	return grandTotal
}

type problem struct {
	operator string
	numbers  []string
}

func part2(lines []string) int {
	// The operations line of my input was missing a trailing space
	lines[len(lines)-1] = lines[len(lines)-1] + " "

	problems := make([]problem, 0)
	currentProblem := problem{}
	for col := len(lines[0]) - 1; col >= 0; col-- {
		n := ""
		for row := 0; row < len(lines)-1; row++ {
			n += string(lines[row][col])
		}
		currentProblem.numbers = append(currentProblem.numbers, n)

		switch lines[len(lines)-1][col] {
		case '+':
			currentProblem.operator = "+"
			problems = append(problems, currentProblem)
			currentProblem = problem{}
			col--
		case '*':
			currentProblem.operator = "*"
			problems = append(problems, currentProblem)
			currentProblem = problem{}
			col--
		}
	}

	grandTotal := 0
	for _, p := range problems {
		result, _ := strconv.Atoi(strings.TrimSpace(p.numbers[0]))

		switch p.operator {
		case "+":
			for i := 1; i < len(p.numbers); i++ {
				n, _ := strconv.Atoi(strings.TrimSpace(p.numbers[i]))
				result += n
			}
		case "*":
			for i := 1; i < len(p.numbers); i++ {
				n, _ := strconv.Atoi(strings.TrimSpace(p.numbers[i]))
				result *= n
			}
		default:
			panic("invalid operator")
		}

		grandTotal += result
	}

	return grandTotal
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	grid := make([][]string, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, " ")
		result := make([]string, 0)
		for j := 0; j < len(numbers); j++ {
			n := strings.TrimSpace(numbers[j])
			if n != "" {
				result = append(result, n)
			}
		}
		grid[i] = result
	}

	fmt.Println(fmt.Sprintf("part 1 grand total sum is %d", part1(grid)))
	fmt.Println(fmt.Sprintf("part 2 grand total sum is %d", part2(lines)))
}
