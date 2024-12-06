package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func travelNorth(grid [][]string, s coord, distinctPos int) (coord, int, bool) {
	for row := s.row; row >= 0; row-- {
		if row-1 < 0 {
			return coord{}, distinctPos, false
		}
		// Obstacle in front
		if grid[row-1][s.col] == "#" {
			return coord{row, s.col}, distinctPos, true
		}

		if grid[row-1][s.col] != "X" {
			distinctPos++
			grid[row-1][s.col] = "X"
		}
	}
	panic("we shouldn't get here")
}

func travelEast(grid [][]string, s coord, distinctPos int) (coord, int, bool) {
	for col := s.col; col < len(grid[s.row]); col++ {
		if col+1 == len(grid[s.row]) {
			return coord{}, distinctPos, false
		}
		// Obstacle in front
		if grid[s.row][col+1] == "#" {
			return coord{s.row, col}, distinctPos, true
		}

		if grid[s.row][col+1] != "X" {
			distinctPos++
			grid[s.row][col+1] = "X"
		}
	}
	panic("we shouldn't get here")
}

func travelSouth(grid [][]string, s coord, distinctPos int) (coord, int, bool) {
	for row := s.row; row < len(grid[s.row]); row++ {
		if row+1 == len(grid) {
			return coord{}, distinctPos, false
		}
		// Obstacle in front
		if grid[row+1][s.col] == "#" {
			return coord{row, s.col}, distinctPos, true
		}

		if grid[row+1][s.col] != "X" {
			distinctPos++
			grid[row+1][s.col] = "X"
		}
	}
	panic("we shouldn't get here")
}

func travelWest(grid [][]string, s coord, distinctPos int) (coord, int, bool) {
	for col := s.col; col >= 0; col-- {
		if col-1 < 0 {
			return coord{}, distinctPos, false
		}
		// Obstacle in front
		if grid[s.row][col-1] == "#" {
			return coord{s.row, col}, distinctPos, true
		}

		if grid[s.row][col-1] != "X" {
			distinctPos++
			grid[s.row][col-1] = "X"
		}
	}
	panic("we shouldn't get here")
}

func part1(grid [][]string, startingPos coord) int {
	currentDirection := "north"
	currentPos := startingPos
	distinctPos := 1
	cont := true
	for {
		switch currentDirection {
		case "north":
			currentPos, distinctPos, cont = travelNorth(grid, currentPos, distinctPos)
			currentDirection = "east"
		case "east":
			currentPos, distinctPos, cont = travelEast(grid, currentPos, distinctPos)
			currentDirection = "south"
		case "south":
			currentPos, distinctPos, cont = travelSouth(grid, currentPos, distinctPos)
			currentDirection = "west"
		case "west":
			currentPos, distinctPos, cont = travelWest(grid, currentPos, distinctPos)
			currentDirection = "north"
		}
		if !cont {
			break
		}
	}
	return distinctPos
}

func travelNorthPart2(grid [][]string, s coord, obstaclesHit map[obstacle]bool) (coord, bool, map[obstacle]bool, bool) {
	for row := s.row; row >= 0; row-- {
		if row-1 < 0 {
			return coord{}, false, obstaclesHit, false
		}
		// Obstacle in front
		if grid[row-1][s.col] == "#" {
			o := obstacle{
				row:       row - 1,
				col:       s.col,
				direction: "north",
			}
			if _, ok := obstaclesHit[o]; ok {
				return coord{}, false, obstaclesHit, true
			}
			obstaclesHit[o] = true
			return coord{row, s.col}, true, obstaclesHit, false
		}
	}
	panic("we shouldn't get here")
}

func travelEastPart2(grid [][]string, s coord, obstaclesHit map[obstacle]bool) (coord, bool, map[obstacle]bool, bool) {
	for col := s.col; col < len(grid[s.row]); col++ {
		if col+1 == len(grid[s.row]) {
			return coord{}, false, obstaclesHit, false
		}
		// Obstacle in front
		if grid[s.row][col+1] == "#" {
			o := obstacle{
				row:       s.row,
				col:       col + 1,
				direction: "east",
			}
			if _, ok := obstaclesHit[o]; ok {
				return coord{}, false, obstaclesHit, true
			}
			obstaclesHit[o] = true
			return coord{s.row, col}, true, obstaclesHit, false
		}
	}
	panic("we shouldn't get here")
}

