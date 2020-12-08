package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type op string

const (
	opAcc = op("acc")
	opJmp = op("jmp")
	opNop = op("nop")
)

type instruction struct {
	op    op
	value int
}

func parseInstruction(s string) instruction {
	parts := strings.Split(s, " ")
	return instruction{
		op:    op(parts[0]),
		value: util.MustAtoi(parts[1]),
	}
}

func runProgram(instructions []instruction) (int, bool) {
	visited := map[int]bool{}

	pc := 0
	acc := 0
	for {
		instr := instructions[pc]

		if v := visited[pc]; v {
			return acc, true
		}
		visited[pc] = true

		switch instr.op {
		case opAcc:
			acc += instr.value
			pc++
		case opJmp:
			pc += instr.value
		case opNop:
			pc++
		}

		if pc >= len(instructions) {
			break
		}

	}

	return acc, false
}

func copyProgram(instrs []instruction) []instruction {
	result := make([]instruction, len(instrs))
	for i, instr := range instrs {
		result[i] = instr
	}

	return result
}

func main() {
	s := util.FileScanner("2020/Day8/input", bufio.ScanLines)

	var instructions []instruction

	for s.Scan() {
		instructions = append(instructions, parseInstruction(s.Text()))
	}

	acc, infinite := runProgram(instructions)
	if !infinite {
		log.Fatal("not infinite for part1")
	}

	log.Printf("infinite loop detected with accumulator (part1): %d", acc)

	for i, instr := range instructions {
		doRun := false
		var program []instruction
		switch instr.op {
		case opJmp:
			program = copyProgram(instructions)
			program[i].op = opNop
			doRun = true
		case opNop:
			program = copyProgram(instructions)
			program[i].op = opJmp
			doRun = true
		}

		if doRun {
			if acc, infinite := runProgram(program); !infinite {
				log.Printf("program exited successfully with accumulator (part2): %d", acc)
				break
			}
		}
	}

}
