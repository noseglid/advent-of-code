package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type dir string

var (
	NW dir = "nw"
	N  dir = "n"
	NE dir = "ne"
	SE dir = "se"
	S  dir = "s"
	SW dir = "sw"
)

func distance(r, q, s int) int {
	return (util.Absolute(q) + util.Absolute(r) + util.Absolute(s)) / 2
}

func parseInput(l string) []dir {
	var r []dir
	for _, e := range strings.Split(l, ",") {
		r = append(r, dir(e))
	}
	return r
}

func main() {
	input := util.GetFile("2017/Day11/input")

	sr, sq, ss := 0, 0, 0
	r, q, s := sr, sq, ss
	md := 0
	for _, d := range parseInput(input) {
		dd := distance(r, q, s)
		if dd > md {
			md = dd
		}
		switch d {
		case NW:
			q--
			s++
		case N:
			s++
			r--
		case NE:
			q++
			r--
		case SE:
			q++
			s--
		case S:
			s--
			r++
		case SW:
			q--
			r++
		}
	}

	fmt.Printf("Distance to child (part1): %d\n", distance(r, q, s)-1)
	fmt.Printf("Max distance child (part2): %d\n", md)
}
