package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func startOfX(l []rune) bool {
	m := map[rune]struct{}{}
	for _, r := range l {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func main() {
	// input := "nppdvjthqldpwncqszvftbrmjlhg"
	// input := util.GetFile("2022/Day6/input")
	input := util.MustDailyInput(2022, 6)

	p, m := false, false
	for i := 3; i < len(input); i++ {
		if !p && startOfX([]rune(input[util.MaxInt(0, i-4+1):i+1])) {
			fmt.Printf("start-of-packet (part1): %d\n", i+1)
			p = true
		}
		if !m && startOfX([]rune(input[util.MaxInt(0, i-14+1):i+1])) {
			fmt.Printf("start-of-message (part2): %d\n", i+1)
			m = true
		}
		if m && p {
			break
		}
	}

}
