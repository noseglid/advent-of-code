package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
	"golang.org/x/exp/constraints"
)

func nbefore(list []int, n int) int {
	c := 0
	for _, l := range list {
		if l <= n {
			c++
		}
	}
	return c
}

func countEmpty(grid [][]rune) ([]int, []int) {
	emptyRows, emptyCols := []int{}, []int{}
Row:
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '.' {
				continue Row
			}
		}
		emptyRows = append(emptyRows, row)
	}

Col:
	for col := range grid[0] {
		for row := range grid {
			if grid[row][col] != '.' {
				continue Col
			}
		}
		emptyCols = append(emptyCols, col)
	}

	return emptyRows, emptyCols
}

func galaxies(grid [][]rune) []util.Point {
	var r []util.Point
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '#' {
				r = append(r, util.Point{X: row, Y: col})
			}
		}
	}

	return r
}

func abs[T constraints.Integer](t T) T {
	if t < 0 {
		return -t
	}
	return t
}

func distanceSum(points []util.Point, emptyRows, emptyCols []int, mult int) int {
	d := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a, b := points[i], points[j]
			rowa := a.X + nbefore(emptyRows, a.X)*(max(1, mult-1))
			cola := a.Y + nbefore(emptyCols, a.Y)*(max(1, mult-1))
			rowb := b.X + nbefore(emptyRows, b.X)*(max(1, mult-1))
			colb := b.Y + nbefore(emptyCols, b.Y)*(max(1, mult-1))
			d += abs(rowa-rowb) + abs(cola-colb)
		}
	}

	return d
}

func main() {
	grid := util.GetFileRuneGrid("2023/Day11/input")
	gx := galaxies(grid)
	emptyRows, emptyCols := countEmpty(grid)
	fmt.Printf("Sum of distance to all galaxies 1 expanded(part1): %d\n", distanceSum(gx, emptyRows, emptyCols, 1))
	fmt.Printf("Sum of distance to all galaxies 1000000 expanded (part2): %d\n", distanceSum(gx, emptyRows, emptyCols, 1000000))
}
