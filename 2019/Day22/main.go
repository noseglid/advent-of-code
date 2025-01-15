package main

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

func cut(cards []int, n int) []int {
	if n < 0 {
		n += len(cards)
	}

	c := make([]int, len(cards))
	copy(c, cards[n:])
	copy(c[len(cards)-n:], cards[:n])
	return c
}

func dealInc(cards []int, inc int) []int {
	c := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		c[i*inc%len(cards)] = cards[i]
	}
	return c
}

func dealNew(cards []int) []int {
	slices.Reverse(cards)
	return cards
}

type deck struct {
	offset    int64
	increment int64
	cards     int64
}

func modPow(x, y, z int64) int64 {
	r := new(big.Int).Exp(big.NewInt(x), big.NewInt(y), big.NewInt(z))
	return r.Int64()
}

func posMod(x, y int64) int64 {
	bx, by := big.NewInt(x), big.NewInt(y)
	return new(big.Int).Mod(bx, by).Int64()
}

func (d *deck) Deal() {
	d.increment *= -1
	d.increment = posMod(d.increment, d.cards)
	d.offset += d.increment
	d.offset = posMod(d.offset, d.cards)
}

func (d *deck) Cut(n int64) {
	d.offset += n * d.increment
	d.offset = posMod(d.offset, d.cards)
}

func (d *deck) DealInc(inc int64) {
	mp := modPow(inc, d.cards-2, d.cards)
	r := big.NewInt(0).Mul(big.NewInt(d.increment), big.NewInt(mp))
	r = r.Mod(r, big.NewInt(d.cards))
	d.increment = r.Int64()
}

func main() {

	lines := util.GetFileStrings("2019/Day22/input")

	cards := util.RangeInt(10007)

	for _, l := range lines {
		if l[0:3] == "cut" {
			n := util.MustAtoi(l[4:])
			cards = cut(cards, n)
		} else if l[0:6] == "deal w" {
			cards = dealInc(cards, util.MustAtoi(l[len("deal with increment "):]))
		} else if l[0:6] == "deal i" {
			cards = dealNew(cards)
		} else if l != "" {
			panic("Bad instruction: " + l)
		}
	}
	for i, c := range cards {
		if c == 2019 {
			fmt.Printf("Position of card 2019 (part1): %d\n", i)
			break
		}
	}

	ncards := int64(119315717514047)
	repeats := int64(101741582076661)
	d := deck{offset: 0, increment: 1, cards: ncards}
	for _, l := range lines {
		if l[0:3] == "cut" {
			n := int64(util.MustAtoi(l[4:]))
			d.Cut(n)
		} else if l[0:6] == "deal w" {
			d.DealInc(int64(util.MustAtoi(l[len("deal with increment "):])))
		} else if l[0:6] == "deal i" {
			d.Deal()
		} else if l != "" {
			panic("Bad instruction: " + l)
		} else if l == "" {
			fmt.Printf("empty\n")
		}
	}

	increment := modPow(d.increment, repeats, d.cards)
	offset_diff := big.NewInt(d.offset)
	one_minus_inc := big.NewInt(1 - increment)
	one_minus_inc_mul := big.NewInt(1 - d.increment)
	one_minus_inc_mul_mod := new(big.Int).Mod(one_minus_inc_mul, big.NewInt(ncards))

	inv := new(big.Int).Exp(one_minus_inc_mul_mod, big.NewInt(ncards-2), big.NewInt(ncards))

	r := big.NewInt(0)
	r = r.Mul(offset_diff, one_minus_inc)
	r = r.Mul(r, inv)
	r = r.Mod(r, big.NewInt(ncards))

	fmt.Printf("Number on card at position 2020 (part2): %d\n", (r.Int64()+increment*2020)%ncards)
}
