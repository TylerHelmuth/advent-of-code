package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	cardValues = map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	cardValues2 = map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 11,
		"K": 12,
		"A": 13,
	}
)

type handType int

const (
	highCard     handType = iota
	onePair      handType = iota
	twoPair      handType = iota
	threeOfAKind handType = iota
	fullHouse    handType = iota
	fourOfAKind  handType = iota
	fiveOfAKind  handType = iota
)

type hand struct {
	cards     []string
	ht        handType
	ht2       handType
	rank, bid int
}

func (h *hand) setType() {
	countsMap := map[string]int{
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"T": 0,
		"J": 0,
		"Q": 0,
		"K": 0,
		"A": 0,
	}
	for _, c := range h.cards {
		countsMap[c]++
	}
	counts := make([]int, 0)
	for _, v := range countsMap {
		if v != 0 {
			counts = append(counts, v)
		}
	}
	sort.Ints(counts)
	slices.Reverse(counts)
	if counts[0] == 5 {
		h.ht = fiveOfAKind
	}
	if counts[0] == 4 {
		h.ht = fourOfAKind
	}
	if counts[0] == 3 {
		if counts[1] == 2 {
			h.ht = fullHouse
		} else {
			h.ht = threeOfAKind
		}
	}
	if counts[0] == 2 {
		if counts[1] == 2 {
			h.ht = twoPair
		} else {
			h.ht = onePair
		}
	}
}

func (h *hand) setType2() {
	countsMap := map[string]int{
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"T": 0,
		"J": 0,
		"Q": 0,
		"K": 0,
		"A": 0,
	}
	for _, c := range h.cards {
		countsMap[c]++
	}
	numJ := countsMap["J"]
	delete(countsMap, "J")
	counts := make([]int, 0)
	for _, v := range countsMap {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	slices.Reverse(counts)

	counts[0] += numJ
	if counts[0] == 5 {
		h.ht = fiveOfAKind
	}
	if counts[0] == 4 {
		h.ht = fourOfAKind
	}
	if counts[0] == 3 {
		if counts[1] == 2 {
			h.ht = fullHouse
		} else {
			h.ht = threeOfAKind
		}
	}
	if counts[0] == 2 {
		if counts[1] == 2 {
			h.ht = twoPair
		} else {
			h.ht = onePair
		}
	}
}

func sortHands(hands []hand) []hand {
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		index := 0
		for a.cards[index] == b.cards[index] {
			index++
		}

		return cardValues[a.cards[index]] < cardValues[b.cards[index]]

	})
	return hands
}

func sortHands2(hands []hand) []hand {
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		index := 0
		for a.cards[index] == b.cards[index] {
			index++
		}

		return cardValues2[a.cards[index]] < cardValues2[b.cards[index]]

	})
	return hands
}

func part1(fiveOfAKinds []hand, fourOfAKinds []hand, fullHouses []hand, threeOfAKinds []hand, twoPairs []hand, onePairs []hand, highCards []hand) int {
	sortedFives := sortHands(fiveOfAKinds)
	sortedFours := sortHands(fourOfAKinds)
	sortedFullHouses := sortHands(fullHouses)
	sortedThrees := sortHands(threeOfAKinds)
	sortedTwos := sortHands(twoPairs)
	sortedPairs := sortHands(onePairs)
	sortedHighCards := sortHands(highCards)

	allHands := append(sortedHighCards, append(sortedPairs, append(sortedTwos, append(sortedThrees, append(sortedFullHouses, append(sortedFours, sortedFives...)...)...)...)...)...)

	sum := 0
	for i, h := range allHands {
		sum += h.bid * (i + 1)
	}
	return sum
}

