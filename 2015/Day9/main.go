package main

import (
	"bufio"
	"log"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

// var re = regexp.MustCompile(`(?P<from>[[:alpha:]]+) to (?P<to>[[:alpha:]]+) = (?P<distance>[[:digit]]+)`)
var re = regexp.MustCompile(`(?P<from>[[:alpha:]]+) to (?P<to>[[:alpha:]]+) = (?P<distance>[[:digit:]]+)`)

func parseDistance(s string) (string, string, int) {
	match := re.FindStringSubmatch(s)
	from := match[re.SubexpIndex("from")]
	to := match[re.SubexpIndex("to")]
	distance := util.MustAtoi(match[re.SubexpIndex("distance")])
	return from, to, distance
}

type route struct {
	from, to string
}

type Permutation []string

func genPermutations(locations []string) []Permutation {
	if len(locations) <= 1 {
		return []Permutation{[]string{locations[0]}}
	}

	result := make([]Permutation, 0, util.Factorial(len(locations)))
	for i, l := range locations {
		subLocations := make([]string, len(locations))
		copy(subLocations, locations)
		permutations := genPermutations(append(subLocations[:i], locations[i+1:]...))
		for _, p := range permutations {
			result = append(result, append([]string{l}, p...))
		}
	}

	return result
}

func calculateDistance(distances map[route]int, permutation Permutation) int {
	distance := 0
	for i := range permutation {
		if i+1 == len(permutation) {
			break
		}

		from := permutation[i]
		to := permutation[i+1]
		distance += distances[route{from, to}]
	}

	return distance
}

func main() {

	s := util.FileScanner("2015/Day9/input", bufio.ScanLines)

	distances := map[route]int{}
	locationSet := map[string]struct{}{}

	for s.Scan() {
		from, to, distance := parseDistance(s.Text())
		locationSet[from] = struct{}{}
		locationSet[to] = struct{}{}
		distances[route{from, to}] = distance
		distances[route{to, from}] = distance
	}

	var locations []string
	for l := range locationSet {
		locations = append(locations, l)
	}

	perms := genPermutations(locations)

	shortest := calculateDistance(distances, perms[0])
	longest := shortest

	for _, p := range perms[1:] {
		d := calculateDistance(distances, p)
		if d < shortest {
			shortest = d
		}
		if d > longest {
			longest = d
		}
	}

	log.Printf("shortest distance (part1): %v", shortest)
	log.Printf("longest distance (part1): %v", longest)
}
