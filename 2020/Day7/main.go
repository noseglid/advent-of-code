package main

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type content struct {
	bag string
	qty int
}

type Graph map[string][]content

var reMain = regexp.MustCompile(`([[:alpha:] ]+)bags contain (.+)`)
var reContain = regexp.MustCompile(`(\d+) ([^ ]+ [^ ]+) bags?`)

func parseRegulation(s string) (string, []content) {
	var contents []content
	m := reMain.FindStringSubmatch(s)
	theBag, contentArr := m[1], m[2]
	for _, contentSpec := range strings.Split(contentArr, ",") {
		inner := reContain.FindStringSubmatch(contentSpec)
		if len(inner) != 3 {
			continue
		}
		contents = append(contents, content{inner[2], util.MustAtoi(inner[1])})
	}
	return strings.Trim(theBag, " "), contents
}

func canEventuallyContain(graph Graph, search string, target string) bool {
	for _, test := range graph[search] {
		if test.bag == target {
			return true
		}
		if canEventuallyContain(graph, test.bag, target) {
			return true
		}
	}

	return false
}

func countBags(graph Graph, target string) int {
	n := 0
	for _, c := range graph[target] {
		n += c.qty * (1 + countBags(graph, c.bag))
	}
	return n
}

func main() {

	s := util.FileScanner("2020/Day7/input", bufio.ScanLines)

	graph := Graph{}

	for s.Scan() {
		bag, allowedContents := parseRegulation(s.Text())
		graph[bag] = allowedContents
	}

	n := 0
	for test := range graph {
		if canEventuallyContain(graph, test, "shiny gold") {
			n++
		}
	}

	log.Printf("can contain shiny gold (part1): %d", n)
	log.Printf("shiny gold nested bags (part2): %d", countBags(graph, "shiny gold"))
}
