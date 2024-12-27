package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type opf func([]int, int, int, int)

func addr(r []int, A, B, C int) {
	r[C] = r[A] + r[B]
}
func addi(r []int, A, B, C int) {
	r[C] = r[A] + B
}
func mulr(r []int, A, B, C int) {
	r[C] = r[A] * r[B]
}
func muli(r []int, A, B, C int) {
	r[C] = r[A] * B
}
func banr(r []int, A, B, C int) {
	r[C] = r[A] & r[B]
}
func bani(r []int, A, B, C int) {
	r[C] = r[A] & B
}
func borr(r []int, A, B, C int) {
	r[C] = r[A] | r[B]
}
func bori(r []int, A, B, C int) {
	r[C] = r[A] | B
}
func setr(r []int, A, B, C int) {
	r[C] = r[A]
}
func seti(r []int, A, B, C int) {
	r[C] = A
}
func gtir(r []int, A, B, C int) {
	if A > r[B] {
		r[C] = 1
	} else {
		r[C] = 0
	}
}
func gtri(r []int, A, B, C int) {
	if r[A] > B {
		r[C] = 1
	} else {
		r[C] = 0
	}
}
func gtrr(r []int, A, B, C int) {
	if r[A] > r[B] {
		r[C] = 1
	} else {
		r[C] = 0
	}
}
func eqir(r []int, A, B, C int) {
	if A == r[B] {
		r[C] = 1
	} else {
		r[C] = 0
	}
}
func eqri(r []int, A, B, C int) {
	if r[A] == B {
		r[C] = 1
	} else {
		r[C] = 0
	}
}
func eqrr(r []int, A, B, C int) {
	if r[A] == r[B] {
		r[C] = 1
	} else {
		r[C] = 0
	}
}

var opmap = map[string]opf{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

type instr struct {
	op      string
	A, B, C int
}

func main() {
	lines := util.GetFileStrings("2018/Day21/input")

	instructions := []instr{}

	bind := 0
	for i, l := range lines {
		if i == 0 {
			fmt.Sscanf(l, "#ip %d\n", &bind)
			continue
		}
		var ins instr
		fmt.Sscanf(l, "%s %d %d %d", &ins.op, &ins.A, &ins.B, &ins.C)
		instructions = append(instructions, ins)
	}

	valid := map[int]bool{}

	ip := 0
	reg := []int{0, 0, 0, 0, 0, 0}

	p := -1
	for ip >= 0 && ip < len(instructions) {
		instr := instructions[ip]
		reg[bind] = ip
		opmap[instr.op](reg, instr.A, instr.B, instr.C)
		ip = reg[bind]
		if ip == 28 {
			if _, ok := valid[reg[4]]; ok {
				fmt.Printf("Most instructions causing halt (part2): %d\n", p)
				break
			}
			if p == -1 {
				fmt.Printf("Fewest instruction causing halt (part1): %d\n", reg[4])
			}
			p = reg[4]
			valid[reg[4]] = true
		}
		ip++
	}
}
