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

func strToInt(l []string) []int {
	ints := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		x, _ := strconv.Atoi(l[i])
		ints[i] = x
	}
	return ints
}

func isSafelyDecreasing(report []int) bool {
	for i := 2; i < len(report); i++ {
		if report[i-1] <= report[i] || report[i-1]-report[i] > 3 {
			return false
		}
	}
	return true
}

func isSafelyIncreasing(report []int) bool {
	for i := 2; i < len(report); i++ {
		if report[i-1] >= report[i] || report[i]-report[i-1] > 3 {
			return false
		}
	}
	return true
}

func part1(lines []string) int {
	numSafe := 0
	for _, l := range lines {

		report := strToInt(strings.Split(l, " "))
		if report[0] == report[1] || math.Abs(float64(report[0]-report[1])) > 3 {
			continue
		}

		if report[0] > report[1] {
			if isSafelyDecreasing(report) {
				numSafe++
			}
		} else if isSafelyIncreasing(report) {
			numSafe++
		}
	}
	return numSafe
}

func isActuallySafe(unsafeReport []int) bool {
	for i := 0; i < len(unsafeReport); i++ {
		temp := make([]int, len(unsafeReport))
		copy(temp, unsafeReport)
		testReport := append(temp[:i], temp[i+1:]...)
		if testReport[0] == testReport[1] || math.Abs(float64(testReport[0]-testReport[1])) > 3 {
			continue
		}
		if testReport[0] > testReport[1] {
			if isSafelyDecreasing(testReport) {
				return true
			}
		} else if isSafelyIncreasing(testReport) {
			return true
		}
	}
	return false
}

func part2(lines []string) int {
	numSafe := 0
	unsafeReports := make([][]int, 0)
	for _, l := range lines {
		report := strToInt(strings.Split(l, " "))
		if report[0] == report[1] || math.Abs(float64(report[0]-report[1])) > 3 {
			unsafeReports = append(unsafeReports, report)
			continue
		}

		if report[0] > report[1] {
			if isSafelyDecreasing(report) {
				numSafe++
			} else {
				unsafeReports = append(unsafeReports, report)
			}
		} else {
			if isSafelyIncreasing(report) {
				numSafe++
			} else {
				unsafeReports = append(unsafeReports, report)
			}
		}
	}

	for _, unsafeReport := range unsafeReports {
		if isActuallySafe(unsafeReport) {
			numSafe++
		}
	}

	return numSafe
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2024/2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 sum is %d", part1(lines)))
	fmt.Println(fmt.Sprintf("part 2 sum is %d", part2(lines)))
}
