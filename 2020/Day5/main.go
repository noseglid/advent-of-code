package main

import (
	"bufio"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func parsePartition(s string) (int, int) {
	rowRange := [2]int{0, 127}

	for _, r := range s[0:7] {
		d := (rowRange[1] - rowRange[0]) / 2
		switch r {
		case 'F':
			rowRange[1] = rowRange[0] + d
		case 'B':
			rowRange[0] = rowRange[1] - d
		}
	}

	colRange := [2]int{0, 7}
	for _, c := range s[7:] {
		d := (colRange[1] - colRange[0]) / 2
		switch c {
		case 'L':
			colRange[1] = colRange[0] + d
		case 'R':
			colRange[0] = colRange[1] - d
		}
	}

	return rowRange[0], colRange[0]
}

func calcSeatID(row, col int) int {
	return row*8 + col
}

func main() {

	s := util.FileScanner("2020/Day5/input", bufio.ScanLines)

	occupied := make([][]bool, 128)
	for row := range occupied {
		occupied[row] = make([]bool, 8)
	}

	maxSeatID := 0
	for s.Scan() {
		row, col := parsePartition(s.Text())
		occupied[row][col] = true

		seatID := calcSeatID(row, col)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	for row := range occupied {
		for col := range occupied[row] {
			if !occupied[row][col] {
				log.Printf("free: %d,%d with id %d", row, col, calcSeatID(row, col))
			}

		}
	}

	log.Printf("max seat id (part1): %d", maxSeatID)

}
