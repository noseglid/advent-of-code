package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

var memo = map[string]int{}

func can(pattern string, towels []string) int {
	if v, ok := memo[pattern]; ok {
		return v
	}
	if len(pattern) == 0 {
		return 1
	}

	n := 0
	for _, t := range towels {
		l := len(t)
		if len(pattern) >= l && pattern[:l] == t {
			c := can(pattern[l:], towels)
			n += c
		}
	}

	memo[pattern] = n
	return n
}

func main() {

	lines := util.GetFileStrings("2024/Day19/input")

	towels := strings.Split(lines[0], ", ")

	possible, combos := 0, 0
	for _, p := range lines[2:] {
		if c := can(p, towels); c > 0 {
			possible++
			combos += c
		}
	}

	fmt.Printf("available patterns (part1): %d\n", possible)
	fmt.Printf("total ways to combine (part2): %d\n", combos)

}
