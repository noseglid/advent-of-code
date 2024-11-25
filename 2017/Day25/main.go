package main

import "fmt"

type state int
type move int

const (
	A state = iota
	B
	C
	D
	E
	F
)

const (
	Left  move = -1
	Right move = 1
)

type action struct {
	write     int
	move      move
	nextState state
}

var m = map[state]map[int]action{
	A: {
		0: {1, Right, B},
		1: {0, Right, F},
	},
	B: {
		0: {0, Left, B},
		1: {1, Left, C},
	},
	C: {
		0: {1, Left, D},
		1: {0, Right, C},
	},
	D: {
		0: {1, Left, E},
		1: {1, Right, A},
	},
	E: {
		0: {1, Left, F},
		1: {0, Left, D},
	},
	F: {
		0: {1, Right, A},
		1: {0, Left, E},
	},
}

func main() {
	tape := map[int]int{}

	state := A
	steps := 12964419
	pos := 0

	for i := 0; i < steps; i++ {
		a := m[state][tape[pos]]
		tape[pos] = a.write
		pos += int(a.move)
		state = a.nextState
	}

	s := 0
	for _, v := range tape {
		s += v
	}

	fmt.Printf("diagnostics (part1): %d\n", s)
}
