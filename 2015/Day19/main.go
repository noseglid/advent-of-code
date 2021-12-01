package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/noseglid/advent-of-code/util"
)

type replacement struct {
	from, to string
}

func parseReplacement(s string) replacement {
	spl := strings.Split(s, " => ")
	return replacement{spl[0], spl[1]}
}

func replaceOnce(molecule string, replacements []replacement) int {
	distinct := map[string]int{}

	for i := 0; i < len(molecule); i++ {
		startIndex := i
		endIndex := i + 1
		if i+1 < len(molecule) && unicode.IsLower(rune(molecule[i+1])) {
			// 2 letter
			endIndex = i + 2
			i++
		}

		single := molecule[startIndex:endIndex]
		for _, r := range replacements {
			if r.from == single {
				distinct[fmt.Sprintf("%s%s%s", molecule[:startIndex], r.to, molecule[endIndex:])]++
			}
		}
	}

	return len(distinct)
}

func main() {

	data := util.GetFile("2015/Day19/input")
	spl := strings.Split(data, "\n\n")
	replacementsSpec, medicineMolecule := spl[0], spl[1]
	var replacements []replacement
	for _, spec := range strings.Split(replacementsSpec, "\n") {
		replacements = append(replacements, parseReplacement(spec))
	}

	log.Printf("distincts with 1 replace (part1): %d", replaceOnce(medicineMolecule, replacements))

	index := 0
	depth := 0
	for {
		log.Printf("depth: %d", depth)
		if index > len(medicineMolecule)-2 {
			break
		}

		if medicineMolecule[index:index+2] == "Rn" {
			index += 2
			depth++
		} else if medicineMolecule[index:index+2] == "Ar" {
			index += 2
			depth--
		} else {
			index += 1
		}
	}

	log.Printf("depth: %d", depth)

}
