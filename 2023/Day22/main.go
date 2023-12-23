package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type Brick struct {
	id         int
	start, end util.Point3D
}

func (b Brick) minz() int {
	return min(b.start.Z, b.end.Z)
}
func (b Brick) maxz() int {
	return max(b.start.Z, b.end.Z)
}

func (b Brick) Label() string {
	if b.id < 26 {
		return string([]rune{rune('A' + b.id)})
	}
	return "*"
}

func (b Brick) String() string {
	return b.Label()
}

func (b Brick) cubes() []util.Point3D {
	r := []util.Point3D{}
	for x := b.start.X; x <= b.end.X; x++ {
		for y := b.start.Y; y <= b.end.Y; y++ {
			for z := b.start.Z; z <= b.end.Z; z++ {
				r = append(r, util.Point3D{X: x, Y: y, Z: z})
			}
		}
	}
	return r
}

func (b Brick) canFall(bricks []Brick) bool {
	blockers, onGround := b.blockedFromFall(bricks)
	return !onGround && len(blockers) == 0
}

func (b Brick) blockedFromFall(bricks []Brick) ([]Brick, bool) {
	var r []Brick
	if b.minz() <= 1 {
		// On ground
		return nil, true
	}

Outer:
	for _, cb := range bricks {
		if b == cb {
			continue
		}
		if cb.minz() > b.minz() {
			// check brick is above b
			continue
		}
		if cb.maxz() != b.minz()-1 {
			// check brick is not immediately below b
			continue
		}

		for _, c1 := range cb.cubes() {
			for _, c2 := range b.cubes() {
				if c2.Z-1 == c1.Z && (c1.X == c2.X && c1.Y == c2.Y) {
					r = append(r, cb)
					continue Outer
				}
			}
		}
	}
	return r, false
}

func dimXZ(bricks []Brick) (int, int, int, int) {
	minx, maxx, minz, maxz := math.MaxInt, 0, math.MaxInt, 0
	for _, b := range bricks {
		if b.start.X < minx {
			minx = b.start.X
		}
		if b.end.X > maxx {
			maxx = b.start.X
		}
		if b.start.Z < minz {
			minz = b.start.Z
		}
		if b.end.Z > maxz {
			maxz = b.end.Z
		}
	}
	return minx, maxx, minz, maxz
}

func gravity(bricks []Brick) int {
	didFall := map[int]bool{}
	for i := 0; i < len(bricks); i++ {
		for bricks[i].canFall(bricks) {
			bricks[i].start.Z--
			bricks[i].end.Z--
			didFall[i] = true
		}
	}
	slices.SortFunc(bricks, sortfn)
	return len(didFall)
}

func sortfn(lhs, rhs Brick) int {
	return lhs.minz() - rhs.minz()
}

func supportMap(bricks []Brick) (map[Brick][]Brick, map[Brick][]Brick) {
	isSupportedBy := map[Brick][]Brick{}
	supports := map[Brick][]Brick{}

	for _, b := range bricks {
		blockers, _ := b.blockedFromFall(bricks)
		for _, bb := range blockers {
			isSupportedBy[b] = append(isSupportedBy[b], bb)
			supports[bb] = append(supports[bb], b)
		}
	}
	return isSupportedBy, supports
}

func cloneBricks(bricks []Brick, skip Brick) []Brick {
	var r []Brick
	for _, b := range bricks {
		if b == skip {
			continue
		}
		r = append(r, b)
	}
	return r
}

func main() {
	lines := util.GetFileStrings("2023/Day22/input")
	parseRegexp := regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)
	var bricks []Brick
	for i, l := range lines {
		m := parseRegexp.FindStringSubmatch(l)
		bricks = append(bricks, Brick{
			id:    i,
			start: util.Point3D{X: util.MustAtoi(m[1]), Y: util.MustAtoi(m[2]), Z: util.MustAtoi(m[3])},
			end:   util.Point3D{X: util.MustAtoi(m[4]), Y: util.MustAtoi(m[5]), Z: util.MustAtoi(m[6])},
		})
	}

	slices.SortFunc(bricks, sortfn)
	gravity(bricks)
	isSupportedBy, supports := supportMap(bricks)

	disintegrations := []Brick{}

	n := 0
	for _, b := range bricks {
		if len(supports[b]) == 0 {
			// Does not support any
			n++
			disintegrations = append(disintegrations, b)
			continue
		}

		allHaveMultipleSupports := true
		for _, s := range supports[b] {
			if len(isSupportedBy[s]) <= 1 {
				allHaveMultipleSupports = false
			}
		}
		if allHaveMultipleSupports {
			disintegrations = append(disintegrations, b)
			n++
		}
	}

	fmt.Printf("Blocks which can be disintegrated (part1): %d\n", n)

	n2 := 0
	for _, b := range bricks {
		if slices.Contains(disintegrations, b) {
			continue
		}
		bb := cloneBricks(bricks, b)
		n2 += gravity(bb)
	}

	fmt.Printf("Sum of dropped for all disintegrations (part2): %d\n", n2)

}
