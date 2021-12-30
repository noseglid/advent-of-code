package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type loc struct {
	x, y  int
	trail string
}

func (l loc) AtVault() bool {
	return l.x == 3 && l.y == 3
}

func hash(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	h := hex.EncodeToString(md5.Sum(nil))
	return h
}

var openChars = "bcdef"

func doorsOpen(s string) (bool, bool, bool, bool) {
	return strings.Contains(openChars, string(s[0])),
		strings.Contains(openChars, string(s[1])),
		strings.Contains(openChars, string(s[2])),
		strings.Contains(openChars, string(s[3]))
}

func moves(l loc, pass string) []loc {
	up, down, left, right := doorsOpen(hash(pass + l.trail))
	var locs []loc
	if up && l.y > 0 {
		locs = append(locs, loc{l.x, l.y - 1, l.trail + "U"})
	}
	if down && l.y < 3 {
		locs = append(locs, loc{l.x, l.y + 1, l.trail + "D"})
	}
	if left && l.x > 0 {
		locs = append(locs, loc{l.x - 1, l.y, l.trail + "L"})
	}
	if right && l.x < 3 {
		locs = append(locs, loc{l.x + 1, l.y, l.trail + "R"})
	}

	return locs
}

func shortestPath(pass string) string {
	p := loc{0, 0, ""}
	current := []loc{p}
	for {
		var next []loc
		for _, c := range current {
			for _, p := range moves(c, pass) {
				if p.AtVault() {
					return p.trail
				}

				next = append(next, p)
			}
		}
		current = next
	}
}

func longestPath(pass string) int {
	p := loc{0, 0, ""}
	current := []loc{p}
	longest := 0
	for steps := 0; len(current) > 0; steps++ {
		var next []loc
		for _, c := range current {
			p := moves(c, pass)
			for _, pp := range p {
				if pp.AtVault() {
					longest = util.MaxInt(longest, steps+1)
				} else {
					next = append(next, pp)
				}
			}
		}
		current = next
	}

	return longest
}

func main() {
	pass := "qtetzkpl"
	log.Printf("Part 1: Shortest path through: %s", shortestPath(pass))
	log.Printf("Part 2: Longest path (steps) through: %d", longestPath(pass))

}
