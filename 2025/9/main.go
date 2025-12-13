package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func area(coord1, coord2 []string) float64 {
	x1, _ := strconv.Atoi(coord1[0])
	y1, _ := strconv.Atoi(coord1[1])
	x2, _ := strconv.Atoi(coord2[0])
	y2, _ := strconv.Atoi(coord2[1])
	return (math.Abs(float64(x2-x1)) + 1) * (math.Abs(float64(y2-y1)) + 1)
}

func part1(lines []string) float64 {
	largestArea := 0.0
	for i := 0; i < len(lines); i++ {
		coord1 := strings.Split(lines[i], ",")
		for j := i + 1; j < len(lines); j++ {
			coord2 := strings.Split(lines[j], ",")
			a := area(coord1, coord2)
			if a > largestArea {
				largestArea = a
			}
		}
	}
	return largestArea
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 largest area rectangle is %f", part1(lines)))
}
