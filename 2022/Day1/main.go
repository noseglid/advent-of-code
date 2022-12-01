package main

import (
	"fmt"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

func main() {
	input := util.GetFileStrings("2022/Day1/input")

	m := []int{}
	s := 0
	for i, l := range input {
		if l != "" {
			s += util.MustAtoi(l)
		}
		if l == "" || i == len(input)-1 {
			m = append(m, s)
			s = 0
			continue
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(m)))

	fmt.Printf("max calories (part1): %d\n", m[0])
	fmt.Printf("top 3 max calories (part2): %d\n", m[0]+m[1]+m[2])
}
