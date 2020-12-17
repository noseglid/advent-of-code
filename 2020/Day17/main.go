package main

import (
	"bufio"
	"fmt"
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

const size = 50

type Grid [size][size][size]int
type Grid4D [size][size][size][size]int

func parseRow(s string) []int {
	var result []int
	for _, r := range s {
		switch r {
		case '#':
			result = append(result, 1)
		case '.':
			result = append(result, 0)
		default:
			panic("bad state")
		}
	}

	return result
}

func printGrid(grid Grid) {
	minZ, maxZ, minY, maxY, minX, maxX := math.MaxInt64, 0, math.MaxInt64, 0, math.MaxInt64, 0
	each(grid, func(z, y, x, v int) {
		if v == 0 {
			return
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
	})
	for z := minZ; z <= maxZ; z++ {
		fmt.Printf("z=%d\n", z-size/2)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				switch grid[z][y][x] {
				case 0:
					fmt.Print(".")
				case 1:
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
	}
}
func each4D(grid Grid4D, fn func(w, z, y, x, v int)) {
	for w := range grid {
		for z := range grid[w] {
			for y := range grid[z] {
				for x := range grid[y] {
					fn(w, z, y, x, grid[w][z][y][x])
				}
			}
		}
	}
}

func each(grid Grid, fn func(z, y, x, v int)) {
	for z := range grid {
		for y := range grid[z] {
			for x := range grid[y] {
				fn(z, y, x, grid[z][y][x])
			}
		}
	}
}

func copyGrid4D(grid Grid4D) Grid4D {
	var ngrid Grid4D

	each4D(grid, func(w, z, y, x, v int) {
		ngrid[w][z][y][x] = v
	})

	return ngrid
}

func copyGrid(grid Grid) Grid {
	var ngrid Grid

	each(grid, func(z, y, x, v int) {
		ngrid[z][y][x] = v
	})

	return ngrid
}

var deltas = []struct{ z, y, x int }{}
var deltas4D = []struct{ w, z, y, x int }{}

func init() {
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					deltas4D = append(deltas4D, struct {
						w int
						z int
						y int
						x int
					}{w, z, y, x})
				}
			}
		}
	}
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if z == 0 && y == 0 && x == 0 {
					continue
				}
				deltas = append(deltas, struct {
					z int
					y int
					x int
				}{z, y, x})
			}
		}
	}
}

func iterate4D(grid Grid4D) Grid4D {
	result := copyGrid4D(grid)

	each4D(grid, func(w, z, y, x, v int) {
		active := 0
		for _, d := range deltas4D {
			if w+d.w < 0 || w+d.w >= len(grid) || z+d.z < 0 || y+d.y < 0 || x+d.x < 0 || z+d.z >= len(grid) || y+d.y >= len(grid) || x+d.x >= len(grid) {
				continue
			}

			if grid[w+d.w][z+d.z][y+d.y][x+d.x] == 1 {
				active++
			}
		}

		switch v {
		case 0:
			if active == 3 {
				result[w][z][y][x] = 1
			}
		case 1:
			if active != 2 && active != 3 {
				result[w][z][y][x] = 0
			}
		}
	})

	return result
}

func iterate(grid Grid) Grid {
	result := copyGrid(grid)

	each(grid, func(z, y, x, v int) {
		active := 0
		for _, d := range deltas {
			if z+d.z < 0 || y+d.y < 0 || x+d.x < 0 || z+d.z >= len(grid) || y+d.y >= len(grid) || x+d.x >= len(grid) {
				continue
			}

			if grid[z+d.z][y+d.y][x+d.x] == 1 {
				active++
			}
		}

		switch v {
		case 0:
			if active == 3 {
				result[z][y][x] = 1
			}
		case 1:
			if active != 2 && active != 3 {
				result[z][y][x] = 0
			}
		}
	})

	return result
}

func count4D(grid Grid4D, state int) int {
	n := 0
	each4D(grid, func(_, _, _, _, v int) {
		if v == state {
			n++
		}
	})

	return n
}

func count(grid Grid, state int) int {
	n := 0
	each(grid, func(_, _, _, v int) {
		if v == state {
			n++
		}
	})

	return n
}

func main() {

	s := util.FileScanner("2020/Day17/input", bufio.ScanLines)
	grid := Grid{}
	grid4d := Grid4D{}

	y := size / 2
	for s.Scan() {
		for i, e := range parseRow(s.Text()) {
			grid[size/2][y][size/2+i] = e
			grid4d[size/2][size/2][y][size/2+i] = e
		}
		y++
	}

	var cycles = 6

	for i := 0; i < cycles; i++ {
		grid = iterate(grid)
		grid4d = iterate4D(grid4d)
	}
	log.Printf("active after %d cycles in 3D (part1): %d", cycles, count(grid, 1))
	log.Printf("active after %d cycles in 4D (part2): %d", cycles, count4D(grid4d, 1))
}
