package util

import "fmt"

func PrintRuneGrid(grid [][]rune) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}
