package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func pow(x, y int) int {
	r := 1
	for i := 0; i < y; i++ {
		r *= x
	}
	return r
}

func singleSnafu(r rune) int {
	switch r {
	case '2':
		return 2
	case '1':
		return 1
	case '0':
		return 0
	case '-':
		return -1
	case '=':
		return -2
	}
	panic("bad single snafu")
}

func snafuToDecimal(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		r += singleSnafu(rune(s[i])) * pow(5, len(s)-i-1)
	}
	return r
}

func decimalToSnafu(d int) string {
	var digits []int

	for d > 0 {
		d += 2
		digits = append(digits, d%5)
		d /= 5
	}

	chars := []rune{'=', '-', '0', '1', '2'}

	var sb strings.Builder
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]
		sb.WriteRune(chars[d])
	}

	return sb.String()

}

func main() {

	input := util.GetFileStrings("2022/Day25/input")

	s := 0
	for _, l := range input {
		s += (snafuToDecimal(l))
	}
	fmt.Printf("input (part1): %s\n", decimalToSnafu(s))

}
