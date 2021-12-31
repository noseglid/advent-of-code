package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type operation interface {
	Exec(s string) string
}

type swapOperation struct {
	x, y int
}

func (o swapOperation) Exec(s string) string {
	rr := []rune(s)
	rr[o.x], rr[o.y] = rr[o.y], rr[o.x]
	return string(rr)
}

type swapLetterOperation struct {
	a, b rune
}

func (o swapLetterOperation) Exec(s string) string {
	return swapOperation{strings.IndexRune(s, o.a), strings.IndexRune(s, o.b)}.Exec(s)
}

type reverseOperation struct {
	x, y int
}

func (o reverseOperation) Exec(s string) string {
	rr := []rune(s)[:o.x]
	rr = append(rr, []rune(util.Reverse(s[o.x:o.y+1]))...)
	if o.y < len(s) {
		rr = append(rr, []rune(s[o.y+1:])...)
	}
	return string(rr)
}

type rotateOperation struct {
	d, n int
}

func clamp(n int, s string) int {
	if n >= len(s) {
		n = n % len(s)
	} else if n < 0 {
		for n < 0 {
			n += len(s)
		}
	}
	return n
}

func (o rotateOperation) Exec(s string) string {
	n := make([]rune, len(s))
	for i, r := range s {
		n[clamp(i+o.d*o.n, s)] = r
	}
	return string(n)
}

type moveOperation struct {
	x, y int
}

func (o moveOperation) Exec(s string) string {
	rr := append([]rune(s)[:o.x], []rune(s)[o.x+1:]...)
	rr = append(rr[:o.y+1], rr[o.y:]...)
	rr[o.y] = rune(s[o.x])
	return string(rr)
}

type rotateLetterOperation struct {
	a rune
}

func (o rotateLetterOperation) Exec(s string) string {
	i := strings.IndexRune(s, o.a)
	n := 1 + i
	if i >= 4 {
		n++
	}
	return rotateOperation{1, n}.Exec(s)
}

func parseOp(s string) operation {
	swap := &swapOperation{}
	if n, err := fmt.Sscanf(s, "swap position %d with position %d", &swap.x, &swap.y); err == nil && n == 2 {
		return swap
	}
	swapL := &swapLetterOperation{}
	if n, err := fmt.Sscanf(s, "swap letter %c with letter %c", &swapL.a, &swapL.b); err == nil && n == 2 {
		return swapL
	}
	rotate := &rotateOperation{d: 1}
	var dir string
	var r rune
	if n, _ := fmt.Sscanf(s, "rotate %s %d step%c", &dir, &rotate.n, &r); n >= 2 {
		if dir == "left" {
			rotate.d = -1
		}
		return rotate
	}
	rotateL := &rotateLetterOperation{}
	if n, err := fmt.Sscanf(s, "rotate based on position of letter %c", &rotateL.a); err == nil && n == 1 {
		return rotateL
	}
	reverse := &reverseOperation{}
	if n, err := fmt.Sscanf(s, "reverse positions %d through %d", &reverse.x, &reverse.y); err == nil && n == 2 {
		return reverse
	}
	move := &moveOperation{}
	if n, err := fmt.Sscanf(s, "move position %d to position %d", &move.x, &move.y); err == nil && n == 2 {
		return move
	}

	panic("cannot parse")
}

func scramble(ops []operation, s string) string {
	for _, o := range ops {
		s = o.Exec(s)
	}
	return s
}

func main() {

	// o := moveOperation{2, 5}
	// s := o.Exec("abcdefgh")
	// log.Print(s)
	// log.Print(o.RExec(s))

	input := "2016/Day21/input"
	var ops []operation
	for _, l := range util.GetFileStrings(input) {
		op := parseOp(l)
		ops = append(ops, op)
	}
	log.Printf("Part 1: Scrambled password: %s", scramble(ops, "abcdefgh"))

	search := "fbgdceah"
	util.PermString([]string{"a", "b", "c", "d", "e", "f", "g", "h"}, func(s []string) {
		ss := strings.Join(s, "")
		if search == scramble(ops, ss) {
			log.Printf("Part 2: String that scrambles int %s: %s", search, strings.Join(s, ""))
		}
	})
}
