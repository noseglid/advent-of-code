package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type Triple[T any] struct {
	A, B, C T
}

type P = util.Point

type Portal struct {
	id         string
	entrance   P
	levelDelta int
}

func findHorizontalPortals(grid [][]rune, y, portalOffset, levelDelta int) []Portal {
	portals := []Portal{}
	for x := 0; x < len(grid[y]); x++ {
		if grid[y][x] < 'A' || grid[y][x] > 'Z' {
			continue
		}

		id := fmt.Sprintf("%c%c", grid[y][x], grid[y+1][x])
		portals = append(portals, Portal{id, P{x, y + portalOffset}, levelDelta})
	}
	return portals
}

func findVerticalPortals(grid [][]rune, x, portalOffset, levelDelta int) []Portal {
	var portals []Portal
	for y := 0; y < len(grid); y++ {
		if grid[y][x] < 'A' || grid[y][x] > 'Z' {
			continue
		}

		id := fmt.Sprintf("%c%c", grid[y][x], grid[y][x+1])
		portals = append(portals, Portal{id, P{x + portalOffset, y}, levelDelta})
	}

	return portals
}

func sortNodes[T comparable](nodes map[T]int) func(a, b T) int {
	return func(a, b T) int {
		va, aok := nodes[a]
		vb, bok := nodes[b]
		if aok && bok {
			return va - vb
		} else if aok {
			return va
		} else {
			return vb
		}
	}
}

func canPortal(portals []Portal, p P) (Portal, Portal, bool) {
	i1 := slices.IndexFunc(portals, func(portal Portal) bool { return portal.entrance == p })
	if i1 == -1 {
		return Portal{}, Portal{}, false
	}
	i2 := slices.IndexFunc(portals, func(portal Portal) bool { return portal != portals[i1] && portal.id == portals[i1].id })
	if i2 == -1 {
		return Portal{}, Portal{}, false
	}
	return portals[i1], portals[i2], true
}

func dijk(grid [][]rune, start, end P, portals []Portal) int {

	nodes := map[P]int{
		start: 0,
	}
	unvisited := []P{start}

	for len(unvisited) > 0 {

		slices.SortFunc(unvisited, sortNodes(nodes))

		n := unvisited[0]
		checks := []P{{n.X, n.Y - 1}, {n.X + 1, n.Y}, {n.X, n.Y + 1}, {n.X - 1, n.Y}}
		for _, c := range checks {
			if grid[c.Y][c.X] != '.' {
				continue
			}
			if v, ok := nodes[c]; !ok {
				nodes[c] = nodes[n] + 1
				unvisited = append(unvisited, c)
			} else if nodes[n]+1 < v {
				nodes[c] = nodes[n] + 1
			}
		}

		if _, to, ok := canPortal(portals, n); ok {
			if v, ok := nodes[to.entrance]; !ok {
				nodes[to.entrance] = nodes[n] + 1
				unvisited = append(unvisited, to.entrance)
			} else if nodes[n]+1 < v {
				nodes[to.entrance] = nodes[n] + 1
			}
		}

		unvisited = unvisited[1:]
	}

	return nodes[end]
}

type Node struct {
	position P
	depth    int
}

func dijkRecursive(grid [][]rune, start, end Node, portals []Portal) int {

	nodes := map[Node]int{
		start: 0,
	}

	unvisited := []Node{start}

	for len(unvisited) > 0 && nodes[end] == 0 {
		slices.SortFunc(unvisited, sortNodes(nodes))

		n := unvisited[0]

		checks := []Node{
			{position: P{n.position.X, n.position.Y - 1}, depth: n.depth},
			{position: P{n.position.X + 1, n.position.Y}, depth: n.depth},
			{position: P{n.position.X, n.position.Y + 1}, depth: n.depth},
			{position: P{n.position.X - 1, n.position.Y}, depth: n.depth},
		}
		for _, c := range checks {
			if grid[c.position.Y][c.position.X] != '.' {
				continue
			}
			if v, ok := nodes[c]; !ok {
				nodes[c] = nodes[n] + 1
				unvisited = append(unvisited, c)
			} else if nodes[n]+1 < v {
				nodes[c] = nodes[n] + 1
			}
		}

		if from, to, ok := canPortal(portals, n.position); ok && n.depth+from.levelDelta >= 0 {
			destinationNode := Node{position: to.entrance, depth: n.depth + from.levelDelta}
			if v, ok := nodes[destinationNode]; !ok {
				nodes[destinationNode] = nodes[n] + 1
				unvisited = append(unvisited, destinationNode)
			} else if nodes[n]+1 < v {
				nodes[destinationNode] = nodes[n] + 1
			}
		}

		unvisited = unvisited[1:]
	}

	return nodes[end]
}

func main() {

	portals := []Portal{}

	//grid := util.GetFileRuneGrid("2019/Day20/sample2")
	//horizontals, verticals := []Triple[int]{
	//	{0, 2, -1}, {9, -1, 1}, {26, 2, 1}, {35, -1, -1},
	//}, []Triple[int]{
	//	{0, 2, -1}, {9, -1, 1}, {34, 2, 1}, {43, -1, -1},
	//}

	grid := util.GetFileRuneGrid("2019/Day20/input")
	horizontals, verticals := []Triple[int]{
		{0, 2, -1}, {35, -1, 1}, {84, 2, 1}, {119, -1, -1},
	}, []Triple[int]{
		{0, 2, -1}, {35, -1, 1}, {92, 2, 1}, {127, -1, -1},
	}

	for _, h := range horizontals {
		portals = append(portals, findHorizontalPortals(grid, h.A, h.B, h.C)...)
	}
	for _, v := range verticals {
		portals = append(portals, findVerticalPortals(grid, v.A, v.B, v.C)...)
	}

	s, e := P{}, P{}
	for _, p := range portals {
		if p.id == "AA" {
			s = p.entrance
		}
		if p.id == "ZZ" {
			e = p.entrance
		}
	}

	fmt.Printf("Minimum steps from AA to ZZ (part1): %d\n", dijk(grid, s, e, portals))
	fmt.Printf("Minimum steps from AA to ZZ with recursive (part2): %d\n", dijkRecursive(grid, Node{s, 0}, Node{e, 0}, portals))

}
