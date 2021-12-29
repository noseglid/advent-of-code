package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type instruction interface {
	exec(*program)
}

type program struct {
	registers    map[string]int
	pc           int
	instructions []instruction
}

func (p program) getValue(vr string) int {
	v, err := strconv.Atoi(vr)
	if err == nil {
		return v
	}
	return p.registers[vr]
}

func (p *program) run() {
	for {
		if p.pc >= len(p.instructions) {
			return
		}
		// fmt.Printf("pc=%d, reg=%v, instr: %#v\n", p.pc, p.registers, p.instructions[p.pc])
		p.instructions[p.pc].exec(p)
	}
}

type cpy struct {
	x, y string
}

func (c cpy) exec(p *program) {
	p.registers[c.y] = p.getValue(c.x)
	p.pc++
}

type inc struct {
	x string
}

func (i inc) exec(p *program) {
	p.registers[i.x]++
	p.pc++
}

type dec struct {
	x string
}

func (d dec) exec(p *program) {
	p.registers[d.x]--
	p.pc++
}

type jnz struct {
	x, y string
}

func (c jnz) exec(p *program) {
	if p.getValue(c.x) == 0 {
		p.pc++
		return
	}

	p.pc += p.getValue(c.y)
}

func parseInstruction(s string) instruction {

	parts := strings.Split(s, " ")
	switch parts[0] {
	case "cpy":
		return &cpy{parts[1], parts[2]}
	case "inc":
		return &inc{parts[1]}
	case "dec":
		return &dec{parts[1]}
	case "jnz":
		return &jnz{parts[1], parts[2]}
	}

	panic("bad instr")
}

func main() {
	input := "2016/Day12/input"
	lines := util.GetFileStrings(input)

	instructions := make([]instruction, 0, len(lines))
	for _, l := range lines {
		instructions = append(instructions, parseInstruction(l))
	}

	p := program{
		registers:    map[string]int{},
		pc:           0,
		instructions: instructions,
	}
	p.run()
	log.Printf("Part 1: Register a=%d", p.registers["a"])

	p2 := program{
		registers:    map[string]int{"c": 1},
		pc:           0,
		instructions: instructions,
	}
	p2.run()
	log.Printf("Part 2: Register a=%d", p2.registers["a"])

}
