package main

import (
	"fmt"
)

func p1() int {
	a, b, c, d, e, f, g, h := 0, 0, 0, 0, 0, 0, 0, 0
	nmul := 0

	b = 65
	c = b
	if a != 0 {
		goto I5
	}
	if 1 != 0 {
		goto I9
	}
I5:
	nmul++
	b *= 100
	b -= -100000
	c = b
	c -= -17000
I9:
	f = 1
	d = 2
I11:
	e = 2
I12:
	g = d
	nmul++
	g *= e
	g -= b
	if g != 0 {
		goto I17
	}
	f = 0
I17:
	e -= -1
	g = e
	g -= b
	if g != 0 {
		goto I12
	}
	d -= -1
	g = d
	g -= b
	if g != 0 {
		goto I11
	}
	if f != 0 {
		goto I27
	}
	h -= -1
I27:
	g = b
	g -= c
	if g != 0 {
		goto I31
	}
	if 1 != 0 {
		goto I33
	}
I31:
	b -= -17
	if 1 != 0 {
		goto I9
	}
I33:
	return nmul
}

func p2() int {
	h, b, c := 0, 65*100+100000, 65*100+100000+17000
	for ; ; b += 17 {
		for d := 2; d*d <= b; d++ {
			if b%d == 0 {
				h++
				break
			}
		}
		if b == c {
			break
		}
	}
	return h
}

func main() {
	fmt.Printf("number of multiplications (part1): %d\n", p1())
	fmt.Printf("value of h (part2): %d\n", p2())
}
