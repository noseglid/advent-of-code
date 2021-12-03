package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func part1() {
	s := util.FileScanner("2021/Day2/input", bufio.ScanLines)

	pos := 0
	depth := 0
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		n := util.MustAtoi(p[1])

		switch p[0] {
		case "forward":
			pos += n
		case "down":
			depth += n
		case "up":
			depth -= n
		}
	}

	log.Printf("Part 1: multiplied pos*depth=%d", pos*depth)
}

func part2() {
	s := util.FileScanner("2021/Day2/input", bufio.ScanLines)

	pos := 0
	depth := 0
	aim := 0
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		n := util.MustAtoi(p[1])

		switch p[0] {
		case "forward":
			pos += n
			depth += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}

	log.Printf("Part 2: multiplied pos*depth=%d", pos*depth)
}

func main() {
	part1()
	part2()
}
