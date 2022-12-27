package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func setBits(i byte) (int, []bool) {
	n := 0
	var r []bool
	for s := 0; s < 8; s++ {
		n += int((i >> s) & 1)
		r = append([]bool{(i>>s)&1 == 1}, r...)
	}
	return n, r
}

func PrintGrid(grid [][]bool) {
	for _, r := range grid {
		for _, c := range r {
			if c {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

type region struct {
	points []util.Point
}

func (r region) Belongs(p util.Point) bool {
	for _, pp := range r.points {
		if p == pp {
			return true
		}
	}
	return false
}

func (r *region) AddPoints(p util.Point, grid [][]bool) []util.Point {
	if !grid[p.Y][p.X] {
		// Not set, skip
		return []util.Point{}
	}
	if util.Contains(r.points, p) {
		return []util.Point{}
	}
	added := []util.Point{p}
	r.points = append(r.points, p)

	if p.Y > 0 {
		added = append(added, r.AddPoints(util.Point{Y: p.Y - 1, X: p.X}, grid)...)
	}
	if p.Y < len(grid)-1 {
		added = append(added, r.AddPoints(util.Point{Y: p.Y + 1, X: p.X}, grid)...)
	}
	if p.X > 0 {
		added = append(added, r.AddPoints(util.Point{Y: p.Y, X: p.X - 1}, grid)...)
	}
	if p.X < len(grid[p.Y])-1 {
		added = append(added, r.AddPoints(util.Point{Y: p.Y, X: p.X + 1}, grid)...)
	}
	return added
}

func main() {
	// input := "flqrgnkx"
	input := "oundnydw"

	grid := make([][]bool, 128)
	for y := range grid {
		grid[y] = make([]bool, 128)
	}
	used := 0
	for y := 0; y < 128; y++ {
		h := KnotHash(fmt.Sprintf("%s-%d", input, y))
		for j, b := range h {
			nn, bits := setBits(b)
			used += nn

			for x, ss := range bits {
				grid[y][j*8+x] = ss
			}
		}
	}
	fmt.Printf("used (part1): %d\n", used)

	var regions []region
	var allPoints []util.Point
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if grid[y][x] {
				allPoints = append(allPoints, util.Point{X: x, Y: y})
			}
		}
	}

	for len(allPoints) > 0 {
		p := allPoints[0]
		allPoints = allPoints[1:]
		r := region{}
		added := r.AddPoints(p, grid)
		for _, a := range added {
			allPoints, _ = util.RemoveByValue(allPoints, a)
		}
		regions = append(regions, r)
		// fmt.Printf("added region with %v\n", r.points)
	}

	fmt.Printf("Number of regions (part2): %d\n", len(regions))

}
