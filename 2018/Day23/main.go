package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point3D

type bot struct {
	p P
	s int
}

type dist struct {
	d, c int
}

func main() {
	lines := util.GetFileStrings("2018/Day23/input")

	var bots []bot
	var strongest bot
	for _, l := range lines {
		var b bot
		fmt.Sscanf(l, "pos=<%d,%d,%d>, r=%d", &b.p.X, &b.p.Y, &b.p.Z, &b.s)
		bots = append(bots, b)
		if b.s > strongest.s {
			strongest = b
		}
	}

	n := 0
	for _, b := range bots {
		if strongest.p.Manhattan(b.p) <= strongest.s {
			n++
		}
	}

	fmt.Printf("bots in range of strongest signal bot (part1): %d\n", n)

	dists := []dist{}
	for _, b := range bots {
		d := b.p.Manhattan(P{0, 0, 0})
		dists = append(dists, dist{util.Max(0, d-b.s), 1})
		dists = append(dists, dist{d + b.s + 1, -1})
	}

	c, m, r := 0, 0, 0
	for len(dists) > 0 {
		slices.SortFunc(dists, func(a, b dist) int { return a.d - b.d })
		d := dists[0]
		c += d.c
		if c > m {
			r = d.d
			m = c
		}
		dists = dists[1:]
	}

	fmt.Printf("distance from origin (part2): %d\n", r)
}
