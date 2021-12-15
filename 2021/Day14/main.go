package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func mostCommon(m map[string]int) int {
	mm := 0
	for _, v := range m {
		if v > mm {
			mm = v
		}
	}

	return mm
}

func leastCommon(m map[string]int) int {
	mm := math.MaxInt
	for _, v := range m {
		if v < mm {
			mm = v
		}
	}
	return mm
}

func countElements(template string) map[string]int {
	m := map[string]int{}
	for _, r := range template {
		m[string(r)]++
	}
	return m
}

func part1subtract(template string) int {
	elements := countElements(template)
	return mostCommon(elements) - leastCommon(elements)
}

func parseRules(l []string) map[string]rune {
	rules := map[string]rune{}
	for _, r := range l {
		var from string
		var to rune
		if _, err := fmt.Sscanf(r, "%s -> %c", &from, &to); err != nil {
			panic(err)
		}
		rules[from] = to
	}

	return rules
}

func step(str string, rules map[string]rune) string {
	insert := make([]rune, len(str)-1)
	for i := 0; i < len(str)-1; i++ {
		insert[i] = rules[str[i:i+2]]
	}

	var sb strings.Builder
	for {
		if len(insert) == 0 {
			sb.WriteRune(rune(str[0]))
			break
		}

		sb.WriteRune(rune(str[0]))
		sb.WriteRune(insert[0])
		str = str[1:]
		insert = insert[1:]
	}

	return sb.String()
}

func part2(base string, rules map[string]rune) {
	elementCounts := map[rune]int{}
	polymer := map[string]int{}
	for i, r := range base {
		elementCounts[r]++
		if i < len(base)-1 {
			polymer[base[i:i+2]]++
		}
	}

	for i := 0; i < 40; i++ {
		newPolymer := map[string]int{}
		for s, v := range polymer {
			elementCounts[rules[s]] += v
			newPolymer[fmt.Sprintf("%c%c", s[0], rules[s])] += v
			newPolymer[fmt.Sprintf("%c%c", rules[s], s[1])] += v
		}
		polymer = newPolymer
	}

	mostCommon, leastCommon := 0, math.MaxInt
	for _, c := range elementCounts {
		if c > mostCommon {
			mostCommon = c
		}
		if c < leastCommon {
			leastCommon = c
		}
	}

	log.Printf("Part 2: Most subtracted by least commmon 40 iterations: %d", mostCommon-leastCommon)
}

func main() {
	input := "2021/Day14/input"

	lines := util.GetFileStrings(input)

	template := lines[0]
	rules := parseRules(lines[2:])

	fmt.Printf("Template: %s\n", template)

	for i := 0; i < 10; i++ {
		template = step(template, rules)
	}

	log.Printf("Part 1: Most common, subtract least common 10 iterations: %d", part1subtract(template))

	part2(lines[0], rules)
}
