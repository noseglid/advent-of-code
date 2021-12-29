package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type floor struct {
	generators []string
	chips      []string
}

func (f floor) clone() floor {
	ff := floor{
		generators: make([]string, len(f.generators)),
		chips:      make([]string, len(f.chips)),
	}
	copy(ff.generators, f.generators)
	copy(ff.chips, f.chips)
	return ff
}

func (f floor) empty() bool {
	return len(f.generators) == 0 && len(f.chips) == 0
}

func (f floor) String() string {
	var sb strings.Builder
	for _, g := range f.generators {
		sb.WriteString(g)
		sb.WriteRune('G')
	}

	for _, c := range f.chips {
		sb.WriteString(c)
		sb.WriteRune('M')
	}
	return sb.String()
}

func (f floor) valid() bool {
	chipsWithoutGenerator := []string{}
	for _, c := range f.chips {
		if _, ok := contains(f.generators, c); !ok {
			chipsWithoutGenerator = append(chipsWithoutGenerator, c)
		}
	}
	if len(chipsWithoutGenerator) == 0 {
		return true
	}

	return len(f.generators) == 0
}

type facility struct {
	elevator int
	floors   [4]floor
	steps    int

	allTypes []string
}

func (f facility) clone() facility {
	return facility{
		elevator: f.elevator,
		steps:    f.steps,
		allTypes: f.allTypes,
		floors: [4]floor{
			f.floors[0].clone(),
			f.floors[1].clone(),
			f.floors[2].clone(),
			f.floors[3].clone(),
		},
	}

}

func (f facility) valid() bool {
	for _, fl := range f.floors {
		if !fl.valid() {
			return false
		}
	}

	return true
}

func (f facility) done() bool {
	return f.floors[0].empty() && f.floors[1].empty() && f.floors[2].empty()
}

func (f *facility) moveGenerator(g string, delta int) {
	idx, ok := contains(f.floors[f.elevator].generators, g)
	if !ok {
		panic("generator not at floor")
	}
	f.floors[f.elevator].generators = append(f.floors[f.elevator].generators[:idx], f.floors[f.elevator].generators[idx+1:]...)
	f.floors[f.elevator+delta].generators = append(f.floors[f.elevator+delta].generators, g)
	sort.Strings(f.floors[f.elevator+delta].generators)
}

func (f *facility) moveChip(c string, delta int) {
	idx, ok := contains(f.floors[f.elevator].chips, c)
	if !ok {
		panic("chip not at floor")
	}
	f.floors[f.elevator].chips = append(f.floors[f.elevator].chips[:idx], f.floors[f.elevator].chips[idx+1:]...)
	f.floors[f.elevator+delta].chips = append(f.floors[f.elevator+delta].chips, c)
	sort.Strings(f.floors[f.elevator+delta].chips)
}

func (f facility) key(repl map[string]int) string {
	type pair struct {
		c, g int
	}

	var pairs []pair

	for idx, fl := range f.floors {
		for _, g := range fl.generators {
			p := pair{g: idx + 1}

		Floor:
			for idx2, fl2 := range f.floors {
				for _, c := range fl2.chips {
					if c == g {
						p.c = idx2 + 1
						pairs = append(pairs, p)
						break Floor
					}
				}
			}
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].c < pairs[j].c {
			return true
		}
		if pairs[i].c > pairs[j].c {
			return false
		}
		return pairs[i].g < pairs[j].g
	})

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(f.elevator + 1))
	for _, pp := range pairs {
		sb.WriteString(strconv.Itoa(pp.c))
		sb.WriteString(strconv.Itoa(pp.g))
	}

	return sb.String()
}

func (f *facility) init() {
	f.allTypes = f.calcAllTypes()
	for i := range f.floors {
		sort.Strings(f.floors[i].generators)
		sort.Strings(f.floors[i].chips)
	}
}

func contains(list []string, search string) (int, bool) {
	for i, l := range list {
		if l == search {
			return i, true
		}
	}
	return -1, false
}

type tuple struct {
	v1, v2 string
}

func permutations1(list []string) []tuple {
	var stuples []tuple
	for i, v1 := range list {
		stuples = append(stuples, tuple{v1, ""})
		for j, v2 := range list {
			if j <= i {
				continue
			}
			stuples = append(stuples, tuple{v1, v2})
		}
	}

	return stuples
}

