package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func main() {
	input := "2017/Day5/input"
	instr := util.GetFileNumbers(input)
	instr2 := util.GetFileNumbers(input)
	pc := 0
	n := 0
	for {
		if pc >= len(instr) {
			break
		}
		jump := instr[pc]
		instr[pc]++
		pc += jump
		n++
	}
	log.Printf("Part 1: Escapes in steps: %d", n)

	pc2 := 0
	n2 := 0
	for {
		if pc2 >= len(instr2) {
			break
		}
		jump := instr2[pc2]
		if jump >= 3 {
			instr2[pc2]--
		} else {
			instr2[pc2]++
		}
		pc2 += jump
		n2++

	}
	log.Printf("Part 2: Escapes in steps: %d", n2)

}
