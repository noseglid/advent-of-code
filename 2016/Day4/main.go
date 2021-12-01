package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type room struct {
	letters  map[rune]int
	id       int
	checksum string
}

var re = regexp.MustCompile(`(\d+)\[([a-z]+)\]`)

func parseRoom(s string) room {
	letters := map[rune]int{}
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '-' })
	for _, p := range parts {
		for _, l := range p {
			letters[l]++
		}
	}

	m := re.FindStringSubmatch(parts[len(parts)-1])
	return room{letters, util.MustAtoi(m[1]), m[2]}
}

func isMostUsed(m map[rune]int, l rune) bool {
	max := 0
	r := rune(0)
	for ir, n := range m {
		if n > max {
			max = n
			r = ir
		}
	}

	return r == l
}

func checksumMatches(r room) bool {
	mm := make(map[rune]int, len(r.letters))
	for k, v := range r.letters {
		mm[k] = v
	}

	for _, r := range r.checksum {
		log.Printf("checking if %c is most used in %+v", r, mm)
		if !isMostUsed(mm, r) {
			return false
		}
		delete(mm, r)
	}

	return true
}

func main() {

	r := parseRoom("aaaaa-bbb-z-y-x-123[abxyz]")
	log.Printf("room: %+v", r)
	log.Printf("checksum matches: %v", checksumMatches(r))

}
