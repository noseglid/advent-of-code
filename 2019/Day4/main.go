package main

import "log"

func parts(i int) []int {

	var r []int
	for i > 0 {
		r = append(r, i%10)
		i /= 10
	}

	rr := make([]int, len(r))

	for i := len(r) - 1; i >= 0; i-- {
		rr[i] = r[len(r)-i-1]
	}

	return rr
}

func isIncreasing(p []int) bool {
	for i := range p {
		if i == len(p)-1 {
			return true
		}
		if p[i] > p[i+1] {
			return false
		}
	}

	return false
}

func sameAdjacent(p []int) bool {
	for i := range p {
		if i == len(p)-1 {
			return false
		}

		if p[i] == p[i+1] {
			return true
		}
	}

	return false
}

func sameAdjacentOnce(p []int) bool {
	for i := range p {
		if i == len(p)-2 {
			return p[i] == p[i+1] && p[i] != p[i-1]
		}

		if i == 0 && p[i] == p[i+1] && p[i] != p[i+2] {
			return true
		}

		if p[i] == p[i+1] && p[i] != p[i+2] && p[i-1] != p[i] {
			return true
		}
	}

	return false
}

func is6digits(p []int) bool {
	return len(p) == 6
}

func main() {

	min, max := 172851, 675869
	n := 0
	n2 := 0
	for i := min; i <= max; i++ {
		p := parts(i)
		if isIncreasing(p) && sameAdjacent(p) && is6digits(p) {
			n++
		}
		if isIncreasing(p) && sameAdjacentOnce(p) && is6digits(p) {
			n2++
		}
	}
	log.Printf("matching passwords (part1): %d", n)
	log.Printf("matching passwords small groups (part2): %d", n2)
}
