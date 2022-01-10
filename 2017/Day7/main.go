package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type program struct {
	name        string
	weight      int
	supporting  []*program
	supportedBy *program
}

func (p program) Weight() int {
	sum := p.weight
	for _, s := range p.supporting {
		sum += s.Weight()
	}
	return sum
}

func main() {
	input := "2017/Day7/input"
	lines := util.GetFileStrings(input)

	progMap := map[string]*program{}
	var firstName string

	for _, def := range lines {
		fields := strings.Fields(def)
		name := fields[0]
		if firstName == "" {
			firstName = fields[0]
		}
		progMap[name] = &program{name, util.MustAtoi(fields[1][1 : len(fields[1])-1]), nil, nil}
	}

	for _, def := range lines {
		fields := strings.Fields(def)
		if len(fields) < 4 {
			continue
		}
		rootName := fields[0]
		for _, f := range fields[3:] {
			name := strings.TrimSuffix(f, ",")
			progMap[rootName].supporting = append(progMap[rootName].supporting, progMap[name])
			progMap[name].supportedBy = progMap[rootName]
		}
	}

	root := progMap[firstName]
	for {
		if root.supportedBy == nil {
			break
		}
		root = root.supportedBy
	}

	log.Printf("Part 1: Bottom program: %s", root.name)

	currProg := root
	type ee struct {
		c int
		n []string
	}
	var prevCommonMap, commonMap map[int]ee
	for {
		commonMap = map[int]ee{}
		for _, s := range currProg.supporting {
			w := s.Weight()
			commonMap[w] = ee{commonMap[w].c + 1, append(commonMap[w].n, s.name)}
		}
		foundNew := false
		for _, m := range commonMap {
			if m.c != 1 {
				continue
			}

			currProg = progMap[m.n[0]]
			foundNew = true
		}
		if !foundNew {
			break
		}
		prevCommonMap = commonMap
	}

	var misalignedWeight, otherWeight int
	for n, e := range prevCommonMap {
		if e.n[0] == currProg.name {
			misalignedWeight = n
		} else {
			otherWeight = n
		}
	}

	delta := otherWeight - misalignedWeight

	log.Printf("Part 2: Misaligned node (%s) should weigh: %d", currProg.name, progMap[currProg.name].weight+delta)

}
