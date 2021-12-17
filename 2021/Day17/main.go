package main

import (
	"fmt"
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type area struct {
	x1, x2 int
	y1, y2 int
}

func (a area) Contains(x, y int) bool {
	return x >= a.x1 && x <= a.x2 && y >= a.y1 && y <= a.y2
}

func (a area) Overshoot(x, y int) bool {
	return y < a.y1
}

func FireProbe(target area, vx, vy int) (int, bool) {
	x, y := 0, 0

	ymax := y

	for {
		x += vx
		y += vy
		if vx > 0 {
			vx--
		}
		vy -= 1

		if y > ymax {
			ymax = y
		}

		if target.Contains(x, y) {
			return ymax, true
		}

		if target.Overshoot(x, y) {
			return 0, false
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := "2021/Day17/input"
	f := util.GetFile(input)

	var target area

	if _, err := fmt.Sscanf(f, "target area: x=%d..%d, y=%d..%d", &target.x1, &target.x2, &target.y1, &target.y2); err != nil {
		panic(err)
	}

	xs, xe := min(0, target.x1), max(0, 2*target.x2)
	ys, ye := target.y1, max(-2*target.y1, 2*target.y1)

	max := math.MinInt
	nhits := 0

	for x := xs; x < xe; x++ {
		for y := ys; y < ye; y++ {
			if m, hit := FireProbe(target, x, y); hit {
				nhits++
				if m > max {
					max = m
				}
			}
		}
	}

	log.Printf("Part 1: highest y position: %d", max)
	log.Printf("Part 2: n initial velocities to hit: %d", nhits)

}
