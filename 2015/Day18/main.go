package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

const dim = 100

func parseRow(s string) []int {
	row := make([]int, dim)
	for i, c := range s {
		switch c {
		case '.':
			row[i] = 0
		case '#':
			row[i] = 1
		default:
			panic("invalid state when parsing")
		}
	}
	return row
}

func neighbourState(x, y int, grid [][]int) (int, int) {
	neighbours := []struct{ dx, dy int }{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}

	on := 0
	off := 0
	for _, n := range neighbours {
		if (x+n.dx) < 0 || (x+n.dx) >= len(grid[0]) || (y+n.dy) < 0 || (y+n.dy) >= len(grid) {
			off++
			continue
		}
		if grid[y+n.dy][x+n.dx] == 1 {
			on++
		} else {
			off++
		}
	}

	return on, off
}

func nextState(currentState, nOn, _ int) int {
	switch currentState {
	case 0:
		if nOn == 3 {
			return 1
		} else {
			return 0
		}

	case 1:
		if nOn == 2 || nOn == 3 {
			return 1
		} else {
			return 0
		}
	}

	panic("invalid state")
}

func copyGrid(grid [][]int) [][]int {
	ngrid := make([][]int, len(grid))
	for y := range grid {
		ngrid[y] = make([]int, len(grid[y]))
		for x := range grid[y] {
			ngrid[y][x] = grid[y][x]
		}
	}
	return ngrid
}

func iterate(grid [][]int, cornerOn bool) [][]int {
	ngrid := copyGrid(grid)

	for y := range grid {
		for x := range grid[y] {
			if cornerOn && (y == 0 || y == len(grid)-1) && (x == 0 || x == len(grid[y])-1) {
				ngrid[y][x] = 1
				continue
			}

			on, off := neighbourState(x, y, grid)
			ngrid[y][x] = nextState(grid[y][x], on, off)
		}
	}
	return ngrid
}

func printGrid(grid [][]int) {
	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func lightsOn(grid [][]int) int {
	n := 0
	for _, row := range grid {
		for _, light := range row {
			if light == 1 {
				n++
			}
		}
	}

	return n
}

func main() {
	s := util.FileScanner("2015/Day18/input", bufio.ScanLines)

	iterations := 100
	grid := make([][]int, 0, dim)

	for s.Scan() {
		grid = append(grid, parseRow(s.Text()))
	}

	gridp1 := copyGrid(grid)
	for i := 0; i < iterations; i++ {
		gridp1 = iterate(gridp1, false)
	}

	log.Printf("lights on after %d iterations (part1): %d", iterations, lightsOn(gridp1))

	gridp2 := copyGrid(grid)
	gridp2[0][0] = 1
	gridp2[0][len(gridp2[0])-1] = 1
	gridp2[len(gridp2)-1][0] = 1
	gridp2[len(gridp2)-1][len(gridp2[len(gridp2)-1])-1] = 1
	printGrid(gridp2)
	fmt.Println()
	for i := 0; i < iterations; i++ {
		gridp2 = iterate(gridp2, true)
		printGrid(gridp2)
		fmt.Println()
	}
	log.Printf("lights on after %d iterations with corners on (part2): %d", iterations, lightsOn(gridp2))
}
