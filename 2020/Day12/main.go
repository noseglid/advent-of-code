package main

import (
	"bufio"
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type instr rune

const (
	North   = instr('N')
	South   = instr('S')
	East    = instr('E')
	West    = instr('W')
	Left    = instr('L')
	Right   = instr('R')
	Forward = instr('F')
)

type Direction int32

const (
	DNorth = Direction(0)
	DEast  = Direction(1)
	DSouth = Direction(2)
	DWest  = Direction(3)
)

func parseInstruction(s string) (instr, int) {
	return instr(s[0]), util.MustAtoi(s[1:])
}

func turn(facing Direction, direction instr, degrees int) Direction {
	its := degrees / 90
	if direction == Left {
		its *= -1
	}

	return Direction((int(facing) + its + 4) % 4)
}

func rotate(x, y int, angle float64) (int, int) {
	c := int(math.Cos(angle))
	s := int(math.Sin(angle))
	log.Printf("%d,%d", c, s)
	return x*c - y*s, x*s + y*c
}

func forward(x, y int, n int, dir Direction) (int, int) {
	switch dir {
	case DNorth:
		return x, y - n
	case DEast:
		return x + n, y
	case DSouth:
		return x, y + n
	case DWest:
		return x - n, y
	default:
		panic("bad forward direction")
	}
}

type ship struct {
	x, y int
	dir  Direction
}

type waypoint struct {
	x, y int
}

func p1() {
	s := util.FileScanner("2020/Day12/input_sample", bufio.ScanLines)

	x, y := 0, 0
	dir := DEast
	for s.Scan() {
		ins, n := parseInstruction(s.Text())
		switch ins {
		case North:
			y -= n
		case East:
			x += n
		case South:
			y += n
		case West:
			x -= n
		case Left:
			fallthrough
		case Right:
			dir = turn(dir, ins, n)
		case Forward:
			x, y = forward(x, y, n, dir)
		default:
			panic("bad instruction")
		}
	}

	log.Printf("sum of manhattan position (part1): %.0f", math.Abs(float64(x))+math.Abs(float64(y)))
}

func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func p2() {
	s := util.FileScanner("2020/Day12/input", bufio.ScanLines)

	sh := ship{0, 0, DEast}
	wp := waypoint{10, -1}
	log.Printf("Ship %+v", sh)
	log.Printf("Waypoint %+v", wp)
	for s.Scan() {
		ins, n := parseInstruction(s.Text())
		log.Printf("%v, %d", string(ins), n)
		switch ins {
		case North:
			wp.y -= n
		case East:
			wp.x += n
		case South:
			wp.y += n
		case West:
			wp.x -= n
		case Left:
			wp.x, wp.y = rotate(wp.x, wp.y, -deg2rad(float64(n)))
		case Right:
			wp.x, wp.y = rotate(wp.x, wp.y, deg2rad(float64(n)))
		case Forward:
			sh.x, sh.y = sh.x+n*wp.x, sh.y+n*wp.y
		default:
			panic("bad instruction")
		}
		log.Printf("Ship %+v", sh)
		log.Printf("Waypoint %+v", wp)
		log.Printf("----------")
	}

	log.Printf("sum of manhattan position (%d,%d) (part1): %.0f", sh.x, sh.y, math.Abs(float64(sh.x))+math.Abs(float64(sh.y)))
}

func main() {
	p2()
}
