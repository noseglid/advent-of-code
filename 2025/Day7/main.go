package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point

func findStart(grid [][]rune) P {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				return P{X: x, Y: y}
			}
		}
	}
	panic("no start")
}

func distinct(ps []P) []P {

	set := map[P]struct{}{}
	for _, pp := range ps {
		set[pp] = struct{}{}
	}

	var next []P
	for k := range set {
		next = append(next, k)
	}

	return next
}

var memo = map[P]int{}

func moveTachyon(grid [][]rune, t P) int {
	if v, ok := memo[t]; ok {
		return v
	}
	if t.Y+1 >= len(grid) {
		return 1
	}

	if grid[t.Y+1][t.X] == '^' {
		s := moveTachyon(grid, P{X: t.X - 1, Y: t.Y + 1}) + moveTachyon(grid, P{X: t.X + 1, Y: t.Y + 1})
		memo[t] = s
		return s
	}

	s := moveTachyon(grid, P{X: t.X, Y: t.Y + 1})
	memo[t] = s
	return s
}

func main() {

	grid := util.GetFileRuneGrid("2025/Day7/input")

	start := findStart(grid)

	beams := []P{start}

	n := 0
Outer:
	for {
		var next []P
		for _, b := range beams {
			if b.Y+1 >= len(grid) {
				break Outer
			}

			if grid[b.Y+1][b.X] == '^' {
				n++
				next = append(next, []P{{X: b.X - 1, Y: b.Y + 1}, {X: b.X + 1, Y: b.Y + 1}}...)
			} else {
				next = append(next, P{X: b.X, Y: b.Y + 1})
			}

			next = distinct(next)
		}
		beams = next
	}
	fmt.Printf("Number of splits (part1): %d\n", n)

	fmt.Printf("Number of timelines (part2): %d\n", moveTachyon(grid, start))

}
