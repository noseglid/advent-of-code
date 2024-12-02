package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func getSign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func getDiffs(s []string) []int {
	diffs := []int{}
	for i, n := range s {
		if i == 0 {
			continue
		}
		diffs = append(diffs, util.MustAtoi(n)-util.MustAtoi(s[i-1]))
	}
	return diffs
}

func isSafe(diffs []int) bool {
	sign := 99
	for _, d := range diffs {
		if util.Absolute(d) > 3 {
			return false
		}
		s := getSign(d)
		if sign != 99 && s != sign {
			return false
		}
		if sign == 99 {
			sign = s
		}
	}

	return true
}

func main() {
	lines := util.GetFileStrings("2024/Day2/input")
	safe, safep2 := 0, 0
Outer:
	for _, l := range lines {
		p := strings.Split(l, " ")

		diffs := getDiffs(p)
		if isSafe(diffs) {
			safe++
		}

		for i := range p {
			v := append([]string{}, p[:i]...)
			v = append(v, p[i+1:]...)
			diffs := getDiffs(v)
			if isSafe(diffs) {
				safep2++
				continue Outer
			}
		}
	}
	fmt.Printf("Number safe (part1): %d\n", safe)
	fmt.Printf("Number safe with dampener (part2): %d\n", safep2)
}
