package main

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type action int

func (a *action) fromString(s string) {
	switch s {
	case "turn on":
		*a = TurnOn
	case "turn off":
		*a = TurnOff
	case "toggle":
		*a = Toggle
	default:
		log.Fatalf("invalid action: %s", s)
	}
}

const (
	TurnOn  = action(0)
	TurnOff = action(1)
	Toggle  = action(2)
)

type point struct {
	x, y int
}

type instruction struct {
	action action
	p0, p1 point
}

var re = regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

func parseInstruction(s string) instruction {
	match := re.FindAllStringSubmatch(s, -1)
	i := instruction{
		p0: point{x: util.MustAtoi(match[0][2]), y: util.MustAtoi(match[0][3])},
		p1: point{x: util.MustAtoi(match[0][4]), y: util.MustAtoi(match[0][5])},
	}
	i.action.fromString(match[0][1])
	return i
}

func buildGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}

func applyInstructionPart1(grid [][]int, instr instruction) {
	for x := instr.p0.x; x <= instr.p1.x; x++ {
		for y := instr.p0.y; y <= instr.p1.y; y++ {
			delta := 0
			switch instr.action {
			case Toggle:
				if grid[x][y] == 0 {
					grid[x][y] = 1
				} else {
					grid[x][y] = 0
				}
			case TurnOn:
				grid[x][y] = 1
			case TurnOff:
				grid[x][y] = 0
			default:
				log.Fatal("invalid instruction")
			}
			grid[x][y] = grid[x][y] + delta
		}
	}
}

func applyInstructionPart2(grid [][]int, instr instruction) {
	for x := instr.p0.x; x <= instr.p1.x; x++ {
		for y := instr.p0.y; y <= instr.p1.y; y++ {
			delta := 0
			switch instr.action {
			case Toggle:
				delta = 2
			case TurnOn:
				delta = 1
			case TurnOff:
				if grid[x][y] > 0 {
					delta = -1
				}
			default:
				log.Fatal("invalid instruction")
			}
			grid[x][y] = grid[x][y] + delta
		}
	}
}

func countOn(grid [][]int) int {
	n := 0
	for _, row := range grid {
		for _, on := range row {
			if on == 1 {
				n++
			}
		}
	}
	return n
}

func countBrightness(grid [][]int) int {
	brightness := 0
	for _, row := range grid {
		for _, b := range row {
			brightness += b
		}
	}
	return brightness
}

func main() {
	f, err := os.Open("2015/Day6/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	grid := buildGrid(1000)
	gridp2 := buildGrid(1000)

	for s.Scan() {
		instr := parseInstruction(s.Text())
		applyInstructionPart1(grid, instr)
		applyInstructionPart2(gridp2, instr)
	}

	log.Printf("number of lights on (part1): %d", countOn(grid))
	log.Printf("total brightness (part2): %d", countBrightness(gridp2))
}
