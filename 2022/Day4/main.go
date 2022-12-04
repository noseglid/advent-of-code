package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func contains(a1, a2 assignment) bool {
	if a1.begin >= a2.begin && a1.end <= a2.end {
		return true
	}
	if a2.begin >= a1.begin && a2.end <= a1.end {
		return true
	}
	return false
}

func overlaps1d(a1, a2 assignment) bool {
	if a1.begin >= a2.begin && a1.begin <= a2.end {
		return true
	}
	if a1.end >= a2.begin && a1.end <= a2.end {
		return true
	}
	return false
}

type assignment struct {
	begin, end int
}

type pair struct {
	a1, a2 assignment
}

func main() {
	input := util.GetFileStrings("2022/Day4/sample")

	var pairs []pair
	for _, l := range input {
		var a1, a2 assignment
		_, err := fmt.Sscanf(l, "%d-%d,%d-%d", &a1.begin, &a1.end, &a2.begin, &a2.end)
		if err != nil {
			panic(err)
		}
		pairs = append(pairs, pair{a1, a2})
	}

	n, n2 := 0, 0
	for _, p := range pairs {
		if contains(p.a1, p.a2) {
			n++
		}
		if overlaps1d(p.a1, p.a2) || overlaps1d(p.a2, p.a1) {
			n2++
		}
	}

	fmt.Printf("contained pairs (part1): %d\n", n)
	fmt.Printf("overlapped pairs (part2): %d\n", n2)
}
