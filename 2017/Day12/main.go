package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func parseLine(l string) (int, []int) {
	id := util.MustAtoi(strings.TrimSpace(l[0:strings.Index(l, "<")]))

	var pp []int
	for _, p := range strings.Split(l[strings.LastIndex(l, ">")+1:], ", ") {
		pp = append(pp, util.MustAtoi(strings.TrimSpace(p)))
	}
	return id, pp

}

func partOfAnyGroup(groups [][]int, p int) bool {
	for _, g := range groups {
		for _, c := range g {
			if c == p {
				return true
			}
		}
	}
	return false
}

func main() {

	lines := util.GetFileStrings("2017/Day12/input")

	programs := map[int][]int{}

	for _, l := range lines {
		id, pipes := parseLine(l)
		programs[id] = pipes
	}

	var groups [][]int

	for p := range programs {
		if partOfAnyGroup(groups, p) {
			continue
		}
		included := []int{}
		toCheck := []int{p}
		for len(toCheck) != 0 {

			p := toCheck[0]
			toCheck = toCheck[1:]

			included = append(included, p)

			for _, c := range programs[p] {
				if !util.Contains(included, c) && !util.Contains(toCheck, c) {
					toCheck = append(toCheck, c)
				}
			}
		}
		groups = append(groups, included)
	}

	fmt.Printf("programs reachable (part1): %d\n", len(groups[0]))
	fmt.Printf("number of groups (part2): %d\n", len(groups))
}
