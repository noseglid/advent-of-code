package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func adj(grid [][]rune, x, y int) []rune {
	res := []rune{}
	if x > 0 {
		res = append(res, grid[y][x-1])
	}
	if x > 0 && y > 0 {
		res = append(res, grid[y-1][x-1])
	}
	if y > 0 {
		res = append(res, grid[y-1][x])
	}
	if x < len(grid[y])-1 && y > 0 {
		res = append(res, grid[y-1][x+1])
	}
	if x < len(grid[y])-1 {
		res = append(res, grid[y][x+1])
	}
	if x < len(grid[y])-1 && y < len(grid)-1 {
		res = append(res, grid[y+1][x+1])
	}
	if y < len(grid)-1 {
		res = append(res, grid[y+1][x])
	}
	if x > 0 && y < len(grid)-1 {
		res = append(res, grid[y+1][x-1])
	}
	return res
}

func count(r []rune) (int, int, int) {
	var open, tree, lumberyard int
	for _, rr := range r {
		switch rr {
		case '.':
			open++
		case '|':
			tree++
		case '#':
			lumberyard++
		}
	}
	return open, tree, lumberyard
}

func it(grid [][]rune) [][]rune {
	next := make([][]rune, len(grid))
	for y := range grid {
		next[y] = make([]rune, len(grid[y]))
	}

	for y, row := range grid {
		for x, cell := range row {
			_, tree, lumberyard := count(adj(grid, x, y))
			var nr rune
			switch cell {
			case '.':
				if tree >= 3 {
					nr = '|'
				} else {
					nr = '.'
				}
			case '|':
				if lumberyard >= 3 {
					nr = '#'
				} else {
					nr = '|'
				}
			case '#':
				if lumberyard >= 1 && tree >= 1 {
					nr = '#'
				} else {
					nr = '.'
				}
			}
			next[y][x] = nr
		}
	}
	return next
}

func total(grid [][]rune) (int, int, int) {
	open, tree, lumberyard := 0, 0, 0
	util.Grid(grid).Each(func(x, y int, r rune) {
		switch r {
		case '.':
			open++
		case '|':
			tree++
		case '#':
			lumberyard++
		}
	})
	return open, tree, lumberyard
}

func main() {
	grid := util.GetFileRuneGrid("2018/Day18/input")

	values := map[int][]int{}
Outer:
	for i := 1; ; i++ {
		grid = it(grid)
		_, tree, lumberyard := total(grid)
		rv := tree * lumberyard

		if i == 10 {
			fmt.Printf("Resource value (part1): %d\n", rv)
		}

		values[rv] = append(values[rv], i)

		for k, v := range values {
			if len(v) >= 2 && (1e9-v[0])%(v[1]-v[0]) == 0 {
				fmt.Printf("resource value after 1e9 minutes (part2): %d\n", k)
				break Outer
			}
		}
	}

}
