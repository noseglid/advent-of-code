package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type ingredient struct {
	id       string
	quantity int
}

func (i ingredient) String() string {
	return fmt.Sprintf("%d %s", i.quantity, i.id)
}

type reaction struct {
	input  []ingredient
	output ingredient
}

func (r reaction) String() string {
	inputs := []string{}
	for _, i := range r.input {
		inputs = append(inputs, fmt.Sprintf("%d %s", i.quantity, i.id))
	}

	return fmt.Sprintf("%s => %s", strings.Join(inputs, ", "), r.output)
}

func parseReaction(s string) reaction {
	var r reaction
	io := strings.Split(s, " => ")

	for _, in := range strings.Split(io[0], ", ") {
		var ing ingredient
		fmt.Sscanf(in, "%d %s", &ing.quantity, &ing.id)
		r.input = append(r.input, ing)
	}
	fmt.Sscanf(io[1], "%d %s", &r.output.quantity, &r.output.id)
	return r
}

func updateRequired(ingredients []ingredient) []ingredient {
	var r []ingredient

	for len(ingredients) > 0 {
		i := ingredients[0]

		ning := []ingredient{}
		for _, ii := range ingredients[1:] {
			if i.id == ii.id {
				i.quantity += ii.quantity
			} else {
				ning = append(ning, ii)
			}
		}
		ingredients = ning

		r = append(r, i)
	}

	return r
}

func dumpGraph(reactions []reaction) {
	f, _ := os.Create("graph.dot")
	fmt.Fprintf(f, "digraph {\n")
	for _, r := range reactions {
		for _, i := range r.input {
			fmt.Fprintf(f, "\t%s -> %s", i.id, r.output.id)
		}
	}
	fmt.Fprintf(f, "}")
	f.Close()
}

func oreForFuel(reactions map[string]reaction, fuel int) int {
	surplus := map[string]int{}
	required := []ingredient{{"FUEL", fuel}}

	ore := 0
	for len(required) > 0 {
		required = updateRequired(required)
		req := required[0]

		reaction := reactions[req.id]

		toGenerate := req.quantity - surplus[req.id]
		if toGenerate < 0 {
			surplus[req.id] -= req.quantity
			required = required[1:]
			continue
		}

		surplus[req.id] = 0
		units := toGenerate / reaction.output.quantity
		if toGenerate%reaction.output.quantity != 0 {
			units++
		}
		generated := units * reaction.output.quantity

		for _, input := range reaction.input {
			if input.id == "ORE" {
				ore += units * input.quantity
			} else {
				required = append(required, ingredient{input.id, units * input.quantity})
			}
		}
		sp := generated - toGenerate
		surplus[reaction.output.id] += sp
		required = required[1:]
	}

	return ore
}

func main() {
	lines := util.GetFileStrings("2019/Day14/input")
	reactions := map[string]reaction{}
	for _, l := range lines {
		r := parseReaction(l)
		reactions[r.output.id] = r
	}

	fmt.Printf("Ore required to produce 1 fuel (part1): %d\n", oreForFuel(reactions, 1))

	step := 7500
	lower, upper := 0, 0
	for i := step; i < 1000000000000; i += step {
		o := oreForFuel(reactions, i)
		if o > 1000000000000 {
			lower, upper = i-step, i
			break
		}
	}
	for i := lower; i <= upper; i++ {
		o := oreForFuel(reactions, i)
		if o >= 1000000000000 {
			fmt.Printf("1 trillion ores produce fuel (part2): %d\n", i-1)
			break
		}
	}

}
