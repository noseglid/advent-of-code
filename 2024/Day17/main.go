package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func dv(a, b int) int {
	return int(float64(a) / math.Pow(2, float64(b)))
}

type op interface {
	exec(p *Program)
}

type adv struct {
	arg int
}

func (op adv) exec(p *Program) {
	p.a = dv(p.a, p.GetCombo(op.arg))
}

type bxl struct {
	arg int
}

func (op bxl) exec(p *Program) {
	p.b = p.b ^ op.arg
}

type bst struct {
	arg int
}

func (op bst) exec(p *Program) {
	p.b = p.GetCombo(op.arg) % 8
}

type jnz struct {
	arg int
}

func (op jnz) exec(p *Program) {
	if p.a == 0 {
		return
	}
	p.pc = op.arg - 1
}

type bxc struct {
	arg int
}

func (op bxc) exec(p *Program) {
	p.b = p.b ^ p.c
}

type out struct {
	arg int
}

func (op out) exec(p *Program) {
	p.Output(p.GetCombo(op.arg) % 8)
}

type bdv struct {
	arg int
}

func (op bdv) exec(p *Program) {
	p.b = dv(p.a, p.GetCombo(op.arg))
}

type cdv struct {
	arg int
}

func (op cdv) exec(p *Program) {
	p.c = dv(p.a, p.GetCombo(op.arg))
}
func parseOp(code, arg int) op {
	switch code {
	case 0:
		return &adv{arg}
	case 1:
		return &bxl{arg}
	case 2:
		return &bst{arg}
	case 3:
		return &jnz{arg}
	case 4:
		return &bxc{arg}
	case 5:
		return &out{arg}
	case 6:
		return &bdv{arg}
	case 7:
		return &cdv{arg}
	}
	panic("bad opcode")
}

type Program struct {
	pc      int
	instr   []op
	a, b, c int
	output  []int
}

func (p *Program) GetCombo(i int) int {
	switch i {
	case 1, 2, 3:
		return i
	case 4:
		return p.a
	case 5:
		return p.b
	case 6:
		return p.c
	case 7:
		fallthrough
	default:
		panic("bad combo")
	}
}

func (p *Program) Run() {
	for p.pc >= 0 && p.pc < len(p.instr) {
		op := p.instr[p.pc]
		op.exec(p)
		p.pc++
	}
}

func (p *Program) Output(i int) {
	p.output = append(p.output, i)
}

func join(l []int) string {
	s := make([]string, len(l))
	for i := range l {
		s[i] = fmt.Sprintf("%d", l[i])
	}
	return strings.Join(s, ",")
}

func r(a int) []int {
	out := []int{}
	for a != 0 {
		out = append(out, f(a))
		a >>= 3
	}
	return out
}

func f(a int) int {
	b := a % 8
	b = b ^ 2
	c := a >> b
	b = b ^ c
	b = b ^ 3
	return b % 8
}

func eq(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {

	lines := util.GetFileStrings("2024/Day17/input")
	var a, b, c int
	fmt.Sscanf(lines[0], "Register A: %d", &a)
	fmt.Sscanf(lines[1], "Register B: %d", &b)
	fmt.Sscanf(lines[2], "Register C: %d", &c)

	numbers := []int{}
	for _, s := range strings.Split(lines[4][len("Program: "):], ",") {
		numbers = append(numbers, util.MustAtoi(s))
	}

	instructions := []op{}
	for i := 0; i < len(numbers); i += 2 {
		instructions = append(instructions, parseOp(numbers[i], numbers[i+1]))
	}

	p := Program{0, instructions, a, b, c, nil}
	p.Run()
	fmt.Printf("Ouput (part1): %s\n", join(p.output))

	group := 3
	i := 1
	for group <= 16 {
		o := r(i)
		if len(o) >= 3 && eq(o[:3], numbers[len(numbers)-group:len(numbers)-group+3]) {
			group++
			i <<= 3
		} else {
			i++
		}
	}

	fmt.Printf("Register A so it's a quine (part2): %d\n", i/8)
	fmt.Printf("r(i): %v\n", r(i/8))
}
