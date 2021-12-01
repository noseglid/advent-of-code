package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/noseglid/advent-of-code/util"
)

type wireValueFn func(wire string) (uint16, bool)

type INInstruction struct {
	in1, out string

	outputValue    uint16
	hasOutputValue bool
}

func plainOrRetrieve(fn wireValueFn, s string) (uint16, bool) {
	value, err := strconv.Atoi(s)
	if err == nil {
		return uint16(value), true
	}

	return fn(s)
}

func (i *INInstruction) Apply(fn wireValueFn) {
	inputValue, ok := plainOrRetrieve(fn, i.in1)
	if !ok {
		return
	}

	i.outputValue = inputValue
	i.hasOutputValue = true
}

func (i *INInstruction) Output() (uint16, string, bool) {
	return i.outputValue, i.out, i.hasOutputValue
}

type ANDInstruction struct {
	in1, in2 string
	out      string

	outputValue    uint16
	hasOutputValue bool
}

func (i *ANDInstruction) Apply(fn wireValueFn) {
	v1, ok1 := plainOrRetrieve(fn, i.in1)
	v2, ok2 := plainOrRetrieve(fn, i.in2)
	if !ok1 || !ok2 {
		return
	}

	i.outputValue = v1 & v2
	i.hasOutputValue = true
}

func (i *ANDInstruction) Output() (uint16, string, bool) {
	return i.outputValue, i.out, i.hasOutputValue
}

type ORInstruction struct {
	in1, in2, out string

	outputValue    uint16
	hasOutputValue bool
}

func (i *ORInstruction) Apply(fn wireValueFn) {
	v1, ok1 := plainOrRetrieve(fn, i.in1)
	v2, ok2 := plainOrRetrieve(fn, i.in2)
	if !ok1 || !ok2 {
		return
	}

	i.outputValue = v1 | v2
	i.hasOutputValue = true
}

func (i *ORInstruction) Output() (uint16, string, bool) {
	return i.outputValue, i.out, i.hasOutputValue
}

type LSHIFTInstruction struct {
	in1, in2, out string

	outputValue    uint16
	hasOutputValue bool
}

func (i *LSHIFTInstruction) Apply(fn wireValueFn) {
	v1, ok1 := plainOrRetrieve(fn, i.in1)
	v2, ok2 := plainOrRetrieve(fn, i.in2)
	if !ok1 || !ok2 {
		return
	}

	i.hasOutputValue = true
	i.outputValue = v1 << v2
}

func (i *LSHIFTInstruction) Output() (uint16, string, bool) {
	return i.outputValue, i.out, i.hasOutputValue
}

type RSHIFTInstruction struct {
	in1, in2, out string

	outputValue    uint16
	hasOutputValue bool
}

func (i *RSHIFTInstruction) Apply(fn wireValueFn) {
	v1, ok1 := plainOrRetrieve(fn, i.in1)
	v2, ok2 := plainOrRetrieve(fn, i.in2)
	if !ok1 || !ok2 {
		return
	}

	i.hasOutputValue = true
	i.outputValue = v1 >> v2
}

func (i *RSHIFTInstruction) Output() (uint16, string, bool) {
	return i.outputValue, i.out, i.hasOutputValue
}

type NOTInstruction struct {
	in1, out string

	outputValue    uint16
	hasOutputValue bool
}

func (i *NOTInstruction) Apply(fn wireValueFn) {
	v1, ok1 := plainOrRetrieve(fn, i.in1)
	if !ok1 {
		return
	}

	i.hasOutputValue = true
	i.outputValue = ^v1
}

func (i *NOTInstruction) Output() (uint16, string, bool) {
	return i.outputValue, i.out, i.hasOutputValue
}

// 123 -> x
var reInput = regexp.MustCompile(`^([a-z0-9]+) -> ([a-z]+)$`)

// x AND y -> d
var reAndOr = regexp.MustCompile(`^([a-z0-9]+) (AND|OR) ([a-z0-9]+) -> ([a-z]+)$`)

// lf RSHIFT 2 -> lg
var reShift = regexp.MustCompile(`^([a-z]+) (LSHIFT|RSHIFT) ([a-z0-9]+) -> ([a-z]+)$`)

// NOT gs -> gt
var reNot = regexp.MustCompile(`^NOT ([a-z]+) -> ([a-z]+)$`)

func parseInstruction(s string) Instruction {
	match := reInput.FindAllStringSubmatch(s, 1)
	if len(match) > 0 {
		return &INInstruction{
			in1: match[0][1],
			out: match[0][2],
		}
	}

	match = reAndOr.FindAllStringSubmatch(s, 1)
	if len(match) > 0 {
		if match[0][2] == "AND" {
			return &ANDInstruction{
				in1: match[0][1],
				in2: match[0][3],
				out: match[0][4],
			}
		} else if match[0][2] == "OR" {
			return &ORInstruction{
				in1: match[0][1],
				in2: match[0][3],
				out: match[0][4],
			}
		}
	}

	match = reShift.FindAllStringSubmatch(s, 1)
	if len(match) > 0 {
		if match[0][2] == "LSHIFT" {
			return &LSHIFTInstruction{
				in1: match[0][1],
				in2: match[0][3],
				out: match[0][4],
			}
		} else if match[0][2] == "RSHIFT" {
			return &RSHIFTInstruction{
				in1: match[0][1],
				in2: match[0][3],
				out: match[0][4],
			}
		}
	}

	match = reNot.FindAllStringSubmatch(s, 1)
	if len(match) > 0 {
		return &NOTInstruction{
			in1: match[0][1],
			out: match[0][2],
		}
	}

	panic(fmt.Sprintf("unable to parse instruction: %s", s))
}

type Instruction interface {
	Apply(wireValueFn)
	Output() (uint16, string, bool)
}

func part1() uint16 {
	s := util.FileScanner("2015/Day7/input", bufio.ScanLines)

	var circuit []Instruction
	for s.Scan() {
		circuit = append(circuit, parseInstruction(s.Text()))
	}

	wireValues := map[string]uint16{}

	var wirefn wireValueFn = func(wire string) (uint16, bool) {
		v, ok := wireValues[wire]
		return v, ok
	}

	again := true
	for again {
		again = false
		for _, instr := range circuit {
			instr.Apply(wirefn)
			if outputValue, wire, ok := instr.Output(); ok {
				if _, ok := wireValues[wire]; !ok {
					wireValues[wire] = outputValue
				}
			} else {
				// No value yet, need another loop
				again = true
			}
		}
	}

	log.Printf("value of wire a (part1): %d", wireValues["a"])
	return wireValues["a"]
}

func part2(init uint16) {
	s := util.FileScanner("2015/Day7/input", bufio.ScanLines)

	var circuit []Instruction
	for s.Scan() {
		circuit = append(circuit, parseInstruction(s.Text()))
	}

	wireValues := map[string]uint16{
		"b": init,
	}
	var wirefn wireValueFn = func(wire string) (uint16, bool) {
		v, ok := wireValues[wire]
		return v, ok
	}

	again := true
	for again {
		again = false
		for _, instr := range circuit {
			instr.Apply(wirefn)
			if outputValue, wire, ok := instr.Output(); ok {
				if _, ok := wireValues[wire]; !ok {
					wireValues[wire] = outputValue
				}
			} else {
				// No value yet, need another loop
				again = true
			}
		}
	}

	log.Printf("value of wire a (part2): %d", wireValues["a"])
}

func main() {
	part2(part1())
}
