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

	l := len(grid)
	for row := 0; row < l; row++ {
		for col := 0; col < l; col++ {
			if grid[row][col] == "X" {
				// check forward
				if col+3 < l {
					if grid[row][col+1] == "M" {
						if grid[row][col+2] == "A" {
							if grid[row][col+3] == "S" {
								count++
							}
						}
					}
				}
				// check down and right
				if row+3 < l && col+3 < l {
					if grid[row+1][col+1] == "M" {
						if grid[row+2][col+2] == "A" {
							if grid[row+3][col+3] == "S" {
								count++
							}
						}
					}
				}
				// check down
				if row+3 < l {
					if grid[row+1][col] == "M" {
						if grid[row+2][col] == "A" {
							if grid[row+3][col] == "S" {
								count++
							}
						}
					}
				}
				// check down and left
				if row+3 < l && col-3 >= 0 {
					if grid[row+1][col-1] == "M" {
						if grid[row+2][col-2] == "A" {
							if grid[row+3][col-3] == "S" {
								count++
							}
						}
					}
				}
				// check left
				if col-3 >= 0 {
					if grid[row][col-1] == "M" {
						if grid[row][col-2] == "A" {
							if grid[row][col-3] == "S" {
								count++
							}
						}
					}
				}
				// check left and up
				if row-3 >= 0 && col-3 >= 0 {
					if grid[row-1][col-1] == "M" {
						if grid[row-2][col-2] == "A" {
							if grid[row-3][col-3] == "S" {
								count++
							}
						}
					}
				}
				// check up
				if row-3 >= 0 {
					if grid[row-1][col] == "M" {
						if grid[row-2][col] == "A" {
							if grid[row-3][col] == "S" {
								count++
							}
						}
					}
				}
				// check up and right
				if row-3 >= 0 && col+3 < l {
					if grid[row-1][col+1] == "M" {
						if grid[row-2][col+2] == "A" {
							if grid[row-3][col+3] == "S" {
								count++
							}
						}
					}
				}
			}
		}
	}

	return count
}

func part2(grid [][]string) int {
	count := 0

	l := len(grid)
	for row := 0; row < l; row++ {
		for col := 0; col < l; col++ {
			if grid[row][col] == "A" && row-1 >= 0 && row+1 < l && col-1 >= 0 && col+1 < l {
				// M M on top
				if grid[row-1][col-1] == "M" && grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S" && grid[row+1][col+1] == "S" {
					count++
				}
				// M M on right
				if grid[row-1][col-1] == "S" && grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S" && grid[row+1][col+1] == "M" {
					count++
				}
				// M M on bottom
				if grid[row-1][col-1] == "S" && grid[row-1][col+1] == "S" && grid[row+1][col-1] == "M" && grid[row+1][col+1] == "M" {
					count++
				}
				// M M on left
				if grid[row-1][col-1] == "M" && grid[row-1][col+1] == "S" && grid[row+1][col-1] == "M" && grid[row+1][col+1] == "S" {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2024/4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	grid := make([][]string, len(lines))
	for i, l := range lines {
		grid[i] = strings.Split(l, "")
	}

	fmt.Println(fmt.Sprintf("part 1 count is %d", part1(grid)))
	fmt.Println(fmt.Sprintf("part 2 count is %d", part2(grid)))
}
