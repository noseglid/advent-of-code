package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Pulse int8

const (
	PulseLow Pulse = iota
	PulseHigh
)

func (p Pulse) String() string {
	switch p {
	case PulseLow:
		return "-low-"
	case PulseHigh:
		return "-high-"
	default:
		panic("bad signal")
	}
}

func (p Pulse) StringShort() string {
	switch p {
	case PulseLow:
		return "L"
	case PulseHigh:
		return "H"
	default:
		panic("bad signal")
	}
}

type Signal struct {
	source string
	target string
	pulse  Pulse
}

func (s Signal) String() string {
	return fmt.Sprintf("%s -> %s -> %s", s.source, s.pulse, s.target)
}

type Module interface {
	Reset()
	Name() string
	Outputs() []string
	RecvSignal(pd Signal) []Signal
}

type Broadcast struct {
	name    string
	outputs []string
}

func (b *Broadcast) Reset() {
}

func (b Broadcast) Name() string {
	return b.name
}

func (b *Broadcast) Outputs() []string {
	return b.outputs
}

func (b *Broadcast) RecvSignal(s Signal) []Signal {
	var signals []Signal
	for _, o := range b.outputs {
		signals = append(signals, Signal{
			source: b.name,
			target: o,
			pulse:  s.pulse,
		})
	}
	return signals
}

type FlipFlop struct {
	name    string
	on      bool
	outputs []string
}

func (f *FlipFlop) Reset() {
	f.on = false
}

func (f *FlipFlop) Name() string {
	return f.name
}

func (f *FlipFlop) Outputs() []string {
	return f.outputs
}

func (f *FlipFlop) RecvSignal(s Signal) []Signal {
	if s.pulse == PulseHigh {
		return []Signal{}
	}

	outputPulse := PulseHigh
	if f.on {
		outputPulse = PulseLow
	}

	f.on = !f.on

	var signals []Signal
	for _, o := range f.outputs {
		signals = append(signals, Signal{source: f.name, target: o, pulse: outputPulse})
	}
	return signals
}

type Conjunction struct {
	name    string
	memory  map[string]Pulse
	outputs []string
}

func (c *Conjunction) Reset() {
	for m := range c.memory {
		c.memory[m] = PulseLow
	}
}

func (c Conjunction) Name() string {
	return c.name
}

func (c *Conjunction) Outputs() []string {
	return c.outputs
}

func (c Conjunction) State() string {
	var a []string
	for m := range c.memory {
		a = append(a, m)
	}
	sort.Strings(a)
	var sb strings.Builder
	for _, m := range a {
		sb.WriteString(c.memory[m].StringShort())
	}
	return sb.String()
}

func (c *Conjunction) RecvSignal(s Signal) []Signal {
	c.memory[s.source] = s.pulse
	outPulse := PulseLow
	mm := ""
	for _, p := range c.memory {
		if p == PulseLow {
			mm += "L"
			outPulse = PulseHigh
		} else {
			mm += "H"
		}
	}
	var signals []Signal
	for _, o := range c.outputs {
		signals = append(signals, Signal{
			source: c.name,
			target: o,
			pulse:  outPulse,
		})
	}

	return signals
}

func parseLine(l string) Module {
	spl := strings.SplitN(l, " ", 3)
	var m Module
	switch spl[0][0] {
	case 'b':
		return &Broadcast{
			name:    spl[0],
			outputs: strings.Split(spl[2], ", "),
		}
	case '&':
		return &Conjunction{
			name:    spl[0][1:],
			outputs: strings.Split(spl[2], ", "),
			memory:  make(map[string]Pulse),
		}
	case '%':
		return &FlipFlop{
			name:    spl[0][1:],
			outputs: strings.Split(spl[2], ", "),
		}
	}
	return m
}

func pushButton(modules map[string]Module) (int, int) {
	signals := []Signal{
		{target: "broadcaster", pulse: PulseLow, source: "button"},
	}
	// fmt.Println(signals[0])

	high, low := 0, 1
	for len(signals) > 0 {
		s := signals[0]
		signals = signals[1:]
		module, ok := modules[s.target]
		if !ok {
			// output only module, do nothing
			continue
		}
		outputs := module.RecvSignal(s)
		for _, o := range outputs {
			if o.pulse == PulseHigh {
				high++
			} else {
				low++
			}
			// fmt.Println(o)
		}
		// slices.Reverse(outputs)
		signals = append(signals, outputs...)
	}

	return high, low
}

func state(modules map[string]Module, checks ...string) map[string]string {
	states := map[string]string{}
	for _, m := range checks {
		if c, ok := modules[m].(*Conjunction); ok {
			states[c.name] = c.State()

		} else {
			panic("not conjunction: " + m)
		}
	}
	return states
}

func oneL(state string) bool {
	l := 0
	for _, r := range state {
		if r == 'L' {
			l++
			if l == 2 {
				return false
			}
		}
	}
	return true

}

func main() {

	lines := util.GetFileStrings("2023/Day20/input")

	modules := map[string]Module{}
	for _, l := range lines {
		m := parseLine(l)
		modules[m.Name()] = m
	}

	for _, m := range modules {
		if conj, ok := m.(*Conjunction); ok {
			for _, im := range modules {
				if slices.Contains(im.Outputs(), conj.Name()) {
					conj.memory[im.Name()] = PulseLow
				}
			}
		}
	}

	/*
	              +-rd <- hc
	              |
	              +-bt <- qt
	              |
	   rx <- vd <-+
	              |
	              +-fv <- ck
	              |
	              +-pr <- kb
	*/

	high, low := 0, 0
	for i := 0; i < 1000; i++ {
		nh, nl := pushButton(modules)
		high += nh
		low += nl
	}
	fmt.Printf("Total pulses multiplied (part1) %d*%d: %d\n", low, high, high*low)

	for _, m := range modules {
		m.Reset()
	}
	chc, cqt, cck, ckb := 0, 0, 0, 0
	mhc, mqt, mck, mkb := false, false, false, false
	for i := 0; i < 10000000; i++ {
		pushButton(modules)
		if i <= 10 {
			continue
		}
		states := state(modules, "hc", "qt", "ck", "kb")
		if mhc && states["hc"] == "LLLLLLLL" {
			chc = i + 1
		}
		mhc = chc == 0 && oneL(states["hc"])

		if mqt && states["qt"] == "LLLLLLLL" {
			cqt = i + 1
		}
		mqt = cqt == 0 && oneL(states["qt"])

		if mck && states["ck"] == "LLLLLLLL" {
			cck = i + 1
		}
		mck = cck == 0 && oneL(states["ck"])

		if mkb && states["kb"] == "LLLLLLL" {
			ckb = i + 1
		}
		mkb = ckb == 0 && oneL(states["kb"])

		if chc != 0 && cqt != 0 && cck != 0 && ckb != 0 {
			break
		}
	}

	fmt.Printf("rx gets one low at iteration (part2): %d\n", util.LCM(chc, cqt, cck, ckb))

}
