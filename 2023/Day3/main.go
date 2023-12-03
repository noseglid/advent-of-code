package main

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/noseglid/advent-of-code/util"
)

type Symbol struct {
	symbol rune
	x, y   int
}

type Number struct {
	number int
	x      []int
	y      int
}

var numberRegex = regexp.MustCompile(`\d+`)

func coordAdjacent(x1, y1, x2, y2 int) bool {
	return (x2 == x1-1 && y2 == y1-1) || // Top-left
		(x2 == x1 && y2 == y1-1) || // Top
		(x2 == x1+1 && y2 == y1-1) || // Top-Right
		(x2 == x1-1 && y2 == y1) || // Left
		(x2 == x1+1 && y2 == y1) || // Right
		(x2 == x1-1 && y2 == y1+1) || // Bottom-Left
		(x2 == x1 && y2 == y1+1) || // Bottom
		(x2 == x1+1 && y2 == y1+1) // Bottom-Right
}

func numberAdjacentSymbol(n Number, s Symbol) bool {
	for _, x := range n.x {
		if coordAdjacent(x, n.y, s.x, s.y) {
			return true
		}
	}
	return false
}

func isAdjacent(n Number, symbols []Symbol) (Symbol, bool) {
	for _, s := range symbols {
		if numberAdjacentSymbol(n, s) {
			return s, true
		}
	}
	return Symbol{}, false
}

func main() {
	lines := util.GetFileStrings("2023/Day3/input")

	numbers := []Number{}
	symbols := []Symbol{}
	for y, l := range lines {
		number := 0
		mult := 1
		length := 0
		for x := len(l) - 1; x >= 0; x-- {
			c := rune(l[x])
			if unicode.IsNumber(c) {
				number += mult * int(c-'0')
				length++
				mult *= 10
			} else if length > 0 {
				xcoords := []int{x + 1}
				for i := 0; i < length-1; i++ {
					xcoords = append(xcoords, xcoords[i]+1)
				}
				numbers = append(numbers, Number{number, xcoords, y})
				number = 0
				mult = 1
				length = 0
			}
			if !unicode.IsNumber(c) && c != '.' {
				symbols = append(symbols, Symbol{c, x, y})
			}
		}
		if length > 0 {
			xcoords := []int{0}
			for i := 0; i < length-1; i++ {
				xcoords = append(xcoords, xcoords[i]+1)
			}
			numbers = append(numbers, Number{number, xcoords, y})
		}
	}

	sum := 0
	for _, n := range numbers {
		if _, ok := isAdjacent(n, symbols); ok {
			sum += n.number
		}
	}

	sump2 := 0
	for _, s := range symbols {
		if s.symbol != '*' {
			continue
		}

		gearParts := []int{}
		for _, n := range numbers {
			if numberAdjacentSymbol(n, s) {
				gearParts = append(gearParts, n.number)
			}
		}
		if len(gearParts) == 2 {
			sump2 += gearParts[0] * gearParts[1]
		}
	}

	fmt.Printf("Sum of part numbers (part1): %d\n", sum)
	fmt.Printf("Sum of gear parts (part2): %d\n", sump2)
}
