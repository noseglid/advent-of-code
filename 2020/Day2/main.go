package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type PasswordPolicy struct {
	min  int
	max  int
	char rune
}

var re = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): (.+)$`)

func ParsePasswordPolicyLine(line string) (PasswordPolicy, string) {
	matches := re.FindAllStringSubmatch(line, -1)
	min, err := strconv.Atoi(matches[0][1])
	if err != nil {
		log.Fatal(err)
	}

	max, err := strconv.Atoi(matches[0][2])
	if err != nil {
		log.Fatal(err)
	}

	return PasswordPolicy{min, max, []rune(matches[0][3])[0]}, matches[0][4]
}

func PolicyMatchesPart1(policy PasswordPolicy, pw string) bool {
	nchars := 0

	for _, r := range pw {
		if r == policy.char {
			nchars++
		}
	}
	return nchars >= policy.min && nchars <= policy.max
}

func PolicyMatchesPart2(policy PasswordPolicy, pw string) bool {
	r1, r2 := rune(pw[policy.min-1]), rune(pw[policy.max-1])
	return (r1 == policy.char || r2 == policy.char) && r1 != r2
}

func main() {
	f, err := os.Open("2020/Day2/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	validp1 := 0
	validp2 := 0
	for s.Scan() {
		policy, pw := ParsePasswordPolicyLine(s.Text())
		if PolicyMatchesPart1(policy, pw) {
			validp1++
		}
		if PolicyMatchesPart2(policy, pw) {
			validp2++
		}
	}

	log.Printf("valid passwords (part1): %d", validp1)
	log.Printf("valid passwords (part2): %d", validp2)
}
