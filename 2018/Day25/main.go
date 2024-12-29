package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type P struct {
	x, y, z, w int
}

func (p P) Manhattan(o P) int {
	return util.Absolute(p.x-o.x) + util.Absolute(p.y-o.y) + util.Absolute(p.z-o.z) + util.Absolute(p.w-o.w)
}

func inSameConstellation(p1 P, constellation []P) bool {
	for _, p := range constellation {
		if p1.Manhattan(p) <= 3 {
			return true
		}
	}
	return false
}

func merge(c [][]P, a, b int) [][]P {
	c[a] = append(c[a], c[b]...)
	return util.RemoveByIndex(c, b)
}

func main() {

	var ps []P
	for _, l := range util.GetFileStrings("2018/Day25/input") {
		var p P
		fmt.Sscanf(l, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.w)
		ps = append(ps, p)
	}

	var constellations [][]P
Outer:
	for _, p := range ps {
		for i, c := range constellations {
			if inSameConstellation(p, c) {
				constellations[i] = append(constellations[i], p)
				continue Outer
			}
		}

		constellations = append(constellations, []P{p})
	}

Outer2:
	for {
		for a, c1 := range constellations {
			for _, p := range c1 {
				for b, c2 := range constellations {
					if a != b && inSameConstellation(p, c2) {
						constellations = merge(constellations, a, b)
						continue Outer2
					}
				}
			}
		}
		break
	}

	fmt.Printf("number of constellations (part1): %d\n", len(constellations))

}
