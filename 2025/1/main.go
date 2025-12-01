package main

import (
	"io"
	"log"
	"os"
	"strings"
)

//func part1(lines []string) float64 {
//
//}
//
//func part2(lines []string) float64 {
//
//}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/1/input.txt")
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

	//fmt.Println(fmt.Sprintf("part 1 count is %d", part1(grid)))
	//fmt.Println(fmt.Sprintf("part 2 count is %d", part2(grid)))
}
