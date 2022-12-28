package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Op interface {
	Val(r map[string]int, d int) int
}

type scalarOp struct {
	v int
}

func (s *scalarOp) Val(r map[string]int, d int) int {
	return s.v
}

type registerOp struct {
	r string
}

func (s *registerOp) Val(r map[string]int, d int) int {
	if _, ok := r[s.r]; !ok {
		return d
	}
	return r[s.r]
}

type Instruction interface {
	Apply(c *computer) bool
}

type computer struct {
	id        int
	pc        int
	instr     []Instruction
	registers map[string]int
	played    int
	recovered int
	sendQueue []int
	recvQueue []int
}

func (c *computer) Reset() {
	c.pc = 0
	c.registers = map[string]int{}
	c.sendQueue = []int{}
	c.recvQueue = []int{}
}

type snd struct {
	r Op
}

func (i *snd) Apply(c *computer) bool {
	c.played = i.r.Val(c.registers, c.id)
	c.sendQueue = append(c.sendQueue, i.r.Val(c.registers, c.id))
	c.pc++
	return false
}

type set struct {
	r  string
	o2 Op
}

func (i *set) Apply(c *computer) bool {
	c.registers[i.r] = i.o2.Val(c.registers, c.id)
	c.pc++
	return false
}

type add struct {
	r  string
	o2 Op
}

func (i *add) Apply(c *computer) bool {
	if _, ok := c.registers[i.r]; !ok {
		c.registers[i.r] = c.id
	}

	c.registers[i.r] += i.o2.Val(c.registers, c.id)
	c.pc++
	return false
}

type mul struct {
	r  string
	o2 Op
}

func (i *mul) Apply(c *computer) bool {
	if _, ok := c.registers[i.r]; !ok {
		c.registers[i.r] = c.id
	}

	c.registers[i.r] *= i.o2.Val(c.registers, c.id)
	c.pc++
	return false
}

type mod struct {
	r  string
	o2 Op
}

func (i *mod) Apply(c *computer) bool {
	if _, ok := c.registers[i.r]; !ok {
		c.registers[i.r] = c.id
	}

	c.registers[i.r] %= i.o2.Val(c.registers, c.id)
	c.pc++
	return false
}

type rcv struct {
	r string
}

func (i *rcv) Apply(c *computer) bool {
	if c.registers[i.r] != 0 {
		c.recovered = c.played
	}

	if len(c.recvQueue) == 0 {
		return true
	}
	c.registers[i.r] = c.recvQueue[0]
	c.recvQueue = c.recvQueue[1:]
	c.pc++

	return false
}

type jgz struct {
	op1, op2 Op
}

func (i *jgz) Apply(c *computer) bool {
	if i.op1.Val(c.registers, c.id) > 0 {
		c.pc += i.op2.Val(c.registers, c.id)
	} else {
		c.pc++
	}
	return false
}

func parseOp(s string) Op {
	if i, err := strconv.Atoi(s); err == nil {
		return &scalarOp{v: i}
	} else {
		return &registerOp{r: s}
	}
}

func main() {

	input := util.GetFileStrings("2017/Day18/input")

	c := &computer{
		id:        0,
		registers: map[string]int{},
	}
	c2 := &computer{
		id:        1,
		registers: map[string]int{},
	}
	var instructions []Instruction
	for _, l := range input {
		parts := strings.Split(l, " ")
		switch parts[0] {
		case "snd":
			instructions = append(instructions, &snd{r: parseOp(parts[1])})
		case "set":
			instructions = append(instructions, &set{r: parts[1], o2: parseOp(parts[2])})
		case "add":
			instructions = append(instructions, &add{r: parts[1], o2: parseOp(parts[2])})
		case "mul":
			instructions = append(instructions, &mul{r: parts[1], o2: parseOp(parts[2])})
		case "mod":
			instructions = append(instructions, &mod{r: parts[1], o2: parseOp(parts[2])})
		case "rcv":
			instructions = append(instructions, &rcv{r: parts[1]})
		case "jgz":
			instructions = append(instructions, &jgz{op1: parseOp(parts[1]), op2: parseOp(parts[2])})
		}
	}

	c.instr, c2.instr = instructions, instructions

	for c.pc >= 0 && c.pc < len(c.instr) {
		c.instr[c.pc].Apply(c)
		if c.recovered != 0 {
			fmt.Printf("recovered (part1): %d\n", c.recovered)
			break
		}
	}

	c.Reset()

	n, n2 := 0, 0
	for c.pc >= 0 && c2.pc >= 0 && c.pc < len(c.instr) && c2.pc < len(c2.instr) {
		needData1 := c.instr[c.pc].Apply(c)
		if len(c.sendQueue) > 0 {
			n += len(c.sendQueue)
			c2.recvQueue = append(c2.recvQueue, c.sendQueue...)
			c.sendQueue = c.sendQueue[:0]
		}
		needData2 := c2.instr[c2.pc].Apply(c2)
		if len(c2.sendQueue) > 0 {
			n2 += len(c2.sendQueue)
			c.recvQueue = append(c.recvQueue, c2.sendQueue...)
			c2.sendQueue = c2.sendQueue[:0]
		}

		if needData1 && len(c.recvQueue) == 0 && needData2 && len(c2.recvQueue) == 0 {
			break
		}
	}

	fmt.Printf("program %d,%d sent (part2): %d,%d\n", c.id, c2.id, n, n2)

}