func permutations2(l1, l2 []string) []tuple {
	var stuples []tuple
	for _, v1 := range l1 {
		for _, v2 := range l2 {
			stuples = append(stuples, tuple{v1, v2})
		}
	}
	return stuples
}

func nextStates(f facility, delta int) []facility {
	var facilities1 []facility
	var facilities2 []facility

	for _, tt := range permutations2(f.floors[f.elevator].generators, f.floors[f.elevator].chips) {
		fn := f.clone()
		fn.moveGenerator(tt.v1, delta)
		fn.moveChip(tt.v2, delta)
		if fn.valid() {
			fn.steps += 1
			fn.elevator += delta
			facilities2 = append(facilities2, fn)
		}
	}
	for _, tt := range permutations1(f.floors[f.elevator].generators) {
		fn := f.clone()
		fn.moveGenerator(tt.v1, delta)
		l := &facilities1
		if tt.v2 != "" {
			fn.moveGenerator(tt.v2, delta)
			l = &facilities2
		}
		if fn.valid() {
			fn.steps += 1
			fn.elevator += delta
			*l = append(*l, fn)
		}
	}
	for _, tt := range permutations1(f.floors[f.elevator].chips) {
		fn := f.clone()
		fn.moveChip(tt.v1, delta)
		l := &facilities1
		if tt.v2 != "" {
			fn.moveChip(tt.v2, delta)
			l = &facilities2
		}
		if fn.valid() {
			fn.steps += 1
			fn.elevator += delta
			*l = append(*l, fn)
		}
	}
	switch delta {
	case 1:
		if len(facilities2) > 0 {
			return facilities2
		} else {
			return facilities1
		}
	case -1:
		if len(facilities1) > 0 {
			return facilities1
		} else {
			return facilities2
		}
	}

	panic("no facilities")
}

func minSteps(f facility) (int, bool) {

	currentFacilities := []facility{f}
	seen := map[string]bool{f.key(nil): true}

	steps := 0
	for {
		steps++

		var nextFacilities []facility
		for _, f := range currentFacilities {
			var facilityOptions []facility
			switch f.elevator {
			case 0:
				facilityOptions = append(facilityOptions, nextStates(f, 1)...)
			case 1:
				facilityOptions = append(facilityOptions, nextStates(f, 1)...)
				facilityOptions = append(facilityOptions, nextStates(f, -1)...)
			case 2:
				facilityOptions = append(facilityOptions, nextStates(f, 1)...)
				facilityOptions = append(facilityOptions, nextStates(f, -1)...)
			case 3:
				facilityOptions = append(facilityOptions, nextStates(f, -1)...)
			}

			for _, nf := range facilityOptions {
				if nf.done() {
					return steps, true
				}

				h := nf.key(nil)
				if !seen[h] {
					nextFacilities = append(nextFacilities, nf)
					seen[h] = true
				}
			}
		}
		currentFacilities = nextFacilities
	}
}

func (f facility) calcAllTypes() []string {
	var res []string
	for _, fl := range f.floors {
		res = append(res, fl.generators...)
	}
	return res

}

func main() {

	// Input
	f1 := facility{
		floors: [4]floor{
			{generators: []string{"TH", "PL", "ST"}, chips: []string{"TH"}},
			{generators: []string{}, chips: []string{"PL", "ST"}},
			{generators: []string{"PR", "RU"}, chips: []string{"PR", "RU"}},
			{generators: []string{}, chips: []string{}},
		},
	}
	f1.init()

	// Input part 2
	f2 := facility{
		floors: [4]floor{
			{generators: []string{"TH", "PL", "ST", "EL", "DI"}, chips: []string{"TH", "EL", "DI"}},
			{generators: []string{}, chips: []string{"PL", "ST"}},
			{generators: []string{"PR", "RU"}, chips: []string{"PR", "RU"}},
			{generators: []string{}, chips: []string{}},
		},
	}
	f2.init()

	m1, _ := minSteps(f1)
	log.Printf("Part 1: minimum steps: %d", m1)

	m2, _ := minSteps(f2)
	log.Printf("Part 2: minimum steps with extra elements: %d", m2)
}
