package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

type Range struct {
	s, e int
}

func fitsShip(xranges []Range, boxSide int) bool {
	ll := len(xranges)
	if ll < boxSide {
		return false
	}
	first, last := xranges[ll-boxSide], xranges[ll-1]
	if last.s >= first.s && last.s <= first.e && first.e-last.s+1 >= boxSide {
		return true
	}

	return false
}

func main() {

	src := util.GetCSVFileNumbers("2019/Day19/input")
	program := intcode.NewProgram(src)
	size := 50

	grid := make([][]rune, size)
	for y := range grid {
		grid[y] = make([]rune, size)
	}

	n := 0
	for x := range util.RangeInt(size) {
		for y := range util.RangeInt(size) {
			program.Input <- x
			program.Input <- y
			program.Run()

			r := '.'
			if o := <-program.Output; o == 1 {
				n++
				r = '#'
			}
			grid[y][x] = r
			program.Reset()
		}
	}

	fmt.Printf("Affected points by tractor beam (part1): %d\n", n)

	program.Reset()
	x, y := 0, 1400
	boxSide := 100

	xrange := []Range{}
	for {
		w, sx, ex := 0, 0, 0
		foundBeam := false
		for {
			program.Input <- x
			program.Input <- y
			program.Run()

			o := <-program.Output
			program.Reset()
			if o == 1 {
				if !foundBeam {
					sx = x
				}
				w++
				foundBeam = true
			} else if foundBeam {
				ex = x - 1
				break
			}

			x++
		}

		xrange = append(xrange, Range{sx, ex})
		if fitsShip(xrange, boxSide) {
			fmt.Printf("Ship fits at y=%d with value (part2): %d\n", y, (xrange[len(xrange)-1].s*10000 + (y - boxSide + 1)))
			break
		}
		y++
		x = 0
	}
}
