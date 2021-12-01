package main

import (
	"log"
	"strconv"
	"strings"
)

func lookAndSay(s string) string {

	c := 0
	cr := ' '

	var out strings.Builder

	for i, r := range s {
		if i == 0 {
			cr = r
			c = 1
		} else if r != cr {
			// new rune, output previous run and reset
			out.WriteString(strconv.Itoa(c))
			out.WriteRune(rune(s[i-1]))
			c = 1
			cr = r
		} else {
			// repeated or first, up counter and take next
			c++
		}
	}

	out.WriteString(strconv.Itoa(c))
	out.WriteRune(rune(s[len(s)-1]))
	return out.String()
}

func main() {
	input := "1321131112"
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	log.Printf("length after 40 iterations (part1): %d", len(input))

	for i := 0; i < 10; i++ {
		input = lookAndSay(input)
	}
	log.Printf("length after 50 iterations (part2): %d", len(input))
}
