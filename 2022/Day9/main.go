package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type knot struct {
	p      point
	follow *knot
}

type point struct {
	x, y int
}

func distance(p1, p2 point) point {
	return point{p1.x - p2.x, p1.y - p2.y}
}

func (p *point) move(dir rune) {
	switch dir {
	case 'U':
		p.y--
	case 'R':
		p.x++
	case 'D':
		p.y++
	case 'L':
		p.x--

	case 'G':
		p.x--
		p.y--
	case 'B':
		p.x--
		p.y++
	case 'H':
		p.x++
		p.y--
	case 'N':
		p.x++
		p.y++
	default:
		panic("point: bad move dir")
	}
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

var gridn int

func printGrid(knots []*knot) {
	gridn++
	fmt.Printf("Step %d\n", gridn)
	for y := -5; y < 5; y++ {
		for x := -5; x < 5; x++ {
			r := '.'
			if x == 0 && y == 0 {
				r = 's'
			}
			for i, k := range knots {
				if k.p.x == x && k.p.y == y {
					r = rune(i + '0')
					if i == 0 {
						r = 'H'
					}
					break
				}
			}
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
	fmt.Println()
}

func shouldMove(h, t point) (bool, rune) {
	d := distance(h, t)
	if d.x == 0 && d.y == 0 {
		return false, 'x'
	}

	if d.x == 0 && util.Absolute(d.y) > 1 {
		dir := 'U'
		if d.y > 0 {
			dir = 'D'
		}
		return true, dir
	}

	if d.y == 0 && util.Absolute(d.x) > 1 {
		dir := 'L'
		if d.x > 0 {
			dir = 'R'
		}
		return true, dir
	}

	if util.Absolute(d.x) == 2 || util.Absolute(d.y) == 2 {
		if d.x == -2 && d.y == -1 || d.x == -1 && d.y == -2 || d.x == -2 && d.y == -2 {
			return true, 'G'
		}
		if d.x == -2 && d.y == 1 || d.x == -1 && d.y == 2 || d.x == -2 && d.y == 2 {
			return true, 'B'
		}
		if d.x == 2 && d.y == -1 || d.x == 1 && d.y == -2 || d.x == 2 && d.y == -2 {
			return true, 'H'
		}
		if d.x == 2 && d.y == 1 || d.x == 1 && d.y == 2 || d.x == 2 && d.y == 2 {
			return true, 'N'
		}
	}

	return false, 'x'
}

func iterate(knots []*knot, dir rune, visited map[point]bool) {
	knots[0].p.move(dir)

	for _, k := range knots[1:] {
		if move, mdir := shouldMove(k.follow.p, k.p); move {
			k.p.move(mdir)
		}
	}

	visited[knots[len(knots)-1].p] = true

}

func main() {
	commands := util.GetFileStrings("2022/Day9/input")

	knotsp1 := make([]*knot, 2)
	knotsp2 := make([]*knot, 10)

	for i := 0; i < len(knotsp1); i++ {
		knotsp1[i] = new(knot)
		if i > 0 {
			knotsp1[i].follow = knotsp1[i-1]
		}
	}

	for i := 0; i < len(knotsp2); i++ {
		knotsp2[i] = new(knot)
		if i > 0 {
			knotsp2[i].follow = knotsp2[i-1]
		}
	}

	var (
		visitedp1 = make(map[point]bool)
		visitedp2 = make(map[point]bool)
	)

	for _, c := range commands {
		var (
			dir   rune
			steps int
		)
		if _, err := fmt.Sscanf(c, "%c %d", &dir, &steps); err != nil {
			panic("bad command")
		}

		for i := 0; i < steps; i++ {
			iterate(knotsp1, dir, visitedp1)
			iterate(knotsp2, dir, visitedp2)
		}
	}

	fmt.Printf("Visited nodes (part1): %d\n", len(visitedp1))
	fmt.Printf("Visited nodes (part2): %d\n", len(visitedp2))
}
