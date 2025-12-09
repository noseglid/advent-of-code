package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point

func redTiles(lines []string) []P {
	var ps []P
	for _, l := range lines {
		if l == "" {
			continue
		}
		ps = append(ps, util.PointFrom(l))
	}
	return ps
}

func area(a, b P) int {
	return (util.Absolute(a.X-b.X) + 1) * (util.Absolute(a.Y-b.Y) + 1)
}

type PairArea struct {
	p1, p2 P
	area   int
}

func onPolygonBoundary(polygon []P, p P) bool {
	for i := 0; i < len(polygon); i++ {
		c1, c2 := polygon[i], polygon[(i+1)%len(polygon)]
		if c1.X == c2.X && p.X == c1.X && p.Y >= min(c1.Y, c2.Y) && p.Y <= max(c1.Y, c2.Y) {
			return true
		}
		if c1.Y == c2.Y && p.Y == c1.Y && p.X >= min(c1.X, c2.X) && p.X <= max(c1.X, c2.X) {
			return true
		}
	}
	return false
}

func isPointInPolygon(polygon []P, p P) bool {
	if onPolygonBoundary(polygon, p) {
		return true
	}

	crosses := 0
	for i := 0; i < len(polygon); i++ {
		v1, v2 := polygon[i], polygon[(i+1)%len(polygon)]

		if (v1.Y > p.Y) == (v2.Y > p.Y) {
			continue
		}

		if v1.X+(p.Y-v1.Y)*(v2.X-v1.X)/(v2.Y-v1.Y) > p.X {
			crosses++
		}
	}

	return crosses%2 == 1
}

func segmentIntersects(a1, a2, b1, b2 P) bool {
	d1 := ccw(b1, b2, a1)
	d2 := ccw(b1, b2, a2)
	d3 := ccw(a1, a2, b1)
	d4 := ccw(a1, a2, b2)

	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}

	return false
}

func ccw(a, b, c P) int {
	return (c.X-a.X)*(b.Y-a.Y) - (c.Y-a.Y)*(b.X-a.X)
}

func areaIsEncapsulated(rtiles []P, pa PairArea) bool {
	minX := min(pa.p1.X, pa.p2.X)
	maxX := max(pa.p1.X, pa.p2.X)
	minY := min(pa.p1.Y, pa.p2.Y)
	maxY := max(pa.p1.Y, pa.p2.Y)

	corners := []P{
		{X: minX, Y: minY},
		{X: maxX, Y: minY},
		{X: minX, Y: maxY},
		{X: maxX, Y: maxY},
	}

	for _, corner := range corners {
		if !isPointInPolygon(rtiles, corner) {
			return false
		}
	}

	edges := [][]P{
		{{X: minX, Y: minY}, {X: maxX, Y: minY}},
		{{X: maxX, Y: minY}, {X: maxX, Y: maxY}},
		{{X: maxX, Y: maxY}, {X: minX, Y: maxY}},
		{{X: minX, Y: maxY}, {X: minX, Y: minY}},
	}
	for i := 0; i < len(rtiles); i++ {
		outerEdge := []P{rtiles[i], rtiles[(i+1)%len(rtiles)]}

		for _, e := range edges {
			if segmentIntersects(outerEdge[0], outerEdge[1], e[0], e[1]) {
				return false
			}
		}
	}

	return true
}

func main() {
	lines := util.GetFileStrings("2025/Day9/input")

	rtiles := redTiles(lines)
	pairs := util.MakePairs(rtiles)

	maxArea := 0
	for _, p := range pairs {
		maxArea = max(maxArea, area(p.X, p.Y))
	}
	fmt.Printf("biggest area (part1): %d\n", maxArea)

	var pairAreas []PairArea
	for _, p := range pairs {
		pairAreas = append(pairAreas, PairArea{p1: p.X, p2: p.Y, area: area(p.X, p.Y)})
	}

	slices.SortFunc(pairAreas, func(lhs, rhs PairArea) int {
		return rhs.area - lhs.area
	})

	maxAreaP2 := 0
	for _, pa := range pairAreas {
		if areaIsEncapsulated(rtiles, pa) {
			maxAreaP2 = pa.area
			break
		}
	}

	fmt.Printf("biggest encapsulated area (part2)=%d\n", maxAreaP2)

}
