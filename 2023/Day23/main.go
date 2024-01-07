package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type P struct {
	row, col int
}

func (p P) String() string {
	return fmt.Sprintf("(%d,%d)", p.row, p.col)
}

func edge(grid [][]rune, row int) P {
	for col, r := range grid[row] {
		if r == '.' {
			return P{row: row, col: col}
		}
	}
	panic("no edge")
}

func canMove(grid [][]rune, from, to P) bool {
	var (
		right = from.col+1 == to.col
		left  = from.col == to.col+1
		up    = from.row == to.row+1
		down  = from.row+1 == to.row
		tr    = grid[to.row][to.col]
	)
	return tr == '.' || (right && tr == '>') || (left && tr == '<') || (down && tr == 'v') || (up && tr == '^')
}

func canMoveJunctions(grid [][]rune, junctions []P, target P) func(grid [][]rune, from, to P) bool {
	return func(grid [][]rune, from, to P) bool {
		if slices.Contains(junctions, to) {
			return to == target
		}

		return grid[to.row][to.col] != '#'
	}
}

func canMovep2(grid [][]rune, from, to P) bool {
	return grid[to.row][to.col] != '#'
}

type CanMoveFn func(grud [][]rune, from, to P) bool

func alternatives(grid [][]rune, pos P, fn CanMoveFn) []P {
	var r []P
	if t := (P{row: pos.row, col: pos.col + 1}); t.col < len(grid[t.row]) && fn(grid, pos, t) {
		r = append(r, t)
	}
	if t := (P{row: pos.row, col: pos.col - 1}); t.col > 0 && fn(grid, pos, t) {
		r = append(r, t)
	}
	if t := (P{row: pos.row + 1, col: pos.col}); t.row < len(grid) && fn(grid, pos, t) {
		r = append(r, t)
	}
	if t := (P{row: pos.row - 1, col: pos.col}); t.row > 0 && fn(grid, pos, t) {
		r = append(r, t)
	}
	return r
}

func cloneVisited(visited map[P]bool) map[P]bool {
	r := make(map[P]bool)
	for k, v := range visited {
		r[k] = v
	}
	return r
}

func hike(grid [][]rune, fn CanMoveFn, visited map[P]bool, pos P, end P) (int, bool) {
	if pos == end {
		return 0, true
	}

	visited[pos] = true
	m := 0
	for _, a := range alternatives(grid, pos, fn) {
		if visited[a] {
			continue
		}

		if t, didFinish := hike(grid, fn, cloneVisited(visited), a, end); didFinish && (t+1) > m {
			m = t + 1
		}
	}
	return m, m > 0
}

func print(grid [][]rune, p P) {
	for row := range grid {
		for col := range grid[row] {
			if p.row == row && p.col == col {
				fmt.Printf("O")
			} else {
				fmt.Printf("%c", grid[row][col])
			}
		}
		fmt.Println()
	}
}

func junctions(grid [][]rune) []P {
	var j []P
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '#' {
				continue
			}
			p := P{row: row, col: col}
			a := alternatives(grid, p, canMovep2)
			if len(a) >= 3 {
				j = append(j, p)
			}
		}
	}

	return j
}

type distSpec struct {
	j1, j2 P
}

func longest(dists map[distSpec]int, connections map[P][]P, visited map[P]bool, from, to P) (int, bool) {
	if from == to {
		return 0, true
	}
	visited[from] = true

	m := 0
	for _, c := range connections[from] {
		if _, ok := visited[c]; ok {
			continue
		}
		td := dists[distSpec{j1: from, j2: c}]
		if d, ok := longest(dists, connections, cloneVisited(visited), c, to); ok {
			if d+td > m {
				m = d + td
			}
		}
	}

	return m, m > 0
}

func connections(grid [][]rune, js []P, s, e P) (map[distSpec]int, map[P][]P) {
	dists := map[distSpec]int{}
	connections := map[P][]P{}

	list := append([]P{s, e}, js...)
	for _, c := range list {
		for _, j := range list {
			if c == j {
				continue
			}
			dist, didFinish := hike(grid, canMoveJunctions(grid, js, j), map[P]bool{}, c, j)
			if didFinish {
				connections[c] = append(connections[c], j)
				dists[distSpec{c, j}] = dist
			}
		}
	}

	return dists, connections
}

func main() {
	grid := util.GetFileRuneGrid("2023/Day23/input")
	s, e := edge(grid, 0), edge(grid, len(grid)-1)

	js := append(junctions(grid), e)

	s1, didFinish1 := hike(grid, canMove, map[P]bool{}, s, e)
	fmt.Printf("Hiking from %s to %s (finish=%t) using scenic route (part1): %d\n", s, e, didFinish1, s1)

	dists, connections := connections(grid, js, s, e)
	d, _ := longest(dists, connections, map[P]bool{}, s, e)
	fmt.Printf("Scenic route with slopes ignored (part2): %d\n", d)
}
