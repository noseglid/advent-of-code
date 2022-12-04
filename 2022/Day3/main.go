package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func prio(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}

	panic("bad rune")
}

func part1(lines []string) {
	var chars []rune
loop:
	for _, l := range lines {
		n := len(l) / 2
		p1, p2 := l[0:n], l[n:]

		for _, c := range p1 {
			for _, d := range p2 {
				if c == d {
					chars = append(chars, c)
					continue loop
				}
			}
		}
	}

	s := 0
	for _, c := range chars {
		s += prio(c)
	}

	fmt.Printf("Priority sum (part1): %d\n", s)
}

func part2(lines []string) {
	if len(lines)%3 != 0 {
		panic("not divisible by 3")
	}
	var chars []rune

	n := 0
loop:
	for {
		if n == len(lines) {
			break
		}
		groups := lines[n : n+3]
		for _, a := range groups[0] {
			for _, b := range groups[1] {
				for _, c := range groups[2] {
					if a == b && b == c {
						chars = append(chars, a)
						n += 3
						continue loop
					}
				}
			}
		}
	}

	s := 0
	for _, c := range chars {
		s += prio(c)
	}

	fmt.Printf("Priority sum (part2): %d\n", s)
}

func main() {
	lines := util.GetFileStrings("2022/Day3/input")
	part1(lines)
	part2(lines)
}
