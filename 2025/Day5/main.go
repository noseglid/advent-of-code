package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Range struct {
	s, e int
}

func (r Range) Len() int {
	return r.e - r.s + 1
}

func (r Range) Contains(i int) bool {
	return i >= r.s && i <= r.e
}

func parseRange(l string) Range {
	parts := strings.Split(l, "-")
	return Range{s: util.MustAtoi(parts[0]), e: util.MustAtoi(parts[1])}
}

func main() {
	lines := util.GetFileStrings("2025/Day5/input")

	parseRanges := true
	var ranges []Range
	var ingredients []int

	for _, l := range lines {
		if l == "" {
			parseRanges = false
			continue
		}

		if parseRanges {
			ranges = append(ranges, parseRange(l))
			continue
		}

		ingredients = append(ingredients, util.MustAtoi(l))
	}

	n := 0
	for _, ing := range ingredients {
		for _, r := range ranges {
			if r.Contains(ing) {
				n++
				break
			}
		}
	}
	fmt.Printf("number of fresh ingredients (part1): %d\n", n)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].s < ranges[j].s
	})

	var newRanges []Range
	curr := ranges[0]
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if r.s <= curr.e+1 {
			if r.e > curr.e {
				curr.e = r.e
			}
			continue
		}

		newRanges = append(newRanges, curr)
		curr = r
	}
	newRanges = append(newRanges, curr)

	ss := 0
	for _, r := range newRanges {
		ss += r.Len()
	}

	fmt.Printf("total number of fresh ingredients (part2): %d\n", ss)

}
