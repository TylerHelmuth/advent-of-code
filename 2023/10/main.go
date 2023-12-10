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

func part1Helper(g grid, loop []pipe, cameFrom direction, currentPipe pipe) float64 {
	cont := true
	for cont {
	keepGoing:
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
				if currentPipe.isS {
					cont = false
					break
				}
				loop = append(loop, currentPipe)
				currentPipe = neighbor
				cameFrom = d.opposite()
				goto keepGoing
			}
		}
		cont = false
	}

	return math.Floor(float64(len(loop)) / 2.0)
}

func part1(g grid, sRow, sCol int) float64 {
	sPipe := g[sRow][sCol]

	for startingDir := range sPipe.allowedConnections {
		connects := false
		var neighbor pipe
		switch startingDir {
		case NORTH:
			neighbor = getNeighbor(g, sPipe.r-1, sPipe.c)
			connects = sPipe.hasConnection(NORTH, neighbor)
		case EAST:
			neighbor = getNeighbor(g, sPipe.r, sPipe.c+1)
			connects = sPipe.hasConnection(EAST, neighbor)
		case SOUTH:
			neighbor = getNeighbor(g, sPipe.r+1, sPipe.c)
			connects = sPipe.hasConnection(SOUTH, neighbor)
		case WEST:
			neighbor = getNeighbor(g, sPipe.r, sPipe.c-1)
			connects = sPipe.hasConnection(WEST, neighbor)
		}

		// We can start actually looking for a loop
		if connects {
			loop := []pipe{sPipe}
			return part1Helper(g, loop, startingDir.opposite(), neighbor)
		}
	}
	return -1.0
}

//func part2(lines []string) int {
//	sum := 0
//	return sum
//}

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

	fmt.Println(part1(g, sRow, sCol))

	//fmt.Println(part2(lines))
}
