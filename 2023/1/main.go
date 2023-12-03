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

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	var re = regexp.MustCompile(`[a-zA-Z]`)
	sum := 0
	for _, line := range lines {
		onlyNumbers := strings.Split(re.ReplaceAllLiteralString(line, ""), "")
		firstDigit := onlyNumbers[0]
		lastDigit := onlyNumbers[len(onlyNumbers)-1]
		number, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	return sum
}

var (
	numberWords = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

var (
	numbers = map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
)

func part2() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	sum := 0
	for _, line := range lines {
		chars := strings.Split(line, "")
		firstNumber := -1
		lastNumber := -1
		for i := 0; i < len(chars); i++ {
			char := chars[i]
			// check easy number case
			if foundNumber, ok := numbers[char]; ok {
				if firstNumber == -1 {
					firstNumber = foundNumber
					lastNumber = firstNumber
				} else {
					lastNumber = foundNumber
				}
			} else { // check for words
				for number, numberVal := range numberWords {
					if i+len(number) > len(line) {
						continue
					}
					if number == line[i:i+len(number)] {
						if firstNumber == -1 {
							firstNumber = numberVal
							lastNumber = firstNumber
						} else {
							lastNumber = numberVal
						}
					}
				}
			}
		}
		sum += (10 * firstNumber) + lastNumber
	}
	return sum
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
