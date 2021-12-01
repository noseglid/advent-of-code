package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type relation struct {
	affected  string
	relative  string
	happiness int
}

var re = regexp.MustCompile(`(?P<recv>[[:alpha:]]+) would (?P<mult>gain|lose) (?P<units>\d+) happiness units by sitting next to (?P<relative>[[:alpha:]]+).`)

func parseRelation(s string) relation {
	m := re.FindStringSubmatch(s)
	t := m[re.SubexpIndex("mult")]
	mult := 1
	if t == "lose" {
		mult = -1
	}

	return relation{
		affected:  m[re.SubexpIndex("recv")],
		relative:  m[re.SubexpIndex("relative")],
		happiness: util.MustAtoi(m[re.SubexpIndex("units")]) * mult,
	}
}

type Permutation []string

func genPermutations(attendees []string) []Permutation {
	if len(attendees) <= 1 {
		return []Permutation{[]string{attendees[0]}}
	}

	result := make([]Permutation, 0, util.Factorial(len(attendees)))
	for i, l := range attendees {
		subLocations := make([]string, len(attendees))
		copy(subLocations, attendees)
		permutations := genPermutations(append(subLocations[:i], attendees[i+1:]...))
		for _, p := range permutations {
			result = append(result, append([]string{l}, p...))
		}
	}

	return result
}

func relationHappiness(relations []relation, affected, relative string) int {
	for _, r := range relations {
		if r.affected == affected && r.relative == relative {
			return r.happiness
		}
	}

	panic(fmt.Sprintf("no relation between %s and %s", affected, relative))
}

func permutationHappniess(relations []relation, p Permutation) int {
	total := 0
	for i := range p {
		ii := (i + 1) % len(p)
		s1 := relationHappiness(relations, p[i], p[ii])
		s2 := relationHappiness(relations, p[ii], p[i])
		total += s1 + s2
	}

	return total

}

func maxHappiness(relations []relation, ps []Permutation) int {
	maxHappiness := permutationHappniess(relations, ps[0])
	for _, p := range ps[1:] {
		happiness := permutationHappniess(relations, p)
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}
	return maxHappiness
}

func main() {
	s := util.FileScanner("2015/Day13/input", bufio.ScanLines)

	var relations []relation

	for s.Scan() {
		relations = append(relations, parseRelation(s.Text()))
	}

	matt := map[string]struct{}{}
	for _, r := range relations {
		matt[r.affected] = struct{}{}
		matt[r.relative] = struct{}{}
	}

	attendees := make([]string, 0, len(matt)+1)
	for a := range matt {
		attendees = append(attendees, a)
	}

	part1 := maxHappiness(relations, genPermutations(attendees))
	log.Printf("max happiness (part1): %d", part1)

	// ******************************** //

	for _, a := range attendees {
		relations = append(relations, relation{
			affected:  "me",
			relative:  a,
			happiness: 0,
		})
		relations = append(relations, relation{
			affected:  a,
			relative:  "me",
			happiness: 0,
		})
	}

	attendees = append(attendees, "me")

	part2 := maxHappiness(relations, genPermutations(attendees))
	log.Printf("max happiness (part2): %d", part2)

}
