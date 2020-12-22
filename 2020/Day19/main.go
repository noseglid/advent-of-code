package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type RuleMatcher interface {
	Match(msg string, rules map[int]RuleMatcher) (bool, int)
}

type StaticRule struct {
	char string
}

func (r *StaticRule) Match(msg string, rules map[int]RuleMatcher) (bool, int) {
	// log.Printf("matching static rule, %c == %s", msg[0], r.char)
	if len(msg) == 0 {
		return false, 0
	}
	return msg[:1] == r.char, 1
}

type CompositeRule struct {
	rules []int
}

func (r *CompositeRule) Match(msg string, rules map[int]RuleMatcher) (bool, int) {
	// log.Printf("matching composite rule, with %v", r.rules)
	c := msg
	consumed := 0
	for _, rule := range r.rules {
		// log.Printf("testing with msg=%s", c)
		if m, cc := rules[rule].Match(c, rules); m {
			consumed += cc
			c = c[cc:]
		} else {
			return false, 0
		}
	}

	return true, consumed
}

type OrRule struct {
	lhs, rhs RuleMatcher
}

func (r *OrRule) Match(msg string, rules map[int]RuleMatcher) (bool, int) {
	// log.Printf("matching or rule")
	if m, c := r.lhs.Match(msg, rules); m {
		return true, c
	}

	if m, c := r.rhs.Match(msg, rules); m {
		return true, c
	}

	return false, 0
}

var staticRe = regexp.MustCompile(`^(\d+): "([a-z])"$`)
var compositeRe = regexp.MustCompile(`^(\d+): ([0-9 ]+)$`)
var orRe = regexp.MustCompile(`^(\d+): ([0-9 |]+)$`)

func parseRule(s string) (int, RuleMatcher) {
	if m := staticRe.FindStringSubmatch(s); len(m) > 0 {
		return util.MustAtoi(m[1]), &StaticRule{char: m[2]}
	}

	if m := compositeRe.FindStringSubmatch(s); len(m) > 0 {
		r := &CompositeRule{}
		for _, rr := range strings.Fields(m[2]) {
			r.rules = append(r.rules, util.MustAtoi(rr))
		}
		return util.MustAtoi(m[1]), r
	}

	if m := orRe.FindStringSubmatch(s); len(m) > 0 {
		parts := strings.Split(m[2], "|")
		cr1 := &CompositeRule{}
		cr2 := &CompositeRule{}

		for _, d := range strings.Fields(parts[0]) {
			cr1.rules = append(cr1.rules, util.MustAtoi(d))
		}
		for _, d := range strings.Fields(parts[1]) {
			cr2.rules = append(cr2.rules, util.MustAtoi(d))
		}

		return util.MustAtoi(m[1]), &OrRule{cr1, cr2}
	}

	panic("no matching rule!")
}

const fakeN = 100

func fakeRule8String() string {
	var b strings.Builder
	for i := 0; i < fakeN; i++ {
		if i > 0 {
			b.WriteString("| ")
		}
		for j := 0; j <= i; j++ {
			b.WriteString("42 ")
		}

	}

	return b.String()
}

func fakeRule11String() string {
	var b strings.Builder
	for i := 0; i < fakeN; i++ {
		if i > 0 {
			b.WriteString("| ")
		}

		for j := 0; j <= i; j++ {
			b.WriteString("42 ")
		}

		for j := 0; j <= i; j++ {
			b.WriteString("31 ")
		}
	}

	return b.String()
}

func main() {
	input := util.GetFile("2020/Day19/input")
	chunks := strings.Split(input, "\n\n")

	rules := map[int]RuleMatcher{}

	for _, rule := range strings.Split(chunks[0], "\n") {
		id, rule := parseRule(rule)
		rules[id] = rule
	}

	matched := 0
	for _, msg := range strings.Fields(chunks[1]) {
		if m, c := rules[0].Match(msg, rules); m && c == len(msg) {
			matched++
		}
	}
	_, rule8 := parseRule(fmt.Sprintf("8: %s", fakeRule8String()))
	rules[8] = rule8
	_, rule11 := parseRule(fmt.Sprintf("11: %s", fakeRule11String()))
	rules[11] = rule11

	matchedp2 := 0
	for _, msg := range strings.Fields(chunks[1]) {
		if m, c := rules[0].Match(msg, rules); m && c == len(msg) {
			matchedp2++
		}
	}

	log.Printf("messages matched (part1): %d", matched)
	log.Printf("messages matched (part2): %d", matchedp2)
}
