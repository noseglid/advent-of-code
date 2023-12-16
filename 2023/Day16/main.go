package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type beam struct {
	row, col int
	dir      rune
}

func (b beam) String() string {
	return fmt.Sprintf("(%d,%d):%c", b.row, b.col, b.dir)
}

func step(grid [][]rune, beams []beam) []beam {
	var next []beam
	for _, b := range beams {
		switch grid[b.row][b.col] {
		case '-':
			switch b.dir {
			case 'u', 'd':
				next = append(next, beam{row: b.row, col: b.col + 1, dir: 'r'})
				b.dir = 'l'
				b.col--
			case 'l':
				b.col--
			case 'r':
				b.col++
			}
		case '|':
			switch b.dir {
			case 'l', 'r':
				next = append(next, beam{row: b.row - 1, col: b.col, dir: 'u'})
				b.dir = 'd'
				b.row++
			case 'u':
				b.row--
			case 'd':
				b.row++
			}
		case '\\':
			switch b.dir {
			case 'r':
				b.dir = 'd'
				b.row++
			case 'l':
				b.dir = 'u'
				b.row--
			case 'd':
				b.dir = 'r'
				b.col++
			case 'u':
				b.dir = 'l'
				b.col--
			}
		case '/':
			switch b.dir {
			case 'r':
				b.dir = 'u'
				b.row--
			case 'l':
				b.dir = 'd'
				b.row++
			case 'd':
				b.dir = 'l'
				b.col--
			case 'u':
				b.dir = 'r'
				b.col++
			}
		case '.':
			switch b.dir {
			case 'r':
				b.col++
			case 'l':
				b.col--
			case 'u':
				b.row--
			case 'd':
				b.row++
			}
		}
		next = append(next, b)
	}
	return next
}

func cleanBeams(beams []beam, grid [][]rune, seen map[beam]struct{}) []beam {
	var r []beam
	for _, b := range beams {
		if b.col < 0 || b.col >= len(grid[0]) || b.row < 0 || b.row >= len(grid) {
			continue
		}

		if _, ok := seen[b]; ok {
			continue
		}

		r = append(r, b)
	}
	return r
}

func beamHash(beams []beam) string {
	var sb strings.Builder
	for _, b := range beams {
		sb.WriteString(b.String())
	}
	return sb.String()
}

func cmp(lhs, rhs beam) int {
	if lhs.row != rhs.row {
		return lhs.row - rhs.row
	}
	if lhs.col != rhs.col {
		return lhs.col - rhs.col
	}
	return int(lhs.dir - rhs.dir)
}

func countEn(en [][]bool) int {
	s := 0
	for _, r := range en {
		for _, c := range r {
			if c {
				s++
			}
		}
	}
	return s
}

func simulate(grid [][]rune, beams []beam) int {
	en := make([][]bool, len(grid))
	for i, g := range grid {
		en[i] = make([]bool, len(g))
	}

	seen := map[beam]struct{}{}
	for len(beams) > 0 {
		for _, b := range beams {
			if b.row < 0 || b.row >= len(en) || b.col < 0 || b.col >= len(en[b.row]) {
				continue
			}
			en[b.row][b.col] = true
		}
		beams = step(grid, beams)
		beams = cleanBeams(beams, grid, seen)
		for _, b := range beams {
			seen[b] = struct{}{}
		}
		slices.SortFunc(beams, cmp)
	}

	return countEn(en)
}

func main() {
	grid := util.GetFileRuneGrid("2023/Day16/input")
	beams := []beam{{row: 0, col: 0, dir: 'r'}}
	fmt.Printf("energized (part 1): %d\n", simulate(grid, beams))

	m := 0
	for row := range []int{0, len(grid) - 1} {
		dir := 'd'
		if row != 0 {
			dir = 'u'
		}
		for col := range grid[row] {
			beams = []beam{{row: row, col: col, dir: dir}}
			v := simulate(grid, beams)
			if v > m {
				m = v
			}
		}
	}
	for col := range []int{0, len(grid[0]) - 1} {
		dir := 'r'
		if col != 0 {
			dir = 'l'
		}
		for row := range grid {
			beams = []beam{{row: row, col: col, dir: dir}}
			v := simulate(grid, beams)
			if v > m {
				m = v
			}
		}
	}
	fmt.Printf("energized max (part 2): %d\n", m)
}
