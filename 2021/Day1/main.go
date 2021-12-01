package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func part1(depths []int) {
	c := 0
	n := 0
	for _, d := range depths {
		if c != 0 && d > c {
			n++
		}

		c = d
	}

	log.Printf("part1: increasing depths: %d", n)
}

func part2(depths []int) {
	d := 0
	n := 0
	for i := 0; i < len(depths)-2; i++ {
		ss := depths[i] + depths[i+1] + depths[i+2]
		if d != 0 && ss > d {
			n++
		}
		d = ss
	}
	log.Printf("part 2, increasing depths: %d", n)
}

func main() {

	depths := util.GetFileNumbers("2021/Day1/input")
	part1(depths)

	part2(depths)
}
