package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type op interface {
	Apply(old int) int
}

type mult struct {
	lhsType rune
	lhs     int
	rhsType rune
	rhs     int
}

func (m mult) Apply(old int) int {
	lhs := m.lhs
	if m.lhsType == 'v' {
		lhs = old
	}
	rhs := m.rhs
	if m.rhsType == 'v' {
		rhs = old
	}
	return lhs * rhs
}

type add struct {
	lhsType rune
	lhs     int
	rhsType rune
	rhs     int
}

func (m add) Apply(old int) int {
	lhs := m.lhs
	if m.lhsType == 'v' {
		lhs = old
	}
	rhs := m.rhs
	if m.rhsType == 'v' {
		rhs = old
	}
	return lhs + rhs
}

type monkey struct {
	id          int
	n           int
	items       []int
	testDiv     int
	op          op
	trueMonkey  int
	falseMonkey int
}

func (m monkey) String() string {
	return fmt.Sprintf("Monkey %d inspected %d items. Hold: %v", m.id, m.n, m.items)
}

func parseOp(str string) op {
	parts := strings.Split(str, " ")
	lhs, op, rhs := parts[2], parts[3], parts[4]
	switch op {
	case "*":
		lhsType, rhsType := 'c', 'c'
		lhsv, rhsv := 0, 0
		if lhs == "old" {
			lhsType = 'v'
		} else {
			lhsv = util.MustAtoi(lhs)
		}
		if rhs == "old" {
			rhsType = 'v'
		} else {
			rhsv = util.MustAtoi(rhs)
		}

		return mult{lhsType, lhsv, rhsType, rhsv}
	case "+":
		lhsType, rhsType := 'c', 'c'
		lhsv, rhsv := 0, 0
		if lhs == "old" {
			lhsType = 'v'
		} else {
			lhsv = util.MustAtoi(lhs)
		}
		if rhs == "old" {
			rhsType = 'v'
		} else {
			rhsv = util.MustAtoi(rhs)
		}

		return add{lhsType, lhsv, rhsType, rhsv}
	}

	panic("bad op " + op)
}

func parseMonkey(input []string) monkey {
	i := 0
	var m monkey
	if _, err := fmt.Sscanf(input[i], "Monkey %d:", &m.id); err != nil {
		panic(err)
	}
	i++

	for _, s := range strings.Split(input[i][len("  Starting items: "):], ", ") {
		m.items = append(m.items, util.MustAtoi(s))
	}
	i++

	m.op = parseOp(input[i][len("  Operation: "):])
	i++

	if _, err := fmt.Sscanf(input[i], "  Test: divisible by %d", &m.testDiv); err != nil {
		panic(err)
	}
	i++

	if _, err := fmt.Sscanf(input[i], "    If true: throw to monkey %d", &m.trueMonkey); err != nil {
		panic(err)
	}
	i++

	if _, err := fmt.Sscanf(input[i], "    If false: throw to monkey %d", &m.falseMonkey); err != nil {
		panic(err)
	}
	i++

	return m
}

func run(monkeys []monkey, rounds int, doDiv bool, lcm int) int {
	for i := 0; i < rounds; i++ {
		for id, m := range monkeys {
			for _, item := range m.items {
				monkeys[id].n++
				w := m.op.Apply(item)
				if doDiv {
					w /= 3
				}

				w %= lcm

				recipient := m.falseMonkey
				if w%m.testDiv == 0 {
					recipient = m.trueMonkey
				}
				monkeys[recipient].items = append(monkeys[recipient].items, w)
			}
			monkeys[id].items = m.items[:0]
		}
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].n > monkeys[j].n
	})

	return monkeys[0].n * monkeys[1].n
}

func main() {
	input := util.GetFileStrings("2022/Day11/input")

	var monkeys, monkeysp2 []monkey
	for i := 0; i < len(input); i += 7 {
		monkeys = append(monkeys, parseMonkey(input[i:i+6]))
		monkeysp2 = append(monkeysp2, parseMonkey(input[i:i+6]))
	}

	lcm := 1
	for _, m := range monkeys {
		lcm *= m.testDiv
	}

	p1 := run(monkeys, 20, true, lcm)
	fmt.Printf("Inspections among top monkeys multiplied (part1): %d\n", p1)

	p2 := run(monkeysp2, 10000, false, lcm)
	fmt.Printf("Inspections among top monkeys multiplied (part2): %d\n", p2)
}
