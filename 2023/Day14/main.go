package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func rotateMatrix[T any](matrix [][]T) [][]T {

	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func tiltNorth(grid [][]rune) {
	didMove := true
	for didMove {
		didMove = false
		for row := 1; row < len(grid); row++ {
			for col := range grid[row] {
				if grid[row][col] != 'O' {
					continue
				}

				if grid[row-1][col] != '.' {
					continue
				}

				grid[row-1][col] = 'O'
				grid[row][col] = '.'
				didMove = true
			}
		}
	}
}

func load(grid [][]rune) int {
	load := 0
	height := len(grid)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != 'O' {
				continue
			}
			load += height - row
		}
	}
	return load
}

func key(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		for _, col := range row {
			sb.WriteRune(col)
		}
	}
	return sb.String()
}

func findCycle(loads []int) (int, int) {
	if len(loads) < 100 {
		return -1, -1
	}
	test := loads[len(loads)-1]

	for j := len(loads) - 2; j >= 0; j-- {
		if loads[j] != test {
			continue
		}
		length := len(loads) - 1 - j
		if reflect.DeepEqual(loads[j-length:j], loads[j:j+length]) {
			return j, length
		}
	}

	return -1, -1
}

func main() {
	src := "2023/Day14/input"
	grid := util.GetFileRuneGrid(src)
	tiltNorth(grid)
	fmt.Printf("Load (part1): %d\n", load(grid))

	grid2 := util.GetFileRuneGrid(src)
	loads := []int{}

	cycleLength := -1
	for i := 1; i < 500; i++ {
		tiltNorth(grid2)
		rotateMatrix(grid2)
		tiltNorth(grid2)
		rotateMatrix(grid2)
		tiltNorth(grid2)
		rotateMatrix(grid2)
		tiltNorth(grid2)
		rotateMatrix(grid2)
		loads = append(loads, load(grid2))

		if start, length := findCycle(loads); start != -1 && length != -1 {
			cycleLength = length
		}
	}
	m := (1e9 / cycleLength) - 10
	fmt.Printf("Load after 1e9 cycles (part2): %d\n", loads[1e9-m*cycleLength-1])
}
