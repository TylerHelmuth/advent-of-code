package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type direction string

func (d direction) opposite() direction {
	switch d {
	case NORTH:
		return SOUTH
	case EAST:
		return WEST
	case SOUTH:
		return NORTH
	case WEST:
		return EAST
	}
	panic("this is impossible to get to")
}

const (
	NORTH direction = "N"
	SOUTH direction = "S"
	EAST  direction = "E"
	WEST  direction = "W"
)

type grid [][]pipe

type coord struct {
	r, c int
}

type loopMap map[coord]pipe

type pipe struct {
	allowedConnections map[direction]bool
	r, c               int
	isS                bool
}

func (p *pipe) hasConnection(dir direction, p2 pipe) bool {
	if !p.allowedConnections[dir] {
		return false
	}
	switch dir {
	case NORTH:
		return p.allowedConnections[NORTH] && p2.allowedConnections[SOUTH]
	case EAST:
		return p.allowedConnections[EAST] && p2.allowedConnections[WEST]
	case SOUTH:
		return p.allowedConnections[SOUTH] && p2.allowedConnections[NORTH]
	case WEST:
		return p2.allowedConnections[EAST]
	default:
		return false
	}
}

func newPipe(letter rune, r, c int) pipe {
	p := pipe{
		r:                  r,
		c:                  c,
		allowedConnections: map[direction]bool{},
	}

	switch letter {
	case '|':
		p.allowedConnections[NORTH] = true
		p.allowedConnections[SOUTH] = true
	case '-':
		p.allowedConnections[EAST] = true
		p.allowedConnections[WEST] = true
	case 'L':
		p.allowedConnections[NORTH] = true
		p.allowedConnections[EAST] = true
	case 'J':
		p.allowedConnections[NORTH] = true
		p.allowedConnections[WEST] = true
	case '7':
		p.allowedConnections[WEST] = true
		p.allowedConnections[SOUTH] = true
	case 'F':
		p.allowedConnections[EAST] = true
		p.allowedConnections[SOUTH] = true
	case 'S':
		p.allowedConnections[NORTH] = true
		p.allowedConnections[SOUTH] = true
		p.allowedConnections[EAST] = true
		p.allowedConnections[WEST] = true
		p.isS = true
	}

	return p
}

func getNeighbor(g grid, r, c int) pipe {
	if r < len(g) && c < len(g[r]) {
		return g[r][c]
	}
	return newPipe('.', r, c)
}

func part1(g grid, sRow, sCol int) (float64, []pipe) {
	currentPipe := g[sRow][sCol]
	loop := []pipe{currentPipe}
	var cameFrom direction

	for {
		for d, allowed := range currentPipe.allowedConnections {
			if !allowed || d == cameFrom {
				continue
			}
			connects := false
			var neighbor pipe
			switch d {
			case NORTH:
				neighbor = getNeighbor(g, currentPipe.r-1, currentPipe.c)
				connects = currentPipe.hasConnection(NORTH, neighbor)
			case EAST:
				neighbor = getNeighbor(g, currentPipe.r, currentPipe.c+1)
				connects = currentPipe.hasConnection(EAST, neighbor)
			case SOUTH:
				neighbor = getNeighbor(g, currentPipe.r+1, currentPipe.c)
				connects = currentPipe.hasConnection(SOUTH, neighbor)
			case WEST:
				neighbor = getNeighbor(g, currentPipe.r, currentPipe.c-1)
				connects = currentPipe.hasConnection(WEST, neighbor)
			}
			if connects {
				if currentPipe.isS && len(loop) > 1 {
					return math.Floor(float64(len(loop)) / 2.0), loop
				}
				loop = append(loop, currentPipe)
				currentPipe = neighbor
				cameFrom = d.opposite()
				break
			}
		}
	}
}

func part2(g grid, loop []pipe) int {
	m := make(loopMap, len(loop))
	for _, p := range loop {
		// Actually set S to its pipe type
		if p.isS {
			p.allowedConnections[NORTH] = p.hasConnection(NORTH, getNeighbor(g, p.r-1, p.c))
			p.allowedConnections[EAST] = p.hasConnection(EAST, getNeighbor(g, p.r, p.c+1))
			p.allowedConnections[SOUTH] = p.hasConnection(SOUTH, getNeighbor(g, p.r+1, p.c))
			p.allowedConnections[WEST] = p.hasConnection(WEST, getNeighbor(g, p.r, p.c-1))
		}
		m[coord{r: p.r, c: p.c}] = p
	}

	count := 0
	for r := 0; r < len(g); r++ {
		numPipesCrossed := 0
		for c := 0; c < len(g[r]); c++ {
			currentPipe := g[r][c]
			if _, ok := m[coord{r: currentPipe.r, c: currentPipe.c}]; !ok {
				if numPipesCrossed%2 == 1 {
					count++
				} else {
					continue
				}
			} else {
				numNorth := 0
				numSouth := 0

				if currentPipe.allowedConnections[NORTH] {
					numNorth++
				}
				if currentPipe.allowedConnections[SOUTH] {
					numSouth++
				}

				for !(numNorth == 2 || numSouth == 2 || (numNorth == 1 && numSouth == 1)) {
					c++
					currentPipe = getNeighbor(g, currentPipe.r, currentPipe.c+1)
					if currentPipe.allowedConnections[NORTH] {
						numNorth++
					}
					if currentPipe.allowedConnections[SOUTH] {
						numSouth++
					}
				}

				if numNorth == 2 || numSouth == 2 {
					numPipesCrossed += 2
				} else {
					numPipesCrossed++
				}

			}
		}
	}
	return count
}

func parse(lines []string) (grid, int, int) {
	g := make([][]pipe, len(lines))
	var sRow, sCol int
	for r := range lines {
		g[r] = make([]pipe, len(lines[r]))
		for c := range lines[r] {
			letter := lines[r][c]
			if letter == 'S' {
				sRow = r
				sCol = c
			}
			g[r][c] = newPipe(rune(letter), r, c)
		}
	}
	return g, sRow, sCol
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2023/10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	g, sRow, sCol := parse(lines)

	part1Answer, loop := part1(g, sRow, sCol)

	fmt.Println(part1Answer)

	fmt.Println(part2(g, loop))
}
