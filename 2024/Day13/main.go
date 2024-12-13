package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type config struct {
	ax, ay int64
	bx, by int64
	px, py int64
}

func tokens(cfgs []config) int64 {
	tokens := int64(0)
	for _, c := range cfgs {
		ta := (c.px*c.by - c.bx*c.py) / (c.ax*c.by - c.bx*c.ay)
		tb := (c.ax*c.py - c.px*c.ay) / (c.ax*c.by - c.bx*c.ay)
		if ta*c.ax+tb*c.bx == c.px && ta*c.ay+tb*c.by == c.py {
			tokens += 3*ta + tb
		}
	}

	return tokens
}

func main() {

	lines := util.GetFileStrings("2024/Day13/input")

	cfgs := []config{}
	cfgs2 := []config{}

	for i := 0; i < len(lines); i++ {
		if !strings.HasPrefix(lines[i], "Button A") {
			continue
		}
		var cfg config
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &cfg.ax, &cfg.ay)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &cfg.bx, &cfg.by)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &cfg.px, &cfg.py)
		cfgs = append(cfgs, cfg)
		cfgs2 = append(cfgs2, config{ax: cfg.ax, ay: cfg.ay, bx: cfg.bx, by: cfg.by, px: 10000000000000 + cfg.px, py: 10000000000000 + cfg.py})
	}

	fmt.Printf("Min tokens for win (part1): %d\n", tokens(cfgs))
	fmt.Printf("Min tokens for win (part2): %d\n", tokens(cfgs2))

}
