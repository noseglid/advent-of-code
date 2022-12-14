package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Element int

const (
	Air Element = iota
	Rock
	Sand
)

func (e Element) String() string {
	switch e {
	case Air:
		return "."
	case Rock:
		return "#"
	case Sand:
		return "o"
	}
	panic("bad element")
}

func parseCoords(s string) []util.Point {
	var res []util.Point
	for _, ss := range strings.Split(s, " -> ") {
		var x, y int
		if _, err := fmt.Sscanf(ss, "%d,%d", &x, &y); err != nil {
			panic(err)
		}
		res = append(res, util.Point{X: x, Y: y})

	}
	return res
}

func step(a1, a2 int) int {
	if a1 < a2 {
		return 1
	}
	return -1
}

func generateSandP1(grid [][]Element, source util.Point) bool {
	x, y := source.X, source.Y
	grid[y][x] = Sand
	for {
		if y+1 >= len(grid) {
			// No rest
			return false
		}
		if grid[y+1][x] == Air {
			grid[y][x] = Air
			grid[y+1][x] = Sand
			y++
			continue
		}
		if grid[y+1][x-1] == Air {
			grid[y][x] = Air
			grid[y+1][x-1] = Sand
			y++
			x--
			continue
		}
		if grid[y+1][x+1] == Air {
			grid[y][x] = Air
			grid[y+1][x+1] = Sand
			y++
			x++
			continue
		}

		// At rest
		return true
	}
}
func generateSandP2(grid [][]Element, floorY int, source util.Point) bool {
	x, y := source.X, source.Y
	grid[y][x] = Sand
	for {
		if y+1 == floorY {
			// on the floor, it's at rest
			return true
		}

		if grid[y+1][x] == Air {
			grid[y][x] = Air
			grid[y+1][x] = Sand
			y++
			continue
		}
		if grid[y+1][x-1] == Air {
			grid[y][x] = Air
			grid[y+1][x-1] = Sand
			y++
			x--
			continue
		}
		if grid[y+1][x+1] == Air {
			grid[y][x] = Air
			grid[y+1][x+1] = Sand
			y++
			x++
			continue
		}

		if source.X == x && source.Y == y {
			return false
		}

		// At rest
		return true
	}
}

func main() {

	lines := util.GetFileStrings("2022/Day14/input")

	grid := make([][]Element, 1000)
	gridp2 := make([][]Element, 1000)
	for i := range grid {
		grid[i] = make([]Element, 1000)
		gridp2[i] = make([]Element, 1000)
	}

	miny, maxy, minx, maxx := math.MaxInt, 0, math.MaxInt, 0
	for _, l := range lines {
		coords := parseCoords(l)
		for i := 1; i < len(coords); i++ {
			x1, y1, x2, y2 := coords[i-1].X, coords[i-1].Y, coords[i].X, coords[i].Y
			minx = util.Min(minx, x1, x2)
			maxx = util.Max(maxx, x1, x2)
			miny = util.Min(miny, y1, y2)
			maxy = util.Max(maxy, y1, y2)

			if x1 == x2 {
				for y := util.Min(y1, y2); y <= util.Max(y1, y2); y += step(util.Min(y1, y2), util.Max(y1, y2)) {
					grid[y][x1] = Rock
					gridp2[y][x1] = Rock
				}
			}
			if y1 == y2 {
				for x := util.Min(x1, x2); x <= util.Max(x1, x2); x += step(util.Min(x1, x2), util.Max(x1, x2)) {
					grid[y1][x] = Rock
					gridp2[y1][x] = Rock
				}
			}
		}
	}

	floor := maxy + 2
	for x := 0; x < 1000; x++ {
		gridp2[floor][x] = Rock
	}

	units := 0
	for {
		cameToRest := generateSandP1(grid, util.Point{X: 500, Y: 0})
		if !cameToRest {
			break
		}
		units++
	}

	unitsp2 := 0
	for {
		done := generateSandP2(gridp2, floor, util.Point{X: 500, Y: 0})
		unitsp2++
		if !done {
			break
		}
	}

	fmt.Printf("Units of sand which came to rest (part1): %d\n", units)
	fmt.Printf("Units of sand which came to rest (part2): %d\n", unitsp2)

}