func travelSouthPart2(grid [][]string, s coord, obstaclesHit map[obstacle]bool) (coord, bool, map[obstacle]bool, bool) {
	for row := s.row; row < len(grid[s.row]); row++ {
		if row+1 == len(grid) {
			return coord{}, false, obstaclesHit, false
		}
		// Obstacle in front
		if grid[row+1][s.col] == "#" {
			o := obstacle{
				row:       row + 1,
				col:       s.col,
				direction: "south",
			}
			if _, ok := obstaclesHit[o]; ok {
				return coord{}, false, obstaclesHit, true
			}
			obstaclesHit[o] = true
			return coord{row, s.col}, true, obstaclesHit, false
		}
	}
	panic("we shouldn't get here")
}

func travelWestPart2(grid [][]string, s coord, obstaclesHit map[obstacle]bool) (coord, bool, map[obstacle]bool, bool) {
	for col := s.col; col >= 0; col-- {
		if col-1 < 0 {
			return coord{}, false, obstaclesHit, false
		}
		// Obstacle in front
		if grid[s.row][col-1] == "#" {
			o := obstacle{
				row:       s.row,
				col:       col - 1,
				direction: "west",
			}
			if _, ok := obstaclesHit[o]; ok {
				return coord{}, false, obstaclesHit, true
			}
			obstaclesHit[o] = true
			return coord{s.row, col}, true, obstaclesHit, false
		}
	}
	panic("we shouldn't get here")
}

func part2(grid [][]string, startingPos coord) int {
	numBlockingObstacles := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				grid[i][j] = "#"
				fmt.Println(fmt.Sprintf("testing obstacle at {%d, %d}", i, j))
				numBlockingObstacles += helper(grid, startingPos)
				//fmt.Println(fmt.Sprintf("numBlockingObstacles = %d", numBlockingObstacles))
				grid[i][j] = "."
			}
		}
	}
	return numBlockingObstacles
}

type obstacle struct {
	row, col  int
	direction string
}

func helper(grid [][]string, startingPos coord) int {
	currentDirection := "north"
	currentPos := startingPos
	cont := true
	obstaclesHit := map[obstacle]bool{}
	foundCycle := false
	for {
		switch currentDirection {
		case "north":
			currentPos, cont, obstaclesHit, foundCycle = travelNorthPart2(grid, currentPos, obstaclesHit)
			currentDirection = "east"
		case "east":
			currentPos, cont, obstaclesHit, foundCycle = travelEastPart2(grid, currentPos, obstaclesHit)
			currentDirection = "south"
		case "south":
			currentPos, cont, obstaclesHit, foundCycle = travelSouthPart2(grid, currentPos, obstaclesHit)
			currentDirection = "west"
		case "west":
			currentPos, cont, obstaclesHit, foundCycle = travelWestPart2(grid, currentPos, obstaclesHit)
			currentDirection = "north"
		}
		if !cont {
			break
		}
	}
	if foundCycle {
		return 1
	}
	return 0
}

type coord struct {
	row, col int
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2024/6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	grid := make([][]string, len(lines))
	grid2 := make([][]string, len(lines))
	var startingPos coord
	for i, l := range lines {
		grid[i] = strings.Split(l, "")
		grid2[i] = strings.Split(l, "")
		if col := slices.Index(grid[i], "^"); col > -1 {
			startingPos = coord{
				row: i,
				col: col,
			}
		}
	}
	grid[startingPos.row][startingPos.col] = "X"

	fmt.Println(fmt.Sprintf("part 1 count is %d", part1(grid, startingPos)))
	fmt.Println(fmt.Sprintf("part 2 count is %d", part2(grid2, startingPos)))
}
