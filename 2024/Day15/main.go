package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type move rune

const (
	up    move = '^'
	right move = '>'
	down  move = 'v'
	left  move = '<'
)

type p struct{ x, y int }

func part1(grid util.Grid, moves []move) {

	rx, ry := grid.Find('@')

	for _, m := range moves {
		nx, ny := grid.GetMove(rx, ry, util.VDir(rune(m)))
		switch grid.Get(nx, ny) {
		case '.':
			grid.Switch(rx, ry, nx, ny)
			rx, ry = nx, ny
		case '#':
		case 'O':
			nnx, nny := nx, ny
			coords := []p{{nx, ny}}
		Loop:
			for {
				nnx, nny = grid.GetMove(nnx, nny, util.VDir(rune(m)))
				e := grid.Get(nnx, nny)
				switch e {
				case 'O':
					coords = append(coords, p{nnx, nny})
				case '#':
					break Loop
				case '.':
					coords = append(coords, p{nnx, nny})
					for i := len(coords) - 1; i > 0; i-- {
						grid.Switch(coords[i].x, coords[i].y, coords[i-1].x, coords[i-1].y)
					}
					grid.Switch(rx, ry, nx, ny)
					rx, ry = nx, ny
					break Loop
				}
			}
		}
	}

	s := 0
	grid.Each(func(x, y int, r rune) {
		if r != 'O' {
			return
		}
		s += y*100 + x
	})
	fmt.Printf("sum of all coordinates (part1): %d\n", s)
}

func toWide(grid util.Grid) util.Grid {
	ng := make([][]rune, len(grid))
	for y, row := range grid {
		ng[y] = make([]rune, len(row)*2)
		for x, cell := range row {
			switch cell {
			case '#':
				ng[y][x*2], ng[y][x*2+1] = '#', '#'
			case 'O':
				ng[y][x*2], ng[y][x*2+1] = '[', ']'
			case '@':
				ng[y][x*2], ng[y][x*2+1] = '@', '.'
			case '.':
				ng[y][x*2], ng[y][x*2+1] = '.', '.'
			}
		}
	}
	return ng
}

type box struct {
	x, y int
}

func Box(grid util.Grid, x, y int) box {
	if grid.Get(x, y) == ']' {
		return box{x - 1, y}
	}
	return box{x, y}
}

func (b box) IsBox(grid util.Grid) bool {
	return grid.Get(b.x, b.y) == '['
}

func canMove(grid util.Grid, b box, d util.Dir) (bool, []box) {
	lx, ly := grid.GetMove(b.x, b.y, d)
	rx, ry := grid.GetMove(b.x+1, b.y, d)

	boxes := []box{b}

	ml, mr := false, false
	switch grid.Get(lx, ly) {
	case '.':
		ml = true
	case '#':
		ml = false
	case '[', ']':
		bb := Box(grid, lx, ly)
		if b == bb {
			ml = true
		} else {
			iml, rr := canMove(grid, bb, d)
			ml = iml
			boxes = append(boxes, rr...)
		}
	}
	if Box(grid, rx, ry) == Box(grid, lx, ly) {
		mr = ml
	} else {
		switch grid.Get(rx, ry) {
		case '.':
			mr = true
		case '#':
			mr = false
		case '[', ']':
			bb := Box(grid, rx, ry)
			if b == bb {
				mr = true
			} else {
				imr, rr := canMove(grid, bb, d)
				mr = imr
				boxes = append(boxes, rr...)
			}
		}
	}

	return ml && mr, boxes

}

func part2(grid util.Grid, moves []move) {
	grid = toWide(grid)
	rx, ry := grid.Find('@')

	for _, m := range moves {
		d := util.VDir(rune(m))
		nx, ny := grid.GetMove(rx, ry, d)
		switch grid.Get(nx, ny) {
		case '.':
			grid.Switch(rx, ry, nx, ny)
			rx, ry = nx, ny
		case '#':
		case '[', ']':
			ok, boxes := canMove(grid, Box(grid, nx, ny), d)
			if !ok {
				break
			}
			dx, dy := d.Deltas()
			moved := map[box]bool{}
			for i := len(boxes) - 1; i >= 0; i-- {
				b := boxes[i]
				if moved[b] {
					continue
				}
				moved[b] = true
				if dx < 0 {
					grid.Switch(b.x, b.y, b.x+dx, b.y+dy)
					grid.Switch(b.x+1, b.y, b.x+1+dx, b.y+dy)
				} else {
					grid.Switch(b.x+1, b.y, b.x+1+dx, b.y+dy)
					grid.Switch(b.x, b.y, b.x+dx, b.y+dy)
				}
			}
			grid.Switch(rx, ry, rx+dx, ry+dy)
			rx, ry = rx+dx, ry+dy
		}
	}

	s := 0
	grid.Each(func(x, y int, r rune) {
		if r != '[' {
			return
		}
		s += y*100 + x
	})
	fmt.Printf("sum of all coordinates (part2): %d\n", s)
}

func main() {
	lines := util.GetFileStrings("2024/Day15/input")
	divider := slices.IndexFunc(lines, func(l string) bool { return l == "" })
	moves := []move{}
	for _, l := range lines[divider+1:] {
		moves = append(moves, []move(l)...)
	}

	part1(util.ParseRuneGrid(lines[0:divider]), moves)
	part2(util.ParseRuneGrid(lines[0:divider]), moves)
}
