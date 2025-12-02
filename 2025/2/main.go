package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(ranges []string) int64 {
	sumInvalidIds := int64(0)

	for _, r := range ranges {
		nums := strings.Split(r, "-")
		low, _ := strconv.ParseInt(nums[0], 10, 64)
		high, _ := strconv.ParseInt(nums[1], 10, 64)

		for i := low; i <= high; i++ {
			str := strconv.FormatInt(i, 10)
			middle := len(str) / 2
			if (len(str)%2) == 0 && str[:(middle)] == str[middle:] {
				sumInvalidIds += i
			}
		}
	}

	return sumInvalidIds
}

func part2(ranges []string) int64 {
	sumInvalidIds := int64(0)

	for _, r := range ranges {
		nums := strings.Split(r, "-")
		low, _ := strconv.ParseInt(nums[0], 10, 64)
		high, _ := strconv.ParseInt(nums[1], 10, 64)

		for i := low; i <= high; i++ {
			str := strconv.FormatInt(i, 10)
			for groupSize := 1; groupSize <= len(str)/2; groupSize++ {
				if len(str)%groupSize == 0 {
					matchFound := true
					for j := 0; j < len(str)-groupSize; j += groupSize {
						if str[j:j+groupSize] != str[j+groupSize:j+groupSize+groupSize] {
							matchFound = false
							break
						}
					}
					if matchFound {
						sumInvalidIds += i
						break
					}
				}
			}
		}
	}

	return sumInvalidIds
}

func main() {
	f, err := os.Open("/Users/tylerhelmuth/Projects/advent-of-code/2025/2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	ranges := strings.Split(string(b), ",")

	fmt.Println(fmt.Sprintf("part 1 sum of invalid IDs is %d", part1(ranges)))
	fmt.Println(fmt.Sprintf("part 2 sum of invalid IDs is %d", part2(ranges)))
}
