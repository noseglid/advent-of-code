package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type computer struct {
	id string
	c  []string
}

var memo = map[string][]string{}

func allConnected(computers map[string]*computer, ids []string) bool {
	for _, a := range ids {
		for _, b := range ids {
			if a == b {
				continue
			}

			if !util.Contains(computers[a].c, b) {
				return false
			}
		}
	}
	return true
}

func longest(computers map[string]*computer, current string, visited []string) []string {
	if util.Contains(visited, current) {
		return visited
	}
	if v, ok := memo[current]; ok {
		return v
	}

	nv := append([]string{current}, visited...)

	chain := []string{}
	for id, c := range computers {
		if id == current {
			continue
		}

		if !allConnected(computers, append([]string{id}, nv...)) {
			continue
		}

		t := longest(computers, c.id, nv)
		if len(t) > len(chain) {
			chain = append([]string{}, t...)
		}
	}

	memo[current] = chain
	return chain
}

func (c *computer) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s]:", c.id))
	sep := ""
	for _, s := range c.c {
		sb.WriteString(fmt.Sprintf("%s%s", sep, s))
		sep = ","
	}
	return sb.String()
}
func main() {

	lines := util.GetFileStrings("2024/Day23/input")

	computers := map[string]*computer{}

	for _, l := range lines {
		c1, c2 := l[0:2], l[3:5]
		if _, ok := computers[c1]; ok {
			computers[c1].c = append(computers[c1].c, c2)
		} else {
			computers[c1] = &computer{c1, []string{c2}}
		}
		if _, ok := computers[c2]; ok {
			computers[c2].c = append(computers[c2].c, c1)
		} else {
			computers[c2] = &computer{c2, []string{c1}}
		}
	}

	triplets := map[string][]string{}

	for _, c1 := range computers {
		for _, c2 := range computers {
			if !util.Contains(c1.c, c2.id) {
				continue
			}
			for _, c3 := range computers {
				if util.Contains(c1.c, c3.id) && util.Contains(c2.c, c3.id) {
					ids := []string{c1.id, c2.id, c3.id}
					slices.Sort(ids)
					triplets[strings.Join(ids, ",")] = ids
				}
			}
		}
	}

	n := 0
Outer:
	for _, ids := range triplets {
		for _, id := range ids {
			if id[0] == 't' {
				n++
				continue Outer
			}
		}
	}
	fmt.Printf("Connections with a computer named 't' (part1): %d\n", n)

	var cc []string
	for id := range computers {
		c := longest(computers, id, []string{})
		if len(c) > len(cc) {
			cc = c
		}
	}
	slices.Sort(cc)
	fmt.Printf("Most interconnectd (part2): %s\n", strings.Join(cc, ","))
}
