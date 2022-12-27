package main

import (
	"github.com/noseglid/advent-of-code/util"
)

func parseInput(l string) []int {
	var r []int
	for _, n := range l {
		r = append(r, int(n))
	}
	return r

}

func copySublist(src []int, start, end int) []int {
	if end == start {
		return []int{src[end]}
	} else if end > start {
		ret := make([]int, end-start)
		copy(ret, src[start:end])
		return ret
	} else {
		ret := make([]int, len(src)-start+end)
		copy(ret, src[start:])
		copy(ret[len(src)-start:], src[:start])
		return ret
	}
}

func makeList(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	return list
}

func iterate(lengths []int, list []int, iterations int) {
	skipSize := 0
	position := 0
	for i := 0; i < iterations; i++ {
		for _, e := range lengths {
			start, end := position, (position+e)%len(list)
			sublist := copySublist(list, start, end)
			util.Reverse(sublist)

			j := 0
			for i := start; i != end; i = (i + 1) % len(list) {
				list[i] = sublist[j]
				j++
			}

			position = (position + e + skipSize) % len(list)
			skipSize++
		}
	}
}

func KnotHash(data string) []byte {
	lengths := append(parseInput(data), []int{17, 31, 73, 47, 23}...)
	list := makeList(256)
	iterate(lengths, list, 64)

	xord := make([]byte, 16)
	for i := 0; i < len(list); i += 16 {
		n := list[i]
		for j := 1; j < 16; j++ {
			n ^= list[i+j]
		}
		xord[i/16] = byte(n)

	}
	return xord

}
