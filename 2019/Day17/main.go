package main

import (
	"fmt"
	"sync"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

func alignmentParameters(grid [][]rune) int {
	s := 0
	for y := range grid {
		if y == 0 || y == len(grid)-1 {
			continue
		}
		for x := range grid[y] {
			if x == 0 || x == len(grid[y])-1 {
				continue
			}

			if grid[y][x] == '#' && grid[y-1][x] == '#' && grid[y][x+1] == '#' && grid[y+1][x] == '#' && grid[y][x-1] == '#' {
				s += y * x
			}
		}
	}
	return s
}

func p1(src []int) int {
	program := intcode.NewProgram(src)
	gsrc := map[util.Point]rune{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	x, maxx, y := 0, 0, 0
	go func() {
		for {
			r, ok := <-program.Output
			if !ok {
				wg.Done()
				break
			}
			switch rune(r) {
			default:
				gsrc[util.Point{x, y}] = rune(r)
				x++
				if x > maxx {
					maxx = x
				}
			case '\n':
				x = 0
				y++
			}
		}
	}()

	program.Run()
	wg.Wait()
	grid := make([][]rune, y+1)
	for y := range grid {
		grid[y] = make([]rune, maxx+1)
	}

	for p, r := range gsrc {
		grid[p.Y][p.X] = r
	}
	return alignmentParameters(grid)
}

func write(p *intcode.Program, s string) {
	for _, r := range s {
		p.Input <- int(r)
	}
	p.Input <- '\n'
}

func p2(src []int) int {
	src[0] = 2
	p := intcode.NewProgram(src)

	wg := sync.WaitGroup{}
	go p.Run()

	write(p, "A,B,A,A,B,C,B,C,C,B")
	/*A=*/ write(p, "L,12,R,8,L,6,R,8,L,6")
	/*B=*/ write(p, "R,8,L,12,L,12,R,8")
	/*C=*/ write(p, "L,6,R,6,L,12")
	/*Video=*/ write(p, "n")

	var dustCollected int
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if r := <-p.Output; r < 255 {
				fmt.Printf("%c", r)
			} else {
				dustCollected = r
				return
			}
		}
	}()

	wg.Wait()
	return dustCollected
}

func main() {
	src := util.GetCSVFileNumbers("2019/Day17/input")
	p1, p2 := p1(src), p2(src)
	fmt.Printf("Alignment parameters of intersections (part1): %d\n", p1)
	fmt.Printf("Dust collected (part2): %d\n", p2)
}
