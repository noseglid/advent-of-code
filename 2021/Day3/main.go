package main

import (
	"log"
	"strconv"

	"github.com/noseglid/advent-of-code/util"
)

func binstringmult(s1, s2 string) int {
	n1, err := strconv.ParseInt(s1, 2, 64)
	if err != nil {
		panic(err)
	}
	n2, err := strconv.ParseInt(s2, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(n1 * n2)
}

func part1() {
	lines := util.GetFileStrings("2021/Day3/input")

	gamma, epsilon := "", ""
	for i := 0; i < len(lines[0]); i++ {
		n0, n1 := 0, 0
		for _, s := range lines {
			switch s[i] {
			case '0':
				n0++
			case '1':
				n1++
			}
		}

		if n0 > n1 {
			gamma += "0"
			epsilon += "1"
		} else if n1 > n0 {
			gamma += "1"
			epsilon += "0"
		} else {
			log.Printf("Draw what to do?")
		}
	}

	log.Printf("Part 1: gamma*epsilon: %d", binstringmult(gamma, epsilon))
}

func oxygen(bits []rune) []int {
	n0, n1 := 0, 0
	for _, b := range bits {
		switch b {
		case '0':
			n0++
		case '1':
			n1++
		}
	}

	toKeep := '1'
	if n0 > n1 {
		toKeep = '0'
	}

	var lines []int
	for i, b := range bits {
		if b == toKeep {
			lines = append(lines, i)
		}
	}

	return lines
}

func carbonDioxide(bits []rune) []int {
	n0, n1 := 0, 0
	for _, b := range bits {
		switch b {
		case '0':
			n0++
		case '1':
			n1++
		}
	}

	toKeep := '0'
	if n1 < n0 {
		toKeep = '1'
	}

	var lines []int
	for i, b := range bits {
		if b == toKeep {
			lines = append(lines, i)
		}
	}

	return lines
}

func cpy(list []string) []string {
	r := make([]string, len(list))
	copy(r, list)
	return r
}

func part2() {
	lines := util.GetFileStrings("2021/Day3/input")
	ll := len(lines[0])
	olines := cpy(lines)
	colines := cpy(lines)

	for i := 0; i < ll; i++ {
		if len(olines) > 1 {
			var bits []rune
			for _, s := range olines {
				bits = append(bits, rune(s[i]))
			}
			nn := []string{}
			keep := oxygen(bits)
			for _, l := range keep {
				nn = append(nn, olines[l])
			}
			olines = nn
		}

		if len(colines) > 1 {
			var bits []rune
			for _, s := range colines {
				bits = append(bits, rune(s[i]))
			}
			nn := []string{}
			keep := carbonDioxide(bits)
			for _, l := range keep {
				nn = append(nn, colines[l])
			}
			colines = nn
		}
	}

	log.Printf("Part 2: oxygen * co2 ratings: %d", binstringmult(olines[0], colines[0]))
}

func main() {
	part1()
	part2()
}
