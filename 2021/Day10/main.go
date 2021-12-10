package main

import (
	"log"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

var illegalScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompleteScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func calcAutocompleteScore(s []rune) int {
	score := 0
	for _, r := range s {
		score *= 5
		score += autocompleteScore[r]
	}

	return score
}

func isCorrupt(s string) (bool, rune, []rune) {
	var stack []rune
	for _, r := range s {
		switch r {
		case ')':
			if stack[len(stack)-1] != '(' {
				return true, r, nil
			}
			stack = stack[:len(stack)-1]
		case ']':
			if stack[len(stack)-1] != '[' {
				return true, r, nil
			}
			stack = stack[:len(stack)-1]
		case '}':
			if stack[len(stack)-1] != '{' {
				return true, r, nil
			}
			stack = stack[:len(stack)-1]
		case '>':
			if stack[len(stack)-1] != '<' {
				return true, r, nil
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, r)
		}
	}

	return false, ' ', stack
}

func completionString(stack []rune) []rune {
	var res []rune
	for i := len(stack) - 1; i >= 0; i-- {
		switch stack[i] {
		case '(':
			res = append(res, ')')
		case '[':
			res = append(res, ']')
		case '{':
			res = append(res, '}')
		case '<':
			res = append(res, '>')
		}
	}
	return res
}

func main() {
	input := "2021/Day10/input"

	lines := util.GetFileStrings(input)

	s := 0
	var acScores []int
	for _, l := range lines {
		if corrupt, r, stack := isCorrupt(l); corrupt {
			s += illegalScore[r]
		} else {
			acScores = append(acScores, calcAutocompleteScore(completionString(stack)))
		}
	}

	sort.Ints(acScores)
	log.Printf("Part 1: Syntax error illegal score sum: %d", s)
	log.Printf("Part 2: Middle score: %d", acScores[len(acScores)/2])
}
