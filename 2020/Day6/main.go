package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func groupYes(s string) int {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)

	yeses := map[string]int{}

	for scanner.Scan() {
		r := scanner.Text()
		if rune(r[0]) >= 'a' && rune(r[0]) <= 'z' {
			yeses[r]++
		}

	}

	return len(yeses)
}

func groupAllYes(s string) int {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)

	yeses := map[rune]int{}

	groupSize := 0
	for scanner.Scan() {
		groupSize++
		personAnswer := scanner.Text()
		for _, r := range personAnswer {
			if r >= 'a' && r <= 'z' {
				yeses[r]++
			}
		}
	}

	groupYeses := 0
	for _, c := range yeses {
		if c == groupSize {
			groupYeses++
		}
	}

	return groupYeses
}

func main() {
	input := util.GetFile("2020/Day6/input")

	answers := strings.Split(input, "\n\n")
	totalYeses := 0
	allYeses := 0
	for _, group := range answers {
		totalYeses += groupYes(group)
		allYeses += groupAllYes(group)
	}

	log.Printf("total yeses (part1): %d", totalYeses)
	log.Printf("total all yeses (part2): %d", allYeses)

}
