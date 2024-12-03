package main

import (
	"fmt"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

func main() {
	instr := util.GetFile("2024/Day3/input")
	r := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\))`)

	s, s2 := 0, 0
	enabled := true
	for _, v := range r.FindAllString(instr, -1) {
		if v == "do()" {
			enabled = true
			continue
		}
		if v == "don't()" {
			enabled = false
			continue
		}
		var a, b int
		fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
		s += a * b
		if enabled {
			s2 += a * b
		}
	}

	fmt.Printf("Sum of all muls (part1): %d\n", s)
	fmt.Printf("Sum of all muls with features (part2): %d\n", s2)
}
