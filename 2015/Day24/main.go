package main

import (
	"log"
	"math"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

func sum(sl []int) int {
	s := 0
	for _, ss := range sl {
		s += ss
	}

	return s
}

func pickToSum(picked []int, remaining []int, target int) [][]int {
	s := sum(picked)
	var result [][]int
	// log.Printf("testing with picked %v, remaining %v", picked, remaining)
	for i, r := range remaining {
		if s+r == target {
			rr := []int{}
			rr = append(rr, picked...)
			rr = append(rr, r)
			result = append(result, rr)

		} else {
			rem := []int{}
			rem = append(rem, remaining[:i]...)
			rem = append(rem, remaining[i+1:]...)
			rr := pickToSum(append(picked, r), rem, target)
			result = append(result, rr...)
		}
	}

	return result
}

func buildBag(packages []int, target int) ([]int, []int) {
	s := 0
	var used []int
	var unused []int
	i := 0
	for ; i < len(packages); i++ {
		if s+packages[i] <= target {
			log.Printf("using %d", packages[i])
			used = append(used, packages[i])
			s += packages[i]
		} else {
			unused = append(unused, packages[i])
		}

		if s == target {
			break
		}
	}

	if i < len(packages) {
		log.Printf("adding packages: %v", packages[i+1:])
		unused = append(unused, packages[i+1:]...)
	}
	return used, unused
}

func qe(b []int) int {
	p := 1
	for _, pp := range b {
		p = p * pp
	}
	return p
}

func main() {
	// r := pickToSum([]int{}, []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 20)
	// minLen := math.MaxInt64
	// maxQE := 0
	// for _, c := range r {
	// 	if len(c) < minLen {
	// 		minLen = len(c)
	// 		maxQE = qe(c)
	// 	} else if len(c) == minLen {
	// 		if qe(c) > maxQE {
	// 			maxQE = qe(c)
	// 		}
	// 	}
	// }
	// log.Printf("minLen = %d, maxQE: %d", minLen, maxQE)

	packages := util.GetFileNumbers("2015/Day24/input")
	log.Printf("packages sum: %d", sum(packages))
	perBag := sum(packages) / 3

	sort.Sort(sort.Reverse(sort.IntSlice(packages)))
	log.Printf("%v", packages)
	log.Printf("want %d per bag", perBag)

	r := pickToSum([]int{}, packages, perBag)
	minLen := math.MaxInt64
	maxQE := 0
	for _, c := range r {
		if len(c) < minLen {
			minLen = len(c)
			maxQE = qe(c)
		} else if len(c) == minLen {
			if qe(c) > maxQE {
				maxQE = qe(c)
			}
		}
	}
	// b1, unused := buildBag(packages, perBag)
	// log.Printf("unused for bag2: %v", unused)
	// b2, unused := buildBag(unused, perBag)
	// log.Printf("unused for bag3: %v", unused)
	// b3, unused := buildBag(unused, perBag)
	// if len(unused) != 0 {
	// 	log.Fatalf("bags left: %v", unused)
	// }

	// log.Printf("lens: %d, %d, %d", len(b1), len(b2), len(b3))
	// log.Printf("qe b1 (part1): %d", qe(b1))
	// log.Printf("qe b2 (part1): %d", qe(b2))
	// log.Printf("qe b3 (part1): %d", qe(b3))
}
