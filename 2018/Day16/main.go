package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

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

type opf func([]int, int, int, int)

func test(fs string, a, op, b []int) bool {
	f := opmap[fs]
	input := make([]int, 4)
	copy(input, a)
	f(input, op[1], op[2], op[3])
	return slices.Equal(input, b)
}

func allOne(m map[int][]string) bool {
	for _, v := range m {
		if len(v) != 1 {
			return false
		}
	}

	return true
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

func main() {
	lines := util.GetFileStrings("2018/Day16/input_sample")

	opcodes := map[int][]string{}
	for i := 0; i < 16; i++ {
		opcodes[i] = make([]string, 16)
		copy(opcodes[i], slices.Collect(maps.Keys(opmap)))
	}

	s := 0
	for i := 0; i+2 < len(lines); i += 4 {
		a, op, b := make([]int, 4), make([]int, 4), make([]int, 4)
		if _, err := fmt.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &a[0], &a[1], &a[2], &a[3]); err != nil {
			fmt.Printf("%s\n", lines[i])
			panic(err)
		}
		if _, err := fmt.Sscanf(lines[i+1], "%d %d %d %d", &op[0], &op[1], &op[2], &op[3]); err != nil {
			panic(err)
		}
		if _, err := fmt.Sscanf(lines[i+2], "After: [%d, %d, %d, %d]", &b[0], &b[1], &b[2], &b[3]); err != nil {
			panic(err)
		}

		n := 0
		for _, f := range slices.Collect(maps.Keys(opmap)) {
			if test(f, a, op, b) {
				n++
			}
		}
		if n >= 3 {
			s++
		}

		next := []string{}
		for _, f := range opcodes[op[0]] {
			if test(f, a, op, b) {
				next = append(next, f)
			}
		}
		opcodes[op[0]] = next
	}

	fmt.Printf("samples behaving like 3 or more opcodes (part1): %d\n", s)

	for !allOne(opcodes) {
		for k, v := range opcodes {
			if len(v) != 1 {
				continue
			}

			for ki := range opcodes {
				if k == ki {
					continue
				}
				opcodes[ki], _ = util.RemoveByValue(opcodes[ki], v[0])
			}
		}
	}

	program := util.GetFileStrings("2018/Day16/input_program")
	r := []int{0, 0, 0, 0}
	for _, l := range program {
		op := make([]int, 4)
		fmt.Sscanf(l, "%d %d %d %d", &op[0], &op[1], &op[2], &op[3])
		opmap[opcodes[op[0]][0]](r, op[1], op[2], op[3])
	}
	fmt.Printf("Value in register 0 after program (part2): %d\n", r[0])
}
