package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type orderingRule struct {
	before, after int
}

func (o orderingRule) needSwap(p1, p2 int) bool {
	return p2 == o.before && p1 == o.after
}

func isCorrectUpdate(u []int, rules []orderingRule) bool {
	for _, r := range rules {
		if !util.Contains(u, r.before) || !util.Contains(u, r.after) {
			continue
		}
		ib, ia := slices.Index(u, r.before), slices.Index(u, r.after)
		if ib > ia {
			return false
		}
	}

	return true
}

func orderAndMiddle(u []int, rules []orderingRule) int {
	for {
		didSwap := false
		for i := 0; i < len(u)-1; i++ {
			n1, n2 := u[i], u[i+1]
			for _, r := range rules {
				if r.needSwap(n1, n2) {
					u[i], u[i+1] = u[i+1], u[i]
					didSwap = true

				}
			}
		}
		if !didSwap {
			break
		}
	}
	return u[len(u)/2]
}

func main() {
	lines := util.GetFileStrings("2024/Day5/input")

	rules := []orderingRule{}
	updates := [][]int{}

	for i, l := range lines {
		if l == "" {
			lines = lines[i+1:]
			break
		}

		var or orderingRule
		fmt.Sscanf(l, "%d|%d", &or.before, &or.after)
		rules = append(rules, or)
	}

	for _, l := range lines {
		u := []int{}
		for _, n := range strings.Split(l, ",") {
			u = append(u, util.MustAtoi(n))
		}
		updates = append(updates, u)
	}
	s, s2 := 0, 0
	for _, u := range updates {
		if isCorrectUpdate(u, rules) {
			s += u[len(u)/2]
		} else {
			c := make([]int, len(u))
			copy(c, u)
			s2 += orderAndMiddle(u, rules)
		}
	}

	fmt.Printf("sum of middle updates (part 1): %d\n", s)
	fmt.Printf("sum of middle updates of reordered (part 2): %d\n", s2)
}
