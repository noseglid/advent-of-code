package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Operator interface {
	Operate(program *Program) (int, bool)
}

type Opcode int
type ParamMode int

const (
	Add          = Opcode(1)
	Mult         = Opcode(2)
	Input        = Opcode(3)
	Output       = Opcode(4)
	JITrue       = Opcode(5)
	JIFalse      = Opcode(6)
	LessThan     = Opcode(7)
	Equals       = Opcode(8)
	Relativebase = Opcode(9)
	Halt         = Opcode(99)
)

const (
	ModePosition  = ParamMode(0)
	ModeImmediate = ParamMode(1)
	ModeRelative  = ParamMode(2)
)

type Program struct {
	pc       int
	relative int
	refmem   []int
	memory   []int
	input    chan int
	output   chan int
}

func NewProgram(memory []int) *Program {
	return &Program{
		pc:     0,
		refmem: memory,
		memory: nil,
		input:  make(chan int, 512),
		output: make(chan int, 512),
	}
}

func (p *Program) reset() {
	p.pc = 0
	p.relative = 0
	p.memory = nil
	p.input = make(chan int, 512)
	p.output = make(chan int, 512)
}

func (p *Program) operator(pos int) Operator {
	op := Opcode(p.memory[pos] % 100)
	mode := p.memory[pos] / 100
	switch op {
	case Add:
		return &addop{instr{mode, []int{p.memory[pos+1], p.memory[pos+2], p.memory[pos+3]}}}

	case Mult:
		return &multop{instr{mode, []int{p.memory[pos+1], p.memory[pos+2], p.memory[pos+3]}}}

	case Input:
		return &inputop{instr{mode, []int{p.memory[pos+1]}}}

	case Output:
		return &outputop{instr{mode, []int{p.memory[pos+1]}}}

	case JITrue:
		return &jumpiftrue{instr{mode, []int{p.memory[pos+1], p.memory[pos+2]}}}

	case JIFalse:
		return &jumpiffalse{instr{mode, []int{p.memory[pos+1], p.memory[pos+2]}}}

	case LessThan:
		return &lessthan{instr{mode, []int{p.memory[pos+1], p.memory[pos+2], p.memory[pos+3]}}}

	case Equals:
		return &equals{instr{mode, []int{p.memory[pos+1], p.memory[pos+2], p.memory[pos+3]}}}

	case Relativebase:
		return &relativebase{instr{mode, []int{p.memory[pos+1]}}}

	case Halt:
		return &haltop{}

	default:
		panic("bad op")
	}
}

func (p *Program) printMemory(from int) {
	sep := " "
	for i := from; i < from+10; i++ {
		if i >= len(p.memory) {
			return
		}

		fmt.Printf("%s%d", sep, p.memory[i])
		sep = ", "
	}
	fmt.Println()
}

func (p *Program) run() {
	p.memory = make([]int, len(p.refmem))
	copy(p.memory, p.refmem)

	n := 0
	for {
		log.Printf("pc=%d, rb=%d memsize=%d", p.pc, p.relative, len(p.memory))

		pc, halt := p.operator(p.pc).Operate(p)
		if halt {
			break
		}

		p.pc = pc
		n++
	}

	close(p.output)
	close(p.input)
}

func (p *Program) paramAddr(i int, mode ParamMode) int {
	switch mode {
	case ModePosition:
		p.ensureMemory(i)
		return i
	case ModeImmediate:
		panic("immediate mode when getting address")
	case ModeRelative:
		p.ensureMemory(p.relative + i)
		return p.relative + i
	default:
		panic("bad mode")
	}
}

func (p *Program) programCounter() int {
	return p.pc
}

func (p *Program) ensureMemory(size int) {
	log.Printf("ensuring memory %d (len=%d)", size, len(p.memory))
	if len(p.memory) > size {
		return
	}

	m := make([]int, size+1)
	copy(m, p.memory)
	p.memory = m
}

func (p *Program) value(i int, mode ParamMode) int {
	switch mode {
	case ModePosition:
		p.ensureMemory(i)
		return p.memory[i]
	case ModeImmediate:
		return i
	case ModeRelative:
		p.ensureMemory(p.relative + i)
		return p.memory[p.relative+i]
	default:
		panic("bad mode")
	}
}

