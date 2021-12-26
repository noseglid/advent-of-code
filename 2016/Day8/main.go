package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

const (
	w = 50
	h = 6
)

type Drawer interface {
	Draw(d *display)
}

type display struct {
	grid [h][w]bool
}

func (d display) print() {
	for y := range d.grid {
		for x := range d.grid[y] {
			c := '.'
			if d.grid[y][x] {
				c = '#'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func (d display) pixelsLit() int {
	n := 0
	for y := range d.grid {
		for x := range d.grid[y] {
			if d.grid[y][x] {
				n++
			}
		}
	}
	return n
}

type rect struct {
	a, b int
}

func (r rect) Draw(d *display) {
	for y := 0; y < r.b; y++ {
		for x := 0; x < r.a; x++ {
			d.grid[y][x] = true
		}
	}
}

type rotateRow struct {
	a, b int
}

func (r rotateRow) Draw(d *display) {
	newRow := make([]bool, w)
	for i := 0; i < w; i++ {
		if !d.grid[r.a][i] {
			continue
		}
		newPos := (i + r.b) % w
		newRow[newPos] = d.grid[r.a][i]
	}

	for i, c := range newRow {
		d.grid[r.a][i] = c
	}
}

type rotateCol struct {
	a, b int
}

func (r rotateCol) Draw(d *display) {
	newCol := make([]bool, h)
	for i := 0; i < h; i++ {
		if !d.grid[i][r.a] {
			continue
		}
		newPos := (i + r.b) % h
		newCol[newPos] = d.grid[i][r.a]
	}

	for i, c := range newCol {
		d.grid[i][r.a] = c
	}
}

func parseDrawer(s string) Drawer {
	fs := strings.IndexRune(s, ' ')
	switch s[:fs] {
	case "rect":
		r := rect{}
		if _, err := fmt.Sscanf(s[fs+1:], "%dx%d", &r.a, &r.b); err != nil {
			panic(err)
		}
		return r
	case "rotate":
		switch s[fs+1] {
		case 'r':
			r := rotateRow{}
			if _, err := fmt.Sscanf(s[fs+1:], "row y=%d by %d", &r.a, &r.b); err != nil {
				panic(err)
			}
			return r
		case 'c':
			r := rotateCol{}
			if _, err := fmt.Sscanf(s[fs+1:], "column x=%d by %d", &r.a, &r.b); err != nil {
				panic(err)
			}
			return r
		}
	}
	panic("unknown drawer")
}

func main() {
	input := "2016/Day8/input"
	lines := util.GetFileStrings(input)
	var ops []Drawer
	for _, l := range lines {
		ops = append(ops, parseDrawer(l))
	}

	var d display
	for _, o := range ops {
		o.Draw(&d)
	}

	log.Printf("Part 1: Pixels lit: %d", d.pixelsLit())
	log.Printf("Part 2: printed on next line!")

	d.print()
}
