package main

import (
	"bufio"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type MapEntry string

const (
	Open = MapEntry(".")
	Tree = MapEntry("#")
)

func parseMapRow(s string) []MapEntry {
	mapRow := make([]MapEntry, len(s))
	for i, r := range s {
		mapRow[i] = MapEntry(r)
	}

	return mapRow
}

func mapTraverse(m [][]MapEntry, xdiff, ydiff int) int {
	x, y, maxHeight, repeatWidth, ntrees := 0, 0, len(m), len(m[0]), 0
	for {
		x, y = (x+xdiff)%repeatWidth, y+ydiff
		if y >= maxHeight {
			break
		}
		if m[y][x] == Tree {
			ntrees++
		}
	}
	return ntrees
}

func main() {
	s := util.FileScanner("2020/Day3/input", bufio.ScanLines)

	m := [][]MapEntry{}
	for s.Scan() {
		m = append(m, parseMapRow(s.Text()))
	}

	part1trees := mapTraverse(m, 3, 1)

	log.Printf("Trees encountered (part1): %d", part1trees)

	part2_1 := mapTraverse(m, 1, 1)
	part2_2 := part1trees
	part2_3 := mapTraverse(m, 5, 1)
	part2_4 := mapTraverse(m, 7, 1)
	part2_5 := mapTraverse(m, 1, 2)

	log.Printf("Product of trees encountered (part2): %d", part2_1*part2_2*part2_3*part2_4*part2_5)

}
