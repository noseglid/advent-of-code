package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type Rule interface {
	fmt.Stringer

	Replace(block [][]rune) ([][]rune, bool)
}

type Two2Three struct {
	source [][]rune
	target [][]rune
}

func (t Two2Three) String() string {
	return fmt.Sprintf("%c%c/%c%c => %c%c%c/%c%c%c/%c%c%c",
		t.source[0][0], t.source[0][1],
		t.source[1][0], t.source[1][1],
		t.target[0][0], t.target[0][1], t.target[0][2],
		t.target[1][0], t.target[1][1], t.target[1][2],
		t.target[2][0], t.target[2][1], t.target[2][2],
	)
}

func (t Two2Three) Replace(block [][]rune) ([][]rune, bool) {
	if len(block) != 2 {
		return nil, false
	}
	dd := make([][]rune, len(t.source))
	for i := range t.source {
		dd[i] = make([]rune, len(t.source[i]))
		copy(dd[i], t.source[i])
	}

	for i := 0; i < 4; i++ {
		if blocksEq(block, dd) {
			return t.target, true
		}
		dd = util.Rotate2D(dd)
	}

	return nil, false
}

type Three2Four struct {
	source [][]rune
	target [][]rune
}

func blocksEq(b1, b2 [][]rune) bool {
	for y := 0; y < len(b1); y++ {
		for x := 0; x < len(b1[y]); x++ {
			if b1[y][x] != b2[y][x] {
				return false
			}
		}
	}
	return true
}

func (t Three2Four) String() string {
	return fmt.Sprintf("%c%c%c/%c%c%c/%c%c%c => %c%c%c%c/%c%c%c%c/%c%c%c%c/%c%c%c%c",
		t.source[0][0], t.source[0][1], t.source[0][2],
		t.source[1][0], t.source[1][1], t.source[1][2],
		t.source[2][0], t.source[2][1], t.source[2][2],
		t.target[0][0], t.target[0][1], t.target[0][2], t.target[0][3],
		t.target[1][0], t.target[1][1], t.target[1][2], t.target[1][3],
		t.target[2][0], t.target[2][1], t.target[2][2], t.target[2][3],
		t.target[3][0], t.target[3][1], t.target[3][2], t.target[3][3],
	)
}
func (t Three2Four) Replace(block [][]rune) ([][]rune, bool) {
	if len(block) != 3 {
		return nil, false
	}
	dd := make([][]rune, len(t.source))
	for i := range t.source {
		dd[i] = make([]rune, len(t.source[i]))
		copy(dd[i], t.source[i])
	}

	for i := 0; i < 4; i++ {
		if blocksEq(block, dd) {
			return t.target, true
		}
		dd = util.Rotate2D(dd)
	}

	dd = util.Flip2D(dd)
	for i := 0; i < 4; i++ {
		if blocksEq(block, dd) {
			return t.target, true
		}
		dd = util.Rotate2D(dd)
	}

	return nil, false
}

var twoToThreeRe = regexp.MustCompile(`(\.|#)(\.|#)/(\.|#)(\.|#) => (\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#)`)

var threeToFourRe = regexp.MustCompile(`(\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#) => (\.|#)(\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#)(\.|#)/(\.|#)(\.|#)(\.|#)(\.|#)`)

func countPixels(grid [][]rune) int {
	var n int
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				n++
			}
		}
	}
	return n
}

func ParseRule(s string) Rule {
	m1 := twoToThreeRe.FindStringSubmatch(s)
	if m1 != nil {
		return Two2Three{
			source: [][]rune{
				0: {rune(m1[1][0]), rune(m1[2][0])},
				1: {rune(m1[3][0]), rune(m1[4][0])},
			},
			target: [][]rune{
				0: {rune(m1[5][0]), rune(m1[6][0]), rune(m1[7][0])},
				1: {rune(m1[8][0]), rune(m1[9][0]), rune(m1[10][0])},
				2: {rune(m1[11][0]), rune(m1[12][0]), rune(m1[13][0])},
			},
		}
	}

	m2 := threeToFourRe.FindStringSubmatch(s)
	if m2 != nil {
		return Three2Four{
			source: [][]rune{
				0: {rune(m2[1][0]), rune(m2[2][0]), rune(m2[3][0])},
				1: {rune(m2[4][0]), rune(m2[5][0]), rune(m2[6][0])},
				2: {rune(m2[7][0]), rune(m2[8][0]), rune(m2[9][0])},
			},
			target: [][]rune{
				0: {rune(m2[10][0]), rune(m2[11][0]), rune(m2[12][0]), rune(m2[13][0])},
				1: {rune(m2[14][0]), rune(m2[15][0]), rune(m2[16][0]), rune(m2[17][0])},
				2: {rune(m2[18][0]), rune(m2[19][0]), rune(m2[20][0]), rune(m2[21][0])},
				3: {rune(m2[22][0]), rune(m2[23][0]), rune(m2[24][0]), rune(m2[25][0])},
			},
		}
	}

	panic("no rule match")
}

