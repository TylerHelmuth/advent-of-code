package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(banks []string) int {
	sumMaximumJoltage := 0

	for _, bank := range banks {
		var firstBattery int
		var secondBattery int
		for i := 0; i < len(bank)-1; i++ {
			n, _ := strconv.Atoi(string(bank[i]))
			if n > firstBattery {
				firstBattery = n
				secondBattery = 0
				for j := i + 1; j < len(bank); j++ {
					m, _ := strconv.Atoi(string(bank[j]))
					if m > secondBattery {
						secondBattery = m
					}
				}
			}
		}
		maximumJoltage := (firstBattery * 10) + secondBattery
		fmt.Println(maximumJoltage)
		sumMaximumJoltage += maximumJoltage
	}

	return sumMaximumJoltage
}

func findLargestDigit(bank []int, startingPos, digitPos int) (int, int) {
	largest := 0
	pos := 0
	for i := startingPos; i < len(bank)-11+digitPos; i++ {
		if bank[i] > largest {
			largest = bank[i]
			pos = i
		}
	}
	return largest, pos + 1
}

func part2(banks []string) int {
	sumMaximumJoltage := 0

	for _, bankStr := range banks {
		bank := make([]int, len(bankStr))
		for i, s := range bankStr {
			num, _ := strconv.Atoi(string(s))
			bank[i] = num
		}

		batteries := make([]int, 12)
		var currentPos int
		for i := 0; i < 12; i++ {
			batteries[i], currentPos = findLargestDigit(bank, currentPos, i)
		}

		numString := ""
		for _, i := range batteries {
			numString += strconv.Itoa(i)
		}
		num, _ := strconv.Atoi(numString)
		sumMaximumJoltage += num
	}

	return sumMaximumJoltage
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	banks := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 sum of maximum joltage is %d", part1(banks)))
	fmt.Println(fmt.Sprintf("part 2 sum of maximum joltage is %d", part2(banks)))
}
