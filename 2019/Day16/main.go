package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func pattern(n int) []int {
	var r []int
	for _, k := range []int{0, 1, 0, -1} {
		for i := 0; i < n; i++ {
			r = append(r, k)
		}
	}

	return r
}

func digitize(s string) []int {
	var r []int
	for _, c := range strings.TrimSpace(s) {
		r = append(r, int(c-'0'))
	}
	return r
}

func fft(s []int) []int {
	var next []int
	for didx := range s {
		p := pattern(didx + 1)
		sum := 0
		for ii, s := range s[didx:] {
			sum += s * p[(didx+ii+1)%len(p)]
		}
		next = append(next, util.Absolute(sum)%10)
	}
	return next
}

func fft2(s []int, offset int) []int {
	next := make([]int, len(s))
	next[len(s)-1] = s[len(s)-1]
	sum := 0
	for n := len(s) - 1; n >= offset; n-- {
		sum += s[n]
		next[n] = sum % 10
	}
	return next
}

func join(list []int) string {
	var sb strings.Builder
	for _, i := range list {
		sb.WriteRune(rune('0' + i))
	}
	return sb.String()
}

func repeat(list []int, n int) []int {
	r := make([]int, 0, len(list)*n)
	for i := 0; i < n; i++ {
		r = append(r, list...)
	}
	return r
}

func main() {

	startSignal := digitize(util.GetFile("2019/Day16/input"))

	phases := 100

	signal := make([]int, len(startSignal))
	copy(signal, startSignal)
	for i := 0; i < phases; i++ {
		signal = fft(signal)
	}

	fmt.Printf("First 8 digits after 100 phases (part1): %s\n", join(signal[0:8]))

	copy(signal, startSignal)
	offset, _ := strconv.Atoi(join(startSignal[0:7]))
	signal = repeat(signal, 10000)

	for i := 0; i < phases; i++ {
		signal = fft2(signal, offset)
	}

	fmt.Printf("First 8 digits after 100 w. repeated phases (part2): %s\n", join(signal[offset:offset+8]))
}
