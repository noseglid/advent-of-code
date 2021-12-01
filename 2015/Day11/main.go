package main

import (
	"log"
	"strings"
)

func isCorrectLength(s string) bool {
	return len(s) == 8
}

func hasIncreasingStraight(s string) bool {
	for i := 0; i < len(s); i++ {
		if i > len(s)-3 {
			return false
		}

		d1 := int(s[i+2]) - int(s[i+1])
		d2 := int(s[i+1]) - int(s[i+0])
		if d1 == 1 && d2 == 1 {
			return true
		}

	}

	return false
}

func hasForbidden(s string) bool {
	return strings.ContainsAny(s, "iol")
}

func hasRepeatedChars(s string) bool {
	nrep := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			nrep++
			i++
		}
	}

	return nrep >= 2
}

func passwordIncrement(s string) string {
	out := []rune(s)
	index := len(s) - 1
	for {
		r := rune(s[index])
		if r == 'z' {
			out[index] = 'a'
			index--
		} else {
			out[index] = r + 1
			return string(out)
		}
	}
}

func matchesPasswordPolicy(s string) bool {
	return isCorrectLength(s) && hasIncreasingStraight(s) && !hasForbidden(s) && hasRepeatedChars(s)
}

func main() {
	input := "hepxcrrq"

	for {
		if matchesPasswordPolicy(input) {
			break
		}
		input = passwordIncrement(input)
	}

	log.Printf("next matching password (part1): %v", input)

	input = passwordIncrement(input)
	for {
		if matchesPasswordPolicy(input) {
			break
		}
		input = passwordIncrement(input)
	}
	log.Printf("next matching password (part2): %v", input)

}
