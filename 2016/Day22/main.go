package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type node struct {
	x, y       int
	size, used int
	visited    bool
	distance   int
}

func parseNode(s string) *node {
	f := strings.Fields(s)
	n := node{visited: false, distance: math.MaxInt}
	if _, err := fmt.Sscanf(f[0], "/dev/grid/node-x%d-y%d", &n.x, &n.y); err != nil {
		panic(err)
	}
	if _, err := fmt.Sscanf(f[1], "%dT", &n.size); err != nil {
		panic(err)
	}
	if _, err := fmt.Sscanf(f[2], "%dT", &n.used); err != nil {
		panic(err)
	}
	return &n
}

func neighbours(dn *node, grid [][]*node) []*node {
	var ret []*node
	if dn.x > 0 {
		ret = append(ret, grid[dn.y][dn.x-1])
	}
	if dn.y > 0 {
		ret = append(ret, grid[dn.y-1][dn.x])
	}
	if dn.x < len(grid[dn.y])-1 {
		ret = append(ret, grid[dn.y][dn.x+1])
	}
	if dn.y < len(grid)-1 {
		ret = append(ret, grid[dn.y+1][dn.x])
	}

	return ret
}

func shortestPath(grid [][]*node, start, target *node) int {
	queue := []*node{start}
	start.distance = 0

	for {
		var nextQueue []*node
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].distance < queue[j].distance
		})
		for _, n := range queue {
			if n.visited {
				continue
			}
			for _, nn := range neighbours(n, grid) {
				if nn.visited || nn.used > start.size {
					continue
				}
				nextQueue = append(nextQueue, nn)
				if nn.distance > n.distance+1 {
					nn.distance = n.distance + 1
				}
			}
			n.visited = true
		}
		if len(nextQueue) == 0 {
			break
		}
		queue = nextQueue
	}
	return grid[target.y][target.x].distance
}

func findMovableNode(grid [][]*node, nodes []*node) *node {
	start := grid[0][len(grid[0])-1]
	var movable *node
	for _, n := range nodes {
		if n.size-n.used > start.used {
			movable = n
			break
		}
	}

	return movable
}

func main() {

	input := "2016/Day22/input"
	var nodes []*node
	for _, l := range util.GetFileStrings(input) {
		nodes = append(nodes, parseNode(l))
	}

	n := 0
	maxX, maxY := 0, 0
	for _, A := range nodes {
		if A.x > maxX {
			maxX = A.x
		}
		if A.y > maxY {
			maxY = A.y
		}
		for _, B := range nodes {
			if A.used == 0 {
				continue
			}
			if A == B {
				continue
			}
			if A.used > B.size-B.used {
				continue
			}
			n++
		}
	}
	log.Printf("Part 1: Viable nodes: %d", n)

	grid := make([][]*node, maxY+1)
	for y := range grid {
		grid[y] = make([]*node, maxX+1)
	}
	for _, n := range nodes {
		grid[n.y][n.x] = n
	}

	movable := findMovableNode(grid, nodes)
	shortest := shortestPath(grid, movable, grid[0][len(grid[0])-2])

	log.Printf("Part 2: Minimum number of data copies: %d", shortest+1+(len(grid[0])-2)*5)

}
