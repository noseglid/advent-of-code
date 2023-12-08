package main

import (
	"bufio"
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type Node struct {
	L, R string
}

func main() {
	s := util.FileScanner("2023/Day8/input", bufio.ScanLines)
	s.Scan()

	instr := []rune(s.Text())
	nodes := map[string]Node{}
	s.Scan()
	for s.Scan() {
		var n, l, r string
		fmt.Sscanf(s.Text(), "%s = (%s %s)", &n, &l, &r)
		nodes[n] = Node{l, r}
	}

	steps := 0
	curr := "AAA"
	for i := 0; ; i = (i + 1) % len(instr) {
		if instr[i] == 'L' {
			curr = nodes[curr].L
		} else {
			curr = nodes[curr].R
		}
		steps++

		if curr == "ZZZ" {
			break
		}
	}

	var currs []string
	p2steps := []int{}
	for n := range nodes {
		if n[2] == 'A' {
			currs = append(currs, n)
			p2steps = append(p2steps, 0)
		}
	}

	stepsp2 := 0
Main:
	for i := 0; ; i = (i + 1) % len(instr) {
		for c := 0; c < len(currs); c++ {
			if instr[i] == 'L' {
				currs[c] = nodes[currs[c]].L
			} else {
				currs[c] = nodes[currs[c]].R
			}
		}
		stepsp2++

		for j, c := range currs {
			if c[2] == 'Z' && p2steps[j] == 0 {
				p2steps[j] = stepsp2
			}
		}

		for _, p := range p2steps {
			if p == 0 {
				continue Main
			}
		}

		break
	}

	p := 1
	for _, s := range p2steps {
		p *= s
	}

	fmt.Printf("steps to reach ZZZ (part1): %d\n", steps)
	fmt.Printf("steps to reach __Z (part2): %d\n", util.LCM(p2steps[0], p2steps[1], p2steps[2:]...))

}
