package main

import (
	"fmt"
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
	output       []int
}

func (p program) getValue(vr string) int {
	v, err := strconv.Atoi(vr)
	if err == nil {
		return v
	}
	return p.registers[vr]
}

func (p *program) run() bool {
	for {
		if p.pc >= len(p.instructions) {
			return false
		}
		p.instructions[p.pc].exec(p)
		if !p.validOutput() {
			return false
		}
		if len(p.output) > 50 {
			return true
		}
	}
}

func (p *program) reset() {
	p.registers = make(map[string]int)
	p.pc = 0
	p.output = make([]int, 0)
}

func (p *program) validOutput() bool {
	ll := len(p.output)
	if ll < 2 {
		return true
	}

	c1, c2 := p.output[ll-1], p.output[ll-2]
	if (c1 != 0 && c1 != 1) || (c2 != 0 && c2 != 1) || c1 == c2 {
		return false
	}
	return true
}

func isScalar(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		// successfully converted to int
		return true
	}
	return false
}

type cpy struct {
	x, y string
}

func (c cpy) exec(p *program) {
	defer func() { p.pc++ }()
	if isScalar(c.y) {
		// destination is a value (due to toggle), skip
		return
	}
	p.registers[c.y] = p.getValue(c.x)
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

type out struct {
	x string
}

func (o out) exec(p *program) {
	defer func() { p.pc++ }()
	p.output = append(p.output, p.getValue(o.x))
}

type mul struct {
	a, b string
}

func (m mul) exec(p *program) {
	defer func() { p.pc++ }()
	p.registers[m.a] *= p.getValue(m.b)
}

type add struct {
	a, b string
}

func (a add) exec(p *program) {
	defer func() { p.pc++ }()
	p.registers[a.b] += p.getValue(a.a)
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
	case "mul":
		return &mul{parts[1], parts[2]}
	case "add":
		return &add{parts[1], parts[2]}
	case "out":
		return &out{parts[1]}
	}

	panic(fmt.Sprintf("bad instr: %s", s))
}

func main() {

	input := "2016/Day25/input"
	var instr []instruction
	for _, l := range util.GetFileStrings(input) {
		instr = append(instr, parseInstruction(l))
	}

	prog := &program{
		registers:    map[string]int{},
		pc:           0,
		instructions: instr,
	}
	i := 0
	for {
		prog.reset()
		prog.registers["a"] = i
		if prog.run() {
			log.Printf("Part 1: good signal with input %d", i)
			break
		}
		i++

	}
}
