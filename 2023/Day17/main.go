package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type nodeEntry struct {
	p     util.Point
	dir   rune
	steps int
}

func dijkstra(grid [][]int, start util.Point, nodegen func(c nodeEntry) []nodeEntry) map[nodeEntry]int {
	distances := map[nodeEntry]int{}
	nodes := nodegen(nodeEntry{})
	for _, n := range nodes {
		distances[n] = 0
	}
	queued := map[nodeEntry]bool{}
	visited := map[nodeEntry]bool{}

	for len(nodes) > 0 {
		c := nodes[0]
		visited[c] = true
		nodes = nodes[1:]

		for _, n := range nodegen(c) {
			if c.dir == n.dir {
				n.steps = c.steps + 1
			} else {
				n.steps = 1
			}

			if n.p.X < 0 || n.p.X >= len(grid) || n.p.Y < 0 || n.p.Y >= len(grid[n.p.X]) {
				continue
			}

			dn, exists := distances[n]
			dc := distances[c]
			v := dc + grid[n.p.X][n.p.Y]
			if !exists || dn > v {
				distances[n] = v
				delete(visited, n)
				delete(queued, n)
			}

			if _, isQueued := queued[n]; !isQueued && !visited[n] {
				queued[n] = true
				nodes = append(nodes, n)
			}
		}
	}

	return distances
}

func minDist(grid [][]int, list map[nodeEntry]int) int {
	rows := len(grid) - 1
	cols := len(grid[rows]) - 1
	m := math.MaxInt
	for e, c := range list {
		if e.p.X == rows && e.p.Y == cols {
			if c < m {
				m = c
			}

		}
	}
	return m
}

func allNeighbours(c nodeEntry) []nodeEntry {
	return []nodeEntry{
		{p: util.Point{X: c.p.X, Y: c.p.Y + 1}, dir: 'e'},
		{p: util.Point{X: c.p.X + 1, Y: c.p.Y}, dir: 's'},
		{p: util.Point{X: c.p.X, Y: c.p.Y - 1}, dir: 'w'},
		{p: util.Point{X: c.p.X - 1, Y: c.p.Y}, dir: 'n'},
	}
}

func delByDir(list []nodeEntry, dir ...rune) []nodeEntry {
	var r []nodeEntry
	for _, e := range list {
		if slices.Contains(dir, e.dir) {
			continue
		}
		r = append(r, e)
	}
	return r
}
func keepByDir(list []nodeEntry, dir ...rune) []nodeEntry {
	var r []nodeEntry
	for _, e := range list {
		if !slices.Contains(dir, e.dir) {
			continue
		}
		r = append(r, e)
	}
	return r
}

func dirReverse(dir rune) rune {
	switch dir {
	case 'n':
		return 's'
	case 'e':
		return 'w'
	case 's':
		return 'n'
	case 'w':
		return 'e'
	}
	panic("bad dir")
}

func nodegenP1(c nodeEntry) []nodeEntry {
	if (nodeEntry{}) == c {
		return []nodeEntry{
			{p: util.Point{X: 0, Y: 0}, dir: 'e'},
		}
	}
	neighbours := allNeighbours(c)
	if c.steps == 3 {
		neighbours = delByDir(neighbours, c.dir)
	}

	return delByDir(neighbours, dirReverse(c.dir))
}

func nodegenP2(c nodeEntry) []nodeEntry {
	if (nodeEntry{}) == c {
		return []nodeEntry{
			{p: util.Point{X: 0, Y: 0}, dir: 'e'},
			{p: util.Point{X: 0, Y: 0}, dir: 's'},
		}
	}
	neighbours := allNeighbours(c)
	if c.steps < 4 {
		// only take in the same direction
		neighbours = keepByDir(neighbours, c.dir)
	} else if c.steps < 10 {
		// can go in any direction
	} else if c.steps == 10 {
		// must turn
		switch c.dir {
		case 'n':
			neighbours = delByDir(neighbours, 'n')
		case 'e':
			neighbours = delByDir(neighbours, 'e')
		case 's':
			neighbours = delByDir(neighbours, 's')
		case 'w':
			neighbours = delByDir(neighbours, 'w')
		}
	}

	// fmt.Printf("giving %d neigh\n", len(neighbours))
	return delByDir(neighbours, dirReverse(c.dir))
}

func main() {
	grid := util.GetFileSingleDigitGrid("2023/Day17/input")
	dists := dijkstra(grid, util.Point{X: 0, Y: 0}, nodegenP1)
	fmt.Printf("min (part1): %d\n", minDist(grid, dists))

	dists2 := dijkstra(grid, util.Point{X: 0, Y: 0}, nodegenP2)
	fmt.Printf("min (part2): %d\n", minDist(grid, dists2))
}
