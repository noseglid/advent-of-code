package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type line struct {
	x1, y1 int
	x2, y2 int
}

type grid [][]int

func (g grid) Print() {
	for _, r := range g {
		for _, c := range r {
			if c == 0 {
				fmt.Printf(" .")
			} else {
				fmt.Printf("%2d", c)
			}
		}
		fmt.Printf("\n")
	}
}

func (g grid) CountIntersects(min int) int {
	n := 0
	for _, r := range g {
		for _, c := range r {
			if c >= min {
				n++
			}
		}
	}

	return n
}

func (l line) Print() {
	fmt.Printf("%d,%d -> %d,%d\n", l.x1, l.y1, l.x2, l.y2)
}

func (l line) ApplyToGrid(grid [][]int, diagonal bool) {
	if l.x1 == l.x2 {
		s := util.MinInt(l.y1, l.y2)
		e := util.MaxInt(l.y1, l.y2)

		for i := s; i <= e; i++ {
			grid[i][l.x1] = grid[i][l.x1] + 1
		}
	} else if l.y1 == l.y2 {
		s := util.MinInt(l.x1, l.x2)
		e := util.MaxInt(l.x1, l.x2)

		for i := s; i <= e; i++ {
			grid[l.y1][i] = grid[l.y1][i] + 1
		}
	} else if diagonal {
		xStep := 1
		if l.x2 < l.x1 {
			xStep = -1
		}
		yStep := 1
		if l.y2 < l.y1 {
			yStep = -1
		}

		yy, xx := l.y1, l.x1
		for yy != l.y2 {
			grid[yy][xx] = grid[yy][xx] + 1
			yy += yStep
			xx += xStep
		}
		grid[yy][xx] = grid[yy][xx] + 1
	}
}

func xMax(lines []line) int {
	m := 0
	for _, l := range lines {
		if l.x1 > m {
			m = l.x1
		}
		if l.x2 > m {
			m = l.x2
		}
	}

	return m
}

func yMax(lines []line) int {
	m := 0
	for _, l := range lines {
		if l.y1 > m {
			m = l.y1
		}
		if l.y2 > m {
			m = l.y2
		}
	}

	return m
}

func main() {
	input := "2021/Day5/input"

	s := util.FileScanner(input, bufio.ScanLines)

	var lines []line
	for s.Scan() {
		var l line
		if _, err := fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &l.x1, &l.y1, &l.x2, &l.y2); err != nil {
			panic(err)
		}
		lines = append(lines, l)
	}

	mx := xMax(lines)
	my := yMax(lines)

	var g grid = make([][]int, my+1)
	for i := range g {
		g[i] = make([]int, mx+1)
	}
	var g2 grid = make([][]int, my+1)
	for i := range g2 {
		g2[i] = make([]int, mx+1)
	}

	for _, l := range lines {
		l.ApplyToGrid(g, false)
		l.ApplyToGrid(g2, true)
	}

	log.Printf("Part 1: At least 2 intersects: %d", g.CountIntersects(2))
	log.Printf("Part 2: At least 2 intersects: %d", g2.CountIntersects(2))
}
