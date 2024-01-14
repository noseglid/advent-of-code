package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/2019/Day13/program"
	"github.com/noseglid/advent-of-code/util"
)

type tile struct {
	id   int
	x, y int
}

func readSegment(ch <-chan int) ([]int, bool) {
	s := []int{}
	for e := range ch {
		s = append(s, e)
		if len(s) == 3 {
			return s, true
		}
	}
	return nil, false
}

func buildTiles(ch <-chan int) ([]tile, int) {
	tiles := []tile{}
	for {
		segment, ok := readSegment(ch)
		if !ok {
			break
		}
		tiles = append(tiles, tile{
			x:  segment[0],
			y:  segment[1],
			id: segment[2],
		})
	}
	return tiles, 0
}

func main() {
	s := util.GetFile("2019/Day13/input")

	var memory []int
	for _, op := range strings.Split(s, ",") {
		memory = append(memory, util.MustAtoi(strings.TrimSpace(op)))
	}
	memory2 := append([]int{}, memory...)
	memory2[0] = 2

	p := program.NewProgram(memory)
	go p.Run()

	tiles, _ := buildTiles(p.Output)
	nblock := 0
	for _, t := range tiles {
		if t.id == 2 {
			nblock++
		}
	}
	fmt.Printf("Number of block tiles (part1): %d\n", nblock)

	p2 := program.NewProgram(memory2)
	go p2.Run()
	bx, px := 0, 0

	score := 0
Outer:
	for {
		s, ok := readSegment(p2.Output)
		if !ok {
			if p2.Halted {
				break
			}
			break Outer
		}
		if s[0] == -1 && s[1] == 0 {
			score = s[2]
		}
		if s[2] == 3 {
			px = s[0]
		}
		if s[2] == 4 {
			bx = s[0]
			if bx < px {
				p2.WriteInput(-1)
				px--
			} else if bx > px {
				p2.WriteInput(1)
				px++
			} else {
				p2.WriteInput(0)
			}
		}
	}

	fmt.Printf("Score after beating the game (part2): %d\n", score)
}
