package main

import (
	"log"
	"math"
)

var containers = []int{
	11,
	30,
	47,
	31,
	32,
	36,
	3,
	1,
	5,
	3,
	32,
	36,
	15,
	11,
	46,
	26,
	28,
	1,
	19,
	3,
}
var target = 150

func volumeSumIsTarget(containers []int, bitset int) (bool, int) {
	totalSum := 0
	used := 0

	for i, v := range containers {
		if bitset>>i&1 == 1 {
			totalSum += v
			used++
		}
	}

	return totalSum == target, used
}

func main() {
	combinations := int(math.Pow(2, float64(len(containers)))) - 1

	nper := map[int]int{}

	n := 0
	for i := 0; i < combinations; i++ {
		if exactly, used := volumeSumIsTarget(containers, i); exactly {
			n++
			nper[used]++
		}
	}

	log.Printf("container combos (part1): %d", n)

	minn := len(containers)
	minc := combinations
	for used, c := range nper {
		if used < minn {
			minn = used
			minc = c
		}
	}
	log.Printf("min containers (part2): %d", minc)

}
