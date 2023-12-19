package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type attr rune

const (
	None attr = 0
	X    attr = 'x'
	M    attr = 'm'
	A    attr = 'a'
	S    attr = 's'
)

type cond rune

const (
	none cond = 0
	lt   cond = '<'
	gt   cond = '>'
)

var condOp = map[cond]func(int, int) bool{
	lt: func(i1, i2 int) bool { return i1 < i2 },
	gt: func(i1, i2 int) bool { return i1 > i2 },
}

type rule struct {
	Attr   attr
	Cond   cond
	Value  int
	Target string
}

func (r rule) String() string {
	if r.Attr != None && r.Cond != none {
		return fmt.Sprintf("%c%c%d:%s", r.Attr, r.Cond, r.Value, r.Target)
	}
	return r.Target
}

func (r rule) Matches(p part) bool {
	if r.Attr == None {
		return true
	}

	cmp := 0
	switch r.Attr {
	case X:
		cmp = p.x
	case M:
		cmp = p.m
	case A:
		cmp = p.a
	case S:
		cmp = p.s
	default:
		panic("bad attr")
	}
	return condOp[r.Cond](cmp, r.Value)
}

func (r rule) Split(p rangePart) (rangePart, rangePart) {
	rest := p
	switch r.Attr {
	case X:
		if r.Cond == '<' {
			p.endx = r.Value - 1
			rest.beginx = r.Value
		} else {
			p.beginx = r.Value + 1
			rest.endx = r.Value
		}
	case M:
		if r.Cond == '<' {
			p.endm = r.Value - 1
			rest.beginm = r.Value
		} else {
			p.beginm = r.Value + 1
			rest.endm = r.Value
		}
	case A:
		if r.Cond == '<' {
			p.enda = r.Value - 1
			rest.begina = r.Value
		} else {
			p.begina = r.Value + 1
			rest.enda = r.Value
		}
	case S:
		if r.Cond == '<' {
			p.ends = r.Value - 1
			rest.begins = r.Value
		} else {
			p.begins = r.Value + 1
			rest.ends = r.Value
		}
	case None:
		// don't change, any value hits target
	default:
		panic("bad attr")
	}
	return p, rest
}

type Workflow struct {
	Name  string
	Rules []rule
}

func (w Workflow) next(p part) string {
	for _, r := range w.Rules {
		if r.Matches(p) {
			return r.Target
		}
	}
	panic("no matching rule")
}

type rsplit struct {
	p      rangePart
	target string
}

func (w Workflow) split(p rangePart) []rsplit {
	var splits []rsplit
	for _, r := range w.Rules {
		p1, p2 := r.Split(p)
		splits = append(splits, rsplit{p: p1, target: r.Target})
		p = p2
	}
	return splits
}

type part struct {
	x, m, a, s int
}

type rangePart struct {
	beginx, endx int
	beginm, endm int
	begina, enda int
	begins, ends int
}

func (r rangePart) distinct() int {
	// return (r.endx - r.beginx) * (r.endm - r.beginm) * (r.enda - r.begina) * (r.ends - r.begins)
	return (r.endx - r.beginx + 1) * (r.endm - r.beginm + 1) * (r.enda - r.begina + 1) * (r.ends - r.begins + 1)
}

var ruleRegexp = regexp.MustCompile(`(x|m|a|s)(<|>)(\d+):([a-zAR]+)`)
var workflowRegexp = regexp.MustCompile(`([a-z]+){([^}]+)}`)

func parseRule(str string) rule {
	mm := ruleRegexp.FindStringSubmatch(str)
	if len(mm) == 0 {
		return rule{Target: str}
	}

	return rule{
		Attr:   attr(mm[1][0]),
		Cond:   cond(mm[2][0]),
		Value:  util.MustAtoi(mm[3]),
		Target: mm[4],
	}
}

func parseWorkflow(line string) Workflow {
	m := workflowRegexp.FindStringSubmatch(line)

	var rules []rule
	for _, r := range strings.Split(m[2], ",") {
		rules = append(rules, parseRule(r))
	}

	return Workflow{
		Name:  m[1],
		Rules: rules,
	}
}

var partRegexp = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

func parsePart(line string) part {
	m := partRegexp.FindStringSubmatch(line)

	return part{
		x: util.MustAtoi(m[1]),
		m: util.MustAtoi(m[2]),
		a: util.MustAtoi(m[3]),
		s: util.MustAtoi(m[4]),
	}
}

func recurseParts(workflows map[string]Workflow, target string, p rangePart) int {
	switch target {
	case "R":
		return 0
	case "A":
		return p.distinct()
	}

	split := workflows[target].split(p)
	s := 0
	for _, ss := range split {
		s += recurseParts(workflows, ss.target, ss.p)
	}
	return s
}

func main() {
	lines := util.GetFileStrings("2023/Day19/input")

	parts := []part{}
	workflows := map[string]Workflow{}
	parseWorkflows := true
	for _, l := range lines {
		if l == "" {
			parseWorkflows = false
			continue
		}
		if parseWorkflows {
			wf := parseWorkflow(l)
			workflows[wf.Name] = wf
		} else {
			parts = append(parts, parsePart(l))
		}
	}

	accepted := []part{}

	for _, p := range parts {
		current := workflows["in"]
		for {
			next := current.next(p)
			if next == "A" {
				accepted = append(accepted, p)
				break
			} else if next == "R" {
				break
			}
			current = workflows[next]
		}
	}

	s := 0
	for _, a := range accepted {
		s += a.x + a.m + a.a + a.s
	}
	fmt.Printf("Sum of accepted parts (part1): %d\n", s)

	start := rangePart{
		beginx: 1, endx: 4000,
		beginm: 1, endm: 4000,
		begina: 1, enda: 4000,
		begins: 1, ends: 4000,
	}
	fmt.Printf("Distinct parts (part2): %d\n", recurseParts(workflows, "in", start))
}
