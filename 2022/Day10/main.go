package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Op interface {
	Apply(x *int)
	Cycles() int
	fmt.Stringer
}

type OpAdd struct{ v int }

func (o OpAdd) Apply(x *int)   { *x = *x + o.v }
func (OpAdd) Cycles() int      { return 2 }
func (o OpAdd) String() string { return fmt.Sprintf("addx %v", o.v) }

type OpNoop struct{}

func (o OpNoop) Apply(x *int)   {}
func (o OpNoop) Cycles() int    { return 1 }
func (o OpNoop) String() string { return "noop" }

func ParseOp(l string) Op {
	parts := strings.Split(l, " ")
	switch parts[0] {
	case "addx":
		return OpAdd{util.MustAtoi(parts[1])}
	case "noop":
		return OpNoop{}
	}

	panic("bad op: " + parts[0])
}

func printCRT(crt [][]rune) {
	for _, r := range crt {
		for _, c := range r {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func main() {
	var ops []Op
	for _, instr := range util.GetFileStrings("2022/Day10/input") {
		ops = append(ops, ParseOp(instr))
	}

	crt := make([][]rune, 6)
	for r := range crt {
		crt[r] = make([]rune, 40)
		for c := range crt[r] {
			crt[r][c] = '.'
		}
	}

	var x, cycle, signalStrength = 1, 0, 0
	crtx, crty := 0, 0

	for _, o := range ops {
		for i := 0; i < o.Cycles(); i++ {
			cycle++
			if crtx == x-1 || crtx == x || crtx == x+1 {
				crt[crty][crtx] = '#'
			}
			crtx = (crtx + 1) % 40
			if crtx == 0 {
				crty++
			}

			if cycle == 20 || (cycle-20)%40 == 0 {
				signalStrength += cycle * x
			}
		}

		o.Apply(&x)
	}

	fmt.Printf("Signal Strenght (part1): %d\n", signalStrength)

	printCRT(crt)

}
