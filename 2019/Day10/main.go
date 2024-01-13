package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type asteroid struct {
	row, col int
	dists    []dist
}

func (a asteroid) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Asteroid(%d,%d):\n", a.row, a.col))

	for _, d := range a.dists {
		sb.WriteRune('\t')
		sb.WriteString(d.String())
		sb.WriteRune('\n')
	}

	return sb.String()
}

type dist struct {
	arow, acol int
	dist       float64
	angle      float64
}

func (d dist) String() string {
	return fmt.Sprintf("(%d,%d):%.2f:%.2f", d.arow, d.acol, d.dist, d.angle)
}

func n_visible(a asteroid) int {
	used_angles := []float64{}
	for _, d := range a.dists {
		if !util.Contains(used_angles, d.angle) {
			used_angles = append(used_angles, d.angle)
		}
	}
	return len(used_angles)
}

func main() {
	grid := util.GetFileRuneGrid("2019/Day10/input")

	var asteroids []asteroid
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '#' {
				asteroids = append(asteroids, asteroid{row: row, col: col})
			}
		}
	}
	for si, source := range asteroids {
		for ti, test := range asteroids {
			if si == ti {
				continue
			}
			rdiff, cdiff := test.row-source.row, test.col-source.col
			d := dist{
				arow:  test.row,
				acol:  test.col,
				angle: math.Atan2(float64(cdiff), -float64(rdiff)),
				dist:  math.Sqrt(float64(rdiff*rdiff + cdiff*cdiff)),
			}
			if d.angle < 0 {
				d.angle += 2 * math.Pi
			}
			asteroids[si].dists = append(asteroids[si].dists, d)
		}
		sort.Slice(asteroids[si].dists, func(i, j int) bool {
			dd := asteroids[si].dists
			if dd[i].angle == dd[j].angle {
				return dd[i].dist < dd[j].dist
			}
			return dd[i].angle < dd[j].angle
		})
	}

	max, station := 0, asteroid{}
	for _, a := range asteroids {
		if n := n_visible(a); n > max {
			max = n
			station = a
		}
	}
	fmt.Printf("Max visible (part1): %d\n", max)

	n := 0
	index := 0
	for {
		angle := station.dists[index].angle
		r, c := station.dists[index].arow, station.dists[index].acol
		station.dists = util.RemoveByIndex(station.dists, index)
		if len(station.dists) == 0 {
			break
		}
		index = index % len(station.dists)
		n++
		if n == 200 {
			fmt.Printf("Asteroid 200 removed at row=%d,col=%d for (part2): %d\n", r, c, c*100+r)
			break
		}
		for station.dists[index].angle == angle {
			index = (index + 1) % len(station.dists)
			if index == 0 {
				break
			}
		}
	}
}
