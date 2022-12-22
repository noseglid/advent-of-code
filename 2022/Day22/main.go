package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type foldEntry struct {
	p         util.Point
	newFacing rune
}

var foldStrategyp1 map[util.Point]foldEntry = make(map[util.Point]foldEntry)
var foldStrategyp2 map[util.Point]foldEntry = make(map[util.Point]foldEntry)

func initp1() {
	for i := 0; i < 50; i++ {
		e1 := util.Point{X: 150, Y: i}
		d1 := util.Point{X: 49, Y: i}
		foldStrategyp1[e1] = foldEntry{p: util.Point{X: d1.X + 1, Y: d1.Y}, newFacing: 'R'}
		foldStrategyp1[d1] = foldEntry{p: util.Point{X: e1.X - 1, Y: e1.Y}, newFacing: 'L'}

		a2 := util.Point{X: 100, Y: 50 + i}
		b1 := util.Point{X: 49, Y: 50 + i}
		foldStrategyp1[a2] = foldEntry{p: util.Point{X: b1.X + 1, Y: b1.Y}, newFacing: 'R'}
		foldStrategyp1[b1] = foldEntry{p: util.Point{X: a2.X - 1, Y: a2.Y}, newFacing: 'L'}

		e2 := util.Point{X: 100, Y: 100 + i}
		d2 := util.Point{X: -1, Y: 100 + i}
		foldStrategyp1[e2] = foldEntry{p: util.Point{X: d2.X + 1, Y: d2.Y}, newFacing: 'R'}
		foldStrategyp1[d2] = foldEntry{p: util.Point{X: e2.X - 1, Y: e2.Y}, newFacing: 'L'}

		c2 := util.Point{X: 50, Y: 150 + i}
		f2 := util.Point{X: -1, Y: 150 + i}
		foldStrategyp1[c2] = foldEntry{p: util.Point{X: f2.X + 1, Y: f2.Y}, newFacing: 'R'}
		foldStrategyp1[f2] = foldEntry{p: util.Point{X: c2.X - 1, Y: c2.Y}, newFacing: 'L'}

		b2 := util.Point{X: i, Y: 99}
		g2 := util.Point{X: i, Y: 200}
		foldStrategyp1[b2] = foldEntry{p: util.Point{X: g2.X, Y: g2.Y - 1}, newFacing: 'U'}
		foldStrategyp1[g2] = foldEntry{p: util.Point{X: b2.X, Y: b2.Y + 1}, newFacing: 'D'}

		f1 := util.Point{X: 50 + i, Y: -1}
		c1 := util.Point{X: 50 + i, Y: 150}
		foldStrategyp1[f1] = foldEntry{p: util.Point{X: c1.X, Y: c1.Y - 1}, newFacing: 'U'}
		foldStrategyp1[c1] = foldEntry{p: util.Point{X: f1.X, Y: f1.Y + 1}, newFacing: 'D'}

		g1 := util.Point{X: 100 + i, Y: -1}
		a1 := util.Point{X: 100 + i, Y: 50}
		foldStrategyp1[g1] = foldEntry{p: util.Point{X: a1.X, Y: a1.Y - 1}, newFacing: 'U'}
		foldStrategyp1[a1] = foldEntry{p: util.Point{X: g1.X, Y: g1.Y + 1}, newFacing: 'D'}
	}
}

