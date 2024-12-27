package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point

type move struct {
	a, b P
}

func extractParenthesis(s string) string {
	d := 0
	for i, v := range s {
		switch v {
		case '(':
			d++
		case ')':
			d--
			if d == 0 {
				return s[1:i]
			}
		}
	}

	panic("no parenthesis")
}

func pipeSplit(s string) []string {
	d := 0
	parts := []string{}

	x := 0
	for i, v := range s {
		switch v {
		case '(':
			d++
		case ')':
			d--
		case '|':
			if d == 0 {
				parts = append(parts, s[x:i])
				x = i + 1
			}
		}
	}
	parts = append(parts, s[x:])
	return parts
}

func isUselessPart(s string) bool {
	sl := []rune(s)
	slices.Sort(sl)
	return len(sl) == 0
}

func traverse(regex string, pos P, doors map[move]bool) {
	for i := 0; i < len(regex); i++ {
		dx, dy := 0, 0
		switch regex[i] {
		case '^', '$':
			continue
		case 'N':
			dx, dy = util.N.Deltas()
		case 'E':
			dx, dy = util.E.Deltas()
		case 'S':
			dx, dy = util.S.Deltas()
		case 'W':
			dx, dy = util.W.Deltas()
		case '(':
			parenthesis := extractParenthesis(regex[i:])
			parts := pipeSplit(parenthesis)
			d := 0
			for _, pa := range parts {
				if isUselessPart(pa) {
					continue
				}
				path := fmt.Sprintf("%s%s", pa, regex[i+len(parenthesis)+2:])
				traverse(path, pos, doors)
				d++
			}
			if d == 0 {
				traverse(regex[i+len(parenthesis)+2:], pos, doors)
			}
			return
		default:
			panic(fmt.Sprintf("unexpected regex (%s) for %d: %c", regex, i, regex[i]))
		}
		npos := P{pos.X + dx, pos.Y + dy}
		doors[move{pos, npos}] = true
		pos = npos
	}
}

func buildGrid(doors map[move]bool) ([][]rune, P) {
	minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0
	for m := range doors {
		minx = util.Min(minx, m.a.X, m.b.X)
		maxx = util.Max(maxx, m.a.X, m.b.X)
		miny = util.Min(miny, m.a.Y, m.b.Y)
		maxy = util.Max(maxy, m.a.Y, m.b.Y)
	}

	xoffset, yoffset := 0-minx, 0-miny
	grid := make([][]rune, 2*(yoffset+maxy)+1)
	for y := range grid {
		grid[y] = make([]rune, 2*(xoffset+maxx)+1)
		for x := range grid[y] {
			grid[y][x] = '#'
		}
	}

	for m := range doors {
		grid[2*(m.a.Y+yoffset)][2*(m.a.X+xoffset)] = '.'
		grid[2*(m.b.Y+yoffset)][2*(m.b.X+xoffset)] = '.'
		if m.a.Y == m.b.Y {
			d := m.b.X - m.a.X
			grid[2*(m.a.Y+yoffset)][2*(m.a.X+xoffset)+d] = '|'
		} else {
			d := m.b.Y - m.a.Y
			grid[2*(m.a.Y+yoffset)+d][2*(m.a.X+xoffset)] = '-'
		}
	}

	ng := make([][]rune, len(grid)+2)
	for y := range ng {
		ng[y] = make([]rune, len(grid[0])+2)
		for x := range ng[y] {
			ng[y][x] = '#'
		}
	}

	for y := range grid {
		for x := range grid[y] {
			ng[y+1][x+1] = grid[y][x]
		}
	}
	return ng, P{2*xoffset + 1, 2*yoffset + 1}
}

var roomSteps = map[P]int{}

func mostDoors(grid [][]rune, p P, visited map[P]bool, steps int) int {
	if v, ok := roomSteps[p]; ok && v > steps || !ok {
		roomSteps[p] = steps
	}

	if visited[p] {
		return 0
	}
	visited[p] = true

	dn, de, ds, dw := 0, 0, 0, 0
	if grid[p.Y][p.X+1] == '|' && !visited[P{p.X + 2, p.Y}] {
		ve := util.CopyMap(visited)
		de = 1 + mostDoors(grid, P{p.X + 2, p.Y}, ve, steps+1)
	}
	if grid[p.Y][p.X-1] == '|' && !visited[P{p.X - 2, p.Y}] {
		vw := util.CopyMap(visited)
		dw = 1 + mostDoors(grid, P{p.X - 2, p.Y}, vw, steps+1)
	}
	if grid[p.Y-1][p.X] == '-' && !visited[P{p.X, p.Y - 2}] {
		vn := util.CopyMap(visited)
		dn = 1 + mostDoors(grid, P{p.X, p.Y - 2}, vn, steps+1)
	}
	if grid[p.Y+1][p.X] == '-' && !visited[P{p.X, p.Y + 2}] {
		vs := util.CopyMap(visited)
		ds = 1 + mostDoors(grid, P{p.X, p.Y + 2}, vs, steps+1)
	}

	return util.Max(dn, de, ds, dw)
}

func dijk(grid [][]rune, start P) map[P]int {
	unvisited := []P{start}
	nodes := map[P]int{
		start: 0,
	}

	for len(unvisited) > 0 {
		n := unvisited[0]
		checks := []struct {
			p, n P
			c    rune
		}{
			{P{n.X, n.Y - 1}, P{n.X, n.Y - 2}, '-'},
			{P{n.X, n.Y + 1}, P{n.X, n.Y + 2}, '-'},
			{P{n.X + 1, n.Y}, P{n.X + 2, n.Y}, '|'},
			{P{n.X - 1, n.Y}, P{n.X - 2, n.Y}, '|'},
		}
		for _, c := range checks {
			if grid[c.p.Y][c.p.X] == c.c {

				if v, ok := nodes[c.n]; !ok {
					nodes[c.n] = nodes[n] + 1
					unvisited = append(unvisited, c.n)
				} else if nodes[n]+1 < v {
					nodes[c.n] = v
				}
			}
		}
		unvisited = unvisited[1:]
	}

	return nodes
}

func main() {
	reg := strings.TrimSpace(util.GetFile("2018/Day20/input"))
	doors := map[move]bool{}

	traverse(reg, P{0, 0}, doors)
	grid, p := buildGrid(doors)
	util.PrintRuneGrid(grid)

	n := mostDoors(grid, p, make(map[P]bool), 0)

	nodes := dijk(grid, p)
	s := 0
	for _, doors := range nodes {
		if doors >= 1000 {
			s++
		}
	}
	fmt.Printf("most doors (part1): %d\n", n)
	fmt.Printf("rooms requiring at least 1000 doors (part2): %d\n", s)
}
