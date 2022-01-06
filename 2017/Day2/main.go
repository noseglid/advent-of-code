package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func minmax(data []int) (int, int) {
	min, max := data[0], data[0]
	for _, d := range data[1:] {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
	}
	return min, max
}

func checksum(data [][]int) int {
	s := 0
	for _, r := range data {
		min, max := minmax(r)
		s += max - min
	}
	return s
}
func checksum2(data [][]int) int {
	s := 0
	for _, r := range data {
		for i, a := range r {
			for j, b := range r {
				if b == 0 || i == j {
					continue
				}
				if a%b == 0 {
					s += a / b
				}
			}
		}
	}
	return s
}

func main() {
	input := "2017/Day2/input"
	lines := util.GetFileStrings(input)
	var data [][]int
	for _, l := range lines {
		var row []int
		for _, n := range strings.Fields(l) {
			row = append(row, util.MustAtoi(n))
		}
		data = append(data, row)
	}

	log.Printf("Part 1: checksum %d", checksum(data))
	log.Printf("Part 2: checksum %d", checksum2(data))

}
