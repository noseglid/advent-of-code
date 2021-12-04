package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type board struct {
	grid   [5][5]int
	marked [5][5]bool

	hasWon bool
}

func NewBoard(lines []string) *board {
	var b board

	for i, l := range lines {
		numbers := strings.Split(l, " ")
		j := 0
		for _, n := range numbers {
			if n == "" {
				continue
			}
			b.grid[i][j] = util.MustAtoi(n)
			j++
		}
	}

	return &b
}

func (b *board) String() string {
	var sb strings.Builder
	var mb strings.Builder
	for i, r := range b.grid {
		for j, c := range r {
			if b.marked[i][j] {
				mb.WriteString(" X")
			} else {
				mb.WriteString(" 0")
			}
			mb.WriteRune(' ')
			sb.WriteString(fmt.Sprintf("%02d", c))
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
		mb.WriteRune('\n')
	}
	return fmt.Sprintf("%s\n%s\n", sb.String(), mb.String())
}

func (b *board) Mark(in int) {
	for i, r := range b.grid {
		for j, c := range r {
			if c == in {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *board) Bingo() bool {
	for i := 0; i < 5; i++ {
		r := true
		c := true
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				r = false
			}

			if !b.marked[j][i] {
				c = false
			}
		}

		if r || c {
			return true
		}
	}

	return false
}

func (b *board) WinScore(lastDraw int) int {
	sum := 0
	for i, r := range b.marked {
		for j, c := range r {
			if !c {
				sum += b.grid[i][j]
			}
		}
	}

	return sum * lastDraw
}

func main() {
	file := "2021/Day4/input"
	lines := util.GetFileStrings(file)

	var numbers []int
	for _, nn := range strings.Split(lines[0], ",") {
		numbers = append(numbers, util.MustAtoi(nn))
	}
	boards := []*board{}

	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, NewBoard(lines[i:i+5]))
	}

	firstWon := false
	nWon := 0

Outer:
	for _, n := range numbers {
		for i, b := range boards {
			b.Mark(n)
			hasBingo := b.Bingo()
			if hasBingo && !firstWon {
				log.Printf("Part1: board %d won when %d was drawn for score: %d", i, n, b.WinScore(n))
				firstWon = true
			}

			if hasBingo && !b.hasWon {
				nWon++
				b.hasWon = true
			}

			if nWon == len(boards) {
				log.Printf("Part2: board %d won last when %d was drawn for score: %d", i, n, b.WinScore(n))
				break Outer
			}

		}
	}
}
