package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func clean(s string) (string, int) {
	removed := 0
	var sb strings.Builder
	inGarbage, escaped := false, false
	for _, c := range s {
		if escaped {
			escaped = false
			continue
		}

		switch c {
		case '!':
			escaped = true
		case '<':
			if inGarbage {
				removed++
			}
			inGarbage = true
		case '>':
			if inGarbage {
				inGarbage = false
			}
		default:
			if !inGarbage {
				sb.WriteRune(c)
			} else {
				removed++
			}
		}
	}

	return sb.String(), removed
}

func score(s string) int {
	score := 0
	n := 1
	for _, c := range s {
		switch c {
		case '{':
			score += n
			n++
		case '}':
			n--
		}
	}

	return score
}

func main() {
	input := "2017/Day9/input"
	data := util.GetFile(input)

	c, r := clean(data)
	s := score(c)
	fmt.Printf("Part 1: %d\n", s)
	fmt.Printf("Part 2: %d\n", r)
}
