package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type op string

const (
	ophlf = op("hlf")
	optpl = op("tpl")
	opinc = op("inc")
	opjmp = op("jmp")
	opjie = op("jie")
	opjio = op("jio")
)

type instruction struct {
	op  op
	r   string
	val int
}

func parseInstruction(s string) instruction {
	fields := strings.Fields(s)
	switch op(fields[0]) {
	case ophlf, optpl, opinc:
		return instruction{op(fields[0]), fields[1], 0}
	case opjmp:
		return instruction{op(fields[0]), "", util.MustAtoi(fields[1])}
	case opjie, opjio:
		return instruction{op(fields[0]), fields[1][:len(fields[1])-1], util.MustAtoi(fields[2])}
	default:
		panic("invalid instruction")
	}
}

func run(instructions []instruction, registers map[string]int) {
	pc := 0
	for {
		if pc >= len(instructions) {
			break
		}

		instr := instructions[pc]
		switch instr.op {
		case ophlf:
			registers[instr.r] /= 2
			pc++
		case optpl:
			registers[instr.r] *= 3
			pc++
		case opinc:
			registers[instr.r] += 1
			pc++
		case opjmp:
			pc += instr.val
		case opjie:
			if registers[instr.r]%2 == 0 {
				pc += instr.val
			} else {
				pc++
			}
		case opjio:
			if registers[instr.r] == 1 {
				pc += instr.val
			} else {
				pc++
			}
		default:
			panic("bad instruction")
		}
	}
}

func main() {

	s := util.FileScanner("2015/Day23/input", bufio.ScanLines)
	var instructions []instruction

	for s.Scan() {
		instructions = append(instructions, parseInstruction(s.Text()))
	}

	registers := map[string]int{"a": 0, "b": 0}
	run(instructions, registers)
	log.Printf("reg b (part1): %d", registers["b"])

	registers = map[string]int{"a": 1, "b": 0}
	run(instructions, registers)
	log.Printf("reg b (part1): %d", registers["b"])
}
