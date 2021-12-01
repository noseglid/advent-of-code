package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var vowels = "aeiou"
var forbidden = []string{"ab", "cd", "pq", "xy"}

type letterPair struct {
	start int
	end   int
	c1    rune
	c2    rune
}

func (lp letterPair) String() string {
	return fmt.Sprintf("%d-%d: %c%c", lp.start, lp.end, lp.c1, lp.c2)
}

func hasVowels(s string) bool {
	sum := 0
	for _, v := range vowels {
		sum += strings.Count(s, string(v))
	}
	return sum >= 3
}

func hasConsecutive(s string) bool {
	l1 := rune(s[0])
	for _, l2 := range s[1:] {
		if l1 == l2 {
			return true
		}
		l1 = l2
	}

	return false
}

func hasForbidden(s string) bool {
	for _, f := range forbidden {
		if strings.Contains(s, f) {
			return true
		}
	}

	return false
}

func hasPair(s string) bool {
	pairs := []letterPair{}
	l1 := rune(s[0])
	for i, l2 := range s[1:] {
		pairs = append(pairs, letterPair{i, i + 1, l1, l2})
		l1 = l2
	}

	for _, po := range pairs {
		for _, pi := range pairs {
			if po.c1 == pi.c1 && po.c2 == pi.c2 && po.end != pi.start && po.start != pi.end && po != pi {
				return true
			}
		}
	}

	return false
}

func hasSkippedRepeat(s string) bool {
	l1, l2 := rune(s[0]), rune(s[1])
	for _, l3 := range s[2:] {
		if l1 == l3 {
			return true
		}
		l1 = l2
		l2 = l3
	}
	return false
}

func main() {
	f, err := os.Open("2015/Day5/input")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("hasSkippedRepeate: %v", hasSkippedRepeat("uurcxstgmygtbstg"))

	s := bufio.NewScanner(f)

	npart1 := 0
	npart2 := 0
	for s.Scan() {
		line := s.Text()
		if hasVowels(line) && hasConsecutive(line) && !hasForbidden(line) {
			npart1++
		}

		if hasPair(line) && hasSkippedRepeat(line) {
			npart2++
		}
	}

	log.Printf("Nice strings (part1): %d", npart1)
	log.Printf("Nice strings (part2): %d", npart2)
}
