package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

//go:embed input
var realInput string

type jets struct {
	pattern string
	pos     int
}

func (j *jets) Left() bool {
	left := j.pattern[j.pos] == '<'
	j.pos = (j.pos + 1) % len(j.pattern)
	return left
}

var patterns = []string{
	`
####

`,
	`
 #
###
 #

`,
	`
###
  #
  #

`,
	`
#
#
#
#
`,
	`
##
##

`,
}

type rock struct {
	layout [][]rune
	x, y   int
}

func (r rock) String() string {
	var sb strings.Builder
	for _, row := range r.layout {
		for _, c := range row {
			sb.WriteRune(c)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (r rock) Dimension() (int, int) {
	maxy, maxx := 0, 0
	for y, row := range r.layout {
		for x, c := range row {
			if c != '#' {
				continue
			}
			maxy = util.Max(maxy, y+1)
			maxx = util.Max(maxx, x+1)
		}
	}
	return maxx, maxy
}

func (r rock) Covers(x, y int) bool {
	offx, offy := x-r.x, y-r.y
	if offy < 0 || offy >= len(r.layout) || offx < 0 || offx >= len(r.layout[offy]) {
		return false
	}
	return r.layout[offy][offx] == '#'
}

type rockGenerator struct {
	rocks []rock
	pos   int
}

func (r *rockGenerator) Next() *rock {
	rr := &rock{layout: r.rocks[r.pos].layout}
	r.pos = (r.pos + 1) % len(r.rocks)
	return rr
}

type chamber struct {
	grid        [][]rune
	rg          *rockGenerator
	jets        *jets
	currentRock *rock
	rocksAtRest int
}

func NewChamber(jets *jets, rg *rockGenerator) *chamber {
	g := make([][]rune, 1000000)
	for y := range g {
		g[y] = make([]rune, 7)
		for x := range g[y] {
			g[y][x] = '.'
		}
	}
	return &chamber{
		grid: g,
		jets: jets,
		rg:   rg,
	}
}

func (c chamber) HighestOpenLine() int {
Row:
	for y := range c.grid {
		for _, c := range c.grid[y] {
			if c == '#' {
				continue Row
			}
		}
		// Checked all cols in this row, none occupied, this is the first free line
		return y
	}

	panic("no open lines")
}

func (c chamber) Print(rows int) {
	height := c.TowerHeight()
	for y := height + rows; y >= util.Max(0, height-rows); y-- {
		fmt.Printf("%03d |", y)
		for x := 0; x < len(c.grid[y]); x++ {
			if c.currentRock != nil && c.currentRock.Covers(x, y) {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", c.grid[y][x])
			}
		}
		fmt.Println("|")
	}
	fmt.Printf("     +%s+\n", strings.Repeat("-", len(c.grid[0])))
}

func (c *chamber) AddRock() {
	r := c.rg.Next()
	l := c.HighestOpenLine()
	r.x = 2
	r.y = l + 3
	c.currentRock = r
}

func (c *chamber) jetEject() {
	w, h := c.currentRock.Dimension()
	if c.jets.Left() {
		if c.currentRock.x == 0 {
			return
		}
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				tx, ty := c.currentRock.x+x-1, c.currentRock.y+y
				if c.currentRock.Covers(tx+1, ty) && c.grid[ty][tx] == '#' {
					// blocked
					return
				}
			}
		}
		// fmt.Printf("push left\n")
		c.currentRock.x--
	} else {
		if c.currentRock.x+w >= len(c.grid[0]) {
			return
		}
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				tx, ty := c.currentRock.x+x+1, c.currentRock.y+y
				if c.currentRock.Covers(tx-1, ty) && c.grid[ty][tx] == '#' {
					// blocked
					return
				}
			}
		}
		// fmt.Printf("push right\n")
		c.currentRock.x++
	}
}

func (c *chamber) currentRockToRest() {
	// fmt.Printf("Come to rest!\n")
	w, h := c.currentRock.Dimension()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if c.currentRock.Covers(c.currentRock.x+x, c.currentRock.y+y) {
				c.grid[c.currentRock.y+y][c.currentRock.x+x] = '#'
			}
		}
	}

	c.rocksAtRest++
	c.currentRock = nil
}

func (c *chamber) Step() {
	// Add rock if necessary
	if c.currentRock == nil {
		c.AddRock()
		// c.Print(50)
		// fmt.Println("==========")
	}

	w, h := c.currentRock.Dimension()

	// Make jets do its job
	c.jetEject()

	if c.currentRock.y == 0 {
		c.currentRockToRest()
		return
	}

	// Fall down ...
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if !c.currentRock.Covers(c.currentRock.x+x, c.currentRock.y+y) {
				continue
			}
			testx, testy := c.currentRock.x+x, c.currentRock.y+y-1
			if c.grid[testy][testx] == '#' {
				c.currentRockToRest()
				return
			}

		}
		// if c.grid[c.currentRock.y-1][c.currentRock.x+x] == '#' {
		// 	// Blocked! Can't move down
		// 	c.currentRockToRest()
		// 	return
		// }
	}

	// ... nothing blocked, move down
	c.currentRock.y--
}

func (c chamber) TowerHeight() int {
Rows:
	for y := 0; y < len(c.grid); y++ {
		for x := 0; x < len(c.grid[y]); x++ {
			if c.grid[y][x] == '#' {
				continue Rows
			}
		}
		return y
	}
	panic("no empty rows")
}

func main() {
	// input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	input := strings.TrimSpace(realInput)
	fmt.Printf("inputlen: %d\n", len(input))

	rocks := make([]rock, 5)
	for i, p := range patterns {
		p = strings.TrimPrefix(p, "\n")
		p = strings.TrimSuffix(p, "\n")
		ll := make([][]rune, 4)
		for y, row := range strings.Split(p, "\n") {
			ll[y] = make([]rune, 4)
			for x, c := range row {
				ll[y][x] = c
			}
		}
		rocks[i] = rock{layout: ll}
	}

	jj := &jets{pattern: input}
	rg := &rockGenerator{rocks: rocks, pos: 0}
	chamber := NewChamber(jj, rg)

	for chamber.rocksAtRest < 2022 {
		chamber.Step()
	}

	chamber.Print(10)
	fmt.Printf("Height of tower after 2022 rocks resting (part1): %d\n", chamber.TowerHeight())

	lcm := LCM(len(input), len(rocks))
	fmt.Printf("lcm=%d\n", lcm)

	jj2 := &jets{pattern: input}
	rg2 := &rockGenerator{rocks: rocks, pos: 0}
	chamber2 := NewChamber(jj2, rg2)
	prev := 0
	first := 0
	for chamber2.rocksAtRest < 1000000 {
		if chamber2.rocksAtRest%lcm == 0 && chamber2.TowerHeight()-prev != 0 {
			fmt.Printf("rocksAtRest=%d, diff=%d\n", chamber2.rocksAtRest, chamber2.TowerHeight()-prev)
			prev = chamber2.TowerHeight()

			if first == 0 {
				first = chamber2.TowerHeight()
			}
		}
		chamber2.Step()
	}
}
