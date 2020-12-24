package main

import (
	"bufio"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type Direction string

const (
	E  = Direction("e")
	SE = Direction("se")
	SW = Direction("sw")
	W  = Direction("w")
	NW = Direction("nw")
	NE = Direction("ne")
)

var allDirs = []Direction{E, SE, SW, W, NW, NE}

const size = 1000

func parseDirections(s string) []Direction {
	var dirs []Direction
	for i := 0; i < len(s); i++ {
		start := i
		end := i + 1
		if s[i] == 'n' || s[i] == 's' {
			end = i + 2
			i++
		}
		dirs = append(dirs, Direction(s[start:end]))
	}

	return dirs
}

func makeGrid() [][]int {
	g := make([][]int, size)
	for y := 0; y < size; y++ {
		g[y] = make([]int, size)
	}

	return g
}

func move(x, y int, d Direction) (int, int) {
	switch d {
	case E:
		x++
	case SE:
		if y%2 != 0 {
			//odd row
			x++
		}
		y++
	case SW:
		if y%2 == 0 {
			//even row
			x--
		}
		y++
	case W:
		x--
	case NW:
		if y%2 == 0 {
			//even row
			x--
		}
		y--
	case NE:
		if y%2 != 0 {
			//odd row
			x++
		}
		y--
	}

	return x, y
}

func turnTiles(dirs []Direction, rx, ry int, grid [][]int) {
	tx, ty := rx, ry
	for _, d := range dirs {
		tx, ty = move(tx, ty, d)
	}

	grid[ty][tx] = (grid[ty][tx] + 1) % 2
}

func countTiles(grid [][]int, v int) int {
	n := 0
	for _, row := range grid {
		for _, gv := range row {
			if gv == v {
				n++
			}
		}
	}

	return n
}

func copyGrid(grid [][]int) [][]int {
	ng := make([][]int, len(grid))
	for y, row := range grid {
		ng[y] = make([]int, len(row))
		for x, v := range row {
			ng[y][x] = v
		}
	}

	return ng
}

func adjacentBlack(x, y int, grid [][]int) int {
	n := 0
	for _, d := range allDirs {
		tx, ty := move(x, y, d)
		if tx < 0 || tx >= len(grid) || ty < 0 || ty >= len(grid[0]) {
			// Not black outside
			continue
		}

		if grid[ty][tx] == 1 {
			n++
		}
	}

	return n
}

func gameOfLife(grid [][]int) [][]int {
	ng := copyGrid(grid)
	for y := range grid {
		for x := range grid[y] {
			adj := adjacentBlack(x, y, grid)
			if grid[y][x] == 1 {
				// black
				if adj == 0 || adj > 2 {
					ng[y][x] = 0
				}
			} else {
				// white
				if adj == 2 {
					ng[y][x] = 1
				}
			}
		}
	}

	return ng
}

func main() {
	s := util.FileScanner("2020/Day24/input", bufio.ScanLines)

	rx, ry := size/2, size/2
	grid := makeGrid()

	for s.Scan() {
		dirs := parseDirections(s.Text())
		turnTiles(dirs, rx, ry, grid)
	}
	log.Printf("blacks in tile pattern (part1): %d", countTiles(grid, 1))

	for i := 0; i < 100; i++ {
		grid = gameOfLife(grid)
	}

	log.Printf("black tiles after game of life (part2): %d", countTiles(grid, 1))
}
