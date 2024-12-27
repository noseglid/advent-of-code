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

func DumpReg(r []int) {
	for i := 0; i < 6; i++ {
		fmt.Printf("%d=%010d ", i, r[i])
	}
	fmt.Println()
}

/*
 0: addi 5 16 5 // r[5] += 16
 1: seti 1 2 2  // r[2] = 1
 2: seti 1 0 4  // r[4] = 1
 3: mulr 2 4 3  // r[3] = r[2] * r[4]
 4: eqrr 3 1 3  // r[3] = r[3] == r[1] ? 1 : 0
 5: addr 3 5 5  // r[5] += r[3]
 6: addi 5 1 5  // r[5] += 1
 7: addr 2 0 0  // r[0] += r[2]
 8: addi 4 1 4  // r[4] += 1
 9: gtrr 4 1 3  // r[3] = r[4] > r[1] ? 1 : 0
10: addr 5 3 5  // r[5] += r[3]
11: seti 2 4 5  // r[5] = 2
12: addi 2 1 2  // r[2] += 1
13: gtrr 2 1 3  // r[3] = r[2] > r[1] ? 1 : 0
14: addr 3 5 5  // r[5] += r[3]
15: seti 1 1 5  // r[5] = 1
16: mulr 5 5 5  // BREAKOUT r[5] *= r[5]
*/

func inner(r []int) {
	for r[2] = 1; r[2] <= r[1]; r[2]++ {
		for r[4] = 1; r[4] <= r[1]; r[4]++ {
			r[3] = r[2] * r[4]
			if r[3] == r[1] {
				r[0] += r[2]
			}
		}
	}
}

func factorSum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			s += i
		}
	}
	return s
}

func main() {

	lines := util.GetFileStrings("2018/Day19/input")

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

	ip := 0
	reg := []int{0, 0, 0, 0, 0, 0}
	for ip >= 0 && ip < len(instructions) {
		instr := instructions[ip]
		reg[bind] = ip
		opmap[instr.op](reg, instr.A, instr.B, instr.C)
		ip = reg[bind]
		ip++
	}

	fmt.Printf("Value of register 0 after halt (part 1): %d\n", reg[0])

	// n := 0
	// ip = 0
	// reg = []int{1, 0, 0, 0, 0, 0}
	// for n < 200 && ip >= 0 && ip < len(instructions) {
	// 	instr := instructions[ip]
	// 	reg[bind] = ip
	// 	fmt.Printf("ip=%02d %v %s %d %d %d ", ip, reg, instr.op, instr.A, instr.B, instr.C)
	// 	// fmt.Printf("ip=%02d ", ip+1)
	// 	// DumpReg(reg)
	// 	opmap[instr.op](reg, instr.A, instr.B, instr.C)
	// 	fmt.Printf("%v\n", reg)
	// 	ip = reg[bind]
	// 	ip++
	// 	n++
	// }

	// reg = []int{0, 10551383, 0, 10550400, 0, 1}
	// inner(reg)
	fmt.Printf("Value of register 0 after halt (part 2): %d\n", factorSum(10551383))

}
