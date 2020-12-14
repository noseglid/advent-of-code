package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type bus struct {
	id        int
	inService bool
}

func parseInput(s string) (int, []bus) {
	s1 := strings.Split(s, "\n")

	var buses []bus
	for _, r := range strings.Split(s1[1], ",") {
		id, err := strconv.Atoi(r)
		if err != nil {
			buses = append(buses, bus{0, false})
		} else {
			buses = append(buses, bus{id, true})
		}
	}

	return util.MustAtoi(s1[0]), buses
}

func main() {
	f := util.GetFile("2020/Day13/input")

	timestamp, buses := parseInput(f)

	test := timestamp
outer:
	for {
		for _, bus := range buses {
			if bus.inService && test%bus.id == 0 {
				log.Printf("found bus %d at timestamp %d, id*waittime (part1): %d", bus.id, test, bus.id*(test-timestamp))
				break outer
			}
		}
		test++
	}

	n, diff := 0, 1
	for i, bus := range buses {
		if !bus.inService {
			continue
		}

		for (n+i)%bus.id != 0 {
			n += diff
		}

		diff *= bus.id
	}

	log.Printf("consecutive timestamp (part2): %d", n)
}
