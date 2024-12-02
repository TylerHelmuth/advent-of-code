package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(lines []string) float64 {
	left := make([]float64, len(lines))
	right := make([]float64, len(lines))

	for i, l := range lines {
		locs := strings.Split(l, "   ")
		x, _ := strconv.Atoi(locs[0])
		y, _ := strconv.Atoi(locs[1])
		left[i] = float64(x)
		right[i] = float64(y)
	}

	sort.Float64s(left)
	sort.Float64s(right)

	sum := 0.0

	for i := 0; i < len(lines); i++ {
		sum += math.Abs(left[i] - right[i])
	}
	return sum
}

func part2(lines []string) float64 {
	left := make([]float64, len(lines))
	right := make([]float64, len(lines))

	for i, l := range lines {
		locs := strings.Split(l, "   ")
		x, _ := strconv.Atoi(locs[0])
		y, _ := strconv.Atoi(locs[1])
		left[i] = float64(x)
		right[i] = float64(y)
	}

	sum := 0.0
	for i := 0; i < len(lines); i++ {
		n := left[i]
		count := 0.0
		for j := 0; j < len(lines); j++ {
			if right[j] == n {
				count++
			}
		}
		sum += n * count
	}
	return sum
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 sum is %f", part1(lines)))
	fmt.Println(fmt.Sprintf("part 2 sum is %f", part2(lines)))
}
