package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point

func check(p P, w, h int) []P {
	var ps []P
	for i := p.Y - 1; i <= p.Y+1; i++ {
		for j := p.X - 1; j <= p.X+1; j++ {
			ps = append(ps, P{X: j, Y: i})
		}
	}

	var filtered []P
	for _, c := range ps {
		if c.X < 0 || c.Y < 0 || c.X >= w || c.Y >= h || (p == c) {
			continue
		}
		filtered = append(filtered, c)
	}
	return filtered
}

func dup(r [][]rune) [][]rune {
	duplicate := make([][]rune, len(r))
	for i := range r {
		duplicate[i] = make([]rune, len(r[i]))
		copy(duplicate[i], r[i])
	}
	return duplicate
}

func removeRolls(grid [][]rune) (int, [][]rune) {
	n := 0
	w, h := len(grid), len(grid[0])

	var dup = dup(grid)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '@' {
				continue
			}

			adj := 0
			for _, c := range check(P{x, y}, w, h) {
				if grid[c.Y][c.X] == '@' {
					adj++
				}
			}

			if adj < 4 {
				n++
				dup[y][x] = '.'
			}
		}
	}
	return n, dup

}

func main() {
	grid := util.GetFileRuneGrid("2025/Day4/input")

	p1, _ := removeRolls(grid)
	fmt.Printf("Rolls which can be moved (part1): %d\n", p1)

	n := 0

	for {
		count, next := removeRolls(grid)
		n += count
		if count == 0 {
			break
		}
		grid = next
	}

	fmt.Printf("total removed by iterating (part2): %d\n", n)
}
