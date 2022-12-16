package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type shape struct {
	top, right, bottom, left util.Point
}

func (s shape) Edge() []util.Point {
	var points []util.Point
	curr := s.top
	for {
		points = append(points, util.Point{X: curr.X, Y: curr.Y - 1})
		if curr == s.right {
			break
		}
		curr.X++
		curr.Y++
	}
	for {
		points = append(points, util.Point{X: curr.X + 1, Y: curr.Y})
		if curr == s.bottom {
			break
		}
		curr.X--
		curr.Y++
	}
	for {
		points = append(points, util.Point{X: curr.X, Y: curr.Y + 1})
		if curr == s.left {
			break
		}
		curr.X--
		curr.Y--
	}
	for {
		points = append(points, util.Point{X: curr.X - 1, Y: curr.Y})
		if curr == s.top {
			break
		}
		curr.X++
		curr.Y--
	}

	return points
}

func manhattanDistance(p1, p2 util.Point) int {
	return util.Absolute(p1.X-p2.X) + util.Absolute(p1.Y-p2.Y)
}

type Beacon struct {
	position util.Point
}

type Sensor struct {
	position      util.Point
	closestBeacon *Beacon
}

func (s Sensor) BeaconDistance() int {
	return manhattanDistance(s.position, s.closestBeacon.position)
}

func (s Sensor) CoversPoint(p util.Point) bool {
	return manhattanDistance(s.position, p) <= s.BeaconDistance()
}

func (s Sensor) Shape() shape {
	d := s.BeaconDistance()
	return shape{
		top:    util.Point{X: s.position.X, Y: s.position.Y - d},
		right:  util.Point{X: s.position.X + d, Y: s.position.Y},
		bottom: util.Point{X: s.position.X, Y: s.position.Y + d},
		left:   util.Point{X: s.position.X - d, Y: s.position.Y},
	}
}

func (s Sensor) CoversOnY(y int) []util.Point {
	var points []util.Point
	for x := -s.BeaconDistance(); x <= s.BeaconDistance(); x++ {
		test := util.Point{X: s.position.X + x, Y: y}
		if s.CoversPoint(test) {
			// fmt.Printf("%v covered\n", test)
			points = append(points, test)
		}
	}

	return points
}

func main() {
	lines := util.GetFileStrings("2022/Day15/input")
	// lostBeaconMax := util.Point{X: 20, Y: 20}
	lostBeaconMax := util.Point{X: 4000000, Y: 4000000}

	// line := 10
	line := 2000000

	var sensors []Sensor
	coveredPoints := map[util.Point]rune{}
	for _, l := range lines {
		var ps, pb util.Point
		if _, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &ps.X, &ps.Y, &pb.X, &pb.Y); err != nil {
			panic(err)
		}
		coveredPoints[ps] = 's'
		coveredPoints[pb] = 'b'
		sensors = append(sensors, Sensor{
			position: ps,
			closestBeacon: &Beacon{
				position: pb,
			}})
	}

	for _, s := range sensors {
		for _, p := range s.CoversOnY(line) {
			if _, ok := coveredPoints[p]; !ok {
				coveredPoints[p] = 'e'
			}
		}
	}

	n := 0
	for p, v := range coveredPoints {
		if p.Y == line && v == 'e' {
			n++
		}
	}

	fmt.Printf("covered points on line %d (part1): %d\n", line, n)

	// grid := make([][]int, 25)
	// for i := range grid {
	// 	grid[i] = make([]int, 25)
	// }
	// for _, p := range sensors[6].Shape().Edge() {
	// 	grid[p.Y+3][p.X+3] = 1
	// }

	// for _, r := range grid {
	// 	for _, i := range r {
	// 		if i == 0 {
	// 			fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf("#")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	var missingBeacon util.Point
	cp, cc := 0, 0
Loop:
	for _, s := range sensors {
	PointLoop:
		for _, p := range s.Shape().Edge() {
			cp++
			for _, s2 := range sensors {
				cc++
				if s2.CoversPoint(p) {
					continue PointLoop
				}
			}
			if p.X > lostBeaconMax.X || p.X < 0 || p.Y > lostBeaconMax.Y || p.Y < 0 {
				continue PointLoop
			}
			// No sensor covers the point!
			missingBeacon = p
			break Loop
		}
	}

	fmt.Printf("Missing beacon at %s, with frequency (part2): %d", missingBeacon, missingBeacon.X*4000000+missingBeacon.Y)
	fmt.Printf("checked points=%d, total checks=%d\n", cp, cc)
}
