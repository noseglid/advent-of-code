package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func isSmallCave(cave string) bool {
	return cave == strings.ToLower(cave)
}

func isVisited(visited []string, cave string) bool {
	for _, c := range visited {
		if cave == c {
			return true
		}
	}
	return false
}

func anySmallCaveVisitedTwice(visited []string) bool {
	m := map[string]struct{}{}
	for _, c := range visited {
		if !isSmallCave(c) {
			continue
		}
		if _, ok := m[c]; ok {
			return true
		}
		m[c] = struct{}{}
	}
	return false
}

func canVisitP1(visited []string, cave string) bool {
	if !isSmallCave(cave) {
		return true
	}

	return !isVisited(visited, cave)
}

func canVisitP2(visited []string, cave string) bool {
	if cave == "start" {
		return false
	}

	if !isSmallCave(cave) {
		return true
	}

	if !isVisited(visited, cave) {
		return true
	}

	if anySmallCaveVisitedTwice(visited) {
		return false
	}

	return true
}

func pathThrough2(connections map[string][]string, cave string, visited []string) int {
	visited = append(visited, cave)
	if cave == "end" {
		return 1
	}

	paths := 0
	for _, n := range connections[cave] {
		if !canVisitP2(visited, n) {
			continue
		}

		paths += pathThrough2(connections, n, visited)
	}

	return paths
}

func pathThrough(connections map[string][]string, cave string, visited []string) int {
	visited = append(visited, cave)
	if cave == "end" {
		return 1
	}

	paths := 0
	for _, n := range connections[cave] {
		if !canVisitP1(visited, n) {
			continue
		}

		paths += pathThrough(connections, n, visited)
	}

	return paths
}

func main() {
	input := "2021/Day12/input"

	lines := util.GetFileStrings(input)

	connections := map[string][]string{}

	for _, c := range lines {
		parts := strings.Split(c, "-")
		from, to := parts[0], parts[1]
		connections[from] = append(connections[from], to)
		connections[to] = append(connections[to], from)
	}
	n := pathThrough(connections, "start", []string{})
	log.Printf("Part 1: Connections through system: %d", n)

	n2 := pathThrough2(connections, "start", []string{})
	log.Printf("Part 2: Multi visit one small cave gives: %d", n2)
}
