package main

import (
	"log"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type room struct {
	name     string
	letters  map[rune]int
	sector   int
	checksum string
}

func contains(r rune, l []rune) bool {
	for _, rr := range l {
		if rr == r {
			return true
		}
	}
	return false
}

func nextMost(letters map[rune]int, excl []rune) []rune {
	var mostRune []rune
	n := 0
	for r, c := range letters {
		if contains(r, excl) {
			continue
		}
		if c > n {
			mostRune = []rune{r}
			n = c
		} else if c == n {
			mostRune = append(mostRune, r)
		}
	}
	return mostRune
}

func (r room) ChecksumMatches() bool {
	used := []rune{}
	checksum := r.checksum
	for {
		if len(checksum) == 0 {
			return true
		}

		most := nextMost(r.letters, used)
		if len(most) == 1 {
			if rune(checksum[0]) != most[0] {
				return false
			}
			used = append(used, rune(checksum[0]))
			checksum = checksum[1:]
		} else if len(most) > 1 {
			n := util.MinInt(len(most)-1, len(checksum)-1)
			isSorted := []int{}
			for i := 0; i < n+1; i++ {
				if !contains(rune(checksum[i]), most) {
					return false
				}
				isSorted = append(isSorted, int(checksum[i]))
			}
			if !sort.IntsAreSorted(isSorted) {
				return false
			}
			used = append(used, []rune(checksum[:n+1])...)
			checksum = checksum[n+1:]
		}

	}
}

func rotate(r rune, i int) rune {
	ri := int(r)
	for j := 0; j < i; j++ {
		if rune(ri) == 'z' {
			ri = int('a')
		} else {
			ri++
		}
	}

	return rune(ri)
}

func (r room) Decrypt() string {
	var sb strings.Builder
	for _, ru := range r.name {
		if ru == '-' {
			sb.WriteRune('-')
		} else {
			sb.WriteRune(rotate(ru, r.sector))
		}
	}

	return sb.String()
}

func parseRoom(s string) room {
	dashPos := strings.LastIndex(s, "-")
	bracketPos := strings.LastIndex(s, "[")
	name := s[:dashPos]
	sector := util.MustAtoi(s[dashPos+1 : bracketPos])
	checksum := s[bracketPos+1 : len(s)-1]

	letters := map[rune]int{}
	for _, r := range name {
		if r != '-' {
			letters[r]++
		}
	}

	return room{name, letters, sector, checksum}
}

func main() {
	input := "2016/Day4/input"
	lines := util.GetFileStrings(input)

	sum := 0
	for _, l := range lines {
		if r := parseRoom(l); r.ChecksumMatches() {
			sum += r.sector
			if r.Decrypt() == "northpole-object-storage" {
				log.Printf("Part 2: Storage sector ID: %d", r.sector)
			}
		}
	}
	log.Printf("Part 1: Sum of valid sectors: %d", sum)
}
