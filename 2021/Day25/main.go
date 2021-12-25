package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type point struct {
	x, y int
}

type cucumber rune

var (
	None  cucumber = '.'
	East  cucumber = '>'
	South cucumber = 'v'
)

func (c cucumber) Direction() (int, int) {
	switch c {
	case East:
		return 1, 0
	case South:
		return 0, 1
	}
	panic("move on bad cucumber")
}

type cmap struct {
	cucumbers     [][]cucumber
	width, height int
}

func (m cmap) String() string {
	var sb strings.Builder
	for y := range m.cucumbers {
		for x := range m.cucumbers[y] {
			sb.WriteRune(rune(m.cucumbers[y][x]))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (m *cmap) CanMoveType(tt cucumber) []point {
	var canMove []point
	for y := range m.cucumbers {
		for x := range m.cucumbers[y] {
			c := m.cucumbers[y][x]
			if c != tt {
				continue
			}

			dx, dy := c.Direction()
			if m.cucumbers[(y+dy)%m.height][(x+dx)%m.width] == None {
				canMove = append(canMove, point{x, y})
			}
		}
	}
	return canMove
}

func (m *cmap) Move(points []point) {
	for _, p := range points {
		cc := m.cucumbers[p.y][p.x]
		dx, dy := cc.Direction()
		m.cucumbers[p.y][p.x] = None
		m.cucumbers[(p.y+dy)%m.height][(p.x+dx)%m.width] = cc
	}
}

func (m *cmap) Step() int {
	n := 0
	e := m.CanMoveType(East)
	n += len(e)
	m.Move(e)
	s := m.CanMoveType(South)
	n += len(s)
	m.Move(s)
	return n
}

func parseMap(lines []string) *cmap {
	res := cmap{
		cucumbers: make([][]cucumber, len(lines)),
	}
	for y, l := range lines {
		res.cucumbers[y] = make([]cucumber, 0, len(l))
		for _, r := range l {
			res.cucumbers[y] = append(res.cucumbers[y], cucumber(r))
		}
	}
	res.height = len(res.cucumbers)
	res.width = len(res.cucumbers[0])

	return &res
}

func main() {
	input := "2021/Day25/input"
	lines := util.GetFileStrings(input)

	cmap := parseMap(lines)
	steps := 0
	for {
		steps++
		n := cmap.Step()
		if n == 0 {
			break
		}
	}

	log.Printf("Part 1: No movement after %d steps", steps)

}
