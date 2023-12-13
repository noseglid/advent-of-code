package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func colsMatch(m []string, c1, c2 int) bool {
	for i := 0; i < len(m); i++ {
		if m[i][c1] != m[i][c2] {
			return false
		}
	}
	return true
}

func reflectVertical(m []string, skip int) (int, bool) {
	width := len(m[0])
Outer:
	for i := 0; i < width-1; i++ {
		ti := i
		for j := i + 1; ; j++ {
			if !colsMatch(m, ti, j) {
				continue Outer
			}
			ti--
			if ti < 0 || j >= width-1 {
				if skip == i {
					continue Outer
				}
				return i, true
			}
		}
	}
	return 0, false
}

func reflectHorizontal(m []string, skip int) (int, bool) {
	height := len(m)
Outer:
	for i := 0; i < height-1; i++ {
		ti := i
		for j := i + 1; ; j++ {
			if m[ti] != m[j] {
				continue Outer
			}
			ti--
			if ti < 0 || j >= height-1 {
				if skip == i {
					continue Outer
				}
				return i, true
			}
		}
	}

	return 0, false
}

func swap(i, j int, m []string) {
	replace := "."
	if m[i][j] == '.' {
		replace = "#"
	}
	m[i] = m[i][:j] + replace + m[i][j+1:]
}

func main() {
	rows := util.GetFileStrings("2023/Day13/input")

	var maps [][]string

	var cmap []string
	for _, r := range rows {
		if r == "" {
			maps = append(maps, cmap)
			cmap = []string{}
			continue
		}
		cmap = append(cmap, r)
	}
	maps = append(maps, cmap)

	type ind struct {
		d rune
		l int
	}

	valids := make([]ind, len(maps))
	s := 0
	for mi, m := range maps {
		if v, ok := reflectVertical(m, -1); ok {
			valids[mi] = ind{d: 'v', l: v}
			s += v + 1
		} else if v, ok := reflectHorizontal(m, -1); ok {
			valids[mi] = ind{d: 'h', l: v}
			s += (v + 1) * 100
		} else {
			panic("no match!")
		}
	}

	s2 := 0
Map:
	for mi, m := range maps {
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				swap(i, j, m)
				skipv := -1
				if valids[mi].d == 'v' {
					skipv = valids[mi].l
				}
				skiph := -1
				if valids[mi].d == 'h' {
					skiph = valids[mi].l
				}
				if v, ok := reflectVertical(m, skipv); ok {
					s2 += v + 1
					swap(i, j, m)
					continue Map
				} else if v, ok := reflectHorizontal(m, skiph); ok {
					s2 += (v + 1) * 100
					swap(i, j, m)
					continue Map
				}
				swap(i, j, m)
			}
		}
		panic("all failed!")
	}

	fmt.Printf("Summarized reflections (part1): %d\n", s)
	fmt.Printf("Summarized reflections corrected (part2): %d\n", s2)
}
