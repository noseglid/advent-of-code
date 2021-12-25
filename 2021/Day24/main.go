package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func getNumberOffset(lines []string, offset int) []int {
	var res []int
	for i := 0; i < 14; i++ {
		parts := strings.Split(lines[offset+18*i], " ")
		res = append(res, util.MustAtoi(parts[2]))
	}
	return res
}

func findz0(ix, iy, iz, z2, w int) []int {
	zs := []int{}
	if 0 <= w-ix && w-ix < 26 {
		z0 := z2 * iz
		zs = append(zs, w-ix+z0)
	}

	x := z2 - w - iy
	if x%26 == 0 {
		zs = append(zs, (x/26)*iz)
	}

	return zs
}

func intSliceString(s []int) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(strconv.Itoa(r))
	}
	return sb.String()
}

func intContains(s []int, v int) bool {
	for _, ss := range s {
		if ss == v {
			return true
		}
	}
	return false
}

func solveForRange(addX, addY, divZ, ws []int) string {
	result := map[int][]int{}
	zs := []int{0} // start with 0
	for i := range addX {
		var nextZ []int
		idx := len(addX) - i - 1
		ix, iy, iz := addX[idx], addY[idx], divZ[idx]
		for _, w := range ws {
			for _, z := range zs {
				for _, z0 := range findz0(ix, iy, iz, z, w) {
					if !intContains(nextZ, z0) {
						nextZ = append(nextZ, z0)
					}
					result[z0] = append([]int{w}, result[z]...)
				}
			}
		}
		zs = nextZ
	}
	return intSliceString(result[0])
}

func main() {
	input := "2021/Day24/input"
	lines := util.GetFileStrings(input)

	divZ := getNumberOffset(lines, 4)
	addX := getNumberOffset(lines, 5)
	addY := getNumberOffset(lines, 15)

	p1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	p2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	log.Printf("Part 1: Highest valid model number: %s", solveForRange(addX, addY, divZ, p1))
	log.Printf("Part 2: Lowest valid model number: %s", solveForRange(addX, addY, divZ, p2))

}
