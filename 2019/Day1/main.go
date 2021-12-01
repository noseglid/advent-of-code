package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func extraFuel(fuel int) int {
	s := 0
	for fuel > 8 {
		d := fuel/3 - 2
		s += d
		fuel = d
	}

	return s
}

func main() {
	input := util.GetFileNumbers("2019/Day1/input")

	s := 0
	for _, m := range input {
		s += m/3 - 2
	}
	log.Printf("fuel required (part1): %v", s)

	st := 0
	for _, m := range input {
		st += m/3 - 2 + extraFuel(m/3-2)
	}
	log.Printf("fuel with fueld (part2): %d", st)
}
