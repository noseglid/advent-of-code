package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type p struct {
	x, y int
}

func findAntinodesp2(grid [][]rune, x, y, nx, ny int) []p {
	an := []p{{x, y}, {nx, ny}}
	dx, dy := nx-x, ny-y

	inMap := func(x, y int) bool {
		return x >= 0 && y >= 0 && y < len(grid) && x < len(grid[y])
	}

	c := true
	m := 0
	for c {
		m++
		c = false
		if inMap(x-dx*m, y-dy*m) {
			c = true
			an = append(an, p{x - dx*m, y - dy*m})
		}
		if inMap(nx+dx*m, ny+dy*m) {
			c = true
			an = append(an, p{nx + dx*m, ny + dy*m})
		}
	}

	return an
}

func findAntinodes(grid [][]rune, x, y, nx, ny int) []p {
	an := []p{}
	dx, dy := nx-x, ny-y
	inMap := func(x, y int) bool {
		return x >= 0 && y >= 0 && y < len(grid) && x < len(grid[y])
	}
	if inMap(x-dx, y-dy) {
		an = append(an, p{x - dx, y - dy})
	}
	if inMap(nx+dx, ny+dy) {
		an = append(an, p{nx + dx, ny + dy})
	}
	return an
}

func main() {
	grid := util.GetFileRuneGrid("2024/Day8/input")
	antinodes, antinodesp2 := map[p]bool{}, map[p]bool{}
	for y, row := range grid {
		for x, cell := range row {
			if cell == '.' {
				continue
			}

			for ny := 0; ny < len(grid); ny++ {
				for nx := 0; nx < len(grid[ny]); nx++ {
					if x == nx && y == ny || grid[ny][nx] != cell {
						continue
					}
					for _, an := range findAntinodes(grid, x, y, nx, ny) {
						antinodes[an] = true
					}
					for _, an := range findAntinodesp2(grid, x, y, nx, ny) {
						antinodesp2[an] = true
					}
				}
			}
		}
	}
	fmt.Printf("Total antinodes (part1): %d\n", len(antinodes))
	fmt.Printf("Total antinodes (part2): %d\n", len(antinodesp2))
}
