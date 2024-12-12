package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type p struct{ x, y int }

func perimiterLength(grid [][]rune, pos p, c rune, visited map[p]bool) int {
	if pos.y < 0 || pos.x < 0 || pos.y == len(grid) || pos.x == len(grid[pos.y]) || grid[pos.y][pos.x] != c {
		return 1
	}
	if visited[pos] {
		return 0
	}
	visited[pos] = true
	return perimiterLength(grid, p{pos.x - 1, pos.y}, c, visited) +
		perimiterLength(grid, p{pos.x + 1, pos.y}, c, visited) +
		perimiterLength(grid, p{pos.x, pos.y - 1}, c, visited) +
		perimiterLength(grid, p{pos.x, pos.y + 1}, c, visited)
}

func area(grid [][]rune, pos p, c rune, visited map[p]bool) int {
	if visited[pos] {
		return 0
	}
	if pos.y < 0 || pos.x < 0 || pos.y == len(grid) || pos.x == len(grid[pos.y]) || grid[pos.y][pos.x] != c {
		return 0
	}
	visited[pos] = true
	return 1 + area(grid, p{pos.x - 1, pos.y}, c, visited) +
		area(grid, p{pos.x + 1, pos.y}, c, visited) +
		area(grid, p{pos.x, pos.y - 1}, c, visited) +
		area(grid, p{pos.x, pos.y + 1}, c, visited)

}

func corners(grid [][]rune, pos p, c rune, visited map[p]bool) int {
	if visited[pos] || pos.y < 0 || pos.x < 0 || pos.y == len(grid) || pos.x == len(grid[pos.y]) || grid[pos.y][pos.x] != c {
		return 0
	}
	visited[pos] = true

	x, y := pos.x, pos.y
	n, e, s, w, ne, se, sw, nw := '.', '.', '.', '.', '.', '.', '.', '.'
	if y > 0 {
		n = grid[y-1][x]
	}
	if y < len(grid)-1 {
		s = grid[y+1][x]
	}
	if x > 0 {
		w = grid[y][x-1]
	}
	if x < len(grid[y])-1 {
		e = grid[y][x+1]
	}
	if y > 0 && x < len(grid[y])-1 {
		ne = grid[y-1][x+1]
	}
	if y < len(grid)-1 && x < len(grid[y])-1 {
		se = grid[y+1][x+1]
	}
	if x > 0 && y < len(grid)-1 {
		sw = grid[y+1][x-1]
	}
	if x > 0 && y > 0 {
		nw = grid[y-1][x-1]
	}
	cc := 0
	if n != c && e != c {
		cc++
	}
	if e != c && s != c {
		cc++
	}
	if s != c && w != c {
		cc++
	}
	if w != c && n != c {
		cc++
	}
	if ne != c && n == c && e == c {
		cc++
	}
	if se != c && s == c && e == c {
		cc++
	}
	if sw != c && s == c && w == c {
		cc++
	}
	if nw != c && n == c && w == c {
		cc++
	}
	return cc +
		corners(grid, p{pos.x - 1, y}, c, visited) +
		corners(grid, p{pos.x + 1, y}, c, visited) +
		corners(grid, p{pos.x, y - 1}, c, visited) +
		corners(grid, p{pos.x, y + 1}, c, visited)
}

func main() {
	grid := util.GetFileRuneGrid("2024/Day12/input")

	visited := map[p]bool{}
	sum, sump2 := 0, 0
	for y, row := range grid {
		for x, cell := range row {
			if visited[p{x, y}] {
				continue
			}
			a := area(grid, p{x, y}, cell, visited)
			prm := perimiterLength(grid, p{x, y}, cell, map[p]bool{})
			s := corners(grid, p{x, y}, cell, map[p]bool{})

			sum += a * prm
			sump2 += a * s
		}
	}

	fmt.Printf("Cost for fence (part1): %d\n", sum)
	fmt.Printf("Cost for fence w. discout (part2): %d\n", sump2)

}
