package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type State string

const (
	Floor    = State(".")
	Empty    = State("L")
	Occupied = State("#")
)

func printGrid(grid [][]State) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func neighbourState(x, y int, grid [][]State) (int, int, int) {
	neighbours := []struct{ dx, dy int }{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}

	floor := 0
	empty := 0
	occupied := 0
	for _, n := range neighbours {
		if (x+n.dx) < 0 || (x+n.dx) >= len(grid[0]) || (y+n.dy) < 0 || (y+n.dy) >= len(grid) {
			continue
		}
		state := grid[y+n.dy][x+n.dx]
		if state == Occupied {
			occupied++
		} else if state == Empty {
			empty++
		} else {
			floor++
		}
	}

	return floor, empty, occupied
}

func neighbourStateP2(x, y int, grid [][]State) (int, int, int) {
	neighbours := []struct{ dx, dy int }{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}

	floor := 0
	empty := 0
	occupied := 0
	for _, n := range neighbours {
		state := Floor
		for i := 1; state == Floor; i++ {
			if (x+i*n.dx) < 0 || (x+i*n.dx) >= len(grid[0]) || (y+i*n.dy) < 0 || (y+i*n.dy) >= len(grid) {
				break
			}

			state = grid[y+i*n.dy][x+i*n.dx]
		}

		if state == Occupied {
			occupied++
		} else if state == Empty {
			empty++
		} else {
			floor++
		}
	}

	return floor, empty, occupied
}

func nextState(state State, _, _, occupied, reqOccupied int) State {
	nextState := state
	switch state {
	case Empty:
		if occupied == 0 {
			nextState = Occupied
		}
	case Occupied:
		if occupied >= reqOccupied {
			nextState = Empty
		}
	}

	return nextState
}

func copyGrid(grid [][]State) [][]State {
	ngrid := make([][]State, len(grid))
	for y := range grid {
		ngrid[y] = make([]State, len(grid[y]))
		for x := range grid[y] {
			ngrid[y][x] = grid[y][x]
		}
	}
	return ngrid
}

func iterate(grid [][]State, neighbourStatesfn func(int, int, [][]State) (int, int, int), reqOccupied int) ([][]State, bool) {
	ngrid := copyGrid(grid)
	didChange := false

	for y := range grid {
		for x := range grid[y] {
			floor, empty, occupied := neighbourStatesfn(x, y, grid)
			currentState := grid[y][x]
			ngrid[y][x] = nextState(currentState, floor, empty, occupied, reqOccupied)
			if ngrid[y][x] != currentState {
				didChange = true
			}
		}
	}
	return ngrid, didChange
}

func count(grid [][]State, state State) int {
	n := 0
	for _, row := range grid {
		for _, col := range row {
			if col == state {
				n++
			}
		}
	}
	return n
}

func main() {
	s := util.FileScanner("2020/Day11/input", bufio.ScanLines)

	var grid [][]State

	for s.Scan() {
		var row []State
		for _, r := range s.Text() {
			switch r {
			case 'L':
				row = append(row, Empty)
			case '#':
				row = append(row, Occupied)
			case '.':
				row = append(row, Floor)
			}
		}
		grid = append(grid, row)
	}

	gridp1 := copyGrid(grid)
	for {
		var didChange bool
		gridp1, didChange = iterate(gridp1, neighbourState, 4)
		if !didChange {
			break
		}
	}
	log.Printf("number of occupied seat in stable state (part1): %d", count(gridp1, Occupied))

	gridp2 := copyGrid(grid)
	for {
		var didChange bool
		gridp2, didChange = iterate(gridp2, neighbourStateP2, 5)
		if !didChange {
			break
		}

	}

	// floor, empty, occupied := neighbourStateP2(3, 3, gridp2)
	// log.Printf("floor=%d, empty=%d, occupied=%d", floor, empty, occupied)

	log.Printf("number of occupied seat in stable state (part2): %d", count(gridp2, Occupied))

}
