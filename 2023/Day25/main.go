package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func count(nodes map[string][]string, visited map[string]bool, currNode string) int {
	if visited == nil {
		visited = map[string]bool{}
	}
	if visited[currNode] {
		return 0
	}

	visited[currNode] = true

	s := 1
	for _, n := range nodes[currNode] {
		s += count(nodes, visited, n)
	}

	return s
}

func cut(nodes map[string][]string, n1, n2 string) map[string][]string {
	out := map[string][]string{}
	for k, v := range nodes {
		switch k {
		case n1:
			lr, _ := util.RemoveByValue(v, n2)
			out[k] = lr
		case n2:
			lr, _ := util.RemoveByValue(v, n1)
			out[k] = lr
		default:
			out[k] = append(out[k], v...)
		}
	}
	return out
}

func without(list []string, el string) []string {
	var r []string
	for _, v := range list {
		if v == el {
			continue
		}
		r = append(r, v)
	}
	return r
}

func printGraphviz(nodes map[string][]string) {
	clone := map[string][]string{}
	for k, v := range nodes {
		clone[k] = append(clone[k], v...)
	}

	simple := map[string][]string{}
	for k, v := range clone {
		simple[k] = v
		for ik, iv := range clone {
			clone[ik] = without(iv, k)
		}
	}

	fmt.Printf("graph {\n")
	for k, v := range simple {
		fmt.Printf("\t%s -- { %s }\n", k, strings.Join(v, " "))
	}
	fmt.Printf("}\n")
}

func main() {

	nodes := map[string][]string{}
	simpleNodes := map[string][]string{}

	lines := util.GetFileStrings("2023/Day25/input")
	for _, l := range lines {
		s := strings.SplitN(l, ":", 2)
		subs := strings.Split(strings.TrimSpace(s[1]), " ")
		simpleNodes[s[0]] = subs
		for _, ss := range subs {
			nodes[s[0]] = append(nodes[s[0]], ss)
			nodes[ss] = append(nodes[ss], s[0])
		}
	}

	for k, v := range nodes {
		nodes[k] = util.Unique(v)
	}

	// Lot's of graphviz magic ...
	nodes = cut(nodes, "mrd", "rjs")
	nodes = cut(nodes, "ncg", "gsk")
	nodes = cut(nodes, "ntx", "gmr")

	n1 := count(nodes, nil, "xgn")
	n2 := count(nodes, nil, "kjb")

	fmt.Printf("Product of groupsizes (part1): %d\n", n1*n2)

}
