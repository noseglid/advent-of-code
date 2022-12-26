package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type scanner struct {
	x, y    int
	upwards bool
	height  int
}

func (s *scanner) Step() {
	if s.upwards {
		s.y--
		if s.y == 0 {
			s.upwards = false
		}
	} else {
		s.y++
		if s.y == s.height-1 {
			s.upwards = true
		}
	}
}

func (s *scanner) Reset() {
	s.y = 0
	s.upwards = false
}

func main() {
	lines := util.GetFileStrings("2017/Day13/input")

	width := 0
	var scanners []*scanner
	for _, l := range lines {
		var x, height int
		if _, err := fmt.Sscanf(l, "%d: %d", &x, &height); err != nil {
			panic(err)
		}
		if x > width {
			width = x
		}

		scanners = append(scanners, &scanner{x: x, y: 0, height: height})
	}

	x := -1
	severity := 0
	for i := 0; i <= width; i++ {
		x++
		for _, s := range scanners {
			if s.x == x && s.y == 0 {
				severity += s.height * s.x
			}
		}
		for _, s := range scanners {
			s.Step()
		}
	}
	fmt.Printf("severity (part1): %d\n", severity)

	for _, s := range scanners {
		s.Reset()
	}

	packets := []int{}
Outer:
	for i := 0; ; i++ {

		packets = append(packets, -1)
		for p := range packets {
			packets[p] = packets[p] + 1
		}

		var survivingPackets []int
	PacketLoop:
		for _, p := range packets {
			for _, s := range scanners {
				if s.x == p && s.y == 0 {
					continue PacketLoop
				}
			}
			if p == width {
				fmt.Printf("min delay to not get caught (part2): %d\n", i-width)
				break Outer
			}
			survivingPackets = append(survivingPackets, p)
		}

		packets = survivingPackets
		for _, s := range scanners {
			s.Step()
		}
	}

}
