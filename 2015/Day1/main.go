package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2015/Day1/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	floor := 0
	position := 1
	for s.Scan() {
		switch s.Text() {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 {
			log.Printf("First basement (part2): %d", position)
		}
		position++

	}

	log.Printf("Floor (part1): %d", floor)

}
