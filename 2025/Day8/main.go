package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type P3 = util.Point3D
type JboxPair struct {
	j1, j2 P3
}

type circuit struct {
	jboxes []P3
}

func (c circuit) Contains(jbox P3) bool {
	for _, jj := range c.jboxes {
		if jj == jbox {
			return true
		}
	}
	return false
}

func (c circuit) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Circuit (len=%d)\n", len(c.jboxes))
	for _, jj := range c.jboxes {
		fmt.Fprintf(&sb, "\t%s\n", jj)
	}

	return sb.String()
}

func (c *circuit) Add(jbox ...P3) {
	for _, j := range c.jboxes {
		for _, jj := range jbox {
			if j == jj {
				panic("adding already existing")
			}
		}
	}

	c.jboxes = append(c.jboxes, jbox...)
}

func GetCircuit(circuits []*circuit, jbox P3) (*circuit, int, bool) {

	for i := 0; i < len(circuits); i++ {
		for _, jj := range circuits[i].jboxes {
			if jj == jbox {
				return circuits[i], i, true

			}
		}
	}
	return nil, 0, false
}

func InSameCircuit(circuits []*circuit, b1, b2 P3) (int, bool) {
	for i, c := range circuits {
		if c.Contains(b1) && c.Contains(b2) {
			return i, true
		}
	}
	return 0, false
}

func makePairs(jboxes []P3) []JboxPair {
	var pairs []JboxPair

	for i := 0; i < len(jboxes); i++ {
		for j := i + 1; j < len(jboxes); j++ {
			pairs = append(pairs, JboxPair{j1: jboxes[i], j2: jboxes[j]})
		}
	}
	return pairs
}

func main() {
	lines := util.GetFileStrings("2025/Day8/input")

	jboxes := []util.Point3D{}

	for _, l := range lines {
		jboxes = append(jboxes, util.Point3DFrom(l))
	}

	pairs := makePairs(jboxes)
	slices.SortFunc(pairs, func(lhs, rhs JboxPair) int {
		return int(lhs.j1.Euclidean(lhs.j2)) - int(rhs.j1.Euclidean(rhs.j2))
	})

	var circuits []*circuit
	for i := 0; i < 1000; i++ {
		p := pairs[i]
		if _, ok := InSameCircuit(circuits, p.j1, p.j2); ok {
			continue
		}
		c1, i1, ok1 := GetCircuit(circuits, p.j1)
		c2, _, ok2 := GetCircuit(circuits, p.j2)
		if ok1 && ok2 {
			circuits = util.RemoveByIndex(circuits, i1)
			c2.Add(c1.jboxes...)
		} else if ok1 {
			c1.Add(p.j2)
		} else if ok2 {
			c2.Add(p.j1)
		} else {
			circuits = append(circuits, &circuit{jboxes: []P3{p.j1, p.j2}})
		}
	}
	slices.SortFunc(circuits, func(lhs, rhs *circuit) int {
		return len(rhs.jboxes) - len(lhs.jboxes)
	})
	if len(circuits) >= 3 {
		fmt.Printf("Top 3 circuit sizes (part1): %d\n", len(circuits[0].jboxes)*len(circuits[1].jboxes)*len(circuits[2].jboxes))
	}

	var circuits2 []*circuit
	for _, p := range pairs {
		if _, ok := InSameCircuit(circuits2, p.j1, p.j2); ok {
			continue
		}
		c1, i1, ok1 := GetCircuit(circuits2, p.j1)
		c2, _, ok2 := GetCircuit(circuits2, p.j2)
		if ok1 && ok2 {
			circuits2 = util.RemoveByIndex(circuits2, i1)
			c2.Add(c1.jboxes...)
		} else if ok1 {
			c1.Add(p.j2)
		} else if ok2 {
			c2.Add(p.j1)
		} else {
			circuits2 = append(circuits2, &circuit{jboxes: []P3{p.j1, p.j2}})
		}

		if len(circuits2) == 1 && len(circuits2[0].jboxes) == len(jboxes) {
			fmt.Printf("Connecting pair x product (part2): %d\n", p.j1.X*p.j2.X)
			break
		}
	}
}
