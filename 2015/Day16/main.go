package main

import (
	"bufio"
	"log"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type sue struct {
	index       int
	children    *int
	cats        *int
	samoyeds    *int
	pomeranians *int
	akitas      *int
	vizslas     *int
	goldfish    *int
	trees       *int
	cars        *int
	perfumes    *int
}

var mainRe = regexp.MustCompile(`Sue (\d+): (.+)`)
var attrRe = regexp.MustCompile(`([[:alpha:]]+): (\d+)`)

func parseSue(s string) sue {
	m := mainRe.FindStringSubmatch(s)
	if len(m) != 3 {
	}

	sue := sue{
		index: util.MustAtoi(m[1]),
	}

	ma := attrRe.FindAllStringSubmatch(m[2], -1)
	for _, attrMatch := range ma {
		attribute, n := attrMatch[1], util.MustAtoi(attrMatch[2])
		switch attribute {
		case "children":
			sue.children = &n
		case "cats":
			sue.cats = &n
		case "samoyeds":
			sue.samoyeds = &n
		case "pomeranians":
			sue.pomeranians = &n
		case "akitas":
			sue.akitas = &n
		case "vizslas":
			sue.vizslas = &n
		case "goldfish":
			sue.goldfish = &n
		case "trees":
			sue.trees = &n
		case "cars":
			sue.cars = &n
		case "perfumes":
			sue.perfumes = &n
		default:
			panic("invalid attribute")
		}

	}
	return sue

}

func (targetSue sue) childrenMatch(rhs sue) bool {
	return rhs.children == nil || *rhs.children == *targetSue.children
}

func (targetSue sue) catsMatch(rhs sue) bool {
	return rhs.cats == nil || *rhs.cats > *targetSue.cats
}

func (targetSue sue) samoyedsMatch(rhs sue) bool {
	return rhs.samoyeds == nil || *rhs.samoyeds == *targetSue.samoyeds
}

func (targetSue sue) pomeraniansMatch(rhs sue) bool {
	return rhs.pomeranians == nil || *rhs.pomeranians < *targetSue.pomeranians
}

func (targetSue sue) akitasMatch(rhs sue) bool {
	return rhs.akitas == nil || *rhs.akitas == *targetSue.akitas
}

func (targetSue sue) vizslasMatch(rhs sue) bool {
	return rhs.vizslas == nil || *rhs.vizslas == *targetSue.vizslas
}

func (targetSue sue) goldfishMatch(rhs sue) bool {
	return rhs.goldfish == nil || *rhs.goldfish < *targetSue.goldfish
}

func (targetSue sue) treesMatch(rhs sue) bool {
	return rhs.trees == nil || *rhs.trees > *targetSue.trees
}

func (targetSue sue) carsMatch(rhs sue) bool {
	return rhs.cars == nil || *rhs.cars == *targetSue.cars
}

func (targetSue sue) perfumesMatch(rhs sue) bool {
	return rhs.perfumes == nil || *rhs.perfumes == *targetSue.perfumes
}

func intPtr(i int) *int {
	return &i
}

func main() {
	s := util.FileScanner("2015/Day16/input", bufio.ScanLines)

	targetSue := sue{
		children:    intPtr(3),
		cats:        intPtr(7),
		samoyeds:    intPtr(2),
		pomeranians: intPtr(3),
		akitas:      intPtr(0),
		vizslas:     intPtr(0),
		goldfish:    intPtr(5),
		trees:       intPtr(3),
		cars:        intPtr(2),
		perfumes:    intPtr(1),
	}

	for s.Scan() {
		sue := parseSue(s.Text())
		if targetSue.childrenMatch(sue) &&
			targetSue.catsMatch(sue) &&
			targetSue.samoyedsMatch(sue) &&
			targetSue.pomeraniansMatch(sue) &&
			targetSue.akitasMatch(sue) &&
			targetSue.vizslasMatch(sue) &&
			targetSue.goldfishMatch(sue) &&
			targetSue.treesMatch(sue) &&
			targetSue.carsMatch(sue) &&
			targetSue.perfumesMatch(sue) {
			log.Printf("Gift origin sue index (part2): %d", sue.index)
		}
	}
}
