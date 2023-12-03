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

type partNumber struct {
	id               int
	startingRow      int
	startingCol      int
	length           int
	isTouchingSymbol bool
}

type gear struct {
	partNumbers []partNumber
}

type coord struct {
	row, col int
}

var (
	symbolReg = regexp.MustCompile(`[^.\d]`)

	gearMap map[coord]*gear = make(map[coord]*gear)
)

func newPartNumber(schematicRow []string, col, row, digit int) (partNumber, int) {
	pn := partNumber{
		id:          digit,
		startingRow: row,
		startingCol: col,
	}
	// do some logic as log as we keep seeing digits
	length := 1
	for col+length < len(schematicRow) {
		item := schematicRow[col+length]
		digit, err := strconv.Atoi(item)
		if err != nil {
			break
		}
		pn.id = pn.id*10 + digit
		length++
	}
	pn.length = length
	return pn, col + length
}

func parse(schematic [][]string) []partNumber {
	partNumbers := make([]partNumber, 0)
	for row := 0; row < len(schematic); row++ {
		schematicRow := schematic[row]

		for col := 0; col < len(schematicRow); col++ {
			item := schematicRow[col]
			digit, err := strconv.Atoi(item)
			if err == nil { // this means item was a digit
				var pn partNumber
				pn, col = newPartNumber(schematicRow, col, row, digit)
				partNumbers = append(partNumbers, pn)
			}
		}
	}

	// check if partNumber is next to symbol
	// TODO break early
	for i := range partNumbers {
		pn := &partNumbers[i]
		for r := -1; r < 2; r++ {
			for c := pn.startingCol - 1; c < pn.startingCol+pn.length+1; c++ {
				safeToIndex := pn.startingRow+r >= 0 &&
					pn.startingRow+r < len(schematic) &&
					c >= 0 &&
					c < len(schematic[pn.startingRow+r])

				if safeToIndex {
					potentialSymbol := schematic[pn.startingRow+r][c]
					if symbolReg.MatchString(potentialSymbol) {
						pn.isTouchingSymbol = true

						// do some gear stuff
						if potentialSymbol == "*" {
							existingGear, ok := gearMap[coord{pn.startingRow + r, c}]
							if ok {
								existingGear.partNumbers = append(existingGear.partNumbers, *pn)
							} else {
								g := gear{
									partNumbers: []partNumber{
										*pn,
									},
								}
								gearMap[coord{pn.startingRow + r, c}] = &g
							}
						}
					}
				}
			}
		}
	}
	return partNumbers
}

func part1(partNumbers []partNumber) int {
	sum := 0
	for _, pn := range partNumbers {
		if pn.isTouchingSymbol {
			sum += pn.id
		}
	}
	return sum
}

func part2() int {
	sum := 0
	for _, g := range gearMap {
		if len(g.partNumbers) == 2 {
			sum += g.partNumbers[0].id * g.partNumbers[1].id
		}
	}
	return sum
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/projects/advent-of-code/2023/3/input.txt")
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

	partNumbers := parse(schematic)

	fmt.Println(part1(partNumbers))
	fmt.Println(part2())
}
