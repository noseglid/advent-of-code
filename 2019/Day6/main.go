package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type object struct {
	id     string
	parent *object
}

func (o *object) orbits() int {
	p := o.parent
	n := 0
	for p != nil {
		p = p.parent
		n++
	}

	return n
}

func (o *object) parents() []string {
	var r []string
	p := o.parent
	for p != nil {
		r = append(r, p.id)
		p = p.parent
	}
	return r
}

func main() {
	s := util.FileScanner("2019/Day6/input", bufio.ScanLines)
	objects := map[string]*object{}

	for s.Scan() {
		parts := strings.Split(s.Text(), ")")
		id0, id1 := parts[0], parts[1]
		o0, o1 := objects[id0], objects[id1]
		if o0 == nil {
			o0 = &object{id: id0}
		}
		if o1 == nil {
			o1 = &object{id: id1}
		}
		o1.parent = o0
		objects[id0] = o0
		objects[id1] = o1
	}

	n := 0
	for _, o := range objects {
		n += o.orbits()
	}

	log.Printf("total orbits (part1): %d", n)

outer:
	for i, myParent := range objects["YOU"].parents() {
		for j, sanParent := range objects["SAN"].parents() {
			if myParent == sanParent {
				log.Printf("total transfers (part2): %d", i+j)
				break outer
			}
		}
	}

}
