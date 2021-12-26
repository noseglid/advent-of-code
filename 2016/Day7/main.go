package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func isABBA(s string) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func supportTLS(s string) bool {
	ok := false
	inBracket := false
	for i, r := range s {
		if i+3 == len(s) {
			break
		}
		switch r {
		case '[':
			inBracket = true
			continue
		case ']':
			inBracket = false
			continue
		}

		a := isABBA(s[i : i+4])
		if inBracket && a {
			return false
		}
		if !inBracket && a {
			ok = true
		}
	}

	return ok
}

type block struct {
	r1, r2 rune
}

func isABA(s string) (block, bool) {
	return block{rune(s[0]), rune(s[1])}, s[0] == s[2]
}

func allABAs(s string) []block {
	var res []block
	inBracket := false
	for i, r := range s {
		if i+2 == len(s) {
			break
		}
		switch r {
		case '[':
			inBracket = true
			continue
		case ']':
			inBracket = false
			continue
		}

		if inBracket {
			continue
		}

		if aba, ok := isABA(s[i : i+3]); ok {
			res = append(res, aba)
		}
	}
	return res
}

func allBABs(s string) []block {
	var res []block
	inBracket := false
	for i, r := range s {
		if i+2 == len(s) {
			break
		}
		switch r {
		case '[':
			inBracket = true
			continue
		case ']':
			inBracket = false
			continue
		}

		if !inBracket {
			continue
		}

		if aba, ok := isABA(s[i : i+3]); ok {
			res = append(res, aba)
		}
	}
	return res
}

func supportSSL(s string) bool {
	aa := allABAs(s)
	bb := allBABs(s)
	for _, a := range aa {
		for _, b := range bb {
			if a.r1 == b.r2 && a.r2 == b.r1 {
				return true
			}
		}
	}

	return false
}

func main() {
	input := "2016/Day7/input"
	lines := util.GetFileStrings(input)
	nTLS := 0
	nSSL := 0
	for _, l := range lines {
		if supportTLS(l) {
			nTLS++
		}
		if supportSSL(l) {
			nSSL++
		}
	}

	log.Printf("Part 1: Number of IPs supporting TLS: %d", nTLS)
	log.Printf("Part 2: Number of IPs supporting SSL: %d", nSSL)
}
