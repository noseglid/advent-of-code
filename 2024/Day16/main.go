package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

func canMove(r rune) bool {
	return r == '.' || r == 'E' || r == 'S'
}

type node struct {
	p util.Point
	d util.Dir
}

func (n node) String() string {
	return fmt.Sprintf("%d,%d %v", n.p.X, n.p.Y, n.d)
}

func isNeighbours(grid util.Grid, n1, n2 *node) bool {
	if n1.p == n2.p {
		return n1.d.Turn(util.Left) == n2.d || n1.d.Turn(util.Right) == n2.d
	} else if n1.d == n2.d {
		mx, my := grid.GetMove(n1.p.X, n1.p.Y, n1.d)
		return n2.p.X == mx && n2.p.Y == my
	}
	return false
}

var visitedTiles = map[util.Point]bool{}
var vvv = map[node]bool{}

func tiles(grid util.Grid, nodes map[*node]int, in *node) {
	visitedTiles[in.p] = true
	if vvv[*in] {
		return
	}
	vvv[*in] = true

	candidates := []*node{}
	for n, v := range nodes {
		if n.d == in.d && v+1 == nodes[in] && isNeighbours(grid, n, in) {
			candidates = append(candidates, n)
		}
		if n.p == in.p && n.d.Turn(util.Left) == in.d && v+1000 == nodes[in] && isNeighbours(grid, n, in) {
			candidates = append(candidates, n)
		}
		if n.p == in.p && n.d.Turn(util.Right) == in.d && v+1000 == nodes[in] && isNeighbours(grid, n, in) {
			candidates = append(candidates, n)
		}
	}

	for _, c := range candidates {
		tiles(grid, nodes, c)
	}
}

func dijk(grid util.Grid) {
	sx, sy := grid.Find('S')
	start := &node{util.Point{sx, sy}, util.E}
	nodes := map[*node]int{start: 0}
	refs := map[node]*node{*start: start}
	unvisited := []*node{start}

	for len(unvisited) > 0 {
		slices.SortFunc(unvisited, func(n1, n2 *node) int { return nodes[n1] - nodes[n2] })
		n := unvisited[0]

		mx, my := grid.GetMove(n.p.X, n.p.Y, n.d)
		checkNodes := []struct {
			n node
			c int
		}{
			{n: node{p: n.p, d: n.d.Turn(util.Left)}, c: 1000},
			{n: node{p: n.p, d: n.d.Turn(util.Right)}, c: 1000},
			{n: node{p: util.Point{mx, my}, d: n.d}, c: 1},
		}

		for _, checkNode := range checkNodes {
			if !canMove(grid.Get(checkNode.n.p.X, checkNode.n.p.Y)) {
				continue
			}
			if in, ok := refs[checkNode.n]; !ok {
				cpy := checkNode.n
				refs[cpy] = &cpy
				nodes[&cpy] = checkNode.c + nodes[n]
				unvisited = append(unvisited, &cpy)
			} else {
				if checkNode.c+nodes[n] < nodes[in] {
					nodes[in] = checkNode.c + nodes[n]
				}
			}
		}
		unvisited = unvisited[1:]
	}

	ex, ey := grid.Find('E')
	endPoint := util.Point{ex, ey}
	var end *node
	m := math.MaxInt
	for n, v := range nodes {
		if n.p == endPoint && v < m {
			m = v
			end = n
		}
	}

	tiles(grid, nodes, end)

	fmt.Printf("Minimum score (part1): %d\n", m)
	fmt.Printf("unique tiles (part2): %d\n", len(visitedTiles))
}

func main() {
	grid := util.Grid(util.GetFileRuneGrid("2024/Day16/input"))
	dijk(grid)
}
