package main

import (
	"log"
)

type elf struct {
	idx  int
	next *elf
	prev *elf
}

func NewElf(idx int, next, prev *elf) *elf {
	return &elf{idx, next, prev}
}

func nextNonZero(list []int, start int) (int, bool) {
	for i := start + 1; ; i++ {
		in := i % len(list)
		if in == start {
			return 0, false
		}

		if list[in] > 0 {
			return in, true
		}
	}
}

func part1() {
	elves := make([]int, 3001330)
	for i := range elves {
		elves[i] = 1
	}
	for {
		for i, e := range elves {
			if e == 0 {
				continue
			}

			if idx, ok := nextNonZero(elves, i); ok {
				elves[i] += elves[idx]
				elves[idx] = 0
			} else {
				log.Printf("Part 1: Elf which gets all presents: %d", i+1) // +1 because 0-indexed
				return
			}
		}
	}
}

func part2() {
	size := 3001330
	r, l := make([]int, 0, size/2+size%2), make([]int, 0, size/2)
	for i := 1; i <= size/2; i++ {
		r = append(r, i)
		l = append(l, i+size/2+size%2)
	}
	if size%2 != 0 {
		r = append(r, size/2+1)
	}

	for {
		s1, s2 := len(r), len(l)
		if (s1+s2)%2 == 0 {
			l = l[1:]
		} else {
			r = r[:len(r)-1]
		}
		if len(r)+len(l) == 1 {
			break
		}
		b := l[0]
		l = append(l[1:], r[0])
		r = append(r[1:], b)
	}

	log.Printf("Part 2: Elf which gets all presents: %d", r[0])
}

func main() {
	part1()
	part2()

}
