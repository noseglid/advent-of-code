package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func isVisible(grid [][]int, r, c int) bool {
	height := grid[r][c]
	cr, cc := r, c
	for {
		cr = cr - 1
		if cr < 0 {
			return true
		}

		if grid[cr][cc] >= height {
			break
		}
	}

	cr, cc = r, c
	for {
		cr = cr + 1
		if cr > len(grid)-1 {
			return true
		}

		if grid[cr][cc] >= height {
			break
		}
	}

	cr, cc = r, c
	for {
		cc = cc - 1
		if cc < 0 {
			return true
		}

		if grid[cr][cc] >= height {
			break
		}
	}

	cr, cc = r, c
	for {
		cc = cc + 1
		if cc > len(grid[0])-1 {
			return true
		}

		if grid[cr][cc] >= height {
			break
		}
	}

	return false
}

func viewDistance(grid [][]int, r, c int, fn func(r, c int) (int, int)) int {
	cr, cc, h, s := r, c, grid[r][c], 0
	for {
		cr, cc = fn(cr, cc)
		if cr < 0 || cc < 0 || cr >= len(grid) || cc >= len(grid[0]) {
			break
		}
		s++
		if grid[cr][cc] >= h {
			break
		}
	}

	return s
}

func scenicScore(grid [][]int, r, c int) int {
	s := 1
	s *= viewDistance(grid, r, c, func(r, c int) (int, int) { return r - 1, c })
	s *= viewDistance(grid, r, c, func(r, c int) (int, int) { return r + 1, c })
	s *= viewDistance(grid, r, c, func(r, c int) (int, int) { return r, c - 1 })
	s *= viewDistance(grid, r, c, func(r, c int) (int, int) { return r, c + 1 })

	return s
}

func main() {

	grid := util.GetFileSingleDigitGrid("2022/Day8/input")

	visible := 0
	for ri := range grid {
		for ci := range grid[ri] {
			if isVisible(grid, ri, ci) {
				visible++
			}
		}
	}

	fmt.Printf("n visible (part1): %d\n", visible)

	m := 0
	for r := range grid {
		for c := range grid[r] {
			if s := scenicScore(grid, r, c); s > m {
				m = s
			}
		}
	}

	fmt.Printf("best scenic score (part2): %d\n", m)

}
