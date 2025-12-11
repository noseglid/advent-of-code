package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type device struct {
	id     string
	output []string
}

func parseDevice(s string) device {
	var d device
	parts := strings.Split(s, ":")
	d.id = parts[0]
	d.output = append(d.output, strings.Split(strings.TrimSpace(parts[1]), " ")...)
	return d
}

func paths(devices map[string]device, current, target string, stop []string) int {
	if current == target {
		return 1
	}

	if slices.Contains(stop, current) {
		return -1
	}

	s := 0
	for _, d := range devices[current].output {
		if pp := paths(devices, d, target, stop); pp != -1 {
			s += pp
		}
	}
	return s
}

type step struct {
	sources []string
	targets []string
	stop    []string
}

func (s step) exec(devices map[string]device, stepsSources []int) []int {
	r := make([]int, len(s.targets))
	for i := 0; i < len(s.sources); i++ {
		for j := 0; j < len(s.targets); j++ {
			p := paths(devices, s.sources[i], s.targets[j], s.stop)
			r[j] += p * stepsSources[i]
		}
	}

	return r
}

func main() {

	lines := util.GetFileStrings("2025/Day11/input")

	devices := map[string]device{}

	for _, l := range lines {
		d := parseDevice(l)
		devices[d.id] = d
	}

	g, _ := os.Create("graph.dot")
	g.WriteString("digraph {\n")
	for _, d := range devices {
		fmt.Fprintf(g, "\t%s -> { %s }\n", d.id, strings.Join(d.output, " "))
	}
	for _, a := range []string{"you", "svr", "fft", "dac", "out"} {
		fmt.Fprintf(g, "\t%s [bgcolor=blue, style=filled, color=cyan]\n", a)
	}
	fmt.Fprint(g, "\twjw [style=filled, color=gray]")
	fmt.Fprint(g, "\trpp [style=filled, color=gray]")
	fmt.Fprint(g, "\tnhb [style=filled, color=gray]")
	fmt.Fprint(g, "\txal [style=filled, color=gray]")
	fmt.Fprint(g, "\tydw [style=filled, color=gray]")
	g.WriteString("}\n")

	svr := []string{"svr"}
	i1 := []string{"wjw", "wzb", "ury", "lob", "bzl"}
	fft := []string{"fft"}
	i2 := []string{"rpp", "ych", "vwz", "wef"}
	i3 := []string{"nhb", "pxb", "czt"}
	i4 := []string{"xal", "jnn", "qqs", "fkf", "wtz"}
	dac := []string{"dac"}
	i5 := []string{"ydw", "sbh", "wyy", "gui", "you"}
	out := []string{"out"}

	var steps = []step{
		{sources: svr, targets: i1, stop: i1},
		{sources: i1, targets: fft, stop: i2},
		{sources: fft, targets: i2, stop: i2},
		{sources: i2, targets: i3, stop: i3},
		{sources: i3, targets: i4, stop: i4},
		{sources: i4, targets: dac, stop: i5},
		{sources: dac, targets: i5, stop: i5},
		{sources: i5, targets: out, stop: out},
	}

	var stepSources = []int{1}
	for _, s := range steps {
		stepSources = s.exec(devices, stepSources)
	}

	fmt.Printf("Paths from you to out (part1): %d\n", paths(devices, "you", "out", []string{}))
	fmt.Printf("Paths from srv to out (via fft and dac) (part2): %d\n", stepSources[0])
}
