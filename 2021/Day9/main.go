package main

import (
	"log"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

func pointContains(list []point, p point) bool {
	for _, l := range list {
		if l == p {
			return true
		}
	}
	return false
}

type point struct {
	x, y int
}

func basinWalk(grid [][]int, p point, checked []point) (int, []point) {
	if pointContains(checked, p) ||
		p.y < 0 ||
		p.y >= len(grid) ||
		p.x < 0 ||
		p.x >= len(grid[p.y]) ||
		grid[p.y][p.x] == 9 {
		return 0, checked
	}

	c0 := append(checked, p)
	l, c1 := basinWalk(grid, point{p.x - 1, p.y}, c0)
	r, c2 := basinWalk(grid, point{p.x + 1, p.y}, c1)
	u, c3 := basinWalk(grid, point{p.x, p.y - 1}, c2)
	d, c4 := basinWalk(grid, point{p.x, p.y + 1}, c3)

	return 1 + l + r + u + d, c4
}

func isLowestPoint(grid [][]int, x, y int) (bool, int) {
	checks := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var cmp []int
	for _, c := range checks {
		yc := y + c.y
		xc := x + c.x

		if yc >= 0 && xc >= 0 && yc < len(grid) && xc < len(grid[yc]) {
			cmp = append(cmp, grid[yc][xc])
		}
	}

	p := grid[y][x]

	for _, c := range cmp {
		if p >= c {
			return false, 0
		}
	}

	return true, 1 + grid[y][x]
}

func main() {
	input := "2021/Day9/input"

	lines := util.GetFileStrings(input)

	grid := [][]int{}

	for _, l := range lines {
		var row []int
		for _, r := range l {
			row = append(row, util.MustAtoi(string(r)))
		}
		grid = append(grid, row)
	}

	s := 0
	bs := []int{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if ok, risk := isLowestPoint(grid, x, y); ok {
				s += risk
				size, _ := basinWalk(grid, point{x, y}, []point{})
				bs = append(bs, size)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(bs)))

	log.Printf("Part 1: Sum of risks: %d", s)
	log.Printf("Part 2: Product of top 3 basins: %d", bs[0]*bs[1]*bs[2])
}
