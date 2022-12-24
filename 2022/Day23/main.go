package main

import (
	"fmt"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type dir rune

const (
	North dir = 'N'
	South dir = 'S'
	East  dir = 'E'
	West  dir = 'W'
)

type elf struct {
	test     []dir
	position util.Point
	decided  *util.Point
}

func (e *elf) PonderMove(elves map[util.Point]*elf) {
	// fmt.Printf("%v: Pondering move\n", e.position)
	tt := []util.Point{
		{X: e.position.X - 1, Y: e.position.Y - 1}, {X: e.position.X + 0, Y: e.position.Y - 1}, {X: e.position.X + 1, Y: e.position.Y - 1},
		{X: e.position.X - 1, Y: e.position.Y + 0}, {X: e.position.X + 1, Y: e.position.Y + 0},
		{X: e.position.X - 1, Y: e.position.Y + 1}, {X: e.position.X + 0, Y: e.position.Y + 1}, {X: e.position.X + 1, Y: e.position.Y + 1},
	}

	allEmpty := true
	for _, t := range tt {
		if _, ok := elves[t]; ok {
			// fmt.Printf("not all are empty\n")
			allEmpty = false
		}
	}

	if allEmpty {
		return
	}

	canMove := func(p1, p2, p3 util.Point) bool {
		_, e1 := elves[p1]
		_, e2 := elves[p2]
		_, e3 := elves[p3]
		return !e1 && !e2 && !e3
	}

Outer:
	for _, d := range e.test {
		switch d {
		case North:
			if canMove(tt[0], tt[1], tt[2]) {
				// fmt.Printf("%v: moving north\n", e.position)
				e.decided = &util.Point{X: tt[1].X, Y: tt[1].Y}
				break Outer
			}
		case South:
			if canMove(tt[5], tt[6], tt[7]) {
				e.decided = &util.Point{X: tt[6].X, Y: tt[6].Y}
				break Outer
			}
		case West:
			if canMove(tt[0], tt[3], tt[5]) {
				e.decided = &util.Point{X: tt[3].X, Y: tt[3].Y}
				break Outer
			}
		case East:
			if canMove(tt[2], tt[4], tt[7]) {
				e.decided = &util.Point{X: tt[4].X, Y: tt[4].Y}
				break Outer
			}
		}
	}
}

func (e *elf) MakeMove(elves map[util.Point]*elf) {
	defer func() {
		e.test = []dir{e.test[1], e.test[2], e.test[3], e.test[0]}
	}()

	if e.decided == nil {
		return
	}

	for _, et := range elves {
		if e == et || et.decided == nil {
			continue
		}
		if *e.decided == *et.decided {
			// Another elf going here, don't move
			return
		}
	}
	// fmt.Printf("%v: Moving to %v\n", e.position, *e.decided)
	delete(elves, e.position)
	e.position = util.Point{X: e.decided.X, Y: e.decided.Y}
	elves[e.position] = e
	// e.decided = nil
}

func (e *elf) Reset() {
	e.decided = nil
}

func PrintElves(elves map[util.Point]*elf) {
	minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0

	for _, e := range elves {
		minx = util.Min(minx, e.position.X)
		maxx = util.Max(maxx, e.position.X)
		miny = util.Min(miny, e.position.Y)
		maxy = util.Max(maxy, e.position.Y)
	}

	for y := -5; y <= 20; y++ {
		for x := -5; x <= 20; x++ {
			if _, ok := elves[util.Point{X: x, Y: y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func EmptySlots(elves map[util.Point]*elf) int {
	minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0

	for _, e := range elves {
		minx = util.Min(minx, e.position.X)
		maxx = util.Max(maxx, e.position.X)
		miny = util.Min(miny, e.position.Y)
		maxy = util.Max(maxy, e.position.Y)
	}

	n := 0
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if _, ok := elves[util.Point{X: x, Y: y}]; !ok {
				n++
			}
		}
	}

	return n

}

func main() {
	input := util.GetFileRuneGrid("2022/Day23/input")

	elves := map[util.Point]*elf{}
	var list []*elf
	for y, r := range input {
		for x, c := range r {
			if c == '#' {
				e := &elf{position: util.Point{X: x, Y: y}, test: []dir{North, South, West, East}}
				elves[util.Point{X: x, Y: y}] = e
				list = append(list, e)
			}
		}
	}

	roundNoMoves := -1
	makeMove := false
	for rounds := 1; ; rounds++ {
		for _, e := range list {
			e.PonderMove(elves)
			if e.decided != nil {
				makeMove = true
			}
		}
		if !makeMove {
			roundNoMoves = rounds
			break
		}

		for _, e := range list {
			e.MakeMove(elves)
		}

		for _, e := range list {
			e.Reset()
		}
		if rounds == 10 {
			fmt.Printf("part1: %d\n", EmptySlots(elves))
		}
		makeMove = false

		fmt.Printf("Round %d\n", rounds)
		PrintElves(elves)
	}

	fmt.Printf("part2: %d\n", roundNoMoves)
}
