package main

import (
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

func main() {

	input := "2016/Day6/input"
	lines := util.GetFileStrings(input)

	var res, res2 []rune
	for i := 0; i < len(lines[0]); i++ {
		occ := map[rune]int{}
		for _, l := range lines {
			occ[rune(l[i])]++
		}
		m := 0
		l := math.MaxInt
		ru := '0'
		lru := '0'
		for r, c := range occ {
			if c > m {
				ru = r
				m = c
			}
			if c < l {
				lru = r
				l = c
			}
		}
		res = append(res, ru)
		res2 = append(res2, lru)
	}

	log.Printf("Part 1: Most common: %s", string(res))
	log.Printf("Part 2: Least common: %s", string(res2))

}
