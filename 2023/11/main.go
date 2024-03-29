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
	r, c                 int64
	expandedR, expandedC int64
}

func parse(lines []string) grid {
	g := make(grid, len(lines))
	for r, row := range lines {
		g[r] = strings.Split(row, "")
	}
	return g
}

func expandRows(g grid, galaxies []galaxy, scalar int64) []galaxy {
	for r := 0; r < len(g); r++ {
		row := g[r]
		if !slices.Contains(row, "#") {
			for i := range galaxies {
				ga := &galaxies[i]
				if ga.r > int64(r) {
					ga.expandedR = ga.expandedR + scalar
				}
			}
		}
	}
	return galaxies
}

func expandCols(g grid, galaxies []galaxy, scalar int64) []galaxy {
	for c := 0; c < len(g[0]); c++ {
		hasGalaxy := false
		for r := 0; r < len(g); r++ {
			if g[r][c] == "#" {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for i := range galaxies {
				ga := &galaxies[i]
				if ga.c > int64(c) {
					ga.expandedC = ga.expandedC + scalar
				}
			}
		}
	}
	return galaxies
}

func findGalaxies(g grid) []galaxy {
	galaxies := make([]galaxy, 0)
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] == "#" {
				galaxies = append(galaxies, galaxy{r: int64(r), c: int64(c), expandedR: int64(r), expandedC: int64(c)})
			}
		}
	}
	return galaxies
}

func solve(galaxies []galaxy) float64 {
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
	scalar := int64(1)
	galaxies = expandRows(g, galaxies, scalar)
	galaxies = expandCols(g, galaxies, scalar)
	fmt.Println(fmt.Sprintf("%f", solve(galaxies)))

	galaxies = findGalaxies(g)
	scalar = 1_000_000 - 1
	galaxies = expandRows(g, galaxies, scalar)
	galaxies = expandCols(g, galaxies, scalar)
	fmt.Println(fmt.Sprintf("%f", solve(galaxies)))
}
