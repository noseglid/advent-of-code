package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func parseTriange(s string) []int {
	var r []int
	for _, side := range strings.Fields(s) {
		r = append(r, util.MustAtoi(side))
	}

	return r
}

func isPossible(t []int) bool {
	return t[0]+t[1] > t[2] && t[0]+t[2] > t[1] && t[1]+t[2] > t[0]
}

func main() {
	s := util.FileScanner("2016/Day3/input", bufio.ScanLines)

	rows := [][]int{}
	possible := 0
	possiblep2 := 0
	for s.Scan() {
		t := parseTriange(s.Text())
		if isPossible(t) {
			possible++
		}
		rows = append(rows, t)

		if len(rows) == 3 {
			if isPossible([]int{rows[0][0], rows[1][0], rows[2][0]}) {
				possiblep2++
			}
			if isPossible([]int{rows[0][1], rows[1][1], rows[2][1]}) {
				possiblep2++
			}
			if isPossible([]int{rows[0][2], rows[1][2], rows[2][2]}) {
				possiblep2++
			}

			rows = [][]int{}
		}
	}
	log.Printf("possible (part1): %d", possible)
	log.Printf("possible (part2): %d", possiblep2)

}
