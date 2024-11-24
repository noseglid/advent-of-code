package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Dir int

const (
	Up Dir = iota
	Right
	Down
	Left
)

func (d Dir) String() string {
	switch d {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	}
	panic("bad dir")
}

func findStart(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '|' {
				return x, y
			}
		}
	}
	panic("no start")
}

func walkable(r rune) bool {
	return r == '|' || r == '-' || r == '+' || isLetter(r)
}

func isLetter(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func follow(x, y int, dir Dir, grid [][]rune) (int, int, []rune, int, bool) {
	end := false
	dx, dy, steps := 0, 0, 0
	letters := []rune{}
	switch dir {
	case Up:
		dy = -1
	case Right:
		dx = 1
	case Down:
		dy = 1
	case Left:
		dx = -1
	}

	for {
		if isLetter(grid[y][x]) {
			letters = append(letters, grid[y][x])
		}
		y += dy
		x += dx
		steps++
		if grid[y][x] == '+' {
			break
		}
		if !walkable(grid[y][x]) {
			end = true
			break
		}

	}

	return x, y, letters, steps, end
}

func turn(x, y int, grid [][]rune, dir Dir) Dir {
	switch dir {
	case Up, Down:
		if x > 0 && walkable(grid[y][x-1]) {
			return Left
		}
		if x < len(grid[y])-1 && walkable(grid[y][x+1]) {
			return Right
		}
	case Right, Left:
		if y > 0 && walkable(grid[y-1][x]) {
			return Up
		}
		if y < len(grid)-1 && walkable(grid[y+1][x]) {
			return Down
		}
	}
	panic("no turn")
}

func dumpLetters(letters []rune) string {
	var sb strings.Builder
	sb.WriteRune('[')
	for _, r := range letters {
		sb.WriteRune(r)
	}
	sb.WriteRune(']')
	return sb.String()
}

func main() {
	grid := util.GetFileRuneGrid("2017/Day19/input")
	d, letters := Down, []rune{}
	x, y := findStart(grid)
	steps := 0

	for {
		// fmt.Printf("start: %d,%d => %s (%v)\n", x, y, d, dumpLetters(letters))
		nx, ny, nletters, ns, done := follow(x, y, d, grid)
		x, y, letters, steps = nx, ny, append(letters, nletters...), steps+ns
		if done {
			break
		}

		d = turn(x, y, grid, d)
		fmt.Printf("%d, steps=%d\n", ns, steps)
		// fmt.Printf("end: %d,%d => %s (%v)\n", x, y, d, dumpLetters(letters))
	}

	fmt.Printf("Code (part1): %s", dumpLetters(letters))
	fmt.Printf("Steps (part2): %d", steps)
	// x, y, letters = follow(x, y, d, grid)
	// fmt.Printf("%d,%d => %s (%v)\n", x, y, d, dumpLetters(letters))

	// d = turn(x, y, grid, d)
	// fmt.Printf("%d,%d => %s (%v)\n", x, y, d, dumpLetters(letters))

}
