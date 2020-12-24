package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type cup struct {
	next  *cup
	label int
}

func cupsString(head *cup) string {
	ref := head
	var sb strings.Builder
	for {
		sb.WriteString(strconv.Itoa(head.label))
		if head = head.next; head == ref {
			break
		}
	}

	return sb.String()
}

func printCups(head *cup) {
	ref := head
	for {
		fmt.Printf("%d ", head.label)
		head = head.next
		if head == ref {
			break
		}
	}

	fmt.Println()
}

func findCup(head *cup, label int) (*cup, bool) {
	ref := head
	for {
		if head.label == label {
			return head, true
		}
		if head.next == ref {
			return nil, false
		}
		head = head.next
	}
}

func makeMove(c *cup, maxN int, cups map[int]*cup) *cup {
	c1, c2, c3 := c.next, c.next.next, c.next.next.next
	c.next = c3.next

	s := c.label - 1

	for s == c1.label || s == c2.label || s == c3.label || s == 0 {
		if s == 0 {
			s = maxN
		} else {
			s--
		}
	}

	dest := cups[s]
	c3.next = dest.next
	dest.next = c1

	return c.next
}

func part1() {
	// input := "496138527"
	input := "389125467"

	var head *cup
	var current *cup
	cupMap := map[int]*cup{}

	for _, r := range input {
		dc := &cup{nil, int(r - '0')}
		cupMap[dc.label] = dc
		if head == nil {
			// initial
			dc.next = dc
			head = dc
			current = dc
		} else {
			current.next = dc
			current = dc
		}
	}
	current.next = head

	for i := 0; i < 100; i++ {
		head = makeMove(head, 9, cupMap)
	}

	c, ok := cupMap[1]
	if !ok {
		panic("no cup 1")
	}

	log.Printf("string order after 1 (part1): %s", cupsString(c)[1:])
}

func part2() {
	input := "496138527"

	cupMap := map[int]*cup{}

	var head *cup
	var current *cup
	for _, r := range input {
		dc := &cup{nil, int(r - '0')}
		cupMap[dc.label] = dc
		if head == nil {
			// initial
			dc.next = dc
			head = dc
			current = dc
		} else {
			current.next = dc
			current = dc
		}
	}

	for i := 10; i <= 1000000; i++ {
		dc := &cup{nil, i}
		current.next = dc
		current = dc
		cupMap[i] = dc
	}

	current.next = head
	for i := 0; i < 10000000; i++ {
		head = makeMove(head, 1000000, cupMap)
	}

	c, ok := findCup(head, 1)
	if !ok {
		panic("no cup 1")
	}

	c2, c3 := c.next, c.next.next
	log.Printf("product of clockwise cups (part2): %d", c2.label*c3.label)

}

func main() {
	part1()
	part2()
}
