package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Schema struct {
	target  int
	buttons []int
	joltage []int
}

var reconfig = regexp.MustCompile(`^\[([.#]+)\]`)
var rebuttons = regexp.MustCompile(`\(([\d+,]+)\)`)
var recommalist = regexp.MustCompile(`\d+`)
var joltage = regexp.MustCompile(`{([\d+,]+)}`)
var z3re = regexp.MustCompile(`\(\(+[^\)]+\) (\d+)\)`)

func parseLine(l string) Schema {

	var s Schema
	m := reconfig.FindStringSubmatch(l)
	for i, match := range m[1] {
		if match == '#' {
			s.target |= 1 << i
		}
	}

	buttons := rebuttons.FindAllStringSubmatch(l, -1)
	for _, match := range buttons {
		sm := recommalist.FindAllStringSubmatch(match[1], -1)
		var bb int
		for _, n := range sm {
			bb |= 1 << util.MustAtoi(n[0])
		}

		s.buttons = append(s.buttons, bb)
	}

	joltages := joltage.FindAllStringSubmatch(l, -1)
	for _, match := range joltages {
		sm := recommalist.FindAllStringSubmatch(match[0], -1)
		for _, n := range sm {
			s.joltage = append(s.joltage, util.MustAtoi(n[0]))
		}
	}
	return s
}

func press(current, target, presses int, buttons []int) int {
	if current == target {
		return presses
	}
	if len(buttons) == 0 {
		return math.MaxInt
	}
	with := press(current^buttons[0], target, presses+1, buttons[1:])
	without := press(current, target, presses, buttons[1:])
	return min(with, without)
}

func genz3(s Schema) string {
	var sb strings.Builder

	var buttonssb strings.Builder
	for i := range s.buttons {
		sb.WriteString(fmt.Sprintf("(declare-const b%d Int)\n", i+1))
		buttonssb.WriteString(fmt.Sprintf("b%d", i+1))
		buttonssb.WriteRune(' ')
	}
	for i := range s.buttons {
		sb.WriteString(fmt.Sprintf("(assert (>= b%d 0))\n", i+1))
	}

	for i := range s.joltage {
		var terms []string
		for j, button := range s.buttons {
			if button&(1<<i) != 0 {
				terms = append(terms, fmt.Sprintf("b%d", j+1))
			}
		}

		sum := strings.Join(terms, " ")
		if len(terms) > 1 {
			sum = "(+ " + sum + ")"
		}
		sb.WriteString(fmt.Sprintf("(assert (= %s %d))\n", sum, s.joltage[i]))
	}

	sb.WriteString(fmt.Sprintf("(minimize (+ %s))\n", buttonssb.String()))
	sb.WriteString("(check-sat)\n")
	sb.WriteString("(get-objectives)\n")

	return sb.String()
}

func main() {

	lines := util.GetFileStrings("2025/Day10/input")

	var schemas []Schema

	for _, l := range lines {
		schemas = append(schemas, parseLine(l))
	}

	sum := 0
	for _, s := range schemas {
		sum += press(0, s.target, 0, s.buttons)
	}
	fmt.Printf("minimum number of presses (part1): %d\n", sum)

	defer func() {
		os.Remove("file.smt")
	}()
	sump2 := 0
	for _, s := range schemas {
		z3code := genz3(s)
		os.WriteFile("file.smt", []byte(z3code), 0666)
		cmd := exec.Command("/opt/homebrew/bin/z3", "file.smt")
		output, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		sump2 += util.MustAtoi(z3re.FindStringSubmatch(string(output))[1])
	}
	fmt.Printf("minimum number of presses for joltage (part2): %d\n", sump2)
}
