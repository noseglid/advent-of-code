package main

import (
	"fmt"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type P util.Point

var npadInstr = map[rune]P{
	'0': {1, 3},
	'A': {2, 3},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'F': {0, 3},
}

var dpadInstr = map[rune]P{
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
	'^': {1, 0},
	'A': {2, 0},
	'F': {0, 0},
}

func shortestNpad(x, y int, r rune) []rune {
	var res []rune

	tx, ty := npadInstr[r].X, npadInstr[r].Y
	dx, dy := tx-x, ty-y

	v := util.Repeat(util.Or(dy > 0, 'v', '^'), util.Absolute(dy))
	h := util.Repeat(util.Or(dx > 0, '>', '<'), util.Absolute(dx))

	if y == 3 && tx == 0 {
		res = append(res, v...)
		res = append(res, h...)
	} else if x == 0 && ty == 3 {
		res = append(res, h...)
		res = append(res, v...)
	} else if dx < 0 {
		res = append(res, h...)
		res = append(res, v...)
	} else if dx >= 0 {
		res = append(res, v...)
		res = append(res, h...)
	}

	res = append(res, 'A')

	return res
}

func shortestDpad(x, y int, r rune) []rune {
	var res []rune

	tx, ty := dpadInstr[r].X, dpadInstr[r].Y
	dx, dy := tx-x, ty-y

	v := util.Repeat(util.Or(dy > 0, 'v', '^'), util.Absolute(dy))
	h := util.Repeat(util.Or(dx > 0, '>', '<'), util.Absolute(dx))

	if x == 0 && ty == 0 {
		res = append(res, h...)
		res = append(res, v...)
	} else if y == 0 && tx == 0 {
		res = append(res, v...)
		res = append(res, h...)
	} else if dx < 0 {
		res = append(res, h...)
		res = append(res, v...)
	} else if dx >= 0 {
		res = append(res, v...)
		res = append(res, h...)
	}

	res = append(res, 'A')

	return res
}

func codeToInt(s string) int {
	r := regexp.MustCompile("(^[0-9]+)")
	return util.MustAtoi(r.FindString(s))
}

type mkey struct {
	code string
	kp   int
}

func getCount(code []rune, maxKeypads, keypad int, memo map[mkey]int) int {
	mk := mkey{string(code), keypad - 1}
	if v, ok := memo[mk]; ok {
		return v
	}

	seq := []rune{}
	cx, cy := 2, 0
	for _, r := range code {
		seq = append(seq, shortestDpad(cx, cy, r)...)
		cx, cy = dpadInstr[r].X, dpadInstr[r].Y
	}

	if keypad == maxKeypads {
		return len(seq)
	}

	seqs := [][]rune{}
	current := []rune{}
	for _, r := range seq {
		current = append(current, r)
		if r == 'A' {
			seqs = append(seqs, current)
			current = []rune{}
		}
	}

	sum := 0
	for _, seq := range seqs {
		sum += getCount(seq, maxKeypads, keypad+1, memo)
	}

	memo[mk] = sum
	return sum
}

func main() {
	codes := []string{
		// Sample input
		// "029A",
		// "980A",
		// "179A",
		// "456A",
		// "379A",

		// Real input
		"985A",
		"540A",
		"463A",
		"671A",
		"382A",
	}
	s1, s2 := 0, 0
	for _, code := range codes {

		dpad1Seq := []rune{}
		cx, cy := 2, 3
		for _, r := range code {
			dpad1Seq = append(dpad1Seq, shortestNpad(cx, cy, r)...)
			cx, cy = npadInstr[r].X, npadInstr[r].Y
		}
		n1 := getCount(dpad1Seq, 2, 1, make(map[mkey]int))
		n2 := getCount(dpad1Seq, 25, 1, make(map[mkey]int))

		s1 += n1 * codeToInt(code)
		s2 += n2 * codeToInt(code)
	}

	fmt.Printf("Sum of codes complexity with 2 intermediates (part1): %d\n", s1)
	fmt.Printf("Sum of codes complexity with 25 intermediates (part2): %d\n", s2)
}
