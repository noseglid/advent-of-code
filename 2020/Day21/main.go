package main

import (
	"bufio"
	"log"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type meal struct {
	ingredients []string
	allergens   []string
}

func parseMeal(s string) meal {
	var i, a []string
	ingr := true
	for _, e := range strings.Fields(s) {
		if e == "(contains" {
			ingr = false
			continue
		}

		if ingr {
			i = append(i, strings.TrimRight(e, "),"))
		} else {
			a = append(a, strings.TrimRight(e, "),"))
		}

	}

	return meal{i, a}
}

func unique(l []string) []string {
	m := map[string]struct{}{}

	for _, e := range l {
		m[e] = struct{}{}
	}

	var r []string
	for e := range m {
		r = append(r, e)
	}

	return r
}

func contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}

	}

	return false
}

func without(l []string, s string) []string {
	r := make([]string, 0, len(l)-1)
	for i, e := range l {
		if e == s {
			r = append(r, l[:i]...)
			r = append(r, l[i+1:]...)
			return r
		}
	}

	return l
}

func intersect(entries [][]string) []string {
	list := make([]string, len(entries[0]))
	copy(list, entries[0])

	for _, e := range entries[1:] {
		for i := 0; i < len(list); i++ {
			if !contains(e, list[i]) {
				list = append(list[:i], list[i+1:]...)
				i--
			}
		}
	}

	return list
}

func removeFound(ingr string, potentials map[string][]string) map[string][]string {
	for a, i := range potentials {
		if contains(i, ingr) {
			potentials[a] = without(i, ingr)
		}
	}

	return potentials

}

func reduceOnce(potentials map[string][]string) (map[string][]string, string, string, bool) {
	for a, i := range potentials {
		if len(i) == 1 {
			ret := removeFound(i[0], potentials)
			return ret, a, i[0], true
		}
	}

	return potentials, "", "", false
}

func occurance(ingr string, meals []meal) int {
	n := 0
	for _, m := range meals {
		for _, i := range m.ingredients {
			if i == ingr {
				n++
			}
		}
	}

	return n
}

func canonicalized(certainByIngr map[string]string) string {
	certanByAllergen := map[string]string{}

	var allergens []string
	for i, a := range certainByIngr {
		allergens = append(allergens, a)
		certanByAllergen[a] = i
	}

	sort.Slice(allergens, func(i, j int) bool {
		return allergens[i] < allergens[j]
	})

	var ingredients []string
	for _, a := range allergens {
		ingredients = append(ingredients, certanByAllergen[a])
	}

	return strings.Join(ingredients, ",")
}

func main() {
	s := util.FileScanner("2020/Day21/input", bufio.ScanLines)

	var allergens []string
	var ingredients []string
	var meals []meal
	for s.Scan() {
		m := parseMeal(s.Text())
		meals = append(meals, m)
		allergens = unique(append(allergens, m.allergens...))
		ingredients = unique(append(ingredients, m.ingredients...))
	}

	potentials := map[string][]string{}

	for _, a := range allergens {
		var hasPotentialAllergen [][]string
		for _, m := range meals {
			if contains(m.allergens, a) {
				hasPotentialAllergen = append(hasPotentialAllergen, m.ingredients)
			}
		}
		potentials[a] = intersect(hasPotentialAllergen)
	}

	certainAllergens := map[string]string{}

	for {
		var al, in string
		var ok bool
		potentials, al, in, ok = reduceOnce(potentials)
		if !ok {
			break
		}

		certainAllergens[in] = al
	}

	sum := 0
	for _, ingr := range ingredients {
		if _, ok := certainAllergens[ingr]; ok {
			continue
		}

		sum += occurance(ingr, meals)
	}

	log.Printf("occurance of non-allergens (part1): %d", sum)
	log.Printf("canonicalized (part2): %s", canonicalized(certainAllergens))
}
