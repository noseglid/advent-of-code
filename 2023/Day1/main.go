package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/noseglid/advent-of-code/util"
)

var nm = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func hasDigit(s string, i int) (bool, int) {
	if unicode.IsNumber(rune(s[i])) {
		return true, int(s[i] - '0')
	}

	for digit, v := range nm {
		if len(s) >= i+len(digit) {
			if s[i:i+len(digit)] == digit {
				return true, v
			}
		}
	}

	return false, -1
}

func main() {
	lines := util.GetFileStrings("2023/Day1/input")

	sum := 0
	for _, l := range lines {
		s := strings.IndexFunc(l, unicode.IsNumber)
		e := strings.LastIndexFunc(l, unicode.IsNumber)
		if s == -1 || e == -1 {
			continue
		}
		v, _ := strconv.Atoi(string(l[s]) + string(l[e]))
		sum += v
	}

	sump2 := 0
	for _, l := range lines {
		v1, v2 := -1, -1
	bv1:
		for i := range l {
			if ok, val := hasDigit(l, i); ok {
				v1 = val
				break bv1
			}
		}

	bv2:
		for i := len(l) - 1; i >= 0; i-- {
			if ok, val := hasDigit(l, i); ok {
				v2 = val
				break bv2
			}
		}

		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", v1, v2))
		sump2 += v
	}

	fmt.Printf("sum (part1): %d\n", sum)
	fmt.Printf("sum (part2): %d\n", sump2)
}
