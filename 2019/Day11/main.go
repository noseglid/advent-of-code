package main

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/noseglid/advent-of-code/2019/Day11/program"
	"github.com/noseglid/advent-of-code/util"
)

type dir int

const (
	N dir = 0
	E dir = 1
	S dir = 2
	W dir = 3
)

func nextDir(curDir dir, output int) dir {
	switch curDir {
	case N:
		if output == 0 {
			return W
		} else {
			return E
		}
	case E:
		if output == 0 {
			return N
		} else {
			return S
		}
	case S:
		if output == 0 {
			return E
		} else {
			return W
		}
	case W:
		if output == 0 {
			return S
		} else {
			return N
		}
	}
	panic("no real dir")
}

func doMove(x, y int, d dir) (int, int) {
	switch d {
	case N:
		return x, y - 1
	case E:
		return x + 1, y
	case S:
		return x, y + 1
	case W:
		return x - 1, y
	}
	panic("bad dir")

}

type tile struct {
	x, y int
}

func print(mtiles map[tile]int) {
	var tiles []tile
	minx, maxx, miny, maxy := math.MaxInt, -math.MaxInt, math.MaxInt, -math.MaxInt
	for t, v := range mtiles {
		if v == 0 {
			continue
		}
		tiles = append(tiles, t)
		if t.x < minx {
			minx = t.x
		}
		if t.x > maxx {
			maxx = t.x
		}
		if t.y < miny {
			miny = t.y
		}
		if t.y > maxy {
			maxy = t.y
		}
	}
	offsetx, offsety := -minx, -miny
	minx += offsetx
	maxx += offsetx
	miny += offsety
	maxy += offsety

	grid := make([][]int, maxy+1)
	for i := range grid {
		grid[i] = make([]int, maxx+1)
	}

	for _, t := range tiles {
		grid[t.y+offsety][t.x+offsetx] = 1
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

}

func main() {
	s := util.GetFile("2019/Day11/input")

	var memory []int

	for _, op := range strings.Split(s, ",") {
		memory = append(memory, util.MustAtoi(strings.TrimSpace(op)))
	}

	p := program.NewProgram(memory)

	mu := sync.Mutex{}
	done := false
	go func() {
		p.Run()
		mu.Lock()
		done = true
		mu.Unlock()
	}()

	m := map[tile]int{
		{0, 0}: 1,
	}
	px, py := 0, 0
	d := N
	for {
		pos := tile{x: px, y: py}
		mu.Lock()
		if done {
			break
		}
		mu.Unlock()
		p.Input <- m[pos]

		clr, move := <-p.Output, <-p.Output
		m[pos] = clr
		d = nextDir(d, move)
		px, py = doMove(px, py, d)
	}

	fmt.Printf("Number of tiles painted (part1): %d\n", len(m))

	print(m)

}
