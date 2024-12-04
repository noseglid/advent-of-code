package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func xmasAt(grid [][]rune, x, y int) (int, int) {
	n, m := 0, 0

	canUp := y >= 3
	canRight := x < len(grid[y])-3
	canDown := y < len(grid)-3
	canLeft := x >= 3

	if grid[y][x] == 'X' {
		// horizontal: forwards
		if canRight && grid[y][x+1] == 'M' && grid[y][x+2] == 'A' && grid[y][x+3] == 'S' {
			n++
		}
		// horizontal: backwards
		if canLeft && grid[y][x-1] == 'M' && grid[y][x-2] == 'A' && grid[y][x-3] == 'S' {
			n++
		}
		// vertical: downwards
		if canDown && grid[y+1][x] == 'M' && grid[y+2][x] == 'A' && grid[y+3][x] == 'S' {
			n++
		}
		// vertical: upwards
		if canUp && grid[y-1][x] == 'M' && grid[y-2][x] == 'A' && grid[y-3][x] == 'S' {
			n++
		}
		// diagonal: NE
		if canUp && canRight && grid[y-1][x+1] == 'M' && grid[y-2][x+2] == 'A' && grid[y-3][x+3] == 'S' {
			n++
		}
		// diagonal: SE
		if canDown && canRight && grid[y+1][x+1] == 'M' && grid[y+2][x+2] == 'A' && grid[y+3][x+3] == 'S' {
			n++
		}
		// diagonal: SW
		if canDown && canLeft && grid[y+1][x-1] == 'M' && grid[y+2][x-2] == 'A' && grid[y+3][x-3] == 'S' {
			n++
		}
		// diagonal: NW
		if canUp && canLeft && grid[y-1][x-1] == 'M' && grid[y-2][x-2] == 'A' && grid[y-3][x-3] == 'S' {
			n++
		}
	}

	if grid[y][x] == 'A' && y >= 1 && x >= 1 && y < len(grid)-1 && x < len(grid[y])-1 {
		if grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S' && grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S' {
			m++
		}
		if grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M' && grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S' {
			m++
		}
		if grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M' && grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M' {
			m++
		}
		if grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S' && grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M' {
			m++
		}
	}

	return n, m
}

func main() {
	grid := util.GetFileRuneGrid("2024/Day4/input")
	n, m := 0, 0
	for y, row := range grid {
		for x := range row {
			nn, mm := xmasAt(grid, x, y)
			if mm > 0 {
				fmt.Printf("found at (%d,%d): %d\n", x, y, mm)
			}
			n += nn
			m += mm
		}
	}

	fmt.Printf("n xmas (part 1): %d\n", n)
	fmt.Printf("n x-mas (part 2): %d\n", m)
}
