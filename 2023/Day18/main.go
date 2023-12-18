package main

import (
	"fmt"
	"strconv"

	"github.com/noseglid/advent-of-code/util"
)

type instr struct {
	dir   rune
	steps int
}

func (i instr) String() string {
	return fmt.Sprintf("%c %d", i.dir, i.steps)
}

type filled struct {
	isFilled bool
	isEdge   bool
	rgb      string
}

func parseRGBInstr(l string) instr {
	i := instr{}
	steps, _ := strconv.ParseInt(l[0:5], 16, 64)
	i.steps = int(steps)
	switch l[5] {
	case '0':
		i.dir = 'R'
	case '1':
		i.dir = 'D'
	case '2':
		i.dir = 'L'
	case '3':
		i.dir = 'U'
	}

	return i
}

func parseLine(l string) (instr, instr) {
	i := instr{}
	var rgb string
	fmt.Sscanf(l, "%c %d (#%s)", &i.dir, &i.steps, &rgb)
	return i, parseRGBInstr(rgb)
}

func dimensions(instrs []instr) (int, int, int, int) {
	minw, minh, mw, mh, cw, ch := 0, 0, 0, 0, 0, 0
	for _, i := range instrs {
		switch i.dir {
		case 'U':
			ch -= i.steps
		case 'R':
			cw += i.steps
		case 'D':
			ch += i.steps
		case 'L':
			cw -= i.steps
		}

		if cw < minw {
			minw = cw
		}
		if ch < minh {
			minh = ch
		}
		if cw > mw {
			mw = cw
		}
		if ch > mh {
			mh = ch
		}
	}
	return minw, minh, mw + 1, mh + 1
}

func printGrid(grid [][]filled) {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col].isFilled {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func isInside(grid [][]filled, row, col int) bool {
	if grid[row][col].isEdge {
		return false
	}
	nEdge := 0

	inEdge := false
	var enterRune rune = 0
	for c := col; c < len(grid[row]); c++ {
		if !inEdge && grid[row][c].isEdge {
			// entered edge
			inEdge = true
			if row > 0 && grid[row-1][c].isEdge && c < len(grid[row])-2 && grid[row][c+1].isEdge {
				enterRune = 'B'
			} else if row < len(grid)-1 && grid[row+1][c].isEdge && c < len(grid[row])-2 && grid[row][c+1].isEdge {
				enterRune = 'T'
			} else {
				enterRune = 'M'
			}
		} else if inEdge && !grid[row][c].isEdge {
			// we exited edge
			inEdge = false
			if enterRune == 'M' {
				nEdge++
			} else {
				var exitRune rune
				if row > 0 && grid[row-1][c-1].isEdge && c > 0 && grid[row][c-1].isEdge {
					exitRune = 'B'
				} else if row < len(grid)-1 && grid[row+1][c-1].isEdge && c > 0 && grid[row][c-1].isEdge {
					exitRune = 'T'
				}

				if enterRune != exitRune {
					nEdge++
				}
			}
		}
	}

	if inEdge {
		nEdge++
	}

	return nEdge%2 == 1
}

func fill(grid [][]filled) {
	for row := range grid {
		// row := 159
		for col := range grid[row] {
			// col := 0
			if isInside(grid, row, col) {
				grid[row][col] = filled{
					isFilled: true,
					isEdge:   false,
					rgb:      "",
				}
			}
		}
	}
}

func countFilled(grid [][]filled) int {
	n := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col].isFilled {
				n++
			}
		}
	}
	return n
}

func doDig(instrs []instr) int {
	minWidth, minHeight, maxWidth, maxHeight := dimensions(instrs)
	grid := make([][]filled, maxHeight-minHeight)
	for i := range grid {
		grid[i] = make([]filled, maxWidth-minWidth)
	}
	crow, ccol := -minHeight, -minWidth
	for _, i := range instrs {
		switch i.dir {
		case 'U':
			for j := 0; j < i.steps; j++ {
				grid[crow][ccol] = filled{isFilled: true, isEdge: true}
				crow--
			}
		case 'R':
			for j := 0; j < i.steps; j++ {
				grid[crow][ccol] = filled{isFilled: true, isEdge: true}
				ccol++
			}
		case 'D':
			for j := 0; j < i.steps; j++ {
				grid[crow][ccol] = filled{isFilled: true, isEdge: true}
				crow++
			}
		case 'L':
			for j := 0; j < i.steps; j++ {
				grid[crow][ccol] = filled{isFilled: true, isEdge: true}
				ccol--
			}
		}
	}

	fill(grid)
	return countFilled(grid)
}

func getCoords(instrs []instr) ([]util.Point, int) {
	crow, ccol, perimiter := 0, 0, 0
	var ps []util.Point
	for j := 0; j < len(instrs); j++ {
		i := instrs[j]
		switch i.dir {
		case 'U':
			crow -= i.steps
		case 'R':
			ccol += i.steps
		case 'D':
			crow += i.steps
		case 'L':
			ccol -= i.steps
		}
		perimiter += i.steps
		ps = append(ps, util.Point{X: crow, Y: ccol})
	}

	return ps, perimiter
}

func shoelace(coords []util.Point) int {
	s := 0
	for i := 0; i < len(coords); i++ {
		c1, c2 := coords[i], coords[(i+1)%len(coords)]
		s += (c2.X*c1.Y - c1.X*c2.Y)
	}

	return s
}

func main() {

	lines := util.GetFileStrings("2023/Day18/input")

	var instrs []instr
	var instrsp2 []instr

	for _, l := range lines {
		p1, p2 := parseLine(l)
		instrs = append(instrs, p1)
		instrsp2 = append(instrsp2, p2)
	}
	fmt.Printf("Number dug out (part1): %d\n", doDig(instrs))
	coordsp1, perimiterp1 := getCoords(instrs)
	sp1 := shoelace(coordsp1)
	fmt.Printf("Number dug out (part1) w. shoelace: %d\n", sp1/2+perimiterp1/2+1)

	coords, perimiter := getCoords(instrsp2)
	s := shoelace(coords)

	fmt.Printf("Number dug out (part2): %d\n", s/2+perimiter/2+1)
}