func initp2() {
	for i := 0; i < 50; i++ {
		a1 := util.Point{X: 100 + i, Y: 50}
		a2 := util.Point{X: 100, Y: 50 + i}
		foldStrategyp2[a1] = foldEntry{p: util.Point{X: a2.X - 1, Y: a2.Y}, newFacing: 'L'}
		foldStrategyp2[a2] = foldEntry{p: util.Point{X: a1.X, Y: a1.Y - 1}, newFacing: 'U'}

		b1 := util.Point{X: 49, Y: 50 + i}
		b2 := util.Point{X: i, Y: 99}
		foldStrategyp2[b2] = foldEntry{p: util.Point{X: b1.X + 1, Y: b1.Y}, newFacing: 'R'}
		foldStrategyp2[b1] = foldEntry{p: util.Point{X: b2.X, Y: b2.Y + 1}, newFacing: 'D'}

		c1 := util.Point{X: 50 + i, Y: 150}
		c2 := util.Point{X: 50, Y: 150 + i}
		foldStrategyp2[c1] = foldEntry{p: util.Point{X: c2.X - 1, Y: c2.Y}, newFacing: 'L'}
		foldStrategyp2[c2] = foldEntry{p: util.Point{X: c1.X, Y: c1.Y - 1}, newFacing: 'U'}

		d1 := util.Point{X: 49, Y: i}
		d2 := util.Point{X: -1, Y: 149 - i}
		foldStrategyp2[d1] = foldEntry{p: util.Point{X: d2.X + 1, Y: d2.Y}, newFacing: 'R'}
		foldStrategyp2[d2] = foldEntry{p: util.Point{X: d1.X + 1, Y: d1.Y}, newFacing: 'R'}

		e1 := util.Point{X: 150, Y: 49 - i}
		e2 := util.Point{X: 100, Y: 100 + i}
		foldStrategyp2[e1] = foldEntry{p: util.Point{X: e2.X - 1, Y: e2.Y}, newFacing: 'L'}
		foldStrategyp2[e2] = foldEntry{p: util.Point{X: e1.X - 1, Y: e1.Y}, newFacing: 'L'}

		f1 := util.Point{X: 50 + i, Y: -1}
		f2 := util.Point{X: -1, Y: 150 + i}
		foldStrategyp2[f1] = foldEntry{p: util.Point{X: f2.X + 1, Y: f2.Y}, newFacing: 'R'}
		foldStrategyp2[f2] = foldEntry{p: util.Point{X: f1.X, Y: f1.Y + 1}, newFacing: 'D'}

		g1 := util.Point{X: 100 + i, Y: -1}
		g2 := util.Point{X: i, Y: 200}
		foldStrategyp2[g1] = foldEntry{p: util.Point{X: g2.X, Y: g2.Y - 1}, newFacing: 'U'}
		foldStrategyp2[g2] = foldEntry{p: util.Point{X: g1.X, Y: g1.Y + 1}, newFacing: 'D'}
	}
}

func init() {
	initp1()
	initp2()
}

type square rune

type instruction interface {
	Apply(m *me, mm *mm)
	String() string
}

type rotateInstr struct {
	dir rune
}

func (instr rotateInstr) Apply(m *me, mm *mm) {
	switch instr.dir {
	case 'R':
		switch m.facing {
		case 'U':
			m.facing = 'R'
		case 'R':
			m.facing = 'D'
		case 'D':
			m.facing = 'L'
		case 'L':
			m.facing = 'U'
		}
	case 'L':
		switch m.facing {
		case 'U':
			m.facing = 'L'
		case 'R':
			m.facing = 'U'
		case 'D':
			m.facing = 'R'
		case 'L':
			m.facing = 'D'
		}
	}
}

func (i rotateInstr) String() string {
	return fmt.Sprintf("%c", i.dir)
}

type moveInstr struct {
	steps int
}

func (instr moveInstr) Apply(m *me, mm *mm) {
	next := func(x, y int) (int, int) {
		switch m.facing {
		case 'U':
			return x, y - 1
		case 'R':
			return x + 1, y
		case 'D':
			return x, y + 1
		case 'L':
			return x - 1, y
		}
		panic(fmt.Sprintf("bad facing: %d", m.facing))
	}

	for i := 0; i < instr.steps; i++ {
		nx, ny := next(m.pos.X, m.pos.Y)
		blocked, tx, ty, f := mm.CanMove(nx, ny, m.facing)
		if blocked {
			break
		}
		m.pos.X, m.pos.Y, m.facing = tx, ty, f
	}
}

func (instr moveInstr) String() string {
	return fmt.Sprintf("%d", instr.steps)
}