func combine(blocks [][][]rune) [][]rune {
	if len(blocks[0])%3 == 0 {
		//3x3 blocks
		s := int(math.Sqrt(float64(len(blocks))))
		g := make([][]rune, s*3)
		for i := 0; i < len(g); i++ {
			g[i] = make([]rune, s*3)
		}

		for i, b := range blocks {
			for iy := range b {
				for ix := range b[iy] {
					nx := (i%s)*3 + ix
					ny := (i/s)*3 + iy
					g[ny][nx] = b[iy][ix]
				}
			}
		}
		return g
	}
	if len(blocks[0])%4 == 0 {
		// 4x4 blocks, sqrt
		s := int(math.Sqrt(float64(len(blocks))))
		g := make([][]rune, s*4)
		for i := 0; i < len(g); i++ {
			g[i] = make([]rune, s*4)
		}

		for i, b := range blocks {
			for iy := range b {
				for ix := range b[iy] {
					nx := (i%s)*4 + ix
					ny := (i/s)*4 + iy
					g[ny][nx] = b[iy][ix]
				}
			}
		}
		return g
	}

	panic("combine: size not handled")
}

func iterate(grid [][]rune, rules []Rule) [][]rune {
	var newBlocks [][][]rune
	if len(grid)%2 == 0 {
		t := len(grid) / 2
		for y := 0; y < t; y++ {
			for x := 0; x < t; x++ {
				block := [][]rune{
					grid[y*2][x*2 : x*2+2],
					grid[y*2+1][x*2 : x*2+2],
				}
				ruleFound := false
				for _, rule := range rules {
					nb, ok := rule.Replace(block)
					if !ok {
						continue
					}

					newBlocks = append(newBlocks, nb)
					ruleFound = true
				}
				if !ruleFound {
					panic("iterate: 2x2: no rule found")
				}

			}
		}
	} else if len(grid)%3 == 0 {
		t := len(grid)/3 + 1
		for y := 0; y < t-1; y++ {
			for x := 0; x < t-1; x++ {
				block := [][]rune{
					grid[y*3][x*3 : x*3+3],
					grid[y*3+1][x*3 : x*3+3],
					grid[y*3+2][x*3 : x*3+3],
				}

				ruleFound := false
				for _, rule := range rules {
					nb, ok := rule.Replace(block)
					if ok {
						newBlocks = append(newBlocks, nb)
						ruleFound = true
						break
					}
				}

				if !ruleFound {
					panic("iterate: 3x3: no rule found")
				}
			}
		}
	} else {
		panic("bad size")
	}

	return combine(newBlocks)
}

func main() {
	lines := util.GetFileStrings("2017/Day21/input")

	rules := []Rule{}

	for _, l := range lines {
		rules = append(rules, ParseRule(l))
	}

	grid := [][]rune{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}
	// grid := [][]rune{
	// 	{'.', '#', '.', '.', '#', '.'},
	// 	{'.', '.', '#', '.', '.', '#'},
	// 	{'#', '#', '#', '#', '#', '#'},
	// 	{'.', '#', '.', '.', '#', '.'},
	// 	{'.', '.', '#', '.', '.', '#'},
	// 	{'#', '#', '#', '#', '#', '#'},
	// }

	iterations := 18

	util.PrintRuneGrid(grid)
	fmt.Printf("-----\n")
	for i := 0; i < iterations; i++ {
		grid = iterate(grid, rules)
		// util.PrintRuneGrid(grid)
		// fmt.Printf("-----\n")
	}

	fmt.Printf("Pixels on after %d iterations (part1): %d\n", iterations, countPixels(grid))

}
