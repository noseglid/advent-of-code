package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func score(opp, me rune) int {
	switch opp {
	case 'A':
		switch me {
		case 'X':
			return 3 + 1
		case 'Y':
			return 6 + 2
		case 'Z':
			return 0 + 3
		}
	case 'B':
		switch me {
		case 'X':
			return 0 + 1
		case 'Y':
			return 3 + 2
		case 'Z':
			return 6 + 3
		}
	case 'C':
		switch me {
		case 'X':
			return 6 + 1
		case 'Y':
			return 0 + 2
		case 'Z':
			return 3 + 3
		}
	}

	panic("bad data")
}

func score2(opp, me rune) int {
	switch opp {
	case 'A':
		switch me {
		case 'X':
			return 0 + 3
		case 'Y':
			return 3 + 1
		case 'Z':
			return 6 + 2
		}
	case 'B':
		switch me {
		case 'X':
			return 0 + 1
		case 'Y':
			return 3 + 2
		case 'Z':
			return 6 + 3
		}
	case 'C':
		switch me {
		case 'X':
			return 0 + 2
		case 'Y':
			return 3 + 3
		case 'Z':
			return 6 + 1
		}
	}

	panic("bad data")
}
func main() {
	lines := util.GetFileStrings("2022/Day2/input")

	s1, s2 := 0, 0
	for _, l := range lines {
		var opp, me rune
		fmt.Sscanf(l, "%c %c", &opp, &me)
		s1 += score(opp, me)
		s2 += score2(opp, me)
	}
	fmt.Printf("Score (part1): %d\n", s1)
	fmt.Printf("Score (part2): %d\n", s2)

}
