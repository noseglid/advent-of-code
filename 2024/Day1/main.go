package main

import (
	"fmt"
	"sort"

	"github.com/noseglid/advent-of-code/util"
	"golang.org/x/exp/constraints"
)

func abs[T constraints.Integer](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

func ntimes(n int, l []int) int {
	t := 0
	for _, v := range l {
		if n == v {
			t++
		}
	}
	return t
}

func main() {
	list := util.GetFileStrings("2024/Day1/input")

	var lhs, rhs []int

	for _, l := range list {
		var v1, v2 int
		fmt.Sscanf(l, "%d  %d", &v1, &v2)
		lhs = append(lhs, v1)
		rhs = append(rhs, v2)
	}

	sort.Ints(lhs)
	sort.Ints(rhs)

	s := 0
	for i := 0; i < len(lhs); i++ {
		s += abs(lhs[i], rhs[i])
	}

	fmt.Printf("Sum of diffs (part1): %d\n", s)

	s2 := 0
	for _, v := range lhs {
		t := ntimes(v, rhs)
		s2 += t * v
	}

	fmt.Printf("Similarity score (part2): %d\n", s2)

}
