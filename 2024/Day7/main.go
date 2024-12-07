package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type op struct {
	v int
	n []int
}

func NewOp(s string) op {
	var o op
	ss := strings.Split(s, ":")
	o.v = util.MustAtoi(ss[0])
	for _, nn := range strings.Split(strings.TrimSpace(ss[1]), " ") {
		o.n = append(o.n, util.MustAtoi(nn))
	}
	return o
}

func isValid(o op, wc bool) bool {
	if len(o.n) == 1 {
		return o.v == o.n[0]
	}
	sum, mul, conc := []int{}, []int{}, []int{}
	sum = append(sum, append([]int{o.n[0] + o.n[1]}, o.n[2:]...)...)
	mul = append(mul, append([]int{o.n[0] * o.n[1]}, o.n[2:]...)...)
	conc = append(conc, append([]int{util.MustAtoi(fmt.Sprintf("%d%d", o.n[0], o.n[1]))}, o.n[2:]...)...)

	return isValid(op{o.v, sum}, wc) || isValid(op{o.v, mul}, wc) || wc && isValid(op{o.v, conc}, wc)
}

func main() {
	lines := util.GetFileStrings("2024/Day7/input")
	ops := []op{}
	for _, l := range lines {
		ops = append(ops, NewOp(l))
	}

	s, s2 := 0, 0
	for _, o := range ops {
		if isValid(o, false) {
			s += o.v
		}
		if isValid(o, true) {
			s2 += o.v
		}
	}

	fmt.Printf("Sum of valid calibrations: %d\n", s)
	fmt.Printf("Sum of valid calibrations with ||: %d\n", s2)
}
