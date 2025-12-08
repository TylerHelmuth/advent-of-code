package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type junctionBox struct {
	x, y, z float64
	circuit *map[*junctionBox]bool
}

func distance(a, b *junctionBox) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
}

type pair struct {
	a, b *junctionBox
	dist float64
}

func equals(a, b *map[*junctionBox]bool) bool {
	if len(*a) != len(*b) {
		return false
	}

	for k, v := range *a {
		v2, ok := (*b)[k]
		if !ok || v != v2 {
			return false
		}
	}

	return true
}

func part1(junctionBoxes []*junctionBox) int {
	distances := make([]pair, 0)

	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			dist := distance(junctionBoxes[i], junctionBoxes[j])
			distances = append(distances, pair{junctionBoxes[i], junctionBoxes[j], dist})
		}
	}

	slices.SortFunc(distances, func(a, b pair) int {
		if a.dist < b.dist {
			return -1
		}
		if a.dist > b.dist {
			return 1
		}
		return 0
	})

	currentPairIndex := 0
	circuits := make([]*map[*junctionBox]bool, 0)
	for itr := 0; itr < 1000; itr++ {
		p := distances[currentPairIndex]

		// a and b are both in a circuit
		if p.a.circuit != nil && p.b.circuit != nil {
			// a and b are in the same circuit, do nothing
			if equals(p.a.circuit, p.b.circuit) {
				currentPairIndex++
				continue
			}

			// a and b are in different circuits, merge
			c1 := *p.a.circuit
			c2 := *p.b.circuit
			for k, v := range c2 {
				c1[k] = v
				k.circuit = p.a.circuit
			}
			p.b.circuit = p.a.circuit

			for k, _ := range c2 {
				delete(c2, k)
			}

			currentPairIndex++
			continue
		}

		// a and b are not in a circuit
		if p.a.circuit == nil && p.b.circuit == nil {
			newCircuit := make(map[*junctionBox]bool)
			newCircuit[p.a] = true
			newCircuit[p.b] = true
			ptr := &newCircuit
			p.a.circuit = ptr
			p.b.circuit = ptr
			circuits = append(circuits, ptr)
			currentPairIndex++
			continue
		}

		// a is not in a circuit
		if p.a.circuit == nil && p.b.circuit != nil {
			c := *p.b.circuit
			c[p.a] = true
			p.a.circuit = p.b.circuit
			currentPairIndex++
			continue
		}

		// b is not in a circuit
		if p.b.circuit == nil && p.a.circuit != nil {
			c := *p.a.circuit
			c[p.b] = true
			p.b.circuit = p.a.circuit
			currentPairIndex++
			continue
		}
	}

	slices.SortFunc(circuits, func(a, b *map[*junctionBox]bool) int {
		return len(*b) - len(*a)
	})

	return len(*circuits[0]) * len(*circuits[1]) * len(*circuits[2])
}

func part2(junctionBoxes []*junctionBox) float64 {
	distances := make([]pair, 0)

	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			dist := distance(junctionBoxes[i], junctionBoxes[j])
			distances = append(distances, pair{junctionBoxes[i], junctionBoxes[j], dist})
		}
	}

	slices.SortFunc(distances, func(a, b pair) int {
		if a.dist < b.dist {
			return -1
		}
		if a.dist > b.dist {
			return 1
		}
		return 0
	})

	currentPairIndex := 0
	circuits := make([]*map[*junctionBox]bool, 0)
	for {
		p := distances[currentPairIndex]

		// a and b are both in a circuit
		if p.a.circuit != nil && p.b.circuit != nil {
			// a and b are in the same circuit, do nothing
			if equals(p.a.circuit, p.b.circuit) {
				currentPairIndex++
				continue
			}

			// a and b are in different circuits, merge
			c1 := *p.a.circuit
			c2 := *p.b.circuit
			for k, v := range c2 {
				c1[k] = v
				k.circuit = p.a.circuit
			}
			p.b.circuit = p.a.circuit

			for k, _ := range c2 {
				delete(c2, k)
			}

			currentPairIndex++

			// we're done
			if len(c1) == len(junctionBoxes) {
				return p.a.x * p.b.x
			}
			continue
		}

		// a and b are not in a circuit
		if p.a.circuit == nil && p.b.circuit == nil {
			newCircuit := make(map[*junctionBox]bool)
			newCircuit[p.a] = true
			newCircuit[p.b] = true
			ptr := &newCircuit
			p.a.circuit = ptr
			p.b.circuit = ptr
			circuits = append(circuits, ptr)
			currentPairIndex++
			continue
		}

		// a is not in a circuit
		if p.a.circuit == nil && p.b.circuit != nil {
			c := *p.b.circuit
			c[p.a] = true
			p.a.circuit = p.b.circuit
			currentPairIndex++

			// we're done
			if len(c) == len(junctionBoxes) {
				return p.a.x * p.b.x
			}

			continue
		}

		// b is not in a circuit
		if p.b.circuit == nil && p.a.circuit != nil {
			c := *p.a.circuit
			c[p.b] = true
			p.b.circuit = p.a.circuit
			currentPairIndex++

			// we're done
			if len(c) == len(junctionBoxes) {
				return p.a.x * p.b.x
			}

			continue
		}
	}
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	junctionBoxes := make([]*junctionBox, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(nums[0], 64)
		y, _ := strconv.ParseFloat(nums[1], 64)
		z, _ := strconv.ParseFloat(nums[2], 64)
		j := junctionBox{
			x: x,
			y: y,
			z: z,
		}
		junctionBoxes[i] = &j
	}

	//fmt.Println(fmt.Sprintf("part 1 product 3 largest circuits is %d", part1(junctionBoxes)))
	fmt.Println(fmt.Sprintf("part 2 product x coords of last 2 junction boxes is %f", part2(junctionBoxes)))
}
