package main

import (
	"log"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type entry struct {
	input  []string
	output []string
}

func NewEntry(line string) entry {
	pp := strings.Split(line, "|")

	var e entry
	for _, i := range strings.Split(strings.TrimSpace(pp[0]), " ") {
		e.input = append(e.input, string(runeSort([]rune(i))))
	}

	for _, o := range strings.Split(strings.TrimSpace(pp[1]), " ") {
		e.output = append(e.output, string(runeSort([]rune(o))))
	}

	return e
}

func (e entry) Num1478() int {
	n := 0
	for _, o := range e.output {
		if len(o) == 2 || len(o) == 3 || len(o) == 4 || len(o) == 7 {
			n++
		}
	}

	return n
}

func (e entry) uniqueSignalWithLength(length int) string {
	found := false
	var signal string
	for _, i := range e.input {
		if len(i) == length {
			if found {
				panic("already found!")
			}
			signal = i
			found = true
		}
	}

	if !found {
		panic("no signal")
	}

	return signal
}

func (e entry) inputWithLength(length int) []string {
	var res []string
	for _, i := range e.input {
		if len(i) == length {
			res = append(res, i)
		}
	}

	return res
}

func (e entry) inputWithLengthButNot(length int, exclude []string) []string {
	haystack := e.inputWithLength(length)
	var res []string
Outer:
	for _, a := range haystack {
		for _, e := range exclude {
			if a == e {
				continue Outer
			}

		}
		res = append(res, a)
	}

	return res
}

func disjoint(l1, l2 string) string {

	m := map[rune]struct{}{}
	for _, r := range l1 {
		m[r] = struct{}{}
	}

	var result []rune
	for _, r := range l2 {
		if _, ok := m[r]; !ok {
			result = append(result, r)
		}
	}

	return string(result)
}

func runeSort(l []rune) []rune {
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})

	return l
}

func stringSum(s1, s2 string) string {
	for _, s := range s2 {
		if !strings.ContainsRune(s1, s) {
			s1 += string(s)
		}
	}

	return string(runeSort([]rune(s1)))
}

func stringWithOneMore(base string, test []string) string {
outer:
	for _, t := range test {
		f := false
		bi := 0
		ii := 0
		for {
			if bi >= len(base) {
				return t
			}

			if base[bi] == t[ii] {
				bi++
				ii++
				continue
			}

			if f {
				continue outer
			}
			f = true
			ii++
		}
	}

	panic("no string has one more")
}

func stringContainAll(l string, set string) bool {
	for _, s := range set {
		if !strings.ContainsRune(l, s) {
			return false
		}
	}
	return true
}

func stringsContainAll(l []string, signals string) string {
	c := []string{}
	for _, ll := range l {
		if stringContainAll(ll, signals) {
			c = append(c, ll)
		}
	}
	if len(c) != 1 {
		panic("non unique stringsContainAll")
	}

	return c[0]
}

func (e entry) OutputValue() int {
	num1 := e.uniqueSignalWithLength(2)
	num4 := e.uniqueSignalWithLength(4)
	num7 := e.uniqueSignalWithLength(3)
	num8 := e.uniqueSignalWithLength(7)
	num9 := stringWithOneMore(stringSum(num4, num7), e.inputWithLength(6))
	num3 := stringsContainAll(e.inputWithLength(5), num1)
	num0 := stringsContainAll(e.inputWithLengthButNot(6, []string{num9}), num1)
	num6 := e.inputWithLengthButNot(6, []string{num0, num9})[0]
	num2 := ""
	num5 := ""

	num25 := e.inputWithLengthButNot(5, []string{num3})
	if len(disjoint(num25[0], num6)) == 1 {
		num2 = num25[1]
		num5 = num25[0]
	} else {
		num2 = num25[0]
		num5 = num25[1]
	}

	str := ""
	for _, o := range e.output {
		switch o {
		case num1:
			str += "1"
		case num2:
			str += "2"
		case num3:
			str += "3"
		case num4:
			str += "4"
		case num5:
			str += "5"
		case num6:
			str += "6"
		case num7:
			str += "7"
		case num8:
			str += "8"
		case num9:
			str += "9"
		case num0:
			str += "0"
		}
	}

	return util.MustAtoi(str)
}

func main() {
	input := "2021/Day8/input"

	lines := util.GetFileStrings(input)

	var entries []entry
	for _, l := range lines {
		entries = append(entries, NewEntry(l))
	}

	n := 0
	on := 0
	for _, e := range entries {
		n += e.Num1478()
		on += e.OutputValue()
	}

	log.Printf("Part 1: number of 1, 4, 7, 8: %d", n)
	log.Printf("Part 2: sum of output values: %d", on)

}
