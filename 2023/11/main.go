package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type grid [][]string

type galaxy struct {
	r, c int
}

func parse(lines []string) grid {
	g := make(grid, len(lines))
	for r, row := range lines {
		g[r] = strings.Split(row, "")
	}
	return g
}

func expandRows(g grid) grid {
	expandedLine := ""
	for range g[0] {
		expandedLine += "."
	}

	for r := 0; r < len(g); r++ {
		row := g[r]
		if !slices.Contains(row, "#") {
			g = append(g[:r+1], g[r:]...)
			r++
		}
	}
	return g
}

func transpose(g grid) grid {
	newGrid := make(grid, len(g[0]))
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			newGrid[c] = append(newGrid[c], g[r][c])
		}
	}
	return newGrid
}

func findGalaxies(g grid) []galaxy {
	galaxies := make([]galaxy, 0)
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] == "#" {
				galaxies = append(galaxies, galaxy{r, c})
			}
		}
	}
	return galaxies
}

func part1(galaxies []galaxy) float64 {
	sumOfShortestLengths := 0.0
	for i, currentGalaxy := range galaxies {
		for _, compGalaxy := range galaxies[i+1:] {
			sumOfShortestLengths += math.Abs(float64(currentGalaxy.r-compGalaxy.r)) + math.Abs(float64(currentGalaxy.c-compGalaxy.c))
		}
	}
	return sumOfShortestLengths
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2023/11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	g := parse(lines)
	g = expandRows(g)
	g = transpose(g)
	g = expandRows(g)

	fmt.Println("")
	for r := 0; r < len(g); r++ {
		fmt.Println(g[r])
	}
	fmt.Println("")

	galaxies := findGalaxies(g)

	fmt.Println(fmt.Sprintf("%f", part1(galaxies)))

	//fmt.Println("")
	//for r := 0; r < len(g); r++ {
	//	fmt.Println(g[r])
	//}

	//parse(lines)
	//g, sRow, sCol := parse(lines)

	//part1Answer, loop := part1(g, sRow, sCol)

	//fmt.Println(part1Answer)
	//
	//fmt.Println(part2(g, loop))
}
