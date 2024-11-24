package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type P struct {
	x, y int
}

type Dir int

const (
	Up    Dir = 0
	Right Dir = 1
	Down  Dir = 2
	Left  Dir = 3
)

type NodeState int

const (
	Clean    NodeState = 0
	Weakened NodeState = 1
	Infected NodeState = 2
	Flagged  NodeState = 3
)

type Virus struct {
	pos P
	dir Dir
}

func (v *Virus) Turn(d Dir) {
	switch d {
	case Left:
		v.dir = (v.dir + 3) % 4
	case Right:
		v.dir = (v.dir + 1) % 4
	default:
		panic("bad turn dir")
	}
}

func (v *Virus) Step() {
	switch v.dir {
	case Up:
		v.pos.y -= 1
	case Right:
		v.pos.x += 1
	case Down:
		v.pos.y += 1
	case Left:
		v.pos.x -= 1
	}
}

func main() {
	m := map[P]bool{}
	m2 := map[P]NodeState{}
	grid := util.GetFileRuneGrid("2017/Day22/input")
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				m[P{x: x - len(grid[0])/2, y: y - len(grid)/2}] = true
				m2[P{x: x - len(grid[0])/2, y: y - len(grid)/2}] = Infected
			}
		}
	}
	v := Virus{}

	bursts := 10000
	nInfections := 0
	for i := 0; i < bursts; i++ {
		if m[v.pos] {
			v.Turn(Right)
		} else {
			v.Turn(Left)
		}
		m[v.pos] = !m[v.pos]
		if m[v.pos] {
			nInfections++
		}
		v.Step()
	}
	fmt.Printf("bursts causing infections (part1): %d\n", nInfections)

	v2 := Virus{}
	burstsp2 := 10000000
	nInfectionsp2 := 0
	for i := 0; i < burstsp2; i++ {
		switch m2[v2.pos] {
		case Clean:
			v2.Turn(Left)
			m2[v2.pos] = Weakened
		case Weakened:
			// no turn
			m2[v2.pos] = Infected
			nInfectionsp2++
		case Infected:
			v2.Turn(Right)
			m2[v2.pos] = Flagged
		case Flagged:
			v2.Turn(Left)
			v2.Turn(Left)
			m2[v2.pos] = Clean
		}

		v2.Step()
	}

	fmt.Printf("bursts causing infections (part2): %d\n", nInfectionsp2)

}
