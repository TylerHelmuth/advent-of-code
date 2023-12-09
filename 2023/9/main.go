package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type history struct {
	initial []int

	predicitonMap map[int][]int
}

func (h *history) calcNextValue() int {
	for i := len(h.predicitonMap) - 2; i >= 0; i-- {
		currentRow := h.predicitonMap[i]
		previousRow := h.predicitonMap[i+1]

		newVal := currentRow[len(currentRow)-1] + previousRow[len(previousRow)-1]
		currentRow = append(currentRow, newVal)
		h.predicitonMap[i] = currentRow
	}
	return h.predicitonMap[0][len(h.predicitonMap[0])-1]
}

func (h *history) calcPreviousValue() int {
	for i := len(h.predicitonMap) - 2; i >= 0; i-- {
		currentRow := h.predicitonMap[i]
		previousRow := h.predicitonMap[i+1]

		newVal := currentRow[0] - previousRow[0]
		currentRow = append([]int{newVal}, currentRow...)
		h.predicitonMap[i] = currentRow
	}
	return h.predicitonMap[0][0]
}

func part1(histories []history) int {
	sum := 0
	for _, h := range histories {
		sum += h.calcNextValue()
	}
	return sum
}

func part2(histories []history) int {
	sum := 0
	for _, h := range histories {
		sum += h.calcPreviousValue()
	}
	return sum
}

func buildPredictionMap(initial []int) map[int][]int {
	predictionMap := make(map[int][]int)
	predictionMap[0] = initial
	currentRowIndex := 0
DONE:
	for {
		currentRow := predictionMap[currentRowIndex]
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
		predictionMap[currentRowIndex] = nextRow
		if atZero {
			break DONE
		}
	}
	return predictionMap
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

		histories[i] = history{
			initial:       nums,
			predicitonMap: buildPredictionMap(nums),
		}
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
