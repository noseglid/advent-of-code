package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type tile rune

var (
	SafeTile tile = '.'
	TrapTile tile = '^'
)

func countSafe(s string) int {
	n := 0
	for _, r := range s {
		if r == rune(SafeTile) {
			n++
		}
	}
	return n
}

func next(s string) string {
	row := make([]rune, len(s))
	for i := range s {
		left := SafeTile
		if i > 0 {
			left = tile(s[i-1])
		}
		center := tile(s[i])
		right := SafeTile
		if i < len(s)-1 {
			right = tile(s[i+1])
		}

		if left == TrapTile && center == TrapTile && right == SafeTile ||
			left == SafeTile && center == TrapTile && right == TrapTile ||
			left == TrapTile && center == SafeTile && right == SafeTile ||
			left == SafeTile && center == SafeTile && right == TrapTile {
			row[i] = rune(TrapTile)

		} else {
			row[i] = rune(SafeTile)
		}
	}

	return string(row)
}

func main() {
	input := "2016/Day18/input"
	row := util.GetFile(input)

	c := row[:len(row)-1]
	safeTiles := countSafe(c)
	for i := 0; i < 40-1; i++ {
		c = next(c)
		safeTiles += countSafe(c)
	}
	log.Printf("Part 1: Safe tiles: %d", safeTiles)

	c = row[:len(row)-1]
	safeTiles = countSafe(c)
	for i := 0; i < 400000-1; i++ {
		c = next(c)
		safeTiles += countSafe(c)
	}
	log.Printf("Part 2: Safe tiles: %d", safeTiles)
}
