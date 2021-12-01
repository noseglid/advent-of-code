package main

import (
	"bufio"
	"log"
	"os"
)

type direction string

const (
	north = direction("^")
	east  = direction(">")
	south = direction("v")
	west  = direction("<")
)

type coord struct {
	x int
	y int
}

type deliverer struct {
	c coord
}

func part1() {
	f, err := os.Open("2015/Day3/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	santa := deliverer{coord{0, 0}}
	current := &santa

	visited := map[coord]bool{}
	visited[coord{0, 0}] = true

	houses := 1

	eval := func(d *deliverer) {
		c := coord{d.c.x, d.c.y}
		if !visited[c] {
			houses++
		}
		visited[c] = true
	}

	for s.Scan() {
		switch direction(s.Text()) {
		case north:
			current.c.y--
		case east:
			current.c.x++
		case south:
			current.c.y++
		case west:
			current.c.x--
		}
		eval(current)
	}

	log.Printf("visited houses (part2): %d", houses)
}

func part2() {
	f, err := os.Open("2015/Day3/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	santa := deliverer{coord{0, 0}}
	robo := deliverer{coord{0, 0}}
	current := &santa

	visited := map[coord]bool{}
	visited[coord{0, 0}] = true

	houses := 1

	eval := func(d *deliverer) {
		c := coord{d.c.x, d.c.y}
		if !visited[c] {
			houses++
		}
		visited[c] = true
	}

	for s.Scan() {
		switch direction(s.Text()) {
		case north:
			current.c.y--
		case east:
			current.c.x++
		case south:
			current.c.y++
		case west:
			current.c.x--
		}
		eval(current)
		if current == &santa {
			current = &robo
		} else {
			current = &santa
		}

	}

	log.Printf("visited houses (part2): %d", houses)
}

func main() {
	part1()
	part2()
}
