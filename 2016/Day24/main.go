package main

import (
	"bytes"
	"log"
	"math"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

type point struct {
	x, y int
}

type Map [][]byte

func (m Map) IsWall(x, y int) bool {
	return m[y][x] == '#'
}

func (m Map) GetStart() point {
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '0' {
				return point{x, y}
			}
		}
	}

	panic("no start position")
}

func (m Map) GetDestinations() []point {
	var p []point
	for y := range m {
		for x := range m[y] {
			c := m[y][x]
			if c == '.' || c == '#' || c == '0' {
				continue
			}
			p = append(p, point{x, y})
		}
	}
	return p
}

func (m Map) Moves(p point) []point {
	var pp []point
	if p.x > 0 && !m.IsWall(p.x-1, p.y) {
		pp = append(pp, point{p.x - 1, p.y})
	}
	if p.y > 0 && !m.IsWall(p.x, p.y-1) {
		pp = append(pp, point{p.x, p.y - 1})
	}
	if p.x < len(m[p.y]) && !m.IsWall(p.x+1, p.y) {
		pp = append(pp, point{p.x + 1, p.y})
	}
	if p.y < len(m) && !m.IsWall(p.x, p.y+1) {
		pp = append(pp, point{p.x, p.y + 1})
	}

	return pp
}

type dijkNode struct {
	p       point
	dist    int
	visited bool
}

type dijkNodeList []*dijkNode

func (d dijkNodeList) Len() int           { return len(d) }
func (d dijkNodeList) Less(i, j int) bool { return d[i].dist < d[j].dist }
func (d dijkNodeList) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func (m Map) ShortestPath(a, b point) int {
	nodes := make([][]*dijkNode, len(m))
	for y := range nodes {
		nodes[y] = make([]*dijkNode, len(m[y]))
		for x := range nodes[y] {
			nodes[y][x] = &dijkNode{point{x, y}, math.MaxInt, false}
		}
	}
	firstNode := nodes[a.y][a.x]
	firstNode.dist = 0

	queue := dijkNodeList{firstNode}
	for {
		var nextQueue dijkNodeList
		sort.Sort(queue)
		for _, n := range queue {
			if n.visited {
				continue
			}
			for _, m := range m.Moves(n.p) {
				mn := nodes[m.y][m.x]
				if mn.visited {
					continue
				}
				nextQueue = append(nextQueue, mn)
				if n.dist+1 < mn.dist {
					mn.dist = n.dist + 1
				}
			}
			n.visited = true
		}
		if len(nextQueue) == 0 {
			break
		}
		queue = nextQueue
	}

	return nodes[b.y][b.x].dist
}

func toInterface(pps []point) []interface{} {
	var dests []interface{}
	for _, d := range pps {
		dests = append(dests, d)
	}
	return dests
}

func fromInterface(ifs []interface{}) []point {
	var pps []point
	for _, i := range ifs {
		pps = append(pps, i.(point))
	}
	return pps
}

func stepsForOrder(start point, dests []point, m Map) int {
	pos := start
	steps := 0
	for _, d := range dests {
		var ss int
		if css, ok := cache[cacheKey{pos, d}]; ok {
			ss = css
		} else {
			ss = m.ShortestPath(pos, d)
			cache[cacheKey{pos, d}] = ss
			cache[cacheKey{d, pos}] = ss
		}
		pos = d
		steps += ss
	}
	return steps
}

type cacheKey struct {
	p1, p2 point
}

var cache map[cacheKey]int

func init() {
	cache = map[cacheKey]int{}
}

func main() {
	input := "2016/Day24/input"
	var m Map = bytes.Fields([]byte(util.GetFile(input)))
	minSteps := math.MaxInt
	util.Perm(toInterface(m.GetDestinations()), func(i []interface{}) {
		steps := stepsForOrder(m.GetStart(), fromInterface(i), m)
		if steps < minSteps {
			minSteps = steps
		}
	})
	log.Printf("Part 1: Shortest path to all destination: %d", minSteps)

	minSteps2 := math.MaxInt
	util.Perm(toInterface(m.GetDestinations()), func(i []interface{}) {
		order := fromInterface(i)
		steps := stepsForOrder(m.GetStart(), append(order, m.GetStart()), m)
		if steps < minSteps2 {
			minSteps2 = steps
		}
	})
	log.Printf("Part 2: Shortest path to all destination returning to 0: %d", minSteps2)
}
