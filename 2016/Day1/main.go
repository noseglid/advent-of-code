package main

import (
	"log"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type p struct {
	x, y int
}

func absSum(a, b int) int {
	return int(math.Abs(float64(a))) + int(math.Abs(float64(b)))
}

func main() {
	input := util.GetFile("2016/Day1/input")

	visited := map[p]struct{}{}

	p2distance := 0

	facing := 0 // 0 - up, 1 - right, 2 - down, 3 - left
	x, y := 0, 0
	for _, s := range strings.Split(input, ", ") {
		steps := util.MustAtoi(strings.TrimSpace(s[1:]))
		switch s[0] {
		case 'L':
			facing = (facing - 1 + 4) % 4
		case 'R':
			facing = (facing + 1) % 4
		}

		dx, dy := 0, 0
		switch facing {
		case 0:
			dy = -1
		case 1:
			dx = 1
		case 2:
			dy = 1
		case 3:
			dx = -1
		}

		for i := 0; i < steps; i++ {
			x += dx
			y += dy

			if p2distance == 0 {
				if _, hasVisited := visited[p{x, y}]; hasVisited {
					p2distance = absSum(x, y)
				}
			}
			visited[p{x, y}] = struct{}{}
		}
	}

	log.Printf("taxi distance (part1): %d", absSum(x, y))
	log.Printf("taxi distance (part2): %d", p2distance)

}
