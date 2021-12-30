package main

import (
	"log"
)

type point struct{ x, y int }

const nn = 1352 // input

func setBits(n int) int {
	s := 0
	for n > 0 {
		if n&1 == 1 {
			s++
		}
		n = n >> 1
	}
	return s
}

func isWall(p point) bool {
	a := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + nn
	return setBits(a)%2 != 0
}

func moves(o point) []point {
	return []point{
		{o.x - 1, o.y},
		{o.x + 1, o.y},
		{o.x, o.y - 1},
		{o.x, o.y + 1},
	}
}

func minSteps(p, t point) int {
	steps := 0
	currentPos := []point{p}
	seen := map[point]bool{}
	for {
		steps++
		nextPos := []point{}

		for _, c := range currentPos {
			for _, o := range moves(c) {
				if o == t {
					return steps
				}
				if o.x < 0 || o.y < 0 || isWall(o) || seen[o] {
					continue
				}
				nextPos = append(nextPos, o)
				seen[o] = true
			}
			currentPos = nextPos

		}
		if steps == 50 {
			log.Printf("Part 2: distinct locations after 50 steps: %d", len(seen))
		}
	}
}

func main() {
	log.Printf("Part 1: minimum steps to reach target: %d", minSteps(point{1, 1}, point{31, 39}))

}
