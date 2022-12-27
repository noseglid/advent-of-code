package main

import "fmt"

type generator struct {
	factor uint64
	value  uint64
	modulo uint64
}

func (g *generator) Next() uint64 {
	g.value = (g.value * g.factor) % 2147483647
	return g.value
}

func (g *generator) NextPicky() uint64 {
	for {
		g.value = (g.value * g.factor) % 2147483647
		if g.value%g.modulo == 0 {
			return g.value
		}
	}
}

func low16Match(a, b uint64) bool {
	return a&0xFFFF == b&0xFFFF
}

func main() {

	var va, vb uint64 = 65, 8921
	// var va, vb uint64 = 116, 299

	genA := generator{factor: 16807, value: va, modulo: 4}
	genB := generator{factor: 48271, value: vb, modulo: 8}
	n := 0
	for i := 0; i < 40_000_000; i++ {
		a, b := genA.Next(), genB.Next()
		if low16Match(a, b) {
			n++
		}
	}
	fmt.Printf("Matched numbers (part1): %d\n", n)

	genA.value, genB.value = va, vb
	n2 := 0
	for i := 0; i < 5_000_000; i++ {
		a, b := genA.NextPicky(), genB.NextPicky()
		if low16Match(a, b) {
			n2++
		}
	}
	fmt.Printf("Matched numbers (part2): %d\n", n2)

}
