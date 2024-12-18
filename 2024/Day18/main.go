package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type node struct {
	x, y int
}

func dijk(grid util.Grid) int {
	unvisited := []node{{0, 0}}
	nodes := map[node]int{}

	for len(unvisited) > 0 {
		n := unvisited[0]

		for _, m := range []util.Dir{util.N, util.E, util.S, util.W} {
			rx, ry := grid.GetMove(n.x, n.y, m)
			if grid.InBounds(rx, ry) && grid.Get(rx, ry) == '.' {
				in := node{rx, ry}
				if v, ok := nodes[in]; !ok {
					nodes[in] = nodes[n] + 1
					unvisited = append(unvisited, in)
				} else if nodes[n]+1 < v {
					nodes[in] = v
				}
			}
		}
		unvisited = unvisited[1:]
	}

	ex, ey := len(grid)-1, len(grid[len(grid)-1])-1
	return nodes[node{ex, ey}]
}

func main() {

	lines := util.GetFileStrings("2024/Day18/input")
	w, h, bytes := 71, 71, 1024

	grid := make([][]rune, w)
	for y := range grid {
		grid[y] = make([]rune, h)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for i, l := range lines {
		if i == bytes {
			break
		}
		var x, y int
		fmt.Sscanf(l, "%d,%d", &x, &y)
		grid[y][x] = '#'
	}

	steps := dijk(grid)
	fmt.Printf("Shortest path after %d bytes (part1): %d\n", bytes, steps)

	for i := bytes; i < len(lines); i++ {
		var x, y int
		fmt.Sscanf(lines[i], "%d,%d", &x, &y)
		grid[y][x] = '#'
		if n := dijk(grid); n == 0 {
			fmt.Printf("First byte to cut off path (part2): %d,%d\n", x, y)
			break
		}
	}
}
