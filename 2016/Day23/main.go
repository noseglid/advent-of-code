package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type instruction interface {
	exec(*program)
	tgl() instruction
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
		p.instructions[p.pc].exec(p)
	}
}

type cpy struct {
	x, y string
}

func isScalar(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		// successfully converted to int
		return true
	}
	return false
}

func (c cpy) exec(p *program) {
	defer func() { p.pc++ }()
	if isScalar(c.y) {
		// destination is a value (due to toggle), skip
		return
	}
	p.registers[c.y] = p.getValue(c.x)
}

func (c cpy) tgl() instruction {
	return &jnz{c.x, c.y}
}

type inc struct {
	x string
}

func (i inc) exec(p *program) {
	if isScalar(i.x) {
		return
	}

	defer func() { p.pc++ }()
	p.registers[i.x]++
}

func (i inc) tgl() instruction {
	return &dec{i.x}
}

type dec struct {
	x string
}

func (d dec) exec(p *program) {
	defer func() { p.pc++ }()
	if isScalar(d.x) {
		return
	}
	p.registers[d.x]--
}

func (d dec) tgl() instruction {
	return &inc{d.x}
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

func (c jnz) tgl() instruction {
	return &cpy{c.x, c.y}
}

type tgl struct {
	x string
}

func (t tgl) exec(p *program) {
	defer func() { p.pc++ }()
	loc := p.pc + p.getValue(t.x)
	if loc >= len(p.instructions) {
		return
	}

	p.instructions[loc] = p.instructions[loc].tgl()
}

func (t tgl) tgl() instruction {
	return &inc{t.x}
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
	case "tgl":
		return &tgl{parts[1]}
	}

	panic("bad instr")
}

func main() {

	input := "2016/Day23/input"
	var instr1, instr2 []instruction
	for _, l := range util.GetFileStrings(input) {
		instr1 = append(instr1, parseInstruction(l))
		instr2 = append(instr2, parseInstruction(l))
	}

	prog1 := &program{
		registers:    map[string]int{"a": 7},
		pc:           0,
		instructions: instr1,
	}
	prog1.run()
	log.Printf("Part 1: Value of register a: %d", prog1.registers["a"])

	prog2 := &program{
		registers:    map[string]int{"a": 12},
		pc:           0,
		instructions: instr2,
	}
	prog2.run()
	log.Printf("Part 2: Value of register a: %d", prog2.registers["a"])
}
