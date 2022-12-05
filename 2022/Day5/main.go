package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
	"golang.org/x/exp/slices"
)

type stack struct {
	blocks []rune
}

func (s stack) String() string {
	var sb strings.Builder
	sep := ""
	for _, b := range s.blocks {
		sb.WriteString(fmt.Sprintf("%s%c", sep, b))
		sep = ", "
	}
	return sb.String()
}

func main() {

	input := util.GetFileStrings("2022/Day5/input")
	divider := slices.Index(input, "")

	var stacks []stack
	for _, e := range input[divider-1] {
		if e != ' ' {
			stacks = append(stacks, stack{[]rune{}})
		}
	}

	for _, l := range util.Reverse(input[:divider-1]) {
		for i, r := range l {
			if r < 'A' || r > 'Z' {
				continue
			}
			index := 1
			if i > 1 {
				index = (i-1)/4 + 1
			}
			stacks[index-1].blocks = append(stacks[index-1].blocks, r)
		}
	}
	stacksp2 := make([]stack, len(stacks))
	for i, s := range stacks {
		stacksp2[i] = stack{blocks: make([]rune, len(s.blocks))}
		copy(stacksp2[i].blocks, s.blocks)
	}

	for _, l := range input[divider+1:] {
		var n, from, to int
		if _, err := fmt.Sscanf(l, "move %d from %d to %d", &n, &from, &to); err != nil {
			panic(err)
		}

		fromBlocks := stacks[from-1].blocks
		moved := fromBlocks[len(fromBlocks)-n:]
		stacks[from-1].blocks = fromBlocks[:len(fromBlocks)-n]
		stacks[to-1].blocks = append(stacks[to-1].blocks, util.Reverse(moved)...)

		fromBlocksp2 := stacksp2[from-1].blocks
		movedp2 := fromBlocksp2[len(fromBlocksp2)-n:]
		stacksp2[from-1].blocks = fromBlocksp2[:len(fromBlocksp2)-n]
		stacksp2[to-1].blocks = append(stacksp2[to-1].blocks, movedp2...)
	}

	var arrangement string
	for _, s := range stacks {
		arrangement += string(s.blocks[len(s.blocks)-1])
	}
	var arrangement2 string
	for _, s := range stacksp2 {
		arrangement2 += string(s.blocks[len(s.blocks)-1])
	}
	fmt.Printf("Arrangement (part1): %s\n", arrangement)
	fmt.Printf("Arrangement (part2): %s\n", arrangement2)

}