func part2(fiveOfAKinds []hand, fourOfAKinds []hand, fullHouses []hand, threeOfAKinds []hand, twoPairs []hand, onePairs []hand, highCards []hand) int {
	sortedFives := sortHands2(fiveOfAKinds)
	sortedFours := sortHands2(fourOfAKinds)
	sortedFullHouses := sortHands2(fullHouses)
	sortedThrees := sortHands2(threeOfAKinds)
	sortedTwos := sortHands2(twoPairs)
	sortedPairs := sortHands2(onePairs)
	sortedHighCards := sortHands2(highCards)

	allHands := append(sortedHighCards, append(sortedPairs, append(sortedTwos, append(sortedThrees, append(sortedFullHouses, append(sortedFours, sortedFives...)...)...)...)...)...)

	sum := 0
	for i, h := range allHands {
		sum += h.bid * (i + 1)
	}
	return sum
}

func parse(lines []string) ([]hand, []hand, []hand, []hand, []hand, []hand, []hand) {
	fiveOfAKinds := make([]hand, 0)
	fourOfAKinds := make([]hand, 0)
	fullHouses := make([]hand, 0)
	threeOfAKinds := make([]hand, 0)
	twoPairs := make([]hand, 0)
	onePairs := make([]hand, 0)
	highCards := make([]hand, 0)

	for _, l := range lines {
		split := strings.Split(strings.Trim(l, " "), " ")
		b, _ := strconv.Atoi(split[1])
		h := hand{
			cards: strings.Split(split[0], ""),
			bid:   b,
		}

		h.setType()

		switch h.ht {
		case fiveOfAKind:
			fiveOfAKinds = append(fiveOfAKinds, h)
		case fourOfAKind:
			fourOfAKinds = append(fourOfAKinds, h)
		case fullHouse:
			fullHouses = append(fullHouses, h)
		case threeOfAKind:
			threeOfAKinds = append(threeOfAKinds, h)
		case twoPair:
			twoPairs = append(twoPairs, h)
		case onePair:
			onePairs = append(onePairs, h)
		case highCard:
			highCards = append(highCards, h)
		}
	}

	return fiveOfAKinds, fourOfAKinds, fullHouses, threeOfAKinds, twoPairs, onePairs, highCards
}

func parse2(lines []string) ([]hand, []hand, []hand, []hand, []hand, []hand, []hand) {
	fiveOfAKinds := make([]hand, 0)
	fourOfAKinds := make([]hand, 0)
	fullHouses := make([]hand, 0)
	threeOfAKinds := make([]hand, 0)
	twoPairs := make([]hand, 0)
	onePairs := make([]hand, 0)
	highCards := make([]hand, 0)

	for _, l := range lines {
		split := strings.Split(strings.Trim(l, " "), " ")
		b, _ := strconv.Atoi(split[1])
		h := hand{
			cards: strings.Split(split[0], ""),
			bid:   b,
		}

		h.setType2()

		switch h.ht {
		case fiveOfAKind:
			fiveOfAKinds = append(fiveOfAKinds, h)
		case fourOfAKind:
			fourOfAKinds = append(fourOfAKinds, h)
		case fullHouse:
			fullHouses = append(fullHouses, h)
		case threeOfAKind:
			threeOfAKinds = append(threeOfAKinds, h)
		case twoPair:
			twoPairs = append(twoPairs, h)
		case onePair:
			onePairs = append(onePairs, h)
		case highCard:
			highCards = append(highCards, h)
		}
	}

	return fiveOfAKinds, fourOfAKinds, fullHouses, threeOfAKinds, twoPairs, onePairs, highCards
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2023/7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	fiveOfAKinds, fourOfAKinds, fullHouses, threeOfAKinds, twoPairs, onePairs, highCards := parse(lines)

	fmt.Println(part1(fiveOfAKinds, fourOfAKinds, fullHouses, threeOfAKinds, twoPairs, onePairs, highCards))

	fiveOfAKinds, fourOfAKinds, fullHouses, threeOfAKinds, twoPairs, onePairs, highCards = parse2(lines)

	fmt.Println(part2(fiveOfAKinds, fourOfAKinds, fullHouses, threeOfAKinds, twoPairs, onePairs, highCards))
}
