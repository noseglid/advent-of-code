package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("2020/Day1/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	var numbers []int

	for s.Scan() {
		number, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	part1(numbers)
	part2(numbers)

}

func part1(numbers []int) {
	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				log.Printf("part1 (of %d and %d): %d", i, j, i*j)
				return
			}
		}
	}
}

func part2(numbers []int) {
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					log.Printf("part2 (of %d, %d and %d): %d", i, j, k, i*j*k)
					return
				}
			}
		}
	}
}
