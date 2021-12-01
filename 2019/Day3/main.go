package main

import (
	"log"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Dir rune

const (
	Up    = Dir('U')
	Right = Dir('R')
	Down  = Dir('D')
	Left  = Dir('L')
)

type step struct {
	dir    Dir
	length int
}

type point struct {
	x, y int
}

func parseRow(s string) []step {
	var r []step
	for _, op := range strings.Split(s, ",") {
		r = append(r, step{Dir(op[0]), util.MustAtoi(op[1:])})
	}
	return r
}

func deltas(s step) (int, int) {
	switch s.dir {
	case 'U':
		return 0, -1
	case 'R':
		return 1, 0
	case 'D':
		return 0, 1
	case 'L':
		return -1, 0
	default:
		panic("bad dir")
	}
}

func main() {
	input := util.GetFile("2019/Day3/input")

	rows := strings.Split(input, "\n")

	w1 := parseRow(rows[0])
	w2 := parseRow(rows[1])
	visited := map[point]int{}

	var p1 point
	stepsp1 := 0
	for _, s := range w1 {
		dx, dy := deltas(s)
		for i := 0; i < s.length; i++ {
			stepsp1++
			p1.x += dx
			p1.y += dy
			visited[p1] = stepsp1
		}
	}

	d := math.MaxInt32
	totalSteps := math.MaxInt32
	var p2 point
	stepsp2 := 0
	for _, s := range w2 {
		dx, dy := deltas(s)
		for i := 0; i < s.length; i++ {
			stepsp2++
			p2.x += dx
			p2.y += dy
			if stepsp1, ok := visited[p2]; ok {
				distance := int(math.Abs(float64(p2.x))) + int(math.Abs(float64(p2.y)))
				if distance < d {
					d = distance
				}

				if totalSteps > stepsp1+stepsp2 {
					totalSteps = stepsp1 + stepsp2
				}
			}
		}
	}

	log.Printf("min distance on crossing (part1): %d", d)
	log.Printf("min steps on crossing (part2): %d", totalSteps)
}
