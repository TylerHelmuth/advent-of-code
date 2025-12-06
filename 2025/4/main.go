package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func part1(grid [][]string) int {
	count := 0

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			pos := grid[row][col]
			if pos == "@" {
				paperCount := 0
				if grid[row-1][col-1] == "@" || grid[row-1][col-1] == "x" {
					paperCount++
				}
				if grid[row-1][col] == "@" || grid[row-1][col] == "x" {
					paperCount++
				}
				if grid[row-1][col+1] == "@" || grid[row-1][col+1] == "x" {
					paperCount++
				}
				if grid[row][col-1] == "@" || grid[row][col-1] == "x" {
					paperCount++
				}
				if grid[row][col+1] == "@" || grid[row][col+1] == "x" {
					paperCount++
				}
				if grid[row+1][col-1] == "@" || grid[row+1][col-1] == "x" {
					paperCount++
				}
				if grid[row+1][col] == "@" || grid[row+1][col] == "x" {
					paperCount++
				}
				if grid[row+1][col+1] == "@" || grid[row+1][col+1] == "x" {
					paperCount++
				}

				if paperCount < 4 {
					count++
					grid[row][col] = "x"
				}
			}
		}
	}

	return count
}

func part2(grid [][]string) int {
	rollsRemoved := 0
	result := -1
	for result != 0 {
		result = part1(grid)
		rollsRemoved += result
		for row := 1; row < len(grid)-1; row++ {
			for col := 1; col < len(grid[row])-1; col++ {
				if grid[row][col] == "x" {
					grid[row][col] = "."
				}
			}
		}
	}
	return rollsRemoved
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	grid := make([][]string, 0)

	padding := make([]string, len(lines[0])+2)
	for i := 0; i < len(padding); i++ {
		padding[i] = "."
	}

	grid = append(grid, padding)
	for _, line := range lines {
		row := make([]string, len(padding))
		row[0] = "."
		for i, r := range line {
			row[i+1] = string(r)
		}
		row[len(row)-1] = "."
		grid = append(grid, row)
	}
	grid = append(grid, padding)

	//fmt.Println(fmt.Sprintf("part 1 count of accessible rolls of paper is %d", part1(grid)))
	fmt.Println(fmt.Sprintf("part 2 count of removed rolls of paper is %d", part2(grid)))
}
