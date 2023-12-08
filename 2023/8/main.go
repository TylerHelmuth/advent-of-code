package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type node struct {
	id   string
	l, r *node
}

var (
	existingNodes = map[string]*node{}
)

func createNode(line string) {
	id := line[0:3]
	lID := line[7:10]
	rID := line[12:15]

	var currentNode *node
	if n, ok := existingNodes[id]; ok {
		currentNode = n
	} else {
		currentNode = &node{
			id: id,
		}
	}

	lNode, ok := existingNodes[lID]
	if !ok {
		lNode = &node{
			id: lID,
		}
		existingNodes[lID] = lNode
	}
	currentNode.l = lNode

	rNode, ok := existingNodes[rID]
	if !ok {
		rNode = &node{
			id: rID,
		}
		existingNodes[rID] = rNode
	}
	currentNode.r = rNode

	existingNodes[id] = currentNode
}

func parse(lines []string) []string {
	instructions := strings.Split(lines[0], "")

	for _, l := range lines[2:] {
		createNode(l)
	}

	return instructions
}

func part1(instructions []string) int {
	currentId := "AAA"
	currentNode := existingNodes[currentId]
	count := 0
DONE:
	for {
		for _, i := range instructions {
			count++
			if i == "R" {
				currentId = currentNode.r.id
				currentNode = currentNode.r
			}
			if i == "L" {
				currentId = currentNode.l.id
				currentNode = currentNode.l
			}
			if currentNode.id == "ZZZ" {
				break DONE
			}
		}
	}
	return count
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2(instructions []string) int {
	currentNodes := make([]*node, 0)
	for id, n := range existingNodes {
		if id[2] == 'A' {
			currentNodes = append(currentNodes, n)
		}
	}

	numStepsToZ := make([]int, len(currentNodes))
	for i, n := range currentNodes {
		currentId := n.id
		currentNode := existingNodes[currentId]
		numSteps := 0
	FOUNDZ:
		for {
			for _, instr := range instructions {
				numSteps++
				if instr == "R" {
					currentId = currentNode.r.id
					currentNode = currentNode.r
				}
				if instr == "L" {
					currentId = currentNode.l.id
					currentNode = currentNode.l
				}
				if currentNode.id[2] == 'Z' {
					break FOUNDZ
				}
			}
		}
		numStepsToZ[i] = numSteps
	}

	return LCM(numStepsToZ[0], numStepsToZ[1], numStepsToZ[2:]...)
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2023/8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	instructions := parse(lines)

	fmt.Println(part1(instructions))

	fmt.Println(part2(instructions))
}
