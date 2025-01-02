package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

type P = util.Point

type Dir int

const (
	N Dir = 1
	S Dir = 2
	W Dir = 3
	E Dir = 4
)

func toUtilDir(d Dir) util.Dir {
	switch d {
	case N:
		return util.N
	case S:
		return util.S
	case W:
		return util.W
	case E:
		return util.E
	}
	panic("bad dir")
}

func coordsAt(p P, d Dir) P {
	switch d {
	case N:
		return P{p.X, p.Y - 1}
	case S:
		return P{p.X, p.Y + 1}
	case W:
		return P{p.X - 1, p.Y}
	case E:
		return P{p.X + 1, p.Y}
	}
	panic("unexpected dir")
}

func nextDir(d Dir) Dir {
	switch d {
	case N:
		return W
	case W:
		return S
	case S:
		return E
	case E:
		return N
	}
	panic("bad dir")
}

func walkMaze(p *intcode.Program, maze map[P]rune, pos P) {
	if _, ok := maze[pos]; ok {
		return
	}
	fmt.Printf("walking from %v\n", pos)

	for _, d := range []Dir{N, S, W, E} {
		p.PushState()
		p.Input <- int(d)
		res := <-p.Output

		pn := pos.Move(toUtilDir(d))
		switch res {
		case 0:
			maze[pn] = '#'
		case 1:
			maze[pos] = '.'
			walkMaze(p, maze, pn)
		case 2:
			maze[pn] = 'E'
			fmt.Printf("found end!\n")
		}
		fmt.Printf("restoring state\n")
		p.PopState()
	}

}

func main() {

	// program := intcode.NewProgram(util.GetCSVFileNumbers("2019/Day15/input"))
	// go func() {
	// 	program.Run()
	// 	fmt.Printf("program halted!\n")
	// }()
	// maze := map[P]rune{}

	// program.Input <- 1
	// fmt.Printf("north res:%d\n", <-program.Output)

	// walkMaze(program, maze, P{0, 0})
	// fmt.Printf("maze=%v\n", maze)

	// minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0
	// for p := range maze {
	// 	minx = util.Min(minx, p.X)
	// 	maxx = util.Max(maxx, p.X)
	// 	miny = util.Min(miny, p.Y)
	// 	maxy = util.Max(maxy, p.Y)
	// }
	// xoff, yoff := 0-minx, 0-miny

	// fmt.Printf("dimensions x=%d->%d, y=%d->%d, xoff,yoff=%d,%d\n", minx, maxx, miny, maxy, xoff, yoff)
	// grid := make([][]rune, maxy+yoff+1)
	// for y := range grid {
	// 	grid[y] = make([]rune, maxx+xoff+1)
	// 	for x := range grid[y] {
	// 		grid[y][x] = '#'
	// 	}
	// }

	// for p, r := range maze {
	// 	grid[p.Y+yoff][p.X+xoff] = r
	// }
	// grid[yoff][xoff] = 'S'

	// util.PrintRuneGrid(grid)

	grid := util.ParseRuneGrid(util.GetFileStrings("2019/Day15/grid"))
	startx, starty := grid.Find('S')
	endx, endy := grid.Find('E')

	canMove := func(x, y int, r rune) bool {
		return r == '.' || r == 'E'
	}

	steps, _, ok := grid.ShortestPath(P{startx, starty}, P{endx, endy}, canMove)
	if !ok {
		panic("no shortest path")
	}
	fmt.Printf("Fewest movements (part1): %d\n", steps)

	longest := 0
	grid.Each(func(x, y int, r rune) {
		if !canMove(x, y, r) {
			return
		}
		steps, _, ok := grid.ShortestPath(P{endx, endy}, P{x, y}, canMove)
		if !ok {
			return
		}
		if steps > longest {
			longest = steps
		}
	})

	fmt.Printf("Time to fill with oxygen (part2): %d\n", longest)

}
