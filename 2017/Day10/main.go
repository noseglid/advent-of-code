package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func parseLengths(l string) []int {
	var r []int
	for _, n := range strings.Split(l, ",") {
		r = append(r, util.MustAtoi(n))
	}

	return r
}

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

func main() {
	input := "206,63,255,131,65,80,238,157,254,24,133,2,16,0,1,3"
	size := 256

	lengths := parseLengths(input)
	list := makeList(size)

	iterate(lengths, list, 1)
	fmt.Printf("Multiply first two numbers (part1): %d\n", list[0]*list[1])

	lengthsp2 := append(parseInput(input), []int{17, 31, 73, 47, 23}...)
	listp2 := makeList(256)
	iterate(lengthsp2, listp2, 64)

	xord := make([]byte, 16)
	for i := 0; i < len(listp2); i += 16 {
		n := listp2[i]
		for j := 1; j < 16; j++ {
			n ^= listp2[i+j]
		}
		xord[i/16] = byte(n)

	}
	fmt.Printf("Hash of '%s' (part2): %s\n", input, hex.EncodeToString(xord))

}
