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

type card struct {
	id             int
	winningNumbers map[string]bool
	numbers        []string
}

func (c *card) worth() float64 {
	numMatches := 0.0
	for _, num := range c.numbers {
		if _, ok := c.winningNumbers[num]; ok {
			numMatches++
		}
	}
	if numMatches > 0 {
		return math.Pow(2, numMatches-1)
	}
	return 0
}

func (c *card) numMatches() int {
	numMatches := 0
	for _, num := range c.numbers {
		if _, ok := c.winningNumbers[num]; ok {
			numMatches++
		}
	}
	return numMatches
}

func parse(lines []string) []card {
	cards := make([]card, 0)
	for _, line := range lines {
		split := strings.Split(line, ":")

		cardSegments := strings.Split(split[0], " ")
		idStr := cardSegments[len(cardSegments)-1]

		split = strings.Split(split[1], "|")

		w := strings.Split(strings.Trim(split[0], " "), " ")
		n := strings.Split(strings.Trim(split[1], " "), " ")

		winningNumbers := map[string]bool{}
		for _, num := range w {
			winningNumbers[num] = true
		}

		numbers := make([]string, 0)
		for _, num := range n {
			if num != "" {
				numbers = append(numbers, num)
			}
		}

		id, _ := strconv.Atoi(idStr)
		c := card{
			id:             id,
			winningNumbers: winningNumbers,
			numbers:        numbers,
		}

		cards = append(cards, c)
	}
	return cards
}

func part1(cards []card) float64 {
	sum := 0.0
	for _, c := range cards {
		sum += c.worth()
	}
	return sum
}

func part2Helper(cardsWon map[int]int, c card, cards map[int]card) {
	cardsWon[c.id]++
	numMatchs := c.numMatches()
	if numMatchs > 0 {
		for i := c.id + 1; i <= c.id+numMatchs; i++ {
			part2Helper(cardsWon, cards[i], cards)
		}
	}
}

func part2(cards map[int]card) int {
	cardsWon := map[int]int{}
	for _, c := range cards {
		part2Helper(cardsWon, c, cards)
	}

	sum := 0
	for _, val := range cardsWon {
		sum += val
	}
	return sum
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/projects/advent-of-code/2023/4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	var schematic [][]string
	for _, line := range lines {
		schematic = append(schematic, strings.Split(line, ""))
	}

	cards := parse(lines)

	cardMap := map[int]card{}
	for _, c := range cards {
		cardMap[c.id] = c
	}

	fmt.Println(part1(cards))
	fmt.Println(part2(cardMap))
}
