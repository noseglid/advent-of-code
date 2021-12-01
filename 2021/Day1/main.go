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

	log.Printf("increasing depths: %d", n)
}

func sum(d [3]int) int {
	return d[0] + d[1] + d[2]
}

func part2(depths []int) {

	w := [3]int{}
	d := 0
	n := 0

	for i := 0; i < len(depths)-2; i++ {
		w = [3]int{
			depths[i],
			depths[i+1],
			depths[i+2],
		}
		ss := sum(w)
		if d != 0 && ss > d {
			n++
		}
		d = ss
		log.Printf("it %d, w=%v, d=%d, ss=%d, n=%d", i, w, d, ss, n)
	}
	log.Printf("part 2, increasing depths: %d", n)
}

func main() {

	depths := util.GetFileNumbers("2021/Day1/input")
	part1(depths)

	part2(depths)
}
