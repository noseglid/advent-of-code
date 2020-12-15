package main

import "log"

func nextNumber(input []int) int {
	search := input[len(input)-1]
	for i := len(input) - 2; i >= 0; i-- {
		if search == input[i] {
			return len(input) - 1 - i
		}
	}

	return 0
}

func main() {
	spoken := map[int]int{
		17: 1,
		1:  2,
		3:  3,
		16: 4,
		19: 5,
	}
	lastSpoken := 0

	for turn := len(spoken) + 2; turn <= 30000000+1; turn++ {
		turnLastSpoken, ok := spoken[lastSpoken]
		if !ok {
			spoken[lastSpoken] = turn - 1
			lastSpoken = 0
		} else {
			spoken[lastSpoken] = turn - 1
			lastSpoken = spoken[lastSpoken] - turnLastSpoken
		}
		if turn == 2020 {
			log.Printf("number after 2020 rounds (part1): %d", lastSpoken)
		} else if turn == 30000000 {
			log.Printf("number after 30000000 rounds (part2): %d", lastSpoken)
		}

	}
}
