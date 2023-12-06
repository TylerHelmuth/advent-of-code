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

type race struct {
	time, dist int
}

func (r *race) numWaysToWin() int {
	numWon := 0
	for held := 1; held < r.time; held++ {
		if held*(r.time-held) > r.dist {
			numWon++
		}
	}
	return numWon
}

func part1(races []race) int {
	product := 1
	for _, r := range races {
		product *= r.numWaysToWin()
	}
	return product
}

var (
	numberReg = regexp.MustCompile(`(\d+)`)
)

func parseFor1(lines []string) []race {
	times := numberReg.FindAllStringSubmatch(lines[0], -1)
	dists := numberReg.FindAllStringSubmatch(lines[1], -1)

	races := make([]race, 0)

	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i][0])
		d, _ := strconv.Atoi(dists[i][0])

		r := race{
			time: t,
			dist: d,
		}
		races = append(races, r)
	}
	return races
}

func parseFor2(lines []string) race {
	time := strings.Replace(strings.Split(lines[0], ":")[1], " ", "", -1)
	dist := strings.Replace(strings.Split(lines[1], ":")[1], " ", "", -1)

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(dist)

	return race{
		time: t,
		dist: d,
	}
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2023/6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	races1 := parseFor1(lines)
	fmt.Println(part1(races1))

	race2 := parseFor2(lines)
	fmt.Println(race2.numWaysToWin())
}
