package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func part1() {
	// steps := 3
	steps := 371
	list := util.NewCircularLinkedList[int]()
	list.Add(0)
	n := list.GetNode(0)
	for i := 1; i < 2018; i++ {
		n = list.StepFromNode(n, steps)
		list.AddAfterNode(i, n)
		n = list.StepFromNode(n, 1)
	}
	n = list.StepFromNode(n, 1)
	fmt.Printf("Value after 2017 (part1): %d\n", list.NodeValue(n))
}

func part2() {
	// steps := 3
	steps := 371
	list := util.NewCircularLinkedList[int]()
	list.Add(0)
	n := list.GetNode(0)
	for i := 1; i < 50_000_000; i++ {
		if i%100000 == 0 {
			fmt.Printf("%d\n", i)
		}
		n = list.StepFromNode(n, steps)
		list.AddAfterNode(i, n)
		n = list.StepFromNode(n, 1)
	}

	zero, ok := list.GetNodeFunc(func(v int) bool { return v == 0 })
	if !ok {
		panic("no 0 value")
	}
	az := list.StepFromNode(zero, 1)

	fmt.Printf("Value after 0 in 50 mill inserts (part2): %d\n", list.NodeValue(az))

}

func main() {
	part1()
	part2()

}
