package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Gate string

const (
	Or  Gate = "OR"
	And Gate = "AND"
	Xor Gate = "XOR"
)

func (g Gate) Output(v1, v2 bool) bool {
	switch g {
	case Or:
		return v1 || v2
	case And:
		return v1 && v2
	case Xor:
		return v1 != v2
	}

	panic("bad gate type")
}

type conn struct {
	l, r, o string
	g       Gate
}

func (c conn) String() string {
	return fmt.Sprintf("%s %s %s -> %s", c.l, c.g, c.r, c.o)
}

func wireDec(wireValues map[string]bool, wire rune) int64 {
	max := 0
	for k := range wireValues {
		if rune(k[0]) == wire {
			if zz := util.MustAtoi(k[1:]); zz > max {
				max = zz
			}
		}
	}

	zv := make([]string, max+1)
	for i := range zv {
		zv[i] = "0"
	}
	for k, v := range wireValues {
		if rune(k[0]) == wire {
			idx := max - util.MustAtoi(k[1:])
			if v {
				zv[idx] = "1"
			}
		}
	}

	v, _ := strconv.ParseInt(strings.Join(zv, ""), 2, 64)
	return v
}

func gen(wireValues map[string]bool, conns []conn) int64 {
	queue := []string{}
	for _, c := range conns {
		if _, ok := wireValues[c.l]; !ok {
			queue = append(queue, c.l)
		}
		if _, ok := wireValues[c.r]; !ok {
			queue = append(queue, c.r)
		}
		if _, ok := wireValues[c.o]; !ok {
			queue = append(queue, c.o)
		}
	}

	for len(queue) > 0 {
		for _, c := range conns {
			if v1, ok := wireValues[c.l]; ok {
				if v2, ok := wireValues[c.r]; ok {
					wireValues[c.o] = c.g.Output(v1, v2)
					queue, _ = util.RemoveByValue(queue, c.o)
				}
			}
		}
	}

	return wireDec(wireValues, 'z')
}

func set(wireValues map[string]bool, width, v int64, wire rune) {
	s := strconv.FormatInt(int64(v), 2)

	for i := int64(0); i < width; i++ {
		key := fmt.Sprintf("%c%02d", wire, i)
		sid := int64(len(s)) - i - 1
		if sid >= 0 {
			wireValues[key] = s[sid] == '1'
		} else {
			wireValues[key] = false
		}
	}
}

func main() {

	input := "2024/Day24/input_fixed"
	wireValues := map[string]bool{}

	ii := 0
	for i, l := range util.GetFileStrings(input) {
		if l == "" {
			ii = i
			break
		}

		wireValues[strings.Split(l, ":")[0]] = util.MustAtoi(strings.TrimSpace(strings.Split(l, ":")[1])) == 1
	}

	conns := []conn{}
	for _, l := range util.GetFileStrings(input)[ii+1:] {
		if l == "" {
			continue
		}
		var c conn
		fmt.Sscanf(l, "%s %s %s -> %s", &c.l, &c.g, &c.r, &c.o)
		conns = append(conns, c)
	}

	slices.SortFunc(conns, func(lhs, rhs conn) int {
		if strings.Compare(lhs.o, rhs.o) == 0 {
			if strings.Compare(lhs.l, rhs.l) == 0 {
				return strings.Compare(lhs.r, rhs.r)
			}
			return strings.Compare(lhs.l, rhs.l)
		}
		return strings.Compare(lhs.o, rhs.o)
	})

	// fmt.Printf("Value %d+%d on z wires (part1): %d\n", wireDec(wireValues, 'x'), wireDec(wireValues, 'y'), gen(wireValues, conns))

	for i := int64(0); i < 45; i++ {
		tcs := []struct{ x, y, s int64 }{
			{1<<i - 1, 1, 1 << i},
			{1 << i, 0, 1 << i},
			{0, 1 << i, 1 << i},
		}
		for _, tc := range tcs {
			wireValues = map[string]bool{}
			// v := int64(1 << (i))
			set(wireValues, 45, tc.x, 'x')
			set(wireValues, 45, tc.y, 'y')
			if g := gen(wireValues, conns); g != tc.s {
				fmt.Printf("invalid value for i=%d, tc = %v, v = %d, got = %d\n", i, tc, tc.s, g)
				break
			}
		}

		// wireValues = map[string]bool{}
		// set(wireValues, 45, v, 'x')
		// set(wireValues, 45, 0, 'y')
		// if g := gen(wireValues, conns); g != v {
		// 	fmt.Printf("invalid value for i = %d, v = %d, got = %d\n", i, v, g)
		// 	break
		// }

		// wireValues = map[string]bool{}
		// set(wireValues, 45, v, 'x')
		// set(wireValues, 45, v, 'y')
		// if g := gen(wireValues, conns); g != v+v {
		// 	fmt.Printf("invalid value for i = %d, v = %d, got = %d\n", i, v, g)
		// 	break
		// }
	}

	var sb strings.Builder
	sb.WriteString("digraph {\n")
	for _, c := range conns {
		fmt.Fprintf(&sb, "  %s -> %s [label=%s]\n", c.r, c.o, c.g[:1])
		fmt.Fprintf(&sb, "  %s -> %s [label=%s]\n", c.l, c.o, c.g[:1])
	}
	sb.WriteString("}")

	// var sb strings.Builder
	// sb.WriteString("digraph {\n")
	// gates := map[string]string{}
	// for _, c := range conns {
	// 	n := rand.Intn(65535)
	// 	g := fmt.Sprintf("%s_%d", c.g, n)
	// 	gates[c.o] = g
	// 	fmt.Fprintf(&sb, "  %s [label=\"%s (%s)\"]\n", g, c.g, c.o)

	// 	cl := c.l
	// 	if v, ok := gates[c.l]; ok {
	// 		cl = v
	// 	}
	// 	cr := c.r
	// 	if v, ok := gates[c.r]; ok {
	// 		cr = v
	// 	}
	// 	fmt.Fprintf(&sb, "  %s -> %s\n", cr, g)
	// 	fmt.Fprintf(&sb, "  %s -> %s\n", cl, g)
	// }
	// sb.WriteString("}")

	f, _ := os.Create("graph.dot")
	f.WriteString(sb.String())

}
