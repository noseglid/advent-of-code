package main

import "log"

func main() {
	mult, div := 252533, 33554393
	trow, tcol := 3010, 3019
	// trow, tcol := 6, 6

	row, col := 1, 1
	value := 20151125
	for {
		// log.Printf("row %d, col %d = %d", row, col, value)
		row--
		if row == 0 {
			row = col + 1
			col = 1
		} else {
			col++
		}

		value = (value * mult) % div
		if row == trow && col == tcol {
			break
		}
	}

	log.Printf("row %d, col %d = %d", row, col, value)
}
