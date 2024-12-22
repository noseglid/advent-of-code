package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

func gen(secret int) int {
	v := secret
	s1 := secret * 64
	v = (s1 ^ v) % 16777216

	s2 := (v / 32)
	v = (s2 ^ v) % 16777216

	s3 := (v * 2048)
	v = (s3 ^ v) % 16777216

	return v
}

func genn(secret, n int) int {
	for i := 0; i < n; i++ {
		secret = gen(secret)
	}
	return secret
}

func gennn(secret, n int) []int {
	secrets := []int{secret}
	for i := 0; i < n; i++ {
		secret = gen(secret)
		secrets = append(secrets, secret)
	}
	return secrets
}

func diffs(l []int) []int {
	d := []int{0}
	for i := 1; i < len(l); i++ {
		d = append(d, l[i]%10-l[i-1]%10)
	}
	return d
}

func bananas(m map[skey]int, a, b, c, d int) int {
	return m[skey{a, b, c, d}]
}

type skey struct {
	a, b, c, d int
}

func sequenceMap(secret int, n int) map[skey]int {
	list := gennn(secret, n)
	dd := diffs(list)
	r := map[skey]int{}

	for i := 4; i < len(dd); i++ {
		sk := skey{dd[i-3], dd[i-2], dd[i-1], dd[i]}
		if _, ok := r[sk]; !ok {
			r[sk] = list[i] % 10
		}
	}

	return r
}

func main() {

	lines := util.GetFileStrings("2024/Day22/input")

	s := 0
	for _, l := range lines {
		s += genn(util.MustAtoi(l), 2000)
	}
	fmt.Printf("Sum of all 2000th secrets (part1): %d\n", s)

	maps := map[int]map[skey]int{}
	for _, l := range lines {
		ll := util.MustAtoi(l)
		maps[ll] = sequenceMap(ll, 2000)
	}

	max := 0
	for a := -9; a <= 9; a++ {
		for b := -9; b <= 9; b++ {
			for c := -9; c <= 9; c++ {
				for d := -9; d <= 9; d++ {
					bananaSeq := 0
					for _, l := range lines {
						bananaSeq += bananas(maps[util.MustAtoi(l)], a, b, c, d)
					}
					if bananaSeq > max {
						max = bananaSeq
					}
				}
			}
		}
	}

	fmt.Printf("Max bananas (part2): %d\n", max)
}
