package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func adjacent(grid [][]rune, x, y int) int {
	bugs := 0
	if y > 0 && grid[y-1][x] == '#' {
		bugs++
	}
	if y < len(grid)-1 && grid[y+1][x] == '#' {
		bugs++
	}
	if x > 0 && grid[y][x-1] == '#' {
		bugs++
	}

	if x < len(grid[y])-1 && grid[y][x+1] == '#' {
		bugs++
	}

	return bugs
}

func hash(grid [][]rune) string {
	var sb strings.Builder
	for y := range grid {
		for x := range grid[y] {
			sb.WriteRune(grid[y][x])
		}
	}
	return sb.String()
}

func biodiversity(grid [][]rune) int {
	m := 0
	bd := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				bd += 1 << m
			}
			m++
		}
	}
	return bd
}

func dup(grid [][]rune) [][]rune {
	r := make([][]rune, len(grid))
	for y := range grid {
		r[y] = make([]rune, len(grid[y]))
		copy(r[y], grid[y])
	}
	return r
}

func dupRecursive(levels map[int][][]rune) map[int][][]rune {
	r := map[int][][]rune{}
	for level, grid := range levels {
		r[level] = dup(grid)
	}
	return r
}

func generation(grid [][]rune) {
	d := dup(grid)
	for y := range d {
		for x := range d[y] {
			bugs := adjacent(d, x, y)
			switch d[y][x] {
			case '.':
				if bugs == 1 || bugs == 2 {
					grid[y][x] = '#'
				}
			case '#':
				if bugs != 1 {
					grid[y][x] = '.'
				}
			}
		}
	}
}

type ct struct{ l, x, y int }

func genCheckRow(level, y int) []ct {
	var r []ct
	for i := 0; i < 5; i++ {
		r = append(r, ct{level, i, y})
	}
	return r
}
func genCheckCol(level, x int) []ct {
	var r []ct
	for i := 0; i < 5; i++ {
		r = append(r, ct{level, x, i})
	}
	return r
}

func adjacentRecursive(levels map[int][][]rune, level, x, y int) int {
	var checks []ct

	switch y {
	case 0: // row 0
		switch x {
		case 0: // A
			checks = []ct{{level - 1, 2, 1}, {level - 1, 1, 2}, {level, x + 1, y}, {level, x, y + 1}}
		case 4: // E
			checks = []ct{{level - 1, 2, 1}, {level, x - 1, y}, {level - 1, 3, 2}, {level, x, y + 1}}
		case 1, 2, 3: // B, C, D
			checks = []ct{{level - 1, 2, 1}, {level, x - 1, y}, {level, x + 1, y}, {level, x, y + 1}}
		default:
			panic("bad pos")
		}
	case 1: // row 1
		switch x {
		case 0: // F
			checks = []ct{{level, x, y - 1}, {level - 1, 1, 2}, {level, x + 1, y}, {level, x, y + 1}}
		case 1, 3: // G, I
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level, x + 1, y}, {level, x, y + 1}}
		case 2: // H
			checks = append([]ct{{level, x, y - 1}, {level, x - 1, y}, {level, x + 1, y}}, genCheckRow(level+1, 0)...)
		case 4: // J
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level - 1, 3, 2}, {level, x, y + 1}}
		default:
			panic("bad pos")
		}
	case 2: // row 2
		switch x {
		case 0: // K
			checks = []ct{{level, x, y - 1}, {level - 1, 1, 2}, {level, x + 1, y}, {level, x, y + 1}}
		case 1: // L
			checks = append([]ct{{level, x, y - 1}, {level, x - 1, y}, {level, x, y + 1}}, genCheckCol(level+1, 0)...)
		case 2: // ?
			//checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level, x + 1, y}, {level, x, y + 1}}
		case 3: // N
			checks = append([]ct{{level, x, y - 1}, {level, x + 1, y}, {level, x, y + 1}}, genCheckCol(level+1, 4)...)
		case 4: // O
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level - 1, 3, 2}, {level, x, y + 1}}
		default:
			panic("bad pos")
		}
	case 3: // row 3
		switch x {
		case 0: // P
			checks = []ct{{level, x, y - 1}, {level - 1, 1, 2}, {level, x + 1, y}, {level, x, y + 1}}
		case 1, 3: // Q, S
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level, x + 1, y}, {level, x, y + 1}}
		case 2: // R
			checks = append([]ct{{level, x - 1, y}, {level, x + 1, y}, {level, x, y + 1}}, genCheckRow(level+1, 4)...)
		case 4: // T
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level - 1, 3, 2}, {level, x, y + 1}}
		default:
			panic("bad pos")
		}
	case 4: // row 4
		switch x {
		case 0: // U
			checks = []ct{{level, x, y - 1}, {level - 1, 1, 2}, {level, x + 1, y}, {level - 1, 2, 3}}
		case 4: // Y
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level - 1, 3, 2}, {level - 1, 2, 3}}
		case 1, 2, 3: // B, C, D
			checks = []ct{{level, x, y - 1}, {level, x - 1, y}, {level, x + 1, y}, {level - 1, 2, 3}}
		default:
			panic("bad pos")
		}
	default:
		panic("bad pos y")
	}

	bugs := 0
	for _, c := range checks {
		//fmt.Printf("checking %+v\n", c)
		if levels[c.l][c.y][c.x] == '#' {
			bugs++
		}
	}
	return bugs
}

