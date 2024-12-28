package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point
type tool rune

const (
	ClimbingGear tool = 'c'
	Torch        tool = 't'
	None         tool = 'n'
)

func (t tool) String() string {
	return string([]rune{rune(t)})
}

func supportsTool(r rune, t tool) bool {
	switch r {
	case '.':
		return t == ClimbingGear || t == Torch
	case '=':
		return t == ClimbingGear || t == None
	case '|':
		return t == Torch || t == None
	}
	panic("bad area")
}

type node struct {
	p P
	t tool
}

func (n node) String() string {
	return fmt.Sprintf("%d,%d:%s", n.p.X, n.p.Y, n.t)
}

func stay(grid [][]rune, n node) (node, bool) {
	switch grid[n.p.Y][n.p.X] {
	case '.':
		switch n.t {
		case ClimbingGear:
			return node{n.p, Torch}, true
		case Torch:
			return node{n.p, ClimbingGear}, true
		}
	case '=':
		switch n.t {
		case ClimbingGear:
			return node{n.p, None}, true
		case None:
			return node{n.p, ClimbingGear}, true
		}
	case '|':
		switch n.t {
		case Torch:
			return node{n.p, None}, true
		case None:
			return node{n.p, Torch}, true
		}
	}

	return node{}, false
}

func move(grid [][]rune, n node, dx, dy int) (node, bool) {
	nx, ny := n.p.X+dx, n.p.Y+dy
	if ny < 0 || nx < 0 || ny >= len(grid) || nx >= len(grid[ny]) {
		return node{}, false
	}
	if !supportsTool(grid[ny][nx], n.t) {
		return node{}, false
	}

	return node{P{nx, ny}, n.t}, true
}

func dijk(grid [][]rune, tx, ty int) int {
	unvisited := []node{{P{0, 0}, Torch}}
	nodes := map[node]int{
		unvisited[0]: 0,
	}

	for len(unvisited) > 0 {
		slices.SortFunc(unvisited, func(a, b node) int {
			if _, ok := nodes[a]; !ok {
				panic("not found")
			}
			if _, ok := nodes[b]; !ok {
				panic("not found")
			}
			return nodes[a] - nodes[b]
		})
		n := unvisited[0]

		if in, ok := stay(grid, n); ok {
			if vi, ok := nodes[in]; ok && vi > nodes[n]+7 {
				nodes[in] = nodes[n] + 7
			} else if !ok {
				nodes[in] = nodes[n] + 7
				unvisited = append(unvisited, in)
			}
		}
		for _, p := range []P{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			if in, ok := move(grid, n, p.X, p.Y); ok {
				if vi, ok := nodes[in]; ok && vi > nodes[n]+1 {
					nodes[in] = nodes[n] + 1
				} else if !ok {
					nodes[in] = nodes[n] + 1
					unvisited = append(unvisited, in)
				}
			}
		}

		unvisited = unvisited[1:]
	}

	return nodes[node{P{tx, ty}, Torch}]
}

func main() {

	lines := util.GetFileStrings("2018/Day22/input")
	var d, tx, ty int
	fmt.Sscanf(lines[0], "depth: %d", &d)
	fmt.Sscanf(lines[1], "target: %d,%d", &tx, &ty)
	gridWidth, gridHeight := tx+50, ty+50

	elvl := make([][]int, gridHeight)
	for y := range elvl {
		elvl[y] = make([]int, gridWidth)
	}
	elvl[0][0] = d % 20183
	elvl[ty][tx] = d % 20183
	for y := range elvl {
		elvl[y][0] = (y*48271 + d) % 20183
	}
	for x := range elvl[0] {
		elvl[0][x] = (x*16807 + d) % 20183
	}

	for y := range elvl {
		for x := range elvl[y] {
			if elvl[y][x] != 0 {
				continue
			}
			elvl[y][x] = (elvl[y-1][x]*elvl[y][x-1] + d) % 20183
		}
	}

	grid := make([][]rune, len(elvl))
	for y := range grid {
		grid[y] = make([]rune, len(elvl[y]))
		for x := range grid[y] {
			var r rune
			switch elvl[y][x] % 3 {
			case 0:
				r = '.'
			case 1:
				r = '='
			case 2:
				r = '|'
			}
			grid[y][x] = r
		}
	}

	risklevel := 0
	util.Grid(grid).Each(func(x, y int, r rune) {
		if x > tx || y > ty {
			return
		}
		switch r {
		case '=':
			risklevel += 1
		case '|':
			risklevel += 2
		}
	})

	fmt.Printf("risk level (part1): %d\n", risklevel)
	fmt.Printf("time to rescue (part2): %d\n", dijk(grid, tx, ty))
}
