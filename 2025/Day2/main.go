package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func idsplit(s string) (int, int) {
	spl := strings.Split(strings.TrimSpace(s), "-")
	return util.MustAtoi(spl[0]), util.MustAtoi(spl[1])
}

func candidatesEq(candidates []string) bool {
	for i := 1; i < len(candidates); i++ {
		if candidates[i] != candidates[0] {
			return false
		}
	}
	return true
}

func idIsInvalid(id int) bool {
	s := fmt.Sprintf("%d", id)

	if len(s)%2 != 0 {
		return false
	}

	return s[0:len(s)/2] == s[len(s)/2:]
}

func idIsInvalidp2(id int) bool {
	s := fmt.Sprintf("%d", id)

	for size := 1; size <= len(s)/2; size++ {
		var candidates []string
		if len(s)%size != 0 {
			continue
		}
		for j := 0; j < len(s)/size; j++ {
			candidates = append(candidates, s[j*size:(j+1)*size])
		}

		if candidatesEq(candidates) {
			return true
		}
	}

	return false
}

func main() {

	input := util.GetFile("2025/Day2/input")

	ids := strings.Split(input, ",")

	sum, sump2 := 0, 0

	for _, id := range ids {
		s, e := idsplit(id)

		for i := s; i <= e; i++ {
			if idIsInvalid(i) {
				sum += i
			}

			if idIsInvalidp2(i) {
				sump2 += i
			}
		}
	}

	fmt.Printf("sum of invalid ids (part1): %d\n", sum)
	fmt.Printf("sum of invalid ids (part2): %d\n", sump2)

}
