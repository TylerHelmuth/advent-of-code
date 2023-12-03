package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	redReg   = regexp.MustCompile(`(\d+) r`)
	greenReg = regexp.MustCompile(`(\d+) g`)
	blueReg  = regexp.MustCompile(`(\d+) b`)
)

func isValidPull(pull string, re *regexp.Regexp, max int) bool {
	if match := re.FindStringSubmatch(pull); len(match) > 1 {
		num, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		return num <= max
	}
	return true
}

func part1(games []string) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	sum := 0
	for _, game := range games {
		split := strings.Split(game, ":")

		gameSegment := split[0]
		pulls := strings.Split(split[1], ";")

		validGame := true
		for _, pull := range pulls {
			validGame = isValidPull(pull, redReg, maxRed) &&
				isValidPull(pull, greenReg, maxGreen) &&
				isValidPull(pull, blueReg, maxBlue)
			if !validGame {
				break
			}
		}
		if validGame {
			id, err := strconv.Atoi(strings.Split(gameSegment, " ")[1])
			if err != nil {
				log.Fatal(err)
			}
			sum += id
		}
	}

	return sum
}

func getColorNumber(pull string, re *regexp.Regexp, currentMax int) int {
	num := 0
	var err error
	if match := re.FindStringSubmatch(pull); len(match) > 1 {
		num, err = strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	if num > currentMax {
		return num
	}
	return currentMax
}

func part2(games []string) int {
	sum := 0
	for _, game := range games {
		split := strings.Split(game, ":")

		pulls := strings.Split(split[1], ";")

		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, pull := range pulls {
			maxRed = getColorNumber(pull, redReg, maxRed)
			maxGreen = getColorNumber(pull, greenReg, maxGreen)
			maxBlue = getColorNumber(pull, blueReg, maxBlue)
		}
		sum += maxRed * maxGreen * maxBlue
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

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
