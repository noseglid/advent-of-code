package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type Component struct {
	p1, p2         int
	usedp1, usedp2 bool
}

func (c *Component) BlockPort(portn int) {
	if portn == 1 {
		c.usedp1 = true
	} else if portn == 2 {
		c.usedp2 = true
	} else {
		panic("BlockPort: bad portn")
	}
}

func (c Component) String() string {
	bp1, bp2 := "", ""
	if c.usedp1 {
		bp1 = "|"
	}
	if c.usedp2 {
		bp2 = "|"
	}

	return fmt.Sprintf("%s%d/%d%s", bp1, c.p1, c.p2, bp2)
}

func (c Component) CanStart() (bool, int) {
	if c.p1 == 0 {
		return true, 1
	} else if c.p2 == 0 {
		return true, 2
	}
	return false, 0
}

func (c Component) Strength() int {
	return c.p1 + c.p2
}

func (c Component) CanConnect(o Component) (bool, int, int) {
	if c.p1 == o.p1 && !c.usedp1 && !o.usedp1 {
		return true, 1, 1
	}
	if c.p1 == o.p2 && !c.usedp1 && !o.usedp2 {
		return true, 1, 2
	}
	if c.p2 == o.p1 && !c.usedp2 && !o.usedp1 {
		return true, 2, 1
	}
	if c.p2 == o.p2 && !c.usedp2 && !o.usedp2 {
		return true, 2, 2
	}
	return false, 0, 0
}

func compCopy(l []Component) []Component {
	n := make([]Component, len(l))
	copy(n, l)
	return n
}

func bridgeStrength(l []Component) int {
	s := 0
	for _, c := range l {
		s += c.Strength()
	}
	return s
}

func calcStrength(bridge, remaining []Component) (int, []Component) {
	connector := bridge[len(bridge)-1]
	m := bridgeStrength(bridge)
	nb := bridge
	for i, c := range remaining {
		canConnect, b1, b2 := connector.CanConnect(c)
		if !canConnect {
			continue
		}
		nc := connector
		nc.BlockPort(b1)
		c2 := c
		c2.BlockPort(b2)
		comp2 := compCopy(remaining)
		comp2 = util.RemoveByIndex(comp2, i)
		if s, b := calcStrength(append(bridge[:len(bridge)-1], nc, c2), comp2); s > m {
			m = s
			nb = b
		}
	}
	return m, nb
}

func calcLength(ibridge, iremaining []Component) (int, []Component) {
	bridge := compCopy(ibridge)
	remaining := compCopy(iremaining)
	connector := bridge[len(bridge)-1]
	ml, mls := len(bridge), bridgeStrength(bridge)
	nb := bridge
	for i, c := range remaining {
		canConnect, b1, b2 := connector.CanConnect(c)
		if !canConnect {
			continue
		}
		nc := connector
		nc.BlockPort(b1)
		c2 := c
		c2.BlockPort(b2)
		comp2 := compCopy(remaining)
		comp2 = util.RemoveByIndex(comp2, i)
		if l, b := calcLength(append(bridge[:len(bridge)-1], nc, c2), comp2); l == ml {
			if bridgeStrength(b) > mls {
				nb = b
			}
		} else if l > ml {
			ml = l
			nb = b
		}
	}
	return ml, nb
}

func main() {
	lines := util.GetFileStrings("2017/Day24/input")
	var components []Component
	for _, l := range lines {
		var c Component
		if _, err := fmt.Sscanf(l, "%d/%d", &c.p1, &c.p2); err != nil {
			panic(err)
		}
		components = append(components, c)
	}

	m := 0
	for i, c := range components {
		canStart, port := c.CanStart()
		if !canStart {
			continue
		}
		c2 := c
		c2.BlockPort(port)

		comp2 := compCopy(components)
		comp2 = util.RemoveByIndex(comp2, i)

		if s, _ := calcStrength([]Component{c2}, comp2); s > m {
			m = s
		}
	}
	fmt.Printf("max strength (part1): %d\n", m)

	ml, mls := 0, 0
	for i, c := range components {
		canStart, port := c.CanStart()
		if !canStart {
			continue
		}
		c2 := c
		c2.BlockPort(port)

		comp2 := compCopy(components)
		comp2 = util.RemoveByIndex(comp2, i)

		if l, b := calcLength([]Component{c2}, comp2); l > ml {
			ml = l
			mls = bridgeStrength(b)
		} else if l == ml {
			nmls := bridgeStrength(b)
			if nmls > mls {
				mls = nmls
			}
		}

	}

	fmt.Printf("strength of longest bridge (part2): %d\n", mls)
}
