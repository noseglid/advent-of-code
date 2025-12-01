package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func main() {

	pos := 50

	instructions := util.GetFileStrings("2025/Day1/input")

	pass, pass2 := 0, 0
	for _, instr := range instructions {
		d := util.MustAtoi(instr[1:])

		dir := 1
		switch instr[0] {
		case 'L':
			dir = -1
		}

		for i := 0; i < d; i++ {
			pos += dir
			if pos < 0 {
				pos += 100
			}
			if pos >= 100 {
				pos -= 100
			}
			if pos == 0 {
				pass2++
			}
		}

		if pos == 0 {
			pass++
		}

		fmt.Printf("ran instruction %s, dial now at  %d\n", instr, pos)

	}

	fmt.Printf("password (part1): %d\n", pass)
	fmt.Printf("password (part2): %d\n", pass2)
}