func generationLevel(levels map[int][][]rune, level int) [][]rune {
	res := dup(levels[level])

	for y, row := range levels[level] {
		for x := range row {
			bugs := adjacentRecursive(levels, level, x, y)
			switch levels[level][y][x] {
			case '.':
				if bugs == 1 || bugs == 2 {
					res[y][x] = '#'
				}
			case '#':
				if bugs != 1 {
					res[y][x] = '.'
				}
			}
		}
	}

	return res

}

func generationRecursive(levels map[int][][]rune, maxLevels int) map[int][][]rune {
	next := dupRecursive(levels)
	for l := range levels {
		if util.Absolute(l) > maxLevels {
			continue
		}
		next[l] = generationLevel(levels, l)
	}
	return next
}

func emptyGrid() [][]rune {
	g := make([][]rune, 5)
	for y := range g {
		g[y] = make([]rune, 5)
		for x := range g[y] {
			g[y][x] = '.'
		}
	}
	return g
}

func printLevels(levels map[int][][]rune) {
	for l := -len(levels) / 2; l <= len(levels)/2; l++ {
		fmt.Printf("Level %d:\n", l)
		util.PrintRuneGrid(levels[l])
		fmt.Println()
	}
}

func countBugs(levels map[int][][]rune) int {
	bugs := 0
	for _, g := range levels {
		for _, row := range g {
			for _, cell := range row {
				if cell == '#' {
					bugs++
				}
			}
		}
	}
	return bugs
}

func main() {
	srcGrid := util.GetFileRuneGrid("2019/Day24/input")

	// ---------------- Part 1 --------------------
	p1grid := dup(srcGrid)
	seen := map[string]bool{
		hash(p1grid): true,
	}

	for {
		generation(p1grid)
		h := hash(p1grid)
		if seen[h] {
			break
		}
		seen[h] = true
	}
	fmt.Printf("Biodiversity of first repeated (part1): %d\n", biodiversity(p1grid))

	// ---------------- Part 2 --------------------
	levels := map[int][][]rune{}
	minutes := 200
	maxLevels := minutes/2 + 1

	for i := 0; i <= maxLevels+1; i++ {
		levels[i] = emptyGrid()
		levels[-i] = emptyGrid()
	}
	levels[0] = dup(srcGrid)

	for i := 0; i < minutes; i++ {
		levels = generationRecursive(levels, maxLevels)
	}

	fmt.Printf("Number of bugs in recursive world (part2): %d\n", countBugs(levels))
}
