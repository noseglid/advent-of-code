package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point

func manhattan(p1, p2 P) int {
	return util.Absolute(p1.X-p2.X) + util.Absolute(p1.Y-p2.Y)
}

func save(path []P, c int) int {
	n := 0
	steps := map[P]int{}
	for i, p := range path {
		steps[p] = i
	}

	for i, p := range path {
		for _, np := range path[i+1:] {
			m := manhattan(p, np)
			if m <= c && steps[np]-steps[p]-m >= 100 {
				n++
			}
		}
	}

	return n
}

func main() {
	grid := util.ParseRuneGrid(util.GetFileStrings("2024/Day20/input"))
	sx, sy := grid.Find('S')
	ex, ey := grid.Find('E')
	_, p, _ := grid.ShortestPath(P{X: sx, Y: sy}, P{X: ex, Y: ey}, func(x, y int, r rune) bool {
		return r == '.' || r == 'S' || r == 'E'
	})

	fmt.Printf("Cheats saving at least 100 ps with 2ps cheat (part1): %d\n", save(p, 2))
	fmt.Printf("Cheats saving at least 100 ps with 20ps cheat (part2): %d\n", save(p, 20))

}
