package main

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func getNthFromRight(s string, n int) (rune, bool) {
	if n > len(s) {
		return 0, false
	}
	return rune(s[len(s)-n]), true
}

func main() {

	lines := util.GetFileStrings("2025/Day6/input")

	var grid [][]string

	var ops []string

	for _, l := range lines[:len(lines)-1] {
		scanner := bufio.NewScanner(bytes.NewReader([]byte(l)))
		scanner.Split(bufio.ScanWords)
		var row []string
		for scanner.Scan() {
			row = append(row, scanner.Text())
		}

		grid = append(grid, row)
	}

	scanner := bufio.NewScanner(bytes.NewReader([]byte(lines[len(lines)-1])))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		ops = append(ops, scanner.Text())
	}

	sum := 0
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "*":
			p := 1
			for j := 0; j < len(grid); j++ {
				p *= util.MustAtoi(grid[j][i])
			}
			sum += p
		case "+":
			s := 0
			for j := 0; j < len(grid); j++ {
				s += util.MustAtoi(grid[j][i])
			}
			sum += s
		}
	}

	fmt.Printf("Grand total score (part1): %d\n", sum)

	sump2 := 0
	var numbers []int
nextcol:
	for col := len(lines[0]) - 1; col >= 0; col-- {
		var digits []rune
		for row := 0; row < len(lines); row++ {
			switch lines[row][col] {
			case ' ':
				continue
			case '+':
				numbers = append(numbers, util.MustAtoi(string(digits)))
				is := 0
				for _, n := range numbers {
					is += n
				}
				sump2 += is
				numbers = numbers[:0]
				continue nextcol

			case '*':
				numbers = append(numbers, util.MustAtoi(string(digits)))
				ip := 1
				for _, n := range numbers {
					ip *= n
				}
				sump2 += ip
				numbers = numbers[:0]
				continue nextcol
			default:
				digits = append(digits, rune(lines[row][col]))
			}
		}
		if len(digits) > 0 {
			numbers = append(numbers, util.MustAtoi(string(digits)))
		}
	}
	fmt.Printf("Grand total score rtl, td (part2): %d\n", sump2)
}
