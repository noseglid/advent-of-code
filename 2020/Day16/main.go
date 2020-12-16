package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type restriction struct {
	name   string
	ranges [][2]int
}

type ticket []int

func parseRestriction(s string) restriction {
	sp := strings.Split(s, ":")
	restr := restriction{
		name: sp[0],
	}
	for _, rr := range strings.Split(sp[1], " or") {
		dashIndex := strings.Index(rr, "-")
		restr.ranges = append(restr.ranges, [2]int{util.MustAtoi(rr[1:dashIndex]), util.MustAtoi(rr[dashIndex+1:])})
	}

	return restr
}

func parseTicket(s string) ticket {
	t := ticket{}

	for _, n := range strings.Split(s, ",") {
		t = append(t, util.MustAtoi(n))
	}

	return t
}

func matchesRestrictionRanges(v int, r [][2]int) bool {
	for _, rr := range r {
		if v >= rr[0] && v <= rr[1] {
			return true
		}
	}
	return false
}

func sumInvalidValues(t ticket, restrictions []restriction) int {
	sum := 0
	for _, v := range t {
		matchesRestriction := false
		for _, r := range restrictions {
			if matchesRestrictionRanges(v, r.ranges) {
				matchesRestriction = true
				break
			}

		}

		if !matchesRestriction {
			sum += v
		}
	}

	return sum
}

func ticketFields(validTickets []ticket, restrictions []restriction) map[string][]int {
	m := map[string][]int{}
	for _, r := range restrictions {
		for tryIndex := 0; tryIndex < len(validTickets[0]); tryIndex++ {
			matchesAll := true
			for _, t := range validTickets {
				if !matchesRestrictionRanges(t[tryIndex], r.ranges) {
					matchesAll = false
				}
			}

			if matchesAll {
				m[r.name] = append(m[r.name], tryIndex)
			}
		}
	}

	return m
}

func contains(i int, l []int) bool {
	for _, e := range l {
		if i == e {
			return true
		}
	}
	return false
}

func findSingleUsage(m map[string][]int, skip []int) (string, int, bool) {
	for name, possibilities := range m {
		if len(possibilities) == 1 {
			if !contains(possibilities[0], skip) {
				return name, possibilities[0], true
			}
		}
	}

	return "", 0, false
}

func reduce(m map[string][]int) map[string][]int {
	skip := []int{}
	for {
		name, index, ok := findSingleUsage(m, skip)
		if !ok {
			break
		}

		skip = append(skip, index)

		for tname, usages := range m {
			if tname != name {
				for i, u := range usages {
					if u == index {
						m[tname] = append(usages[:i], usages[i+1:]...)
						break
					}
				}
			}
		}
	}

	return m
}

func main() {
	parts := strings.Split(util.GetFile("2020/Day16/input"), "\n\n")

	var restrictions []restriction
	for _, srestr := range strings.Split(parts[0], "\n") {
		restrictions = append(restrictions, parseRestriction(srestr))
	}

	myTicket := parseTicket(strings.Split(parts[1], "\n")[1])

	var nearbyTickets []ticket
	for _, sticket := range strings.Split(parts[2], "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(sticket))
	}

	var validTickets []ticket

	invalid := 0
	for _, t := range nearbyTickets {
		sum := sumInvalidValues(t, restrictions)
		if sum == 0 {
			validTickets = append(validTickets, t)
		}
		invalid += sum
	}
	log.Printf("Invalid tickets value sum (part1): %d", invalid)

	prod := 1
	for key, index := range reduce(ticketFields(validTickets, restrictions)) {
		if strings.HasPrefix(key, "departure") {
			prod *= myTicket[index[0]]
		}
	}

	log.Printf("product of departures (part2): %d", prod)
}
