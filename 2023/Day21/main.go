package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type P struct {
	row, col int
}

type MEntry struct {
	row int
	col int
	s   int
}

func start(grid [][]rune) P {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				return P{row: row, col: col}
			}
		}
	}
	panic("no start")
}
func step(grid [][]rune, steps int, p P, memo map[MEntry]bool) {
	e := MEntry{row: p.row, col: p.col, s: steps}
	hasVisited := memo[e]
	memo[e] = true
	if steps == 0 || hasVisited {
		return
	}
	memo[e] = true

	if p.row > 0 && grid[p.row-1][p.col] == '.' {
		step(grid, steps-1, P{row: p.row - 1, col: p.col}, memo)
	}
	if p.row < len(grid)-1 && grid[p.row+1][p.col] == '.' {
		step(grid, steps-1, P{row: p.row + 1, col: p.col}, memo)
	}
	if p.col > 0 && grid[p.row][p.col-1] == '.' {
		step(grid, steps-1, P{row: p.row, col: p.col - 1}, memo)
	}
	if p.col < len(grid[p.row])-1 && grid[p.row][p.col+1] == '.' {
		step(grid, steps-1, P{row: p.row, col: p.col + 1}, memo)
	}
}

func main() {
	grid := util.GetFileRuneGrid("2023/Day21/input")
	s := start(grid)

	memo := map[MEntry]bool{}
	step(grid, 64, s, memo)

	n := 0
	c := map[P]bool{}
	for e := range memo {
		p := P{row: e.row, col: e.col}
		if (e.row+e.col)%2 == 0 && !c[p] {
			c[p] = true
			n++
		}
	}
	fmt.Printf("Total grid reached (part1): %d\n", n)

}
