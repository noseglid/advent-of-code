package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type disc struct {
	id        int
	positions int
	t0        int
}

func (d disc) isAtPosition0(dropTime int) bool {
	return (d.t0+d.id+dropTime)%d.positions == 0
}

func parseDisc(s string) disc {
	d := disc{}
	if _, err := fmt.Sscanf(s, "Disc #%d has %d positions; at time=0, it is at position %d.", &d.id, &d.positions, &d.t0); err != nil {
		panic(err)
	}

	return d
}

func findCapsule(discs []disc) int {
Outer:
	for t := 0; ; t++ {
		for _, d := range discs {
			if !d.isAtPosition0(t) {
				continue Outer
			}
		}
		return t
	}
}

func main() {
	input := "2016/Day15/input"
	lines := util.GetFileStrings(input)

	discs := []disc{}
	for _, l := range lines {
		discs = append(discs, parseDisc(l))
	}

	log.Printf("Part 1: Capsule received by pressing at time: %d", findCapsule(discs))
	log.Printf("Part 2: Capsule received by pressing at time: %d", findCapsule(append(discs, disc{len(discs) + 1, 11, 0})))

}
