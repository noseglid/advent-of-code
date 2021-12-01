package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func operate(pos int, program []int) (int, bool) {
	switch program[pos] {
	case 1:
		// addition
		program[program[pos+3]] = program[program[pos+1]] + program[program[pos+2]]
		return pos + 4, false
	case 2:
		// mult
		program[program[pos+3]] = program[program[pos+1]] * program[program[pos+2]]
		return pos + 4, false
	case 99:
		// done
		return 0, true

	default:
		log.Fatalf("bad op at %d: %d", pos, program[pos])
		panic("")
	}
}

func copyProgram(program []int) []int {
	r := make([]int, len(program))
	for i, n := range program {
		r[i] = n
	}
	return r
}

func runProgram(program []int) int {
	pc := 0
	for {
		var done bool
		pc, done = operate(pc, program)
		if done {
			break
		}
	}

	return program[0]
}

func main() {
	s := util.GetFile("2019/Day2/input")

	var program []int

	for _, op := range strings.Split(s, ",") {
		program = append(program, util.MustAtoi(strings.TrimSpace(op)))
	}

	p1program := copyProgram(program)
	p1program[1] = 12
	p1program[2] = 2

	log.Printf("initial run (part1): %d", runProgram(p1program))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			test := copyProgram(program)
			test[1] = noun
			test[2] = verb
			if v := runProgram(test); v == 19690720 {
				log.Printf("input with determinate output (part2): %d", 100*noun+verb)
			}
		}
	}

}
