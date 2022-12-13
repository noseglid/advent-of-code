package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Entry interface {
	ToList() Entry
	Equal(rhs Entry) bool
	fmt.Stringer
}

type listEntry struct {
	entries []Entry
}

func (e *listEntry) ToList() Entry { return e }
func (e *listEntry) Equal(rhs Entry) bool {
	rhsList, ok := rhs.(*listEntry)
	if !ok || len(rhsList.entries) != len(e.entries) {
		return false
	}
	for i, e := range e.entries {
		if !e.Equal(rhsList.entries[i]) {
			return false
		}
	}
	return true
}

func (e listEntry) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	sep := ""
	for _, t := range e.entries {
		sb.WriteString(sep)
		sb.WriteString(t.String())
		sep = ","
	}
	sb.WriteRune(']')

	return sb.String()
}

type scalarEntry struct {
	value int
}

func (e *scalarEntry) ToList() Entry {
	return &listEntry{
		entries: []Entry{&scalarEntry{value: e.value}},
	}
}

func (e *scalarEntry) Equal(rhs Entry) bool {
	rhsScalar, ok := rhs.(*scalarEntry)
	return ok && e.value == rhsScalar.value
}

func (e scalarEntry) String() string {
	return fmt.Sprintf("%d", e.value)
}

func findEnclose(s string, at int) (int, int) {
	d := 0
	for i := at; i < len(s); i++ {
		r := s[i]
		switch r {
		case ']':
			d--
			if d == 0 {
				return at, i + 1
			}
		case '[':
			d++
		}
	}

	panic(fmt.Sprintf("no enclose in '%s', beginning at %d", s, at))
}

func parseScalarEntry(s string) *scalarEntry {
	// fmt.Printf("parsing scalar from '%s'\n", s)
	return &scalarEntry{value: util.MustAtoi(s)}
}

func parseListEntry(s string) *listEntry {
	// fmt.Printf("parsing list entry from '%s'\n", s)
	entry := &listEntry{}
	for i := 1; i < len(s)-1; i++ {
		r := s[i]
		// fmt.Printf("examining rune at index %d: '%c'\n", i, r)
		switch r {
		case '[':
			b, e := findEnclose(s, i)
			// fmt.Printf("enclose: %d -> %d\n", b, e)
			entry.entries = append(entry.entries, parseListEntry(s[b:e]))
			i = e
		case ']':
			return entry
		default:
			idx := strings.IndexFunc(s[i:], func(r rune) bool { return r == ',' || r == ']' })
			if idx == -1 {
				idx = len(s) - i
			}

			entry.entries = append(entry.entries, parseScalarEntry(s[i:i+idx]))
			i += idx
		}
	}

	return entry
}

func isCorrectOrder(e1, e2 Entry) (bool, bool) {
	l1, okl1 := e1.(*listEntry)
	l2, okl2 := e2.(*listEntry)
	s1, oks1 := e1.(*scalarEntry)
	s2, oks2 := e2.(*scalarEntry)

	if oks1 && oks2 {
		// both are scalars
		if s1.value < s2.value {
			return true, true
		} else if s1.value > s2.value {
			return false, true
		}
		return false, false
	}

	if okl1 && okl2 {
		// both are lists
		for i := range l1.entries {
			if i == len(l2.entries) {
				// Right side ran out of entries
				return false, true
			}
			if correctOrder, determinate := isCorrectOrder(l1.entries[i], l2.entries[i]); determinate {
				return correctOrder, true
			}
		}

		if len(l1.entries) < len(l2.entries) {
			return true, true
		}

		return false, false
	}

	if okl1 && oks2 {
		// first is list, second is scalar
		return isCorrectOrder(l1, s2.ToList())
	}

	if oks1 && okl2 {
		// first is scalar, second is list
		return isCorrectOrder(s1.ToList(), l2)
	}

	panic("Not both list, or both scalar, or one list other scalar, or one scalar other list")
}

func lessFunc(e1, e2 Entry) (bool, bool) {
	l1, okl1 := e1.(*listEntry)
	l2, okl2 := e2.(*listEntry)
	s1, oks1 := e1.(*scalarEntry)
	s2, oks2 := e2.(*scalarEntry)

	// fmt.Printf("comparing %s and %s\n", e1, e2)
	if okl1 && okl2 {
		// both lists
		// fmt.Printf("checking lists:")
		for i := range l1.entries {
			if i == len(l2.entries) {
				// right ran out of entries
				// fmt.Printf("right ran out, less=true, determinate!\n")
				return false, true
			}
			// fmt.Printf("nesting, will now check %s and %s\n", l1.entries[i], l2.entries[i])
			less, determinate := lessFunc(l1.entries[i], l2.entries[i])
			if determinate {
				return less, true
			}
		}

		// same length list, and no differentiating terms
		if len(l1.entries) == len(l2.entries) {
			// fmt.Printf("same length list, no determning yet, indeterminate\n")
			return false, false
		}
		// fmt.Printf("differing lengths, returning %d < %d\n", len(l1.entries), len(l2.entries))
		return len(l1.entries) < len(l2.entries), true
	}

	if oks1 && oks2 {
		// both scalar
		// fmt.Printf("checking scalars: ")
		if s1.value == s2.value {
			// fmt.Printf("equal - non determinate\n")
			return false, false
		}
		// fmt.Printf("testing %d < %d, determinate\n", s1.value, s2.value)
		return s1.value < s2.value, true
	}

	if okl1 && oks2 {
		// list and scalar
		// fmt.Printf("checking list and scalars\n")
		return lessFunc(l1, s2.ToList())
	}

	if oks1 && okl2 {
		// scalar and list
		// fmt.Printf("checking scalars and list\n")
		return lessFunc(s1.ToList(), l2)
	}

	panic("bad in less func")
}

func sortEntries(allEntries []Entry) func(i, j int) bool {
	return func(i, j int) bool {
		less, determinate := lessFunc(allEntries[i], allEntries[j])
		// fmt.Printf("%s < %s: %t\n", allEntries[i], allEntries[j], less)
		if !determinate {
			panic("not determinate in sortEntries")
		}
		return less
	}
}

func main() {
	input := util.GetFileStrings("2022/Day13/input")

	s := 0
	var allEntries []Entry
	for i := 0; i < len(input); i += 3 {
		l1, l2 := parseListEntry(input[i]), parseListEntry(input[i+1])
		correctOrder, _ := isCorrectOrder(l1, l2)
		if correctOrder {
			s += i/3 + 1
		}

		allEntries = append(allEntries, l1, l2)
	}
	fmt.Printf("Sum of indices in right order (part1); %d\n", s)

	allEntries = append(allEntries, parseListEntry("[[2]]"), parseListEntry("[[6]]"))
	// allEntries = []Entry{parseListEntry("[9]"), parseListEntry("[[[]]]")}
	sort.Slice(allEntries, sortEntries(allEntries))

	var indexes []int
	for i, e := range allEntries {
		if e.Equal(parseListEntry("[[2]]")) {
			indexes = append(indexes, i)
		}
		if e.Equal(parseListEntry("[[6]]")) {
			indexes = append(indexes, i)
		}
	}

	fmt.Printf("Product of position for sentinels (part2): %d*%d = %d\n", indexes[0]+1, indexes[1]+1, (indexes[0]+1)*(indexes[1]+1))

}
