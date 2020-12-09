package main

import (
	"bufio"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func hasSum(numbers []int, expected int) bool {
	for i := range numbers {
		for j := range numbers {
			if i == j {
				continue
			}

			if numbers[i]+numbers[j] == expected {
				return true
			}
		}
	}
	return false
}

func contiguous(numbers []int, n int) (int, int, bool) {
	for startIndex := 0; startIndex < len(numbers); startIndex++ {
		for endIndex := startIndex + 1; endIndex < len(numbers); endIndex++ {
			sum := 0
			for i := startIndex; i < endIndex; i++ {
				sum += numbers[i]
			}

			if sum == n {
				return startIndex, endIndex, true
			}
		}
	}

	return 0, 0, false
}

func extremes(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return min, max
}

func main() {
	s := util.FileScanner("2020/Day9/input", bufio.ScanLines)

	var numbers []int

	for s.Scan() {
		numbers = append(numbers, util.MustAtoi(s.Text()))
	}

	invalidNumber := 0
	for i := range numbers[25:] {
		index := i + 25
		if !hasSum(numbers[i:i+25], numbers[index]) {
			invalidNumber = numbers[index]
			log.Printf("First number not matching sum of 25 previous (part1): %d", invalidNumber)
			break
		}
	}

	a, b, ok := contiguous(numbers, invalidNumber)
	if !ok {
		log.Fatal("no contiguous for invalid number")
	}

	min, max := extremes(numbers[a:b])
	log.Printf("Sum of min max for invalid number (part2): %d", min+max)

}
