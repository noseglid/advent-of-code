package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func toKey(list []int) string {
	var sb strings.Builder
	for _, n := range list {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

func largestIndex(list []int) int {
	m := list[0]
	i := 0
	for j, e := range list[1:] {
		if e > m {
			m = e
			i = j + 1
		}
	}
	return i
}

func distribute(list []int, index int) {
	v := list[index]
	list[index] = 0
	ii := (index + 1) % len(list)
	for i := 0; i < v; i++ {
		list[ii]++
		ii = (ii + 1) % len(list)
	}
}

func main() {
	input := "2017/Day6/input"
	var list []int
	for _, d := range strings.Fields(util.GetFile(input)) {
		list = append(list, util.MustAtoi(d))
	}

	seen := map[string]int{}

	n := 0
	lastSeen := 0
	for {
		k := toKey(list)
		if ls, ok := seen[k]; ok {
			lastSeen = ls
			break
		}

		i := largestIndex(list)
		distribute(list, i)
		seen[k] = n
		n++
	}

	log.Printf("Part 1: Steps until infinite loop: %d", n)
	log.Printf("Part 2: Cycles until infinite loop: %d", n-lastSeen)

}
