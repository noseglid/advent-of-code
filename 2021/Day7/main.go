package main

import (
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

type Coster func(a, b int) int

func costConstant(a, b int) int {
	return util.Absolute(a - b)
}

func costLinear(a, b int) int {
	d := util.Absolute(a - b)
	return (d * (d + 1)) / 2
}

func alignCost(pos int, positions []int, coster Coster) int {
	t := 0
	for _, p := range positions {
		t += coster(p, pos)
	}
	return t
}

func main() {
	input := "2021/Day7/input"
	n := util.GetCSVFileNumbers(input)
	min, max := util.MinIntList(n), util.MaxIntList(n)

	minCostp1, minCostp2 := math.MaxInt, math.MaxInt
	for i := min; i <= max; i++ {
		costConstant := alignCost(i, n, costConstant)
		if costConstant < minCostp1 {
			minCostp1 = costConstant
		}

		costLinear := alignCost(i, n, costLinear)
		if costLinear < minCostp2 {
			minCostp2 = costLinear
		}

	}

	log.Printf("Part 1: Minimum fuel alignment cost: %d", minCostp1)
	log.Printf("Part 2: Minimum fuel alignment cost: %d", minCostp2)

}
