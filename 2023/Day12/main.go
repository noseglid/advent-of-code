package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

const verbose = false

var memo = map[string]int{}

func key(springs []rune, groups []int) string {
	return fmt.Sprintf("%s|%v", string(springs), groups)
}

func NConfigurations(springs []rune, groups []int) int {
	ckey := key(springs, groups)
	if v, ok := memo[ckey]; ok {
		return v
	}

	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	domemo := func(v int) int {
		memo[ckey] = v
		return v
	}

	switch springs[0] {
	case '.':
		return domemo(NConfigurations(springs[1:], groups))
	case '?':
		return NConfigurations(append([]rune{'.'}, springs[1:]...), groups) +
			NConfigurations(append([]rune{'#'}, springs[1:]...), groups)
	case '#':
		if len(groups) == 0 || len(springs) < groups[0] {
			// Not enough to build arrangement
			return domemo(0)
		}
		if slices.Contains(springs[0:groups[0]], '.') {
			// N first springs contains a '.' - not a good arrangement
			return domemo(0)
		}

		if len(groups) <= 1 {
			return domemo(NConfigurations(springs[groups[0]:], groups[1:]))
		}

		if len(springs) < groups[0]+1 || springs[groups[0]] == '#' {
			return domemo(0)
		}
		return domemo(NConfigurations(springs[groups[0]+1:], groups[1:]))
	}

	return 0
}

func parseLine(line string) ([]rune, []int) {
	spl := strings.Split(line, " ")

	var groups []int
	for _, n := range strings.Split(spl[1], ",") {
		nn, _ := strconv.Atoi(n)
		groups = append(groups, nn)
	}

	return []rune(spl[0]), groups
}

func unfold(springs []rune, groups []int) ([]rune, []int) {
	var newSprings []rune
	var newGroups []int
	newSprings = append(newSprings, springs...)
	newGroups = append(newGroups, groups...)
	for i := 0; i < 4; i++ {
		newSprings = append(newSprings, '?')
		newSprings = append(newSprings, springs...)
		newGroups = append(newGroups, groups...)
	}
	return newSprings, newGroups
}

func main() {
	lines := util.GetFileStrings("2023/Day12/input")

	s := 0
	s2 := 0
	for _, l := range lines {
		spring, groups := parseLine(l)
		unfoldedSprings, unfoldedGroups := unfold(spring, groups)
		s += NConfigurations(spring, groups)
		s2 += NConfigurations(unfoldedSprings, unfoldedGroups)
	}
	fmt.Printf("Total number of arrangements (part1): %d\n", s)
	fmt.Printf("Total number of unfolded arrangements (part2): %d\n", s2)
}
