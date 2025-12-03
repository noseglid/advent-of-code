package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func maxJoltageN(line string, batteries int) int {
	d := len(line) - batteries

	var result []rune

	for i := 0; i < len(line); i++ {
		for len(result) > 0 && d > 0 && rune(line[i]) > result[len(result)-1] {
			result = result[:len(result)-1]
			d--
		}
		result = append(result, rune(line[i]))
	}

	result = result[:batteries]

	return util.MustAtoi(string(result))
}

func main() {
	lines := util.GetFileStrings("2025/Day3/input")

	sum := 0
	for _, l := range lines {
		sum += maxJoltageN(l, 2)
	}
	fmt.Printf("max joltage (part1): %d\n", sum)

	sump2 := 0
	for _, l := range lines {
		sump2 += maxJoltageN(l, 12)
	}
	fmt.Printf("max joltage with 12 batteries (part2): %d\n", sump2)
}
