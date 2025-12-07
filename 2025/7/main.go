package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func part1(grid [][]string) int {
	splitCount := 0
	for row := 1; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			valueAbove := grid[row-1][col]
			switch valueAbove {
			case "S":
				grid[row][col] = "|"
			case "|":
				if grid[row][col] == "." {
					grid[row][col] = "|"
				} else if grid[row][col] == "^" {
					grid[row][col-1] = "|"
					grid[row][col+1] = "|"
					splitCount++
				}
			default:
			}
		}
	}
	return splitCount
}

type node struct {
	row, col int
}

var memo = make(map[node]int)

func helper(grid [][]string, row, col int) int {
	for row < len(grid) {
		if grid[row][col] == "^" {
			n := node{row: row, col: col}
			if memo[n] != 0 {
				return memo[n]
			}
			result := helper(grid, row, col-1) + helper(grid, row, col+1)
			memo[n] = result
			return result
		}
		row++
	}
	return 1
}

func part2(grid [][]string) int {
	for row := 2; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == "^" {
				return helper(grid, row, col-1) + helper(grid, row, col+1)
			}
		}
	}
	panic("should not reach here")
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/7/input.txt")
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
		grid[i] = strings.Split(line, "")
	}

	fmt.Println(fmt.Sprintf("part 1 total splits is %d", part1(grid)))
	fmt.Println(fmt.Sprintf("part 2 total timelines is %d", part2(grid)))
}
