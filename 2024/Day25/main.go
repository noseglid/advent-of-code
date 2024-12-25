package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func overlaps(key, lock []int) bool {
	for i := 0; i < 5; i++ {
		if key[i]+lock[i] > 5 {
			return true
		}
	}
	return false
}

func main() {
	keys := [][]int{}
	locks := [][]int{}

	lines := util.GetFileStrings("2024/Day25/input")
	for i := 0; i < len(lines); i += 8 {
		blocks := lines[i : i+7]

		ns := []int{}
		for col := 0; col < 5; col++ {
			n := 0
			for row := 0; row < 7; row++ {
				if blocks[row][col] == '#' {
					n++
				}
			}
			ns = append(ns, n-1)
		}

		if blocks[0] == "#####" {
			locks = append(locks, ns)
		} else {
			keys = append(keys, ns)
		}
	}

	s := 0
	for _, key := range keys {
		for _, lock := range locks {
			if !overlaps(key, lock) {
				s++
			}

		}
	}

	fmt.Printf("keys which fit (part1): %d\n", s)
}
