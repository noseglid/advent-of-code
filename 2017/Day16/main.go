package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func dance(programs []rune, moves []string) {
}

func main() {
	// programs := []rune("abcde")
	programs := []rune("abcdefghijklmnop")
	input := util.GetFile("2017/Day16/input")
	its := 1000000000 - 41666666*24
	fmt.Printf("its=%d\n", its)
	for i := 0; i < its; i++ {
		for _, s := range strings.Split(input, ",") {
			s = strings.TrimSpace(s)
			// fmt.Printf("-- start, move: %s, programs: %s\n", s, string(programs))
			switch s[0] {
			case 's':
				v := len(programs) - util.MustAtoi(s[1:])
				// fmt.Printf("Spin %d\n", v)
				end := programs[v:]
				programs = append(end, programs[:v]...)
			case 'x':
				parts := strings.Split(s[1:], "/")
				i1, i2 := util.MustAtoi(parts[0]), util.MustAtoi(parts[1])
				// fmt.Printf("Exchange %d <-> %d\n", i1, i2)
				programs[i1], programs[i2] = programs[i2], programs[i1]
			case 'p':
				parts := strings.Split(s[1:], "/")
				p1, p2 := parts[0], parts[1]
				i1, i2 := strings.Index(string(programs), p1), strings.Index(string(programs), p2)
				// fmt.Printf("Partner %s <-> %s, %d <-> %d\n", p1, p2, i1, i2)
				programs[i1], programs[i2] = programs[i2], programs[i1]
			}
		}

		if string(programs) == "abcdefghijklmnop" {
			fmt.Printf("Back after %d\n", i)
		}
	}

	fmt.Printf("Order of programs (part1): %s\n", string(programs))

}
