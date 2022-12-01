package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func mult(s string) int {
	if s == "inc" {
		return 1
	} else if s == "dec" {
		return -1
	}
	panic("bad mult: " + s)
}

func check(op string, v1, v2 int) bool {
	switch op {
	case "<":
		return v1 < v2
	case "<=":
		return v1 <= v2
	case ">":
		return v1 > v2
	case ">=":
		return v1 >= v2
	case "==":
		return v1 == v2
	case "!=":
		return v1 != v2
	}
	panic("bad op: " + op)
}

func main() {
	input := "2017/Day8/input"
	lines := util.GetFileStrings(input)

	registers := map[string]int{}
	mall := math.MinInt
	for _, l := range lines {
		if l == "" {
			continue
		}
		parts := strings.Split(l, " ")
		rt, mult, val, rc, co, cv := parts[0], mult(parts[1]), util.MustAtoi(parts[2]), parts[4], parts[5], util.MustAtoi(parts[6])
		if !check(co, registers[rc], cv) {
			continue
		}

		registers[rt] += mult * val
		if registers[rt] > mall {
			mall = registers[rt]
		}
	}

	m := math.MinInt
	for _, v := range registers {
		if v > m {
			m = v
		}
	}
	fmt.Printf("Part 1: Largest value in registers at end is %d\n", m)
	fmt.Printf("Part 2: Largest value in registers all time is %d\n", mall)

}
