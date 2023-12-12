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
	r, c                 int
	expandedR, expandedC int
}

func parse(lines []string) grid {
	g := make(grid, len(lines))
	for r, row := range lines {
		g[r] = strings.Split(row, "")
	}
	return g
}

func expandRows(g grid, galaxies []galaxy, scalar int) []galaxy {
	for r := 0; r < len(g); r++ {
		row := g[r]
		if !slices.Contains(row, "#") {
			for i := range galaxies {
				ga := &galaxies[i]
				if ga.r > r {
					ga.expandedR = ga.expandedR + scalar
				}
			}
		}
	}
	return galaxies
}

func transpose(g grid, galaxies []galaxy) (grid, []galaxy) {
	newGrid := make(grid, len(g[0]))
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			newGrid[c] = append(newGrid[c], g[r][c])
		}
	}
	newGalaxies := make([]galaxy, len(galaxies))
	for i, ga := range galaxies {
		newGalaxies[i] = galaxy{r: ga.c, c: ga.r, expandedR: ga.expandedC, expandedC: ga.expandedR}
	}
	return newGrid, newGalaxies
}

func findGalaxies(g grid) []galaxy {
	galaxies := make([]galaxy, 0)
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] == "#" {
				galaxies = append(galaxies, galaxy{r: r, c: c, expandedR: r, expandedC: c})
			}
		}
	}
	return galaxies
}

func part1(galaxies []galaxy) float64 {
	sumOfShortestLengths := 0.0
	for i, currentGalaxy := range galaxies {
		for _, compGalaxy := range galaxies[i+1:] {
			sumOfShortestLengths += math.Abs(float64(currentGalaxy.expandedR-compGalaxy.expandedR)) + math.Abs(float64(currentGalaxy.expandedC-compGalaxy.expandedC))
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
	galaxies := findGalaxies(g)
	galaxies = expandRows(g, galaxies, 1)
	fmt.Println("")
	for r := 0; r < len(g); r++ {
		fmt.Println(g[r])
	}
	fmt.Println("")
	fmt.Println(galaxies)
	g, galaxies = transpose(g, galaxies)
	fmt.Println("")
	for r := 0; r < len(g); r++ {
		fmt.Println(g[r])
	}
	fmt.Println("")
	fmt.Println(galaxies)
	galaxies = expandRows(g, galaxies, 1)
	fmt.Println("")
	for r := 0; r < len(g); r++ {
		fmt.Println(g[r])
	}
	fmt.Println("")
	fmt.Println(galaxies)

	fmt.Println(fmt.Sprintf("%f", part1(galaxies)))
}
