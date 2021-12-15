package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func max(points []point) (int, int) {
	xmax, ymax := 0, 0
	for _, p := range points {
		if p.x > xmax {
			xmax = p.x
		}
		if p.y > ymax {
			ymax = p.y
		}
	}

	return xmax, ymax
}

func countDots(grid [][]bool) int {
	n := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] {
				n++
			}
		}
	}
	return n
}

func trimGrid(grid [][]bool) [][]bool {
	xmax, ymax := 0, 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] {
				if y > ymax {
					ymax = y
				}
				if x > xmax {
					xmax = x
				}
			}
		}
	}

	trimmed := make([][]bool, ymax+1)
	for y := range trimmed {
		trimmed[y] = make([]bool, xmax+1)
		for x := range trimmed[y] {
			trimmed[y][x] = grid[y][x]
		}
	}

	return trimmed
}

func prnt(grid [][]bool) {
	for y := range grid {
		for x := range grid[y] {
			v := "."
			if grid[y][x] {
				v = "#"
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func foldUp(grid [][]bool, foldy int) [][]bool {
	res := make([][]bool, len(grid))
	for y := range grid {
		if y == foldy {
			break
		}
		res[y] = make([]bool, len(grid[y]))
		for x := range grid[y] {
			res[y][x] = grid[y][x] || grid[2*foldy-y][x]
		}
	}

	return res
}

func foldLeft(grid [][]bool, foldx int) [][]bool {
	res := make([][]bool, len(grid))
	for y := range grid {
		res[y] = make([]bool, len(grid[y]))
		for x := range grid[y] {
			if x == foldx {
				break
			}
			res[y][x] = grid[y][x] || grid[y][2*foldx-x]
		}
	}
	return res
}

type point struct {
	x, y int
}

func newPoint(desc string) point {
	var p point
	if _, err := fmt.Sscanf(desc, "%d,%d", &p.x, &p.y); err != nil {
		panic(err)
	}
	return p
}

func main() {
	input := "2021/Day13/input"

	allLines := util.GetFileStrings(input)
	var folds []string
	points := []point{}

	for i, l := range allLines {
		if l == "" {
			folds = allLines[i+1:]
			break
		}
		points = append(points, newPoint(l))
	}

	mx, my := max(points)

	grid := make([][]bool, my+1)
	for y := range grid {
		grid[y] = make([]bool, mx+1)
	}

	for _, p := range points {
		grid[p.y][p.x] = true
	}

	first := true
	for _, f := range folds {
		var axis rune
		var val int
		fmt.Sscanf(f, "fold along %c=%d", &axis, &val)
		switch axis {
		case 'y':
			grid = foldUp(grid, val)
		case 'x':
			grid = foldLeft(grid, val)
		}

		if first {
			first = false
			log.Printf("Part 1: dots after first fold instruction: %d", countDots(grid))
		}
	}
	prnt(trimGrid(grid))
}
