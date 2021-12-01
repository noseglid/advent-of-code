package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

func isFiveZeros(sl [16]byte) bool {
	return sl[0]|0x00 == 0 && sl[1]|0x00 == 0 && sl[2]|0x0f == 0x0f
}

func isSixZeros(sl [16]byte) bool {
	return sl[0]|0x00 == 0 && sl[1]|0x00 == 0 && sl[2]|0x00 == 0x00
}

func main() {
	input := "yzbqklnj"

	foundp1 := false
	foundp2 := false
	n := 0
	for {
		test := fmt.Sprintf("%s%d", input, n)
		sum := md5.Sum([]byte(test))
		// log.Printf("test: %s, sum: %x", test, sum)
		if !foundp1 && isFiveZeros(sum) {
			log.Printf("part1: hash %x: number: %d", sum, n)
			foundp1 = true
		}

		if isSixZeros(sum) {
			log.Printf("parts2: hash %x, number: %d", sum, n)
			foundp2 = true
		}

		if foundp1 && foundp2 {
			break
		}
		n++
	}
}
