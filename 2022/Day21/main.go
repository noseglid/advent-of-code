package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type monkey struct {
	name string
	op   Op
}

func (m monkey) Eval(monkeys map[string]monkey) (int, string) {
	v, s := m.op.Eval(monkeys)
	if m.name == "humn" {
		return v, "x"
	} else {
		return v, s
	}
}

type Op interface {
	Eval(monkeys map[string]monkey) (int, string)
}

type scalar struct {
	v int
}

func (s scalar) Eval(monkeys map[string]monkey) (int, string) {
	return s.v, strconv.Itoa(s.v)
}

type add struct {
	op1, op2 string
}

func (a add) Eval(monkeys map[string]monkey) (int, string) {
	v1, s1 := monkeys[a.op1].Eval(monkeys)
	v2, s2 := monkeys[a.op2].Eval(monkeys)
	if !strings.Contains(s1, "x") {
		s1 = strconv.Itoa(v1)
	}
	if !strings.Contains(s2, "x") {
		s2 = strconv.Itoa(v2)
	}
	return v1 + v2, fmt.Sprintf("(%s) + (%s)", s1, s2)
}

type sub struct {
	op1, op2 string
}

func (a sub) Eval(monkeys map[string]monkey) (int, string) {
	v1, s1 := monkeys[a.op1].Eval(monkeys)
	v2, s2 := monkeys[a.op2].Eval(monkeys)
	if !strings.Contains(s1, "x") {
		s1 = strconv.Itoa(v1)
	}
	if !strings.Contains(s2, "x") {
		s2 = strconv.Itoa(v2)
	}
	return v1 - v2, fmt.Sprintf("(%s) - (%s)", s1, s2)
}

type mult struct {
	op1, op2 string
}

func (a mult) Eval(monkeys map[string]monkey) (int, string) {
	v1, s1 := monkeys[a.op1].Eval(monkeys)
	v2, s2 := monkeys[a.op2].Eval(monkeys)
	if !strings.Contains(s1, "x") {
		s1 = strconv.Itoa(v1)
	}
	if !strings.Contains(s2, "x") {
		s2 = strconv.Itoa(v2)
	}
	return v1 * v2, fmt.Sprintf("(%s) * (%s)", s1, s2)
}

type div struct {
	op1, op2 string
}

func (a div) Eval(monkeys map[string]monkey) (int, string) {
	v1, s1 := monkeys[a.op1].Eval(monkeys)
	v2, s2 := monkeys[a.op2].Eval(monkeys)
	if !strings.Contains(s1, "x") {
		s1 = strconv.Itoa(v1)
	}
	if !strings.Contains(s2, "x") {
		s2 = strconv.Itoa(v2)
	}
	return v1 / v2, fmt.Sprintf("(%s) / (%s)", s1, s2)
}

type eq struct {
	op1, op2 string
}

func (a eq) Eval(monkeys map[string]monkey) (int, string) {
	v1, s1 := monkeys[a.op1].Eval(monkeys)
	v2, s2 := monkeys[a.op2].Eval(monkeys)
	fmt.Printf("v1=%d, v2=%d\n", v1, v2)
	return v1 + v2, fmt.Sprintf("(%s) = (%s)", s1, s2)
}

func main() {
	input := util.GetFileStrings("2022/Day21/input")

	monkeys := map[string]monkey{}
	for _, l := range input {
		parts := strings.Split(l, " ")
		m := monkey{
			name: parts[0][:len(parts[0])-1],
		}
		if m.name == "root" {
			m.op = eq{op1: parts[1], op2: parts[3]}
		} else if len(parts) == 2 {
			m.op = scalar{v: util.MustAtoi(parts[1])}
		} else if len(parts) == 4 {
			switch parts[2] {
			case "+":
				m.op = add{op1: parts[1], op2: parts[3]}
			case "-":
				m.op = sub{op1: parts[1], op2: parts[3]}
			case "*":
				m.op = mult{op1: parts[1], op2: parts[3]}
			case "/":
				m.op = div{op1: parts[1], op2: parts[3]}
			default:
				panic("bad op " + parts[2])
			}
		}

		monkeys[m.name] = m
	}

	v, ops := monkeys["root"].Eval(monkeys)

	fmt.Printf("root yells (part1): %d\n", v)

	fmt.Printf("equation (to wolfram alpha)\n%s\n", ops)

}
