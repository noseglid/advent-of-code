package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type point struct {
	x, y, z int
}

func NewRebootStep(def string) cube {
	s := cube{}
	if _, err := fmt.Sscanf(def,
		"%s x=%d..%d,y=%d..%d,z=%d..%d",
		&s.state, &s.x1, &s.x2, &s.y1, &s.y2, &s.z1, &s.z2,
	); err != nil {
		panic(err)
	}

	return s
}

func (r cube) set(cubes map[point]bool) {
	if r.x1 > 50 || r.y1 > 50 || r.z1 > 50 || r.x2 < -50 || r.y2 < -50 || r.z2 < -50 {
		return
	}
	for x := r.x1; x <= r.x2; x++ {
		for y := r.y1; y <= r.y2; y++ {
			for z := r.z1; z <= r.z2; z++ {
				switch r.state {
				case "on":
					cubes[point{x, y, z}] = true
				case "off":
					delete(cubes, point{x, y, z})
				}
			}
		}
	}
}

type cube struct {
	state  string
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

func (c cube) Volume() int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func (c cube) intersect(d cube) (cube, bool) {
	if c.x2 < d.x1 || c.y2 < d.y1 || c.z2 < d.z1 || c.x1 > d.x2 || c.y1 > d.y2 || c.z1 > d.z2 {
		return cube{}, false
	}

	ns := "on"
	if d.state == "on" {
		ns = "off"
	}

	return cube{
		ns,
		util.MaxInt(c.x1, d.x1), util.MinInt(c.x2, d.x2),
		util.MaxInt(c.y1, d.y1), util.MinInt(c.y2, d.y2),
		util.MaxInt(c.z1, d.z1), util.MinInt(c.z2, d.z2),
	}, true
}

func part2(steps []cube) {
	ac := []cube{}
	for _, s := range steps {
		l := []cube{}
		if s.state == "on" {
			l = append(l, s)
		}
		for _, c := range ac {
			if ic, ok := s.intersect(c); ok {
				l = append(l, ic)
			}
		}
		ac = append(ac, l...)
	}

	v := 0
	for _, c := range ac {
		m := 1
		if c.state == "off" {
			m = -1
		}
		v += c.Volume() * m
	}
	log.Printf("Part 2: Total cubes on: %d", v)
}

func main() {
	input := "2021/Day22/input"
	lines := util.GetFileStrings(input)

	var steps []cube
	for _, l := range lines {
		steps = append(steps, NewRebootStep(l))
	}

	cubes := map[point]bool{}
	for _, s := range steps {
		s.set(cubes)
	}
	log.Printf("Part 1: Enabled cubes: %d", len(cubes))

	part2(steps)

}
