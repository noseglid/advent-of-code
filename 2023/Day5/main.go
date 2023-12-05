package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/noseglid/advent-of-code/util"
)

type Range struct {
	start  int
	length int
}

type AgriMap struct {
	source []Range
	dest   []Range
}

func (a *AgriMap) Add(l []int) {
	a.dest = append(a.dest, Range{start: l[0], length: l[2]})
	a.source = append(a.source, Range{start: l[1], length: l[2]})
}

func (a *AgriMap) SourceToDest(t int) int {
	for i, v := range a.source {
		if t >= v.start && t < v.start+v.length {
			return a.dest[i].start + t - v.start
		}
	}
	return t
}

func numberList(s string) []int {
	numbers := []int{}
	for _, s := range strings.Split(s, " ") {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}

func parse(file *bufio.Scanner) ([]int, map[string]*AgriMap) {
	file.Scan()
	maps := map[string]*AgriMap{}

	seeds := numberList(strings.TrimPrefix(file.Text(), "seeds: "))

	file.Scan() // empty line

	for file.Scan() {
		name := strings.TrimSuffix(file.Text(), " map:")
		m := &AgriMap{}
		for file.Scan() {
			if file.Text() == "" {
				break
			}
			nl := numberList(file.Text())
			m.Add(nl)
		}

		maps[name] = m
	}

	return seeds, maps
}

func main() {
	sc := util.FileScanner("2023/Day5/input", bufio.ScanLines)
	seeds, maps := parse(sc)

	min := math.MaxInt
	for _, s := range seeds {
		n := s
		n = maps["seed-to-soil"].SourceToDest(n)
		n = maps["soil-to-fertilizer"].SourceToDest(n)
		n = maps["fertilizer-to-water"].SourceToDest(n)
		n = maps["water-to-light"].SourceToDest(n)
		n = maps["light-to-temperature"].SourceToDest(n)
		n = maps["temperature-to-humidity"].SourceToDest(n)
		n = maps["humidity-to-location"].SourceToDest(n)
		if n < min {
			min = n
		}
	}
	workers := 200

	wg := sync.WaitGroup{}
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < len(seeds); i += 2 {
			fmt.Printf("Generating %d entries at %d\n", seeds[i+1], i)
			for s := 0; s < seeds[i+1]; s++ {
				ch <- seeds[i] + s
			}
		}
		close(ch)
	}()

	minp2 := math.MaxInt
	locationch := make(chan int, 8)
	go func() {
		r := 0
		for m := range locationch {
			r++
			if r%1000000 == 0 {
				fmt.Printf("processed %d results\n", r)
			}
			if m < minp2 {
				fmt.Printf("new min at %d\n", m)
				minp2 = m
			}
		}
	}()

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(ch <-chan int) {
			for seed := range ch {
				n := seed
				n = maps["seed-to-soil"].SourceToDest(n)
				n = maps["soil-to-fertilizer"].SourceToDest(n)
				n = maps["fertilizer-to-water"].SourceToDest(n)
				n = maps["water-to-light"].SourceToDest(n)
				n = maps["light-to-temperature"].SourceToDest(n)
				n = maps["temperature-to-humidity"].SourceToDest(n)
				n = maps["humidity-to-location"].SourceToDest(n)
				locationch <- n
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()

	fmt.Printf("min location (part1): %d\n", min)
	fmt.Printf("min location (part2): %d\n", minp2)
}
