package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/noseglid/advent-of-code/util"
)

type Op rune
type Evaler func(s string) int

const (
	Add  = Op('+')
	Mult = Op('*')
)

func getNumberAt(s string, index int) (int, int) {
	endIndex := index + 1
	for endIndex < len(s) && unicode.IsDigit(rune(s[endIndex])) {
		endIndex++
	}

	return util.MustAtoi(s[index:endIndex]), endIndex - index
}

var plusRe = regexp.MustCompile(`(\d+) \+ (\d+)`)
var multRe = regexp.MustCompile(`(\d+) \* (\d+)`)

func evalNonParenthesisp2(s string) int {
	for {
		m := plusRe.FindStringSubmatchIndex(s)
		if len(m) == 0 {
			break
		}
		op1 := util.MustAtoi(s[m[2]:m[3]])
		op2 := util.MustAtoi(s[m[4]:m[5]])
		newS := fmt.Sprintf("%s%d%s", s[:m[0]], op1+op2, s[m[1]:])
		s = newS
	}

	for {
		m := multRe.FindStringSubmatchIndex(s)
		if len(m) == 0 {
			break
		}
		op1 := util.MustAtoi(s[m[2]:m[3]])
		op2 := util.MustAtoi(s[m[4]:m[5]])
		newS := fmt.Sprintf("%s%d%s", s[:m[0]], op1*op2, s[m[1]:])
		s = newS
	}

	return util.MustAtoi(s)
}

func evalNonParanthesis(s string) int {
	first := true
	n := 0
	op := Add
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if first {
				var adv int
				n, adv = getNumberAt(s, i)
				i += adv
				first = false
			} else {
				switch op {
				case Add:
					lhs, adv := getNumberAt(s, i)
					i += adv
					n += lhs
				case Mult:
					lhs, adv := getNumberAt(s, i)
					i += adv
					n *= lhs
				}
			}
		case '+', '*':
			op = Op(s[i])
		}
	}

	return n
}

func reduceExpression(s string, eval Evaler) (string, bool) {
	if !strings.ContainsRune(s, '(') {
		return s, false
	}

	start := 0
	for i, r := range s {
		switch r {
		case '(':
			start = i
		case ')':
			return fmt.Sprintf("%s%d%s", s[:start], eval(s[start+1:i]), s[i+1:]), true
		}
	}

	log.Fatalf("invalid reduce expression: %s", s)
	return "", false
}

func evalExpression(s string, eval Evaler) int {
	if !strings.ContainsRune(s, '(') {
		return eval(s)
	}

	for {
		expr, ok := reduceExpression(s, eval)
		if !ok {
			break
		}
		s = expr
	}

	return eval(s)
}

func main() {

	// log.Printf("s=%d", evalExpression("1 + 2 * 3 + 4 * 5 + 6", evalNonParenthesisp2))

	s := util.FileScanner("2020/Day18/input", bufio.ScanLines)
	result := 0
	resultp2 := 0
	for s.Scan() {
		result += evalExpression(s.Text(), evalNonParanthesis)
		resultp2 += evalExpression(s.Text(), evalNonParenthesisp2)
	}

	log.Printf("sum of resulting values (part1): %d", result)
	log.Printf("sum of resulting values (part2): %d", resultp2)

}
