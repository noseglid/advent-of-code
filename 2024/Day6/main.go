package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func findStart(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				return x, y
			}
		}
	}

	panic("no start")
}

type dir int

const (
	up dir = iota
	right
	down
	left
)

type p struct{ x, y int }

func isLooped(grid [][]rune) bool {
	x, y := findStart(grid)
	dir := up
	finished := false
Loop:
	for i := 0; i < len(grid)*len(grid[y]); i++ {
		switch dir {
		case up:
			if y == 0 {
				finished = true
				break Loop
			} else if grid[y-1][x] == '#' {
				dir = right
			} else {
				y--
			}
		case right:
			if x == len(grid[y])-1 {
				finished = true
				break Loop
			} else if grid[y][x+1] == '#' {
				dir = down
			} else {
				x++
			}
		case down:
			if y == len(grid)-1 {
				finished = true
				break Loop
			} else if grid[y+1][x] == '#' {
				dir = left
			} else {
				y++
			}
		case left:
			if x == 0 {
				finished = true
				break Loop
			} else if grid[y][x-1] == '#' {
				dir = up
			} else {
				x--
			}
		}
	}

	return !finished
}

func copyGrid(grid [][]rune) [][]rune {
	cpy := make([][]rune, len(grid))
	for y := range grid {
		cpy[y] = make([]rune, len(grid[y]))
		copy(cpy[y], grid[y])
	}
	return cpy
}

func main() {
	grid := util.GetFileRuneGrid("2024/Day6/input")

	x, y := findStart(grid)
	used := map[p]bool{}

	dir := up
Loop:
	for {
		used[p{x, y}] = true
		switch dir {
		case up:
			if y == 0 {
				break Loop
			} else if grid[y-1][x] == '#' {
				dir = right
			} else {
				y--
			}
		case right:
			if x == len(grid[y])-1 {
				break Loop
			} else if grid[y][x+1] == '#' {
				dir = down
			} else {
				x++
			}
		case down:
			if y == len(grid)-1 {
				break Loop
			} else if grid[y+1][x] == '#' {
				dir = left
			} else {
				y++
			}
		case left:
			if x == 0 {
				break Loop
			} else if grid[y][x-1] == '#' {
				dir = up
			} else {
				x--
			}
		}
	}

	n := 0
	for y := range grid {
		for x := range grid[y] {
			cpy := copyGrid(grid)

			if cpy[y][x] == '^' || cpy[y][x] == '#' {
				continue
			}

			cpy[y][x] = '#'
			if isLooped(cpy) {
				n++
			}
		}
	}

	fmt.Printf("steps (part1): %d\n", len(used))
	fmt.Printf("obstructions (part2): %d\n", n)

}
