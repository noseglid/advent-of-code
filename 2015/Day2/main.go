package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rrp struct {
	length int
	width  int
	height int
}

func parseToRRP(def string) rrp {
	parts := strings.Split(def, "x")
	l, _ := strconv.Atoi(parts[0])
	w, _ := strconv.Atoi(parts[1])
	h, _ := strconv.Atoi(parts[2])
	return rrp{l, w, h}
}

func (r rrp) surfaceArea() int {
	return 2*r.length*r.width + 2*r.width*r.height + 2*r.height*r.length
}

func (r rrp) smallestSideArea() int {
	slice := []int{r.length * r.width, r.width * r.height, r.height * r.length}
	sort.Ints(slice)
	return slice[0]
}

func (r rrp) smallestPerimiter() int {
	slice := []int{r.length, r.width, r.height}
	sort.Ints(slice)
	return 2*slice[0] + 2*slice[1]
}

func (r rrp) volume() int {
	return r.length * r.width * r.height
}

func main() {
	f, err := os.Open("2015/Day2/input")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	paper := 0
	ribbon := 0
	for s.Scan() {
		rrp := parseToRRP(s.Text())

		paper += rrp.surfaceArea() + rrp.smallestSideArea()
		ribbon += rrp.smallestPerimiter() + rrp.volume()
	}

	log.Printf("total wrapping (part1): %d square feet", paper)
	log.Printf("total ribbon (part2): %d feet", ribbon)
}
