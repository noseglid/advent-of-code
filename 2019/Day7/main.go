package main

import (
	"log"
	"math"
	"strings"
	"sync"

	"github.com/noseglid/advent-of-code/util"
)

type Operator interface {
	Operate(program *Program) (int, bool)
}

type Opcode int
type ParamMode int

const (
	Add      = Opcode(1)
	Mult     = Opcode(2)
	Input    = Opcode(3)
	Output   = Opcode(4)
	JITrue   = Opcode(5)
	JIFalse  = Opcode(6)
	LessThan = Opcode(7)
	Equals   = Opcode(8)
	Halt     = Opcode(99)
)

const (
	ModePosition  = ParamMode(0)
	ModeImmediate = ParamMode(1)
)

type Program struct {
	id     int
	pc     int
	refmem []int
	memory []int
	input  chan int
	output chan int
}

func NewProgram(id int, memory []int) *Program {
	return &Program{
		id:     id,
		pc:     0,
		refmem: memory,
		memory: nil,
		input:  make(chan int, 512),
		output: make(chan int, 512),
	}
}

func (p *Program) reset() {
	p.pc = 0
	p.memory = nil
	p.input = nil
	p.output = nil
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

	case Halt:
		return &haltop{}

	default:
		panic("bad op")
	}
}
func (p *Program) run() {
	p.memory = make([]int, len(p.refmem))
	copy(p.memory, p.refmem)

	for {
		log.Printf("pc=%d, memory: %v", p.pc, p.memory)
		pc, halt := p.operator(p.pc).Operate(p)
		if halt {
			break
		}

		p.pc = pc
	}
}

func (p *Program) programCounter() int {
	return p.pc
}

func (p *Program) value(i int, mode ParamMode) int {
	switch mode {
	case ModePosition:
		return p.memory[i]
	case ModeImmediate:
		return i
	default:
		panic("bad mode")
	}
}

func (p *Program) write(i, v int) {
	p.memory[i] = v
}

func (p *Program) getInput() int {
	log.Printf("program %d getting input", p.id)
	return <-p.input
}

func (p *Program) final() int {
	return <-p.output
}

func (p *Program) writeOutput(v int) {
	log.Printf("program %d writing output %d", p.id, v)
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
	log.Printf("doing add with %d + %d", v0, v1)
	p.write(o.params[2], v0+v1)
	return p.programCounter() + 4, false
}

type multop struct {
	instr
}

func (o *multop) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	log.Printf("doing mult %d + %d", v0, v1)
	p.write(o.params[2], v0*v1)
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
	log.Printf("getting input")
	p.write(o.params[0], p.getInput())
	return p.programCounter() + 2, false
}

type outputop struct {
	instr
}

func (o *outputop) Operate(p *Program) (int, bool) {
	v0 := p.value(o.params[0], o.paramMode(0))
	log.Printf("writing output %d from %d (mode=%d)", v0, o.params[0], o.paramMode(0))
	p.writeOutput(v0)
	return p.programCounter() + 2, false
}

type jumpiftrue struct {
	instr
}

func (o *jumpiftrue) Operate(p *Program) (int, bool) {
	log.Printf("jumping if true")
	v0 := p.value(o.params[0], o.paramMode(0))
	if v0 != 0 {
		return p.value(o.params[1], o.paramMode(1)), false
	} else {
		return p.programCounter() + 3, false
	}
}

type jumpiffalse struct {
	instr
}

func (o *jumpiffalse) Operate(p *Program) (int, bool) {
	log.Printf("jumping if false")
	v0 := p.value(o.params[0], o.paramMode(0))
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
	log.Printf("testing less than")
	v0 := p.value(o.params[0], o.paramMode(0))
	v1 := p.value(o.params[1], o.paramMode(1))
	v := 0
	if v0 < v1 {
		v = 1
	}
	p.write(o.params[2], v)
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
	p.write(o.params[2], v)
	return p.programCounter() + 4, false
}

func main() {
	s := util.GetFile("2019/Day7/input")

	var memory []int

	for _, op := range strings.Split(s, ",") {
		memory = append(memory, util.MustAtoi(strings.TrimSpace(op)))
	}

	maxSignal := 0
	util.PermInt([]int{0, 1, 2, 3, 4}, func(phaseSettings []int) {
		input := 0
		for i := 0; i < 5; i++ {
			p := NewProgram(i, memory)
			p.input <- phaseSettings[i]
			p.input <- input
			p.run()
			input = p.final()
		}

		if input > maxSignal {
			maxSignal = input
		}
	})

	log.Printf("max thruster (part1): %d", maxSignal)

	maxSignalFeedback := 0
	util.PermInt([]int{5, 6, 7, 8, 9}, func(phaseSettings []int) {

		var programs []*Program
		var pA, pCurr *Program
		for i := 0; i < 5; i++ {
			pNew := NewProgram(i, memory)
			if pA == nil {
				pA = pNew
			}
			if pCurr == nil {
				pNew.input <- phaseSettings[i]
				pNew.input <- 0
			} else {
				pNew.input = pCurr.output
				pNew.input <- phaseSettings[i]
			}

			pCurr = pNew
			programs = append(programs, pCurr)
		}

		pCurr.output = pA.input

		wg := sync.WaitGroup{}
		wg.Add(len(programs))
		for _, p := range programs {
			go func(p *Program) { p.run(); wg.Done() }(p)
		}

		wg.Wait()
		f := pCurr.final()
		if f > maxSignalFeedback {
			maxSignalFeedback = f
		}
	})

	log.Printf("max thruster with feedback (part2): %d", maxSignalFeedback)
}
