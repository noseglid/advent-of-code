package main

import (
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

func spiralSteps2(n int) {
	size := 65536
	grid := make([][]int, size)
	for r := range grid {
		grid[r] = make([]int, size)
	}

	set := func(x, y int) bool {
		type point struct{ x, y int }
		coords := []point{
			{-1, -1},
			{0, -1},
			{1, -1},
			{-1, 0},
			{1, 0},
			{-1, 1},
			{0, 1},
			{1, 1},
		}

		for _, c := range coords {
			px, py := x+c.x, y+c.y
			if px < 0 || px >= size || py < 0 || py >= size {
				continue
			}
			grid[y][x] += grid[py][px]
		}
		if grid[y][x] > 368078 {
			log.Printf("Part 2: First value larger: %d", grid[y][x])
			return true
		}
		return false
	}

	cx, cy := size/2, size/2
	grid[cy][cx] = 1
	for i := 3; i <= 512; i += 2 {
		cx++
		if set(cx, cy) {
			return
		}
		for a := 0; a < i-2; a++ {
			cy--
			if set(cx, cy) {
				return
			}
		}
		for a := 0; a < i-1; a++ {
			cx--
			if set(cx, cy) {
				return
			}
		}
		for a := 0; a < i-1; a++ {
			cy++
			if set(cx, cy) {
				return
			}
		}
		for a := 0; a < i-1; a++ {
			cx++
			if set(cx, cy) {
				return
			}
		}
	}
}

func spiralSteps(n int) int {
	sq := int(math.Ceil(math.Sqrt(float64(n))))
	if sq%2 == 0 {
		sq++
	}
	steps := sq*sq - n
	side := steps / (sq - 1)
	edgeSteps := steps - (sq-1)*side
	return (sq-1)/2 + util.Absolute(edgeSteps-(sq-1)/2)
}

func main() {
	log.Printf("Part 1: Distance to first square: %d", spiralSteps(368078))
	spiralSteps2(368078)
}

/*

// Next bound:
368449

Steps:
368449 -368078 = 371

y coord: 303
x coord = 303-371 = -68

303 + 68 = 371



*/
