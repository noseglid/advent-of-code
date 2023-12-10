package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

func findStart(grid [][]rune) util.Point {
	for x, row := range grid {
		for y, c := range row {
			if c == 'S' {
				return util.Point{X: x, Y: y}
			}
		}
	}
	panic("no start")
}

func findFirstStep(s util.Point, grid [][]rune) util.Point {
	switch grid[s.X][s.Y-1] {
	case '-', 'F', 'L':
		return util.Point{X: s.X, Y: s.Y - 1}
	}
	switch grid[s.X][s.Y+1] {
	case '-', 'J', '7':
		return util.Point{X: s.X, Y: s.Y + 1}
	}
	switch grid[s.X-1][s.Y] {
	case '|', 'F', '7':
		return util.Point{X: s.X - 1, Y: s.Y}
	}
	switch grid[s.X+1][s.Y] {
	case '|', 'L', 'J':
		return util.Point{X: s.X + 1, Y: s.Y}
	}
	panic("no pipe connected to start")
}

func nextStep(prev, curr util.Point, grid [][]rune) util.Point {
	switch grid[curr.X][curr.Y] {
	case '-':
		if prev.Y == curr.Y-1 { // should move right
			return util.Point{X: curr.X, Y: curr.Y + 1}
		} else { // should move left
			return util.Point{X: curr.X, Y: curr.Y - 1}
		}
	case '|':
		if prev.X == curr.X-1 { // should move down
			return util.Point{X: curr.X + 1, Y: curr.Y}
		} else { // should move up
			return util.Point{X: curr.X - 1, Y: curr.Y}
		}
	case 'F':
		if prev.X == curr.X+1 { // should move down
			return util.Point{X: curr.X, Y: curr.Y + 1}
		} else { // should move right
			return util.Point{X: curr.X + 1, Y: curr.Y}
		}
	case '7':
		if prev.Y == curr.Y-1 { //should move down
			return util.Point{X: curr.X + 1, Y: curr.Y}
		} else { // should move left
			return util.Point{X: curr.X, Y: curr.Y - 1}
		}
	case 'J':
		if prev.X == curr.X-1 { // should move left
			return util.Point{X: curr.X, Y: curr.Y - 1}
		} else { // should move up
			return util.Point{X: curr.X - 1, Y: curr.Y}
		}
	case 'L':
		if prev.X == curr.X-1 { //should move right
			return util.Point{X: curr.X, Y: curr.Y + 1}
		} else { // should move up
			return util.Point{X: curr.X - 1, Y: curr.Y}
		}
	}
	panic("no next step")
}

func tileEnclosed(p util.Point, steps []util.Point, grid [][]rune) bool {
	intersects := 0
	onPipe := false
	var onPipeintersectOn rune
	for p.Y < len(grid[0]) {
		c := grid[p.X][p.Y]
		if slices.Contains(steps, p) {
			if c == '|' {
				onPipe = false
				onPipeintersectOn = 0
				intersects++
			} else if c == 'F' {
				onPipe = true
				onPipeintersectOn = 'J'
			} else if c == 'L' {
				onPipe = true
				onPipeintersectOn = '7'
			} else if onPipe && c == onPipeintersectOn {
				intersects++
				onPipe = false
				onPipeintersectOn = 0
			}
		}
		p.Y++
	}
	return intersects%2 != 0
}

func findEnclosed(steps []util.Point, grid [][]rune) int {
	n := 0
	for x, row := range grid {
		for y := range row {
			p := util.Point{X: x, Y: y}
			if !slices.Contains(steps, p) && tileEnclosed(p, steps, grid) {
				n++
			}
		}
	}
	return n
}

func findReplace(steps []util.Point) rune {
	prev, start, next := steps[len(steps)-1], steps[0], steps[1]
	if prev.Y == start.Y-1 { //from left
		if next.Y == start.Y+1 { // going right
			return '-'
		}
		if next.X == start.X-1 { // going up
			return 'J'
		}
		if next.X == start.X+1 { // going down
			return '7'
		}
	}
	if prev.Y == start.Y+1 { //from right
		if next.Y == start.Y-1 { // going left
			return '-'
		}
		if next.X == start.X-1 { // going up
			return 'L'
		}
		if next.X == start.X+1 { // going down
			return '7'
		}
	}
	if prev.X == start.X-1 { // from above
		if next.X == start.X+1 { // going down
			return '|'
		}
		if next.Y == start.Y-1 { // Going left
			return 'J'
		}
		if next.Y == start.Y+1 { // going right
			return 'L'
		}
	}
	if prev.X == start.X+1 { // from below
		if next.X == start.X-1 { // going up
			return '|'
		}
		if next.Y == start.Y-1 { // going left
			return '7'
		}
		if next.Y == start.Y+1 { // going right
			return 'F'
		}
	}
	panic("no replacement found")
}

func main() {
	grid := util.GetFileRuneGrid("2023/Day10/input")
	ps := findStart(grid)
	pc, pr := findFirstStep(ps, grid), ps
	steps := []util.Point{ps}
	for pc != ps {
		steps = append(steps, pc)
		pr, pc = pc, nextStep(pr, pc, grid)
	}
	grid[ps.X][ps.Y] = findReplace(steps)

	fmt.Printf("Furthest steps away (part1): %d\n", len(steps)/2)
	fmt.Printf("tiles enclosed (part2): %d\n", findEnclosed(steps, grid))

}
