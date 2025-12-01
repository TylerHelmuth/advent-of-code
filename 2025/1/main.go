package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(rotations []string) int64 {
	numTimesHitZero := int64(0)
	currentPosition := int64(50)
	for _, rotation := range rotations {
		num, _ := strconv.ParseInt(rotation[1:], 10, 64)
		if rotation[0] == 'R' {
			currentPosition += num
		} else {
			currentPosition -= num
		}

		if currentPosition < 0 {
			for currentPosition < 0 {
				currentPosition += 100
			}
		} else if currentPosition > 99 {
			for currentPosition > 99 {
				currentPosition -= 100
			}
		}

		if currentPosition == 0 {
			numTimesHitZero++
		}
	}
	return numTimesHitZero
}

func part2(rotations []string) int64 {
	numTimesHitZero := int64(0)
	currentPosition := int64(50)
	for _, rotation := range rotations {
		startingPosition := currentPosition
		num, _ := strconv.ParseInt(rotation[1:], 10, 64)
		if rotation[0] == 'R' {
			currentPosition += num
		} else {
			currentPosition -= num
		}

		if currentPosition < 0 {
			if startingPosition == 0 {
				numTimesHitZero--
			}
			for currentPosition < 0 {
				currentPosition += 100
				numTimesHitZero++
			}
			if currentPosition == 0 {
				numTimesHitZero++
			}
		} else if currentPosition > 99 {
			for currentPosition > 99 {
				currentPosition -= 100
				numTimesHitZero++
			}
		} else if currentPosition == 0 {
			numTimesHitZero++
		}
	}
	return numTimesHitZero
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	rotations := strings.Split(string(b), "\n")

	fmt.Println(fmt.Sprintf("part 1 password is %d", part1(rotations)))
	fmt.Println(fmt.Sprintf("part 2 password is %d", part2(rotations)))
}
