package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func blink(memo map[string]int, stone, it int) int {
	k := fmt.Sprintf("%d-%d", stone, it)
	if v, ok := memo[k]; ok {
		return v
	} else if it == 0 {
		memo[k] = 1
	} else if stone == 0 {
		memo[k] = blink(memo, 1, it-1)
	} else if ss := fmt.Sprintf("%d", stone); len(ss)%2 == 0 {
		memo[k] = blink(memo, util.MustAtoi(ss[0:len(ss)/2]), it-1) + blink(memo, util.MustAtoi(ss[len(ss)/2:]), it-1)
	} else {
		memo[k] = blink(memo, stone*2024, it-1)
	}
	return memo[k]

}

func main() {
	// a := "125 17"
	a := "8069 87014 98 809367 525 0 9494914 5"

	n := []int{}
	for _, aa := range strings.Split(a, " ") {
		n = append(n, util.MustAtoi(aa))
	}

	memo := map[string]int{}
	s1, s2 := 0, 0
	for _, v := range n {
		s1 += blink(memo, v, 25)
		s2 += blink(memo, v, 75)
	}

	fmt.Printf("Number of stones after 25 blinks (part1): %d\n", s1)
	fmt.Printf("Number of stones after 75 blinks (part2): %d\n", s2)

}
