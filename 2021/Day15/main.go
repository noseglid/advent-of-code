package main

import (
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type point struct {
	x, y int
}

type grid struct {
	nodes     [][]*node
	unvisited []*node
}

type node struct {
	x, y      int
	value     int
	tentative int
	visited   bool
}

func NewNode(x, y, value int) *node {
	return &node{
		x:         x,
		y:         y,
		value:     value,
		tentative: math.MaxInt,
		visited:   false,
	}
}

func gridInsert(dst, src [][]*node, x, y int) {
	for yi := range src {
		for xi := range src[yi] {
			dst[y+yi][x+xi] = NewNode(x+xi, y+yi, src[yi][xi].value)
		}
	}
}

func riskIncrease(inodes [][]*node) [][]*node {
	gg := make([][]*node, len(inodes))
	for y := range inodes {
		gg[y] = make([]*node, len(inodes[y]))
		for x := range inodes[y] {
			gg[y][x] = NewNode(x, y, inodes[y][x].value%9+1)
		}
	}

	return gg
}

func (g grid) expand() *grid {
	nodes := make([][]*node, 5*len(g.nodes))
	for y := range nodes {
		nodes[y] = make([]*node, 5*len(g.nodes))
	}

	unit := len(g.nodes)

	gridInsert(nodes, g.nodes, 0, 0)
	g1 := riskIncrease(g.nodes)
	gridInsert(nodes, g1, 1*unit, 0*unit)
	gridInsert(nodes, g1, 0*unit, 1*unit)
	g2 := riskIncrease(g1)
	gridInsert(nodes, g2, 2*unit, 0*unit)
	gridInsert(nodes, g2, 1*unit, 1*unit)
	gridInsert(nodes, g2, 0*unit, 2*unit)
	g3 := riskIncrease(g2)
	gridInsert(nodes, g3, 3*unit, 0*unit)
	gridInsert(nodes, g3, 2*unit, 1*unit)
	gridInsert(nodes, g3, 1*unit, 2*unit)
	gridInsert(nodes, g3, 0*unit, 3*unit)
	g4 := riskIncrease(g3)
	gridInsert(nodes, g4, 4*unit, 0*unit)
	gridInsert(nodes, g4, 3*unit, 1*unit)
	gridInsert(nodes, g4, 2*unit, 2*unit)
	gridInsert(nodes, g4, 1*unit, 3*unit)
	gridInsert(nodes, g4, 0*unit, 4*unit)
	g5 := riskIncrease(g4)
	gridInsert(nodes, g5, 4*unit, 1*unit)
	gridInsert(nodes, g5, 3*unit, 2*unit)
	gridInsert(nodes, g5, 2*unit, 3*unit)
	gridInsert(nodes, g5, 1*unit, 4*unit)
	g6 := riskIncrease(g5)
	gridInsert(nodes, g6, 4*unit, 2*unit)
	gridInsert(nodes, g6, 3*unit, 3*unit)
	gridInsert(nodes, g6, 2*unit, 4*unit)
	g7 := riskIncrease(g6)
	gridInsert(nodes, g7, 4*unit, 3*unit)
	gridInsert(nodes, g7, 3*unit, 4*unit)
	g8 := riskIncrease(g7)
	gridInsert(nodes, g8, 4*unit, 4*unit)

	return NewGridNodes(nodes)
}

func (g *grid) removeVisited(n *node) {
	for i, nn := range g.unvisited {
		if nn == n {
			g.unvisited = append(g.unvisited[:i], g.unvisited[i+1:]...)
		}
	}
}

func (g grid) start() *node {
	return g.nodes[0][0]
}

func (g grid) end() *node {
	return g.nodes[len(g.nodes)-1][len(g.nodes)-1]
}

func (g grid) SmallestUnvisited() *node {
	s := g.unvisited[0]

	for _, n := range g.unvisited[1:] {
		if n.tentative < s.tentative {
			s = n
		}
	}

	return s
}

func (g grid) GetNeighbours(c *node) []*node {
	var nn []*node
	dd := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dd {
		xcheck, ycheck := c.x+d.x, c.y+d.y
		if xcheck < 0 || xcheck >= len(g.nodes[0]) || ycheck < 0 || ycheck >= len(g.nodes) || g.nodes[ycheck][xcheck].visited {
			continue
		}
		nn = append(nn, g.nodes[ycheck][xcheck])
	}
	return nn
}

func NewGridNodes(nodes [][]*node) *grid {
	g := grid{
		nodes: nodes,
	}
	for _, r := range g.nodes {
		g.unvisited = append(g.unvisited, r...)
	}

	return &g
}

func NewGrid(digitGrid [][]int) *grid {
	nodes := make([][]*node, len(digitGrid))
	for y := range digitGrid {
		nodes[y] = make([]*node, len(digitGrid[y]))
		for x := range digitGrid[y] {
			nodes[y][x] = NewNode(x, y, digitGrid[y][x])
		}
	}

	return NewGridNodes(nodes)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dijkstra(gridmap *grid) {
	current := gridmap.start()
	current.tentative = 0

	for {
		if current.visited {
			continue
		}

		nn := gridmap.GetNeighbours(current)

		for _, n := range nn {
			n.tentative = min(n.tentative, n.value+current.tentative)
		}

		current.visited = true
		gridmap.removeVisited(current)

		if len(gridmap.unvisited) == 0 {
			break
		}

		current = gridmap.SmallestUnvisited()
	}
}

func main() {
	input := "2021/Day15/input"

	gridmap := NewGrid(util.GetFileSingleDigitGrid(input))
	fullGridmap := gridmap.expand()
	dijkstra(gridmap)
	dijkstra(fullGridmap)
	log.Printf("Part 1: Shortest path has a risk of %d", gridmap.end().tentative)
	log.Printf("Part 2: Shortest path of full map has a risk of %d", fullGridmap.end().tentative)
}
