package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type octupus struct {
	energy   int
	didFlash bool
}

func (o *octupus) increaseLevel() {
	o.energy += 1
}

func (o *octupus) dispenseEnergy(x, y int, grid [10][10]*octupus) {
	if o.energy < 10 || o.didFlash {
		return
	}
	o.didFlash = true

	type p struct{ x, y int }
	dirs := []p{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, p := range dirs {
		xx := x + p.x
		yy := y + p.y
		if xx < 0 || xx >= 10 || yy < 0 || yy >= 10 {
			continue
		}

		o := grid[yy][xx]
		o.increaseLevel()
		o.dispenseEnergy(xx, yy, grid)
	}
}

func (o *octupus) reset() bool {
	r := o.didFlash
	if o.didFlash {
		o.energy = 0
		o.didFlash = false
	}
	return r
}

func newOctupus(level int) *octupus {
	return &octupus{
		energy: level,
	}
}

func iterate(grid [10][10]*octupus, fn func(x, y int, o *octupus)) {
	for y := range grid {
		for x := range grid[y] {
			fn(x, y, grid[y][x])
		}
	}
}

func main() {
	input := "2021/Day11/input"

	var grid [10][10]*octupus

	for i, l := range util.GetFileStrings(input) {
		for j, c := range l {
			grid[i][j] = newOctupus(util.MustAtoi(string(c)))
		}
	}

	p1done := false
	p2done := false
	step := 0
	flashes := 0
	for !p1done || !p2done {
		step++
		flashesInStep := 0
		iterate(grid, func(x, y int, o *octupus) {
			o.increaseLevel()
		})

		iterate(grid, func(x, y int, o *octupus) {
			o.dispenseEnergy(x, y, grid)
		})

		iterate(grid, func(x, y int, o *octupus) {
			if o.reset() {
				flashesInStep++
				flashes++
			}
		})

		if step == 100 {
			p1done = true
			log.Printf("Part 1: Flashes after 100 steps: %d", flashes)
		}

		if flashesInStep == 100 {
			p2done = true
			log.Printf("Part 2: Step that resulted in all flashes: %d", step)
		}
	}
}
