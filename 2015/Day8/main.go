package main

import (
	"bufio"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type state int

func checkSize(s string) (int, int) {
	memSize := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"':
		case '\\':
			if s[i+1] == 'x' {
				memSize += 1
				i += 3
			} else if s[i+1] == '\\' || s[i+1] == '"' {
				memSize += 1
				i += 1
			}

		default:
			memSize += 1
		}
	}

	return len(s), memSize
}

func encodeSize(s string) (int, int) {
	encSize := 2 // for quotes
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"':
			encSize += 2
		case '\\':
			encSize += 2
		default:
			encSize += 1
		}
	}

	log.Printf("%s : %d -> %d", s, len(s), encSize)
	return len(s), encSize
}

func main() {
	s := util.FileScanner("2015/Day8/input", bufio.ScanLines)

	totalCodeSize := 0
	totalMemSize := 0
	totalEncSize := 0
	for s.Scan() {
		cm, mm := checkSize(s.Text())
		totalCodeSize += cm
		totalMemSize += mm

		_, em := encodeSize(s.Text())
		totalEncSize += em
	}

	log.Printf("code overhead (part1): %d", totalCodeSize-totalMemSize)
	log.Printf("encode overhead (part2): %d", totalEncSize-totalCodeSize)
}
