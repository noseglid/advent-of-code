package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type lanternfish int

func (l *lanternfish) DayPass() bool {
	*l = *l - 1
	if *l == -1 {
		*l = 6
		return true
	}

	return false
}

func NewLanternfish(timer int) *lanternfish {
	l := lanternfish(timer)
	return &l
}

func PrintState(fishes []*lanternfish) {
	sep := ""
	for _, f := range fishes {
		fmt.Printf("%s%d", sep, *f)
		sep = ","
	}
	fmt.Println()
}

func PrintStateMap(mm map[int]int) {
	for i := 0; i <= 8; i++ {
		fmt.Printf("%d: %d\n", i, mm[i])
	}
}

func part2(ff []int, days int) {
	mm := make(map[int]int)
	for _, f := range ff {
		mm[f]++
	}

	for i := 0; i < days; i++ {
		newFishes := mm[0]
		for t := 0; t < 8; t++ {
			mm[t] = mm[t+1]
		}

		mm[8] = newFishes
		mm[6] += newFishes
	}

	s := 0
	for _, v := range mm {
		s += v
	}

	log.Printf("Part 2: Fishes after %d days: %d", days, s)
}

func part1(ff []int, days int) {
	var fishes []*lanternfish
	for _, f := range ff {
		fishes = append(fishes, NewLanternfish(f))
	}

	for i := 0; i < days; i++ {
		newFishes := 0
		for _, f := range fishes {
			if f.DayPass() {
				newFishes++
			}
		}
		for j := 0; j < newFishes; j++ {
			fishes = append(fishes, NewLanternfish(8))
		}
	}

	log.Printf("Part 1: Fishes after %d days: %d", days, len(fishes))
}

func main() {
	input := "2021/Day6/input"
	nn := util.GetCSVFileNumbers(input)

	part1(nn, 80)
	part2(nn, 256)
}
