package main

import (
	"log"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func isAnagram(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	return string(r1) == string(r2)
}

func hasDuplicates(ss []string, cmp func(string, string) bool) bool {
	for i, s1 := range ss {
		for j, s2 := range ss {
			if i == j {
				continue
			}

			if cmp(s1, s2) {
				return true
			}
		}
	}
	return false
}

func main() {
	input := "2017/Day4/input"
	lines := util.GetFileStrings(input)

	n := 0
	n2 := 0
	for _, l := range lines {
		ss := strings.Fields(l)
		if !hasDuplicates(ss, func(s1, s2 string) bool { return s1 == s2 }) {
			n++
		}
		if !hasDuplicates(ss, isAnagram) {
			n2++
		}
	}
	log.Printf("Part 1: Number of valid passphrases: %d", n)
	log.Printf("Part 2: Number of valid passphrases: %d", n2)

}
