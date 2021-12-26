package main

import (
	"log"
	"strings"
	"unicode"

	"github.com/noseglid/advent-of-code/util"
)

func interpretMarker(s string, ps int) (string, int, int) {
	chars, times := "", 0
	pe := ps + strings.Index(s[ps:], ")")
	parts := strings.Split(s[ps:pe], "x")
	chars = s[pe+1 : pe+1+util.MustAtoi(parts[0])]
	times = util.MustAtoi(parts[1])

	return chars, times, pe + len(chars)
}

func decompress(s string) (string, int) {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			continue
		}
		if s[i] == '(' {
			chars, times, ni := interpretMarker(s, i+1)
			i = ni
			for k := 0; k < times; k++ {
				sb.WriteString(chars)
			}
		} else {
			sb.WriteByte(s[i])
		}
	}

	return sb.String(), len(sb.String())
}

type block struct {
	times, chars int
	data         string
}

func hasMarker(s string) bool {
	return strings.Contains(s, "(")
}

func parseInput(s string) (block, string) {
	if unicode.IsSpace(rune(s[0])) {
		return block{}, s[1:]
	}
	ps := strings.IndexRune(s, '(')
	if ps > 0 {
		return block{
			times: 1,
			chars: ps,
			data:  s[:ps],
		}, s[ps:]
	}
	if ps < 0 {
		return block{
			times: 1,
			chars: len(s),
			data:  s,
		}, ""
	}
	chars, times, pe := interpretMarker(s, ps+1)
	return block{
		times: times,
		chars: len(chars),
		data:  s[pe-len(chars)+1 : pe+1],
	}, s[pe+1:]
}

func (i block) len() int {
	l := i.chars
	if hasMarker(i.data) {
		l = iterateInput(i.data)
	}

	return i.times * l
}

func iterateInput(s string) int {
	var blocks []block
	for {
		block, remaining := parseInput(s)
		blocks = append(blocks, block)
		if len(remaining) == 0 {
			break
		}
		s = remaining
	}

	n := 0
	for _, b := range blocks {
		n += b.len()
	}

	return n
}

func main() {
	input := "2016/Day9/input"
	data := util.GetFile(input)
	_, n := decompress(data)
	log.Printf("Part 1: length of decompressed: %d", n)

	// log.Printf("it1=%d", iterateInput("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"))
	log.Printf("Part 2: length of iterated decompressed: %d", iterateInput(data))

}
