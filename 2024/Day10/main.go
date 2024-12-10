package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type p struct {
	x, y int
}

func trailheads(grid [][]rune, x, y int, found map[p]bool) (int, []p) {
	if !found[p{x, y}] && grid[y][x] == '9' {
		found[p{x, y}] = true
		return 1, []p{{x, y}}
	} else if grid[y][x] == '9' {
		return 0, []p{{x, y}}
	}

	ps := []p{}
	n := 0
	if x > 0 && grid[y][x-1] == grid[y][x]+1 {
		v, pp := trailheads(grid, x-1, y, found)
		ps = append(ps, pp...)
		n += v
	}
	if x < len(grid[y])-1 && grid[y][x+1] == grid[y][x]+1 {
		v, pp := trailheads(grid, x+1, y, found)
		ps = append(ps, pp...)
		n += v
	}
	if y > 0 && grid[y-1][x] == grid[y][x]+1 {
		v, pp := trailheads(grid, x, y-1, found)
		ps = append(ps, pp...)
		n += v
	}
	if y < len(grid)-1 && grid[y+1][x] == grid[y][x]+1 {
		v, pp := trailheads(grid, x, y+1, found)
		ps = append(ps, pp...)
		n += v
	}

	return n, ps
}

func main() {

	grid := util.GetFileRuneGrid("2024/Day10/input")
	s := 0
	ps := []p{}
	for y, row := range grid {
		for x, cell := range row {
			if cell == '0' {
				v, pp := trailheads(grid, x, y, map[p]bool{})
				ps = append(ps, pp...)
				s += v
			}
		}
	}
	fmt.Printf("trailheads (part1): %d\n", s)
	fmt.Printf("distinct trails (part2): %d\n", len(ps))
}
