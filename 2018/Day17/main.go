package main

import (
	"fmt"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type P util.Point

func boxed(grid util.Grid, x, y int) ([]P, bool) {
	r := []P{{x, y}}
	for ix := x - 1; grid.InBounds(ix, y); ix-- {
		if r := grid.Get(ix, y); r == '#' {
			break
		}
		if r := grid.Get(ix, y+1); r == '.' {
			return nil, false
		}
		r = append(r, P{ix, y})
	}

	for ix := x + 1; grid.InBounds(ix, y); ix++ {
		if r := grid.Get(ix, y); r == '#' {
			break
		}
		if r := grid.Get(ix, y+1); r == '.' {
			return nil, false
		}
		r = append(r, P{ix, y})
	}

	return r, true
}

func fill(grid util.Grid, x, y int) bool {
	if !grid.InBounds(x, y+1) {
		return true
	}
	grid.Set(x, y, '|')
	if r := grid.Get(x, y+1); r == '.' {
		fill(grid, x, y+1)
	}
	if r := grid.Get(x, y+1); r == '|' {
		return false
	}

	if ps, ok := boxed(grid, x, y); ok {
		for _, p := range ps {
			grid.Set(p.X, p.Y, '~')
		}
	}

	if r := grid.Get(x, y+1); r == '#' || r == '~' {
		if ir := grid.Get(x-1, y); ir == '.' {
			fill(grid, x-1, y)
		}
		if ir := grid.Get(x+1, y); ir == '.' {
			fill(grid, x+1, y)
		}
	}

	return true
}

func reachedTiles(grid util.Grid, miny, maxy int) (int, int) {
	flow, stationary := 0, 0
	util.Grid(grid).Each(func(x, y int, r rune) {
		if y >= miny && y <= maxy {
			if r == '~' {
				stationary++
			}
			if r == '|' {
				flow++
			}
		}
	})

	return flow, stationary
}

func main() {
	lines := util.GetFileStrings("2018/Day17/input")

	grid := make([][]rune, 2000)

	for y := range grid {
		grid[y] = make([]rune, 2000)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0
	for _, l := range lines {
		var s, v1, v2 int
		if _, err := fmt.Sscanf(l, "y=%d, x=%d..%d", &s, &v1, &v2); err == nil {
			miny, maxy, minx, maxx = util.Min(miny, s), util.Max(maxy, s), util.Min(minx, v1, v2), util.Max(maxx, v1, v2)
			for x := v1; x <= v2; x++ {
				grid[s][x] = '#'
			}
		}
		if _, err := fmt.Sscanf(l, "x=%d, y=%d..%d", &s, &v1, &v2); err == nil {
			miny, maxy, minx, maxx = util.Min(miny, v1, v2), util.Max(maxy, v1, v2), util.Min(minx, s), util.Max(maxx, s)
			for y := v1; y <= v2; y++ {
				grid[y][s] = '#'
			}
		}
	}

	fill(util.Grid(grid), 500, 0)

	flow, stationary := reachedTiles(util.Grid(grid), miny, maxy)
	fmt.Printf("Reachable tiles by water (part1): %d\n", flow+stationary)
	fmt.Printf("Stationary water (part2): %d\n", stationary)

}
