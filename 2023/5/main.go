package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type rule struct {
	destRangeStart, sourceRangeStart, length float64
}

type almanacMap struct {
	rules []rule
}

func (a *almanacMap) generateNewSeeds(seeds []float64) []float64 {
	newSeeds := make([]float64, len(seeds))

	for i, seed := range seeds {
		for _, r := range a.rules {
			if seed >= r.sourceRangeStart && seed < r.sourceRangeStart+r.length {
				newSeeds[i] = seed - r.sourceRangeStart + r.destRangeStart
				break
			}
			newSeeds[i] = seed
		}
	}

	return newSeeds
}

var (
	numberReg = regexp.MustCompile(`\d`)
)

func part1(lines []string) float64 {
	seedStrs := strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " ")
	seeds := make([]float64, len(seedStrs))
	for i, s := range seedStrs {
		seedNum, _ := strconv.ParseFloat(s, 64)
		seeds[i] = seedNum
	}

	currentMap := almanacMap{
		rules: make([]rule, 0),
	}
	for _, line := range lines[3:] {
		if line == "" {
			continue
		}

		// When we get in here it is time to send the numbers through a map
		if !numberReg.MatchString(line) {
			seeds = currentMap.generateNewSeeds(seeds)
			currentMap = almanacMap{}
			continue
		}
		numbers := strings.Split(line, " ")

		dest, _ := strconv.ParseFloat(numbers[0], 64)
		source, _ := strconv.ParseFloat(numbers[1], 64)
		length, _ := strconv.ParseFloat(numbers[2], 64)

		currentMap.rules = append(currentMap.rules, rule{
			destRangeStart:   dest,
			sourceRangeStart: source,
			length:           length,
		})
	}
	seeds = currentMap.generateNewSeeds(seeds)

	smallestLocation := seeds[0]
	for _, s := range seeds[1:] {
		if s < smallestLocation {
			smallestLocation = s
		}
	}
	return smallestLocation
}

type tuple struct {
	start float64
	end   float64
}

// Runs in 150 seconds
func part2(lines []string) float64 {
	start := time.Now()

	seedStrs := strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " ")
	seedRanges := make([]float64, len(seedStrs))
	for i, s := range seedStrs {
		seedNum, _ := strconv.ParseFloat(s, 64)
		seedRanges[i] = seedNum
	}

	tuples := make([]tuple, 0)
	for i := 0; i < len(seedRanges)-1; i += 2 {
		tuples = append(tuples, tuple{
			start: seedRanges[i],
			end:   seedRanges[i] + seedRanges[i+1] - 1,
		})
	}

	sort.Slice(tuples, func(i, j int) bool { return tuples[i].start < tuples[j].start })

	for i := 0; i < len(tuples)-1; i++ {
		a := tuples[i]
		b := tuples[i+1]
		// this means there is overlap
		if a.end >= b.start && b.end > a.end {
			n := tuple{
				start: a.start,
				end:   b.end,
			}
			tuples = append(tuples[:i], append([]tuple{n}, tuples[i+2:]...)...)
			continue
		}
		if a.end >= b.start && b.end < a.end {
			tuples = append(tuples[:i], tuples[i+2:]...)
			continue
		}
	}

	capcity := 0.0
	for _, t := range tuples {
		capcity += t.end - t.start + 1
	}

	seeds := make([]float64, int64(capcity))
	index := 0
	for _, t := range tuples {
		for i := 0; float64(i) <= t.end-t.start; i++ {
			seeds[index] = t.start + float64(i)
			index++
		}
	}

	end := time.Now()
	fmt.Println(fmt.Sprintf("generating seeds took %v seconds", end.Sub(start).Seconds()))

	currentMap := almanacMap{
		rules: make([]rule, 0),
	}
	for _, line := range lines[3:] {
		if line == "" {
			continue
		}

		// When we get in here it is time to send the numbers through a map
		if !numberReg.MatchString(line) {
			for i, seed := range seeds {
				for _, r := range currentMap.rules {
					if seed >= r.sourceRangeStart && seed < r.sourceRangeStart+r.length {
						seeds[i] = seed - r.sourceRangeStart + r.destRangeStart
						break
					}
					seeds[i] = seed
				}
			}
			currentMap = almanacMap{}
			continue
		}
		numbers := strings.Split(line, " ")

		dest, _ := strconv.ParseFloat(numbers[0], 64)
		source, _ := strconv.ParseFloat(numbers[1], 64)
		length, _ := strconv.ParseFloat(numbers[2], 64)

		currentMap.rules = append(currentMap.rules, rule{
			destRangeStart:   dest,
			sourceRangeStart: source,
			length:           length,
		})
	}

	for i, seed := range seeds {
		for _, r := range currentMap.rules {
			if seed >= r.sourceRangeStart && seed < r.sourceRangeStart+r.length {
				seeds[i] = seed - r.sourceRangeStart + r.destRangeStart
				break
			}
			seeds[i] = seed
		}
	}

	smallestLocation := seeds[0]
	for _, s := range seeds[1:] {
		if s < smallestLocation {
			smallestLocation = s
		}
	}
	return smallestLocation
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/projects/advent-of-code/2023/5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	start := time.Now()
	fmt.Println(fmt.Sprintf("%f", part1(lines)))
	end := time.Now()
	fmt.Println(fmt.Sprintf("part 1 took %v seconds", end.Sub(start).Seconds()))

	start = time.Now()
	fmt.Println(fmt.Sprintf("%f", part2(lines)))
	end = time.Now()
	fmt.Println(fmt.Sprintf("part 2 took %v seconds", end.Sub(start).Seconds()))
}
