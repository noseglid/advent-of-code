package main

import (
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func main() {

	input := "2017/Day1/input"
	data := util.GetFileSingleDigitGrid(input)[0]

	s := 0
	for i := 0; i < len(data); i++ {
		if data[i] == data[(i+1)%len(data)] {
			s += data[i]
		}
	}
	log.Printf("Part 1: Sum of repeated digits: %d", s)

	s = 0
	for i := 0; i < len(data); i++ {
		if data[i] == data[(i+len(data)/2)%len(data)] {
			s += data[i]
		}
	}
	log.Printf("Part 2: Sum of repeated digits: %d", s)

}
