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

func p2(grid [][]rune) int {

	maxSteps := 26501365
	nodes := []P{start(grid)}
	dirs := []P{
		{row: -1, col: 0}, {row: 0, col: -1}, {row: 0, col: 1}, {row: 1, col: 0},
	}

	polynomial := make([]int, 0)
	for steps := 0; steps < maxSteps; steps++ {
		nextNodes := []P{}
		visited := map[P]bool{}

		for len(nodes) > 0 {
			element := nodes[0]
			nodes = nodes[1:]
			for _, d := range dirs {
				newPos := P{row: element.row + d.row, col: element.col + d.col}

				tr := ((newPos.col % len(grid)) + len(grid)) % len(grid)
				tc := ((newPos.row % len(grid)) + len(grid)) % len(grid)
				if grid[tr][tc] != '#' && !visited[newPos] {
					visited[newPos] = true
					nextNodes = append(nextNodes, newPos)
				}
			}
		}

		nodes = nextNodes
		if (steps+1)%(len(grid)) == maxSteps%len(grid) {
			polynomial = append(polynomial, len(visited))

			if len(polynomial) == 3 {
				p0 := polynomial[0]
				p1 := polynomial[1] - polynomial[0]
				p2 := polynomial[2] - polynomial[1]

				return p0 + (p1 * (maxSteps / len(grid))) + ((maxSteps/len(grid))*((maxSteps/len(grid))-1)/2)*(p2-p1)
			}
		}
	}

	panic("invalid")
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
	fmt.Printf("Total reached in infinite grid (part2): %d\n", p2(grid))

}
