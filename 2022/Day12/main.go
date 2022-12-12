package main

import (
	"fmt"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

func printGrid(g [][]int) {
	for _, r := range g {
		for _, c := range r {
			if c > 99 {
				fmt.Printf("** ")
			} else {
				fmt.Printf("%02d ", c)
			}
		}
		fmt.Println()
	}
}

func stepGrid(elevation [][]int, cost [][]int, p util.Point) {
	currEl := elevation[p.Y][p.X]
	if p.Y > 0 && currEl <= elevation[p.Y-1][p.X]+1 {
		if cost[p.Y-1][p.X] > cost[p.Y][p.X]+1 {
			cost[p.Y-1][p.X] = cost[p.Y][p.X] + 1
			stepGrid(elevation, cost, util.Point{X: p.X, Y: p.Y - 1})
		}
	}

	if p.Y < len(elevation)-1 && currEl <= elevation[p.Y+1][p.X]+1 {
		if cost[p.Y+1][p.X] > cost[p.Y][p.X]+1 {
			cost[p.Y+1][p.X] = cost[p.Y][p.X] + 1
			stepGrid(elevation, cost, util.Point{X: p.X, Y: p.Y + 1})
		}
	}

	if p.X > 0 && currEl <= elevation[p.Y][p.X-1]+1 {
		if cost[p.Y][p.X-1] > cost[p.Y][p.X]+1 {
			cost[p.Y][p.X-1] = cost[p.Y][p.X] + 1
			stepGrid(elevation, cost, util.Point{X: p.X - 1, Y: p.Y})
		}
	}

	if p.X < len(elevation[p.Y])-1 && currEl <= elevation[p.Y][p.X+1]+1 {
		if cost[p.Y][p.X+1] > cost[p.Y][p.X]+1 {
			cost[p.Y][p.X+1] = cost[p.Y][p.X] + 1
			stepGrid(elevation, cost, util.Point{X: p.X + 1, Y: p.Y})
		}
	}
}

func buildCostGrid(elevation [][]int, p util.Point) [][]int {
	ret := make([][]int, len(elevation))
	for i, e := range elevation {
		ret[i] = make([]int, len(e))
		for j := range ret[i] {
			ret[i][j] = math.MaxInt
		}
	}

	ret[p.Y][p.X] = 0

	stepGrid(elevation, ret, p)

	return ret
}

func main() {
	grid := util.GetFileRuneGrid("2022/Day12/input")
	elevation := make([][]int, len(grid))

	var start, end util.Point
	for y, row := range grid {
		elevation[y] = make([]int, len(row))
		for x, p := range row {
			switch p {
			case 'S':
				start.Set(x, y)
				p = 'a'
			case 'E':
				end.Set(x, y)
				p = 'z'
			}
			elevation[y][x] = int(p - 'a')
		}
	}

	costs := buildCostGrid(elevation, end)

	fmt.Printf("Cost to move to end (part1): %d\n", costs[start.Y][start.X])

	min := math.MaxInt
	for y := range elevation {
		for x := range elevation[y] {
			if elevation[y][x] == 0 {
				if costs[y][x] < min {
					min = costs[y][x]
				}
			}
		}
	}

	fmt.Printf("Cost from lowest point with fewest steps (part2): %d", min)
}