func (p *Program) write(i, v int) {
	p.ensureMemory(i)
	p.memory[i] = v
}

func (p *Program) getInput() int {
	log.Printf("program getting input")
	return <-p.input
}

func (p *Program) final() int {
	return <-p.output
}

func (p *Program) writeOutput(v int) {
	log.Printf("program writing output %d", v)
	p.output <- v
}

type instr struct {
	mode   int
	params []int
}

func (i *instr) paramMode(paramIndex int) ParamMode {
	return ParamMode((i.mode / int(math.Pow10(paramIndex))) % 10)
}

type addop struct {
	instr
}

func (o *addop) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	log.Printf("doing add with %d+%d=%d to %d", v0, v1, v0+v1, o.params[2])
	addr := p.paramAddr(o.params[2], o.paramMode(2))
	p.write(addr, v0+v1)
	return p.programCounter() + 4, false
}

type multop struct {
	instr
}

func (o *multop) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	log.Printf("doing mult %d*%d=%d to %d", v0, v1, v0*v1, o.params[2])
	addr := p.paramAddr(o.params[2], o.paramMode(2))
	p.write(addr, v0*v1)
	return p.programCounter() + 4, false
}

type haltop struct {
	instr
}

func (o *haltop) Operate(p *Program) (int, bool) {
	log.Printf("halting")
	return p.programCounter() + 1, true
}

type inputop struct {
	instr
}

func (o *inputop) Operate(p *Program) (int, bool) {
	input := p.getInput()

	addr := p.paramAddr(o.params[0], o.paramMode(0))
	p.write(addr, input)
	log.Printf("getting input %d to addr=%d", input, addr)
	return p.programCounter() + 2, false
}

type outputop struct {
	instr
}

func (o *outputop) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	log.Printf("writing output %d from (param=%d,mode=%d)", v0, o.params[0], o.paramMode(0))
	p.writeOutput(v0)
	return p.programCounter() + 2, false
}

type jumpiftrue struct {
	instr
}

func (o *jumpiftrue) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	log.Printf("jumping to %d if %d != 0", v1, v0)
	if v0 != 0 {
		return v1, false
	} else {
		return p.programCounter() + 3, false
	}
}

type jumpiffalse struct {
	instr
}

func (o *jumpiffalse) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	log.Printf("jumping if false: %d", v0)
	if v0 == 0 {
		return p.value(o.params[1], o.paramMode(1)), false
	} else {
		return p.programCounter() + 3, false
	}
}

type lessthan struct {
	instr
}

func (o *lessthan) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	v := 0
	if v0 < v1 {
		v = 1
	}
	log.Printf("lessthan %d < %d, writing %d to %d", v0, v1, v, o.params[2])
	addr := p.paramAddr(o.params[2], o.paramMode(2))
	p.write(addr, v)
	return p.programCounter() + 4, false
}

type equals struct {
	instr
}

func (o *equals) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	v := 0
	if v0 == v1 {
		v = 1
	}
	log.Printf("testing equals %d == %d, storing %d at %d", v0, v1, v, o.params[2])
	addr := p.paramAddr(o.params[2], o.paramMode(2))
	p.write(addr, v)
	return p.programCounter() + 4, false
}

type relativebase struct {
	instr
}

func (o relativebase) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	log.Printf("modifying relative base (param=%d, mode=%d) with %d", o.params[0], o.paramMode(0), v0)
	p.relative += v0
	return p.programCounter() + 2, false
}

func main() {
	s := util.GetFile("2019/Day9/input")

	var memory []int

	for _, op := range strings.Split(s, ",") {
		memory = append(memory, util.MustAtoi(strings.TrimSpace(op)))
	}

	program := NewProgram(memory)
	program.input <- 1
	program.run()

	log.Printf("BOOST keycode in test mode (part1): %d", <-program.output)

	program.reset()
	program.input <- 2
	program.run()
	log.Printf("Distress signal coordinates (part2): %d", <-program.output)
}
