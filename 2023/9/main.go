package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type history map[int][]int

func calcNextValue(h history) int {
	for i := len(h) - 2; i >= 0; i-- {
		currentRow := h[i]
		previousRow := h[i+1]

		newVal := currentRow[len(currentRow)-1] + previousRow[len(previousRow)-1]
		currentRow = append(currentRow, newVal)
		h[i] = currentRow
	}
	return h[0][len(h[0])-1]
}

func calcPreviousValue(h history) int {
	for i := len(h) - 2; i >= 0; i-- {
		currentRow := h[i]
		previousRow := h[i+1]

		newVal := currentRow[0] - previousRow[0]
		currentRow = append([]int{newVal}, currentRow...)
		h[i] = currentRow
	}
	return h[0][0]
}

func part1(histories []history) int {
	sum := 0
	for _, h := range histories {
		sum += calcNextValue(h)
	}
	return sum
}

func part2(histories []history) int {
	sum := 0
	for _, h := range histories {
		sum += calcPreviousValue(h)
	}
	return sum
}

func buildPredictionMap(initial []int) history {
	h := make(history)
	h[0] = initial
	currentRowIndex := 0
DONE:
	for {
		currentRow := h[currentRowIndex]
		nextRow := make([]int, len(currentRow)-1)
		atZero := true
		for i := 0; i < len(currentRow)-1; i++ {
			newNum := currentRow[i+1] - currentRow[i]
			nextRow[i] = newNum
			if atZero {
				atZero = newNum == 0
			}
		}

		currentRowIndex++
		h[currentRowIndex] = nextRow
		if atZero {
			break DONE
		}
	}
	return h
}

func parse(lines []string) []history {
	histories := make([]history, len(lines))
	for i, l := range lines {
		split := strings.Split(l, " ")
		nums := make([]int, 0)
		for _, n := range split {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}
		histories[i] = buildPredictionMap(nums)
	}
	return histories
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2023/9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	histories := parse(lines)

	fmt.Println(part1(histories))

	fmt.Println(part2(histories))
}
