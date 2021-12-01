package main

import (
	"bufio"
	"log"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type state string

const (
	flying  = state("flying")
	resting = state("resting")
)

type reindeer struct {
	name              string
	speed             int
	flyTime, restTime int

	state       state
	timeInState int
	distance    int
	points      int
}

var re = regexp.MustCompile(`([[:alpha:]]+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func parseReindeer(s string) reindeer {
	m := re.FindStringSubmatch(s)
	if len(m) != 5 {
		panic("invalid match")
	}

	return reindeer{
		name:     m[1],
		state:    state("flying"),
		speed:    util.MustAtoi(m[2]),
		flyTime:  util.MustAtoi(m[3]),
		restTime: util.MustAtoi(m[4]),
	}
}

func tickSecond(list []reindeer) {
	for i := range list {
		r := &list[i]
		r.timeInState = r.timeInState + 1

		switch r.state {
		case flying:
			r.distance = r.distance + r.speed
			if r.timeInState == r.flyTime {
				r.state = resting
				r.timeInState = 0
			}

		case resting:
			if r.timeInState == r.restTime {
				r.state = flying
				r.timeInState = 0
			}
		}
	}
}

func reindeerContains(list []reindeer, name string) bool {
	for _, r := range list {
		if r.name == name {
			return true
		}
	}
	return false
}

func awardPoints(list []reindeer) {
	leaders := leaders(list)
	for i := range list {
		r := &list[i]

		if reindeerContains(leaders, r.name) {
			r.points = r.points + 1
		}

	}
}

func print(list []reindeer) {
	for _, r := range list {
		log.Printf("%s is %s (has been for %d s), reached %d km", r.name, r.state, r.timeInState, r.distance)
	}
}

func leaders(list []reindeer) []reindeer {
	winners := []reindeer{list[0]}
	for _, r := range list[1:] {
		if r.distance == winners[0].distance {
			winners = append(winners, r)
		} else if r.distance > winners[0].distance {
			winners = []reindeer{r}
		}
	}

	return winners
}

func main() {
	s := util.FileScanner("2015/Day14/input", bufio.ScanLines)
	totalTime := 2503

	var reindeers []reindeer

	for s.Scan() {
		reindeers = append(reindeers, parseReindeer(s.Text()))
	}

	for i := 0; i < totalTime; i++ {
		tickSecond(reindeers)
		awardPoints(reindeers)
	}

	leaders := leaders(reindeers)
	log.Printf("winning distance (part1): %d", leaders[0].distance)

	maxPoints := reindeers[0].points
	for _, r := range reindeers[1:] {
		if r.points > maxPoints {
			maxPoints = r.points
		}
	}

	log.Printf("winning points (part2): %d", maxPoints)

}
