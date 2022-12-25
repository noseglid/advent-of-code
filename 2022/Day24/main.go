package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Entry rune

const (
	Open          Entry = '.'
	Wall          Entry = '#'
	BlizzardNorth Entry = '^'
	BlizzardEast  Entry = '>'
	BlizzardSouth Entry = 'v'
	BlizzardWest  Entry = '<'
)

type grid struct {
	iteration   int
	expeditions []util.Point
	entries     [][][]Entry
}

func (g grid) HasAnyEntry(x, y int, en ...Entry) bool {
	for _, e := range g.entries[y][x] {
		for _, en := range en {
			if e == en {
				return true
			}
		}
	}
	return false
}

func (g grid) CanMoveTo(entries []Entry) bool {
	for _, e := range entries {
		if e != Open {
			return false
		}
	}
	return true
}

func (g grid) Blizzards(x, y int) []Entry {
	var ret []Entry
	for _, b := range g.entries[y][x] {
		switch b {
		case BlizzardNorth, BlizzardEast, BlizzardSouth, BlizzardWest:
			ret = append(ret, b)
		}
	}
	return ret
}

func (g grid) AddEntry(dst [][][]Entry, x, y int, en Entry) bool {
	dst[y][x] = append(dst[y][x], en)
	return true
}

func NewGrid(width, height int) *grid {
	g := &grid{}
	g.entries = make([][][]Entry, height)
	for y := range g.entries {
		g.entries[y] = make([][]Entry, width)
	}
	return g
}

func CopyEntries(src [][][]Entry) [][][]Entry {
	n := make([][][]Entry, len(src))
	for y, r := range src {
		n[y] = make([][]Entry, len(r))
		for x, c := range r {
			for _, e := range c {
				switch e {
				case Open, Wall:
					n[y][x] = append(n[y][x], e)
				}
			}
		}
	}

	return n
}

func (g *grid) DestroyExpeditions() {
	expSet := map[util.Point]struct{}{}
	for _, e := range g.expeditions {
		if !g.HasAnyEntry(e.X, e.Y, BlizzardNorth, BlizzardEast, BlizzardSouth, BlizzardWest) {
			expSet[e] = struct{}{}
		}
	}

	newExpeditions := []util.Point{}
	for e := range expSet {
		newExpeditions = append(newExpeditions, e)
	}
	g.expeditions = newExpeditions
}

func (g *grid) Iterate(target util.Point) bool {
	g.DestroyExpeditions()
	newGrid := CopyEntries(g.entries)
	for y, r := range g.entries {
		for x, c := range r {
			for _, en := range c {
				switch en {
				case '^':
					if g.HasAnyEntry(x, y-1, Wall) {
						g.AddEntry(newGrid, x, len(g.entries)-2, BlizzardNorth)
					} else {
						g.AddEntry(newGrid, x, y-1, BlizzardNorth)
						// g.DestroyExpedition(x, y-1)
					}
				case '>':
					if g.HasAnyEntry(x+1, y, Wall) {
						g.AddEntry(newGrid, 1, y, BlizzardEast)
					} else {
						g.AddEntry(newGrid, x+1, y, BlizzardEast)
						// g.DestroyExpedition(x+1, y)
					}
				case 'v':
					if g.HasAnyEntry(x, y+1, Wall) {
						g.AddEntry(newGrid, x, 1, BlizzardSouth)
					} else {
						g.AddEntry(newGrid, x, y+1, BlizzardSouth)
						// g.DestroyExpedition(x, y+1)
					}
				case '<':
					if g.HasAnyEntry(x-1, y, Wall) {
						g.AddEntry(newGrid, len(g.entries[0])-2, y, BlizzardWest)
					} else {
						g.AddEntry(newGrid, x-1, y, BlizzardWest)
						// g.DestroyExpedition(x-1, y)
					}
				}
			}
		}
	}
	g.entries = newGrid

	newExpeditions := []util.Point{}
	foundEnd := false
	for _, exp := range g.expeditions {
		for _, m := range PossibleMoves(exp) {
			if m.Y < 0 || m.Y >= len(newGrid) || m.X < 0 || m.X >= len(newGrid[m.Y]) {
				continue
			}
			if g.CanMoveTo(newGrid[m.Y][m.X]) {
				if m == target {
					foundEnd = true
				}
				newExpeditions = append(newExpeditions, m)
			}
		}
	}
	g.expeditions = newExpeditions

	g.iteration++

	return foundEnd
}

func PossibleMoves(src util.Point) []util.Point {
	return []util.Point{
		{X: src.X, Y: src.Y - 1},
		{X: src.X + 1, Y: src.Y},
		{X: src.X, Y: src.Y + 1},
		{X: src.X - 1, Y: src.Y},
		src,
	}
}

func (g *grid) String() string {
	var sb strings.Builder
	for y, r := range g.entries {
		for x, c := range r {
			foundExpedition := false
			for _, e := range g.expeditions {
				if e.X == x && e.Y == y {
					foundExpedition = true
					sb.WriteRune('E')
					break
				}
			}
			if foundExpedition {
				continue
			}

			if blizzards := g.Blizzards(x, y); len(blizzards) > 0 {
				if len(blizzards) > 1 {
					sb.WriteString(strconv.Itoa(len(blizzards)))
				} else if len(blizzards) == 1 {
					sb.WriteRune(rune(c[1]))
				}
			} else if g.HasAnyEntry(x, y, Wall) {
				sb.WriteRune('#')
			} else if g.HasAnyEntry(x, y, Open) {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func main() {
	input := util.GetFileRuneGrid("2022/Day24/input")
	g := NewGrid(len(input[0]), len(input))
	end, start := util.Point{}, util.Point{}
	for y, r := range input {
		for x, c := range r {
			switch c {
			case '.':
				p := util.Point{X: x, Y: y}
				if y == 0 {
					start = p
				}
				if y == len(input)-1 {
					end = p
				}
				g.entries[y][x] = append(g.entries[y][x], Open)
			case '^':
				g.entries[y][x] = append(g.entries[y][x], Open, BlizzardNorth)
			case '>':
				g.entries[y][x] = append(g.entries[y][x], Open, BlizzardEast)
			case 'v':
				g.entries[y][x] = append(g.entries[y][x], Open, BlizzardSouth)
			case '<':
				g.entries[y][x] = append(g.entries[y][x], Open, BlizzardWest)
			case '#':
				g.entries[y][x] = append(g.entries[y][x], Open, Wall)
			}
		}
	}
	g.expeditions = append(g.expeditions, start)

	for {
		if g.Iterate(end) {
			fmt.Printf("Found end after %d minutes (part1)\n", g.iteration)
			break
		}
	}

	targets := []util.Point{end, start, end}
	s := 0
	for len(targets) > 0 {
		if g.Iterate(targets[0]) {
			fmt.Printf("Found target after %d\n", g.iteration)
			s = g.iteration
			g.expeditions = []util.Point{targets[0]}
			targets = targets[1:]
		}
	}

	fmt.Printf("There, back, then there again (part2): %d\n", s)

}