const (
	Outside square = ' '
	Wall    square = '#'
	Open    square = '.'
)

type mm struct {
	grid         [][]square
	instructions []instruction
	foldStrategy map[util.Point]foldEntry
}

func (m mm) CanMove(xto, yto int, facing rune) (bool, int, int, rune) {
	if t, ok := m.foldStrategy[util.Point{X: xto, Y: yto}]; ok {
		xto = t.p.X
		yto = t.p.Y
		facing = t.newFacing
	}

	switch m.grid[yto][xto] {
	case Open:
		return false, xto, yto, facing
	case Wall:
		return true, xto, yto, facing
	default:
		panic("bad grid value")
	}
}

type me struct {
	pos    util.Point
	facing rune
}

func (m me) Password() int {
	r := 1000 * (m.pos.Y + 1)
	c := 4 * (m.pos.X + 1)
	switch m.facing {
	case 'R':
		return r + c + 0
	case 'D':
		return r + c + 1
	case 'L':
		return r + c + 2
	case 'U':
		return r + c + 3
	}
	panic("no password")
}

func (m mm) String(me *me) string {
	var sb strings.Builder

	offset := 35
	ymin, ymax := me.pos.Y-offset, me.pos.Y+offset

	for y := ymin; y <= ymax; y++ {
		if y < 0 || y >= len(m.grid) {
			fmt.Println()
			continue
		}
		for x := range m.grid[y] {
			if me.pos.X == x && me.pos.Y == y {
				sb.WriteString("\033[31;1;4m")
				switch me.facing {
				case 'U':
					sb.WriteRune('^')
				case 'R':
					sb.WriteRune('>')
				case 'D':
					sb.WriteRune('v')
				case 'L':
					sb.WriteRune('<')
				}
				sb.WriteString("\033[00m")
			} else {
				sb.WriteRune(rune(m.grid[y][x]))
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func NewMM(rows, cols int) *mm {
	g := make([][]square, rows)
	for y := range g {
		g[y] = make([]square, cols)
		for x := range g[y] {
			g[y][x] = ' '
		}
	}

	return &mm{grid: g}
}

func (m mm) Start() util.Point {
	for y := range m.grid {
		for x := range m.grid[y] {
			if m.grid[y][x] == Open {
				return util.Point{X: x, Y: y}
			}
		}
	}
	panic("no start found")
}

func main() {
	input := util.GetFile("2022/Day22/input")
	mr, mc := 0, 0
	parts := strings.Split(input, "\n\n")
	for _, l := range strings.Split(parts[0], "\n") {
		mr = mr + 1
		mc = util.Max(mc, len(l))
	}
	mm := NewMM(mr, mc)

	for y, l := range strings.Split(parts[0], "\n") {
		for x, s := range l {
			mm.grid[y][x] = square(s)
		}
	}

	strinstr := strings.TrimSpace(parts[1])
	var instructions []instruction
	for i := 0; i < len(strinstr); i++ {
		switch strinstr[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			endIndex := strings.IndexFunc(strinstr[i:], func(r rune) bool { return r == 'R' || r == 'L' })
			if endIndex == -1 {
				endIndex = len(strinstr)
			} else {
				endIndex += i
			}
			instructions = append(instructions, moveInstr{util.MustAtoi(strinstr[i:endIndex])})
			i = endIndex - 1
		case 'L', 'R':
			instructions = append(instructions, rotateInstr{dir: rune(strinstr[i])})
		}

	}
	m := &me{
		pos:    mm.Start(),
		facing: 'R',
	}

	mm.foldStrategy = foldStrategyp1
	for _, i := range instructions {
		i.Apply(m, mm)
	}
	fmt.Printf("password (part1): %d\n", m.Password())

	m2 := &me{
		pos:    mm.Start(),
		facing: 'R',
	}
	mm.foldStrategy = foldStrategyp2
	for ii, i := range instructions {
		_ = ii
		i.Apply(m2, mm)
	}
	fmt.Printf("password (part2): %d\n", m2.Password())

}
