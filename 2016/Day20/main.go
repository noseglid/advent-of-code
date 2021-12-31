package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/noseglid/advent-of-code/util"
)

type blocked struct {
	s, e int
}

type blockList []blocked

func (b blockList) Len() int { return len(b) }
func (b blockList) Less(i, j int) bool {
	if b[i].s == b[j].s {
		return b[i].e < b[j].e
	}
	return b[i].s < b[j].s
}
func (b blockList) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b blockList) Sum() int {
	n := 0
	for _, e := range b {
		n += e.e - e.s + 1
	}
	return n
}

func collapse(a, b blocked) blockList {
	if a.s > b.e+1 || b.s > a.e+1 {
		return blockList{a, b}
	}
	return blockList{blocked{util.MinInt(a.s, b.s), util.MaxInt(a.e, b.e)}}
}

func reduce(b blockList) blockList {
	for i := 1; i < len(b); {
		cc := collapse(b[i-1], b[i])
		if len(cc) == 2 {
			i++
			continue
		}

		b[i] = cc[0]
		b = append(b[:i-1], b[i:]...)
	}

	return b
}

func main() {

	input := "2016/Day20/input"
	lines := util.GetFileStrings(input)

	var list blockList
	for _, l := range lines {
		b := blocked{}
		fmt.Sscanf(l, "%d-%d", &b.s, &b.e)
		list = append(list, b)
	}
	sort.Sort(list)
	log.Print(list[0])
	m := 0
	for _, e := range list {
		if e.s > m {
			break
		}
		m = e.e + 1
	}
	log.Printf("Part 1: Minimum allowed IP: %d", m)
	log.Printf("Part 2: Total allowed IPs: %d", math.MaxUint32-reduce(list).Sum()+1)
}
