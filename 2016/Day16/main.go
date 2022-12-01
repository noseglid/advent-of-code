package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func checksum(s string) string {
	if len(s)%2 != 0 {
		panic("not even length string")
	}

	var cs []rune
	for i := 0; i < len(s); i += 2 {
		c := '0'
		if s[i] == s[i+1] {
			c = '1'
		}
		cs = append(cs, c)
	}
	if len(cs)%2 == 0 {
		return checksum(string(cs))
	}
	return string(cs)
}

func dragon(s string) string {
	b := util.ReverseString(s)
	var n []rune
	for _, r := range b {
		n = append(n, rune(util.Absolute(int(r-'0'-1))+'0'))
	}
	return fmt.Sprintf("%s0%s", s, string(n))
}

func sizedDragon(s string, size int) string {
	for {
		if len(s) >= size {
			break
		}
		s = dragon(s)
	}
	return s[:size]
}

func main() {
	//input
	initial := "00101000101111010"

	log.Printf("Part 1: Checksum of size 272: %s", checksum(sizedDragon(initial, 272)))
	log.Printf("Part 2: Checksum of size 35651584: %s", checksum(sizedDragon(initial, 35651584)))

}
