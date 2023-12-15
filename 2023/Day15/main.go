package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func hash(l string) int {
	current := 0
	for _, c := range l {
		current += int(c)
		current *= 17
		current %= 256
	}

	return current
}

func remove(lenses []lens, label string) []lens {
	var res []lens
	for _, l := range lenses {
		if l.label == label {
			continue
		}
		res = append(res, l)
	}
	return res
}

func add(lenses []lens, label string, focal int) []lens {
	var res []lens
	didOverwrite := false
	for _, l := range lenses {
		if l.label == label {
			res = append(res, lens{label: label, focal: focal})
			didOverwrite = true
		} else {
			res = append(res, l)
		}
	}

	if !didOverwrite {
		res = append(res, lens{label: label, focal: focal})
	}
	return res
}

var re = regexp.MustCompile(`([a-z]+)(=|-)(\d+)?`)

type lens struct {
	label string
	focal int
}

func main() {
	input := util.GetFile("2023/Day15/input")
	s := 0
	boxes := map[int][]lens{}
	for _, e := range strings.Split(strings.TrimSpace(input), ",") {
		s += hash(e)

		m := re.FindStringSubmatch(e)
		label, op := m[1], m[2]
		b := hash(label)
		switch op {
		case "=":
			boxes[b] = add(boxes[b], label, util.MustAtoi(m[3]))
		case "-":
			boxes[b] = remove(boxes[b], label)
		default:
			panic("bad op: " + op)
		}
	}

	s2 := 0
	for boxIndex, lenses := range boxes {
		for lensIndex, lens := range lenses {
			v := (boxIndex + 1) * (lensIndex + 1) * lens.focal
			s2 += v
		}
	}

	fmt.Printf("Hash verification (part1): %d\n", s)
	fmt.Printf("Focusing power (part2): %d\n", s2)

}
