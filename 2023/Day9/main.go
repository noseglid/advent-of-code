package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func allZero(r []int) bool {
	for _, n := range r {
		if n != 0 {
			return false
		}
	}
	return true
}

func nextRow(r []int) []int {
	var result []int
	for i := 0; i < len(r)-1; i++ {
		result = append(result, r[i+1]-r[i])
	}
	return result
}

func diffRows(r []int) [][]int {
	diffRows := [][]int{}
	diffRows = append(diffRows, r)

	i := 0
	for !allZero(r) {
		r = nextRow(r)
		diffRows = append(diffRows, r)
		i++
	}
	return diffRows
}

func addValues(r [][]int) {
	r[len(r)-1] = append(r[len(r)-1], 0)
	for i := len(r) - 2; i >= 0; i-- {
		r[i] = append(r[i], r[i][len(r[i])-1]+r[i+1][len(r[i+1])-1])
	}
}

func addValuesPre(r [][]int) {
	for i := range r {
		r[i] = append(r[i], r[i][len(r[i])-1])
		for j := len(r[i]) - 1; j >= 0; j-- {
			if j == 0 {
				r[i][j] = 0
			} else {
				r[i][j] = r[i][j-1]
			}
		}
	}

	for i := len(r) - 2; i >= 0; i-- {
		r[i][0] = r[i][1] - r[i+1][0]
	}
}

func main() {
	var rows [][]int
	for _, l := range util.GetFileStrings("2023/Day9/input") {
		rows = append(rows, util.NumberList(l))
	}

	s := 0
	for _, r := range rows {
		dr := diffRows(r)
		addValues(dr)
		s += dr[0][len(dr[0])-1]
	}

	s2 := 0
	for _, r := range rows {
		dr := diffRows(r)
		addValuesPre(dr)
		s2 += dr[0][0]
	}
	fmt.Printf("sum of added values (part1): %d\n", s)
	fmt.Printf("sum of added values pre (part2): %d\n", s2)

}
