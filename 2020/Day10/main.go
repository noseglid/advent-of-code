package main

import (
	"bufio"
	"log"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

func adapterDiffs(adapters []int) ([]int, int) {
	var diffs []int
	cjoltage := 0
	remainingAdapters := adapters

	for {
		index, minDiff, diff := -1, -1, -1

		for i, a := range remainingAdapters {
			d := a - cjoltage
			// log.Printf("checking adapter %d with currentj %d, got diff %d", a, cjoltage, d)
			if d >= 1 && d <= 3 {
				if minDiff == -1 || d < minDiff {
					// log.Printf("got candidate, adapterj = %d", a)
					diff = d
					minDiff = d
					index = i
				}
			}
		}

		if index == -1 {
			if len(remainingAdapters) > 0 {
				log.Fatalf("adapters left: %v", remainingAdapters)
			}
			break
		}
		cjoltage = remainingAdapters[index]
		remainingAdapters = append(remainingAdapters[:index], remainingAdapters[index+1:]...)
		diffs = append(diffs, diff)
	}

	return diffs, cjoltage
}

func countDiffs(diffs []int, s int) int {
	n := 0
	for _, d := range diffs {
		if d == s {
			n++
		}
	}
	return n
}

func countArrangements(sortedAdapters []int, target int) int {
	arrangements := map[int]int{}
	arrangements[0] = 1

	for _, a := range sortedAdapters {
		arrangements[a] = arrangements[a-1] + arrangements[a-2] + arrangements[a-3]
	}

	return arrangements[target]
}

func cpy(adapters []int) []int {
	result := make([]int, len(adapters))
	copy(result, adapters)
	return result
}

func main() {

	s := util.FileScanner("2020/Day10/input", bufio.ScanLines)

	var adapters []int

	for s.Scan() {
		adapters = append(adapters, util.MustAtoi(s.Text()))
	}

	sort.Ints(adapters)
	diffs, deviceJoltage := adapterDiffs(cpy(adapters))
	diffs = append(diffs, 3) // add the device
	log.Printf("min and max multiplied (part1): %d", countDiffs(diffs, 1)*countDiffs(diffs, 3))

	log.Printf("arrangements (part2): %d", countArrangements(append(cpy(adapters), deviceJoltage), deviceJoltage))
}
