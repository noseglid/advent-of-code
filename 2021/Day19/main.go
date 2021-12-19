package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type matrix struct {
	a, b, c int
	d, e, f int
	g, h, i int
}

func (m matrix) String() string {
	return fmt.Sprintf("%d, %d, %d\n%d, %d, %d\n%d, %d, %d\n", m.a, m.b, m.c, m.d, m.e, m.f, m.g, m.h, m.i)
}

func (m matrix) mul(p point) point {
	return point{
		m.a*p.x + m.b*p.y + m.c*p.z,
		m.d*p.x + m.e*p.y + m.f*p.z,
		m.g*p.x + m.h*p.y + m.i*p.z,
	}
}

func (m matrix) mmul(r matrix) matrix {
	return matrix{
		(m.a*r.a + m.b*r.d + m.c*r.g), (m.a*r.b + m.b*r.e + m.c*r.h), (m.a*r.c + m.b*r.f + m.c*r.i),
		(m.d*r.a + m.e*r.d + m.f*r.g), (m.d*r.b + m.e*r.e + m.f*r.h), (m.d*r.c + m.e*r.f + m.f*r.i),
		(m.g*r.a + m.h*r.d + m.i*r.g), (m.g*r.b + m.h*r.e + m.i*r.h), (m.g*r.c + m.h*r.f + m.i*r.i),
	}
}

type point struct {
	x, y, z int
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

var identity = matrix{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

var rotx90 = matrix{
	1, 0, 0,
	0, 0, -1,
	0, 1, 0,
}
var rotx180 = matrix{
	1, 0, 0,
	0, -1, 0,
	0, 0, -1,
}
var rotx270 = matrix{
	1, 0, 0,
	0, 0, 1,
	0, -1, 0,
}

var roty90 = matrix{
	0, 0, 1,
	0, 1, 0,
	-1, 0, 0,
}
var roty180 = matrix{
	-1, 0, 0,
	0, 1, 0,
	0, 0, -1,
}

var roty270 = matrix{
	0, 0, -1,
	0, 1, 0,
	1, 0, 0,
}

var rotz90 = matrix{
	0, -1, 0,
	1, 0, 0,
	0, 0, 1,
}
var rotz180 = matrix{
	-1, 0, 0,
	0, -1, 0,
	0, 0, 1,
}
var rotz270 = matrix{
	0, 1, 0,
	-1, 0, 0,
	0, 0, 1,
}

var allRots map[matrix]point

func init() {
	allRots = make(map[matrix]point)
	for i, xrot := range []matrix{identity, rotx90, rotx180, rotx270} {
		for j, yrot := range []matrix{identity, roty90, roty180, roty270} {
			for k, zrot := range []matrix{identity, rotz90, rotz180, rotz270} {
				m := xrot.mmul(yrot).mmul(zrot)
				allRots[m] = point{i, j, k}
			}
		}
	}
}

type scanner struct {
	id      int
	posRel0 point
	probes  []point
}

func (s scanner) String() string {
	return fmt.Sprintf("Scanner %d", s.id)
}

func NewScanner(lines []string) *scanner {
	s := scanner{}

	if _, err := fmt.Sscanf(lines[0], "--- scanner %d ---", &s.id); err != nil {
		panic(err)
	}

	for _, l := range lines[1:] {
		p := point{}
		if _, err := fmt.Sscanf(l, "%d,%d,%d", &p.x, &p.y, &p.z); err != nil {
			panic(err)
		}
		s.probes = append(s.probes, p)
	}

	return &s
}

func (s *scanner) distance(rhs *scanner) int {
	dx, dy, dz := s.posRel0.x-rhs.posRel0.x, s.posRel0.y-rhs.posRel0.y, s.posRel0.z-rhs.posRel0.z
	return util.Absolute(dx) + util.Absolute(dy) + util.Absolute(dz)
}

func (s *scanner) realign(rot matrix, pos point) {
	for i := 0; i < len(s.probes); i++ {
		s.probes[i] = rot.mul(s.probes[i]).add(pos)
	}
}

func (p point) add(pp point) point {
	return point{p.x + pp.x, p.y + pp.y, p.z + pp.z}
}

func pointAlign(lhs, rhs point) point {
	return point{lhs.x - rhs.x, lhs.y - rhs.y, lhs.z - rhs.z}
}

func unrotatedOverlap(l1, l2 []point) (point, []point, bool) {
	candidates := map[point]struct{}{}

	for _, ll1 := range l1 {
		for _, ll2 := range l2 {
			candidates[pointAlign(ll1, ll2)] = struct{}{}
		}
	}

	for c := range candidates {
		n := 0
		resPoints := map[point]struct{}{}
		for _, ll1 := range l1 {
			for _, ll2 := range l2 {
				if ll2.add(c) == ll1 {
					resPoints[ll1] = struct{}{}
					n++
				}
			}
		}
		if n >= 12 {
			var ret []point
			for p := range resPoints {
				ret = append(ret, p)
			}
			return c, ret, true
		}
	}
	return point{}, nil, false
}

func (s scanner) overlapping(rhs *scanner) (matrix, point, bool) {
	for rot := range allRots {
		var aligned []point
		for _, p := range rhs.probes {
			aligned = append(aligned, rot.mul(p))
		}

		if c, _, ok := unrotatedOverlap(s.probes, aligned); ok {
			return rot, c, true
		}

	}

	return matrix{}, point{}, false
}

func main() {
	input := "2021/Day19/input"
	lines := util.GetFileStrings(input)

	var scanners []*scanner

	var def []string
	for _, l := range lines {
		if l == "" {
			scanners = append(scanners, NewScanner(def))
			def = def[:0]
		} else {
			def = append(def, l)
		}
	}
	scanners = append(scanners, NewScanner(def))

	s0normalized := map[point]struct{}{}
	normalizedScanners := []*scanner{scanners[0]}
	originalScanners := scanners[1:]
	scanners[0].posRel0 = point{0, 0, 0}

	// Add all probes for scanner 0, this is the reference
	for _, p := range scanners[0].probes {
		s0normalized[p] = struct{}{}
	}

Outer:
	for len(originalScanners) > 0 {
		log.Printf("looping outer, originalScanners=%d, normalizedScanners=%d", len(originalScanners), len(normalizedScanners))
		for _, s := range normalizedScanners {
			for j, o := range originalScanners {
				if rot, align, ok := s.overlapping(o); ok {
					o.realign(rot, align)
					o.posRel0 = align
					normalizedScanners = append(normalizedScanners, o)
					originalScanners = append(originalScanners[:j], originalScanners[j+1:]...)
					for _, p := range o.probes {
						s0normalized[p] = struct{}{}
					}
					continue Outer
				}
			}
		}
	}
	log.Printf("Part 1: n probes: %d", len(s0normalized))

	md := 0
	for _, s1 := range normalizedScanners {
		for _, s2 := range normalizedScanners {
			if s1 == s2 {
				continue
			}
			if d := s1.distance(s2); d > md {
				md = d
			}
		}
	}

	log.Printf("Part 2: max manhattan distance: %d", md)
}
