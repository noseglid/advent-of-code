package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type UnitType rune

const (
	Elf    UnitType = 'E'
	Goblin UnitType = 'G'
)

type unit struct {
	t      UnitType
	x, y   int
	grid   util.Grid
	hp, ap int
}

func (u unit) String() string {
	return fmt.Sprintf("{%c;(%d,%d)HP:%d}", u.t, u.x, u.y, u.hp)
}

func (u unit) Adjacent() []util.Point {
	p := []util.Point{}
	for _, m := range []util.Dir{util.N, util.W, util.E, util.S} {
		rx, ry := u.grid.GetMove(u.x, u.y, m)
		if u.grid.Get(rx, ry) == '.' {
			p = append(p, util.Point{rx, ry})
		}
	}
	return p
}

func (u unit) AdjacentEnemies(units []*unit) []*unit {
	ret := []*unit{}

	for _, v := range units {
		if v.t == u.t {
			continue
		}

		if v.y == u.y && util.Absolute(v.x-u.x) == 1 || v.x == u.x && util.Absolute(v.y-u.y) == 1 {
			ret = append(ret, v)
		}
	}
	return ret
}

func minReadingOrder(lhs, rhs util.Point) (util.Point, bool) {
	if lhs.Y == rhs.Y {
		if lhs.X < rhs.X {
			return lhs, true
		}
		return rhs, false
	}
	if lhs.Y < rhs.Y {
		return lhs, true
	}
	return rhs, false
}

func (u *unit) Move(grid util.Grid, units []*unit) bool {
	if len(u.AdjacentEnemies(units)) > 0 {
		return false
	}
	min, target, path := math.MaxInt, util.Point{}, []util.Point{}
	for _, c := range units {
		if u.t == c.t {
			continue
		}

		for _, adj := range c.Adjacent() {
			d, p, ok := grid.ShortestPath(util.Point{u.x, u.y}, adj, func(x, y int, r rune) bool { return r == '.' })
			if !ok {
				continue
			} else if d < min {
				min = d
				target = adj
				path = p
			} else if d == min {
				itarget, nochange := minReadingOrder(target, adj)
				if !nochange {
					target = itarget
					path = p
				}
			}
		}
	}

	if len(path) == 0 {
		return false
	}
	s := path[1]
	grid.Switch(u.x, u.y, s.X, s.Y)
	u.x, u.y = s.X, s.Y

	return true
}

func (u *unit) Attack(units []*unit) bool {
	var target *unit
	for _, e := range u.AdjacentEnemies(units) {
		if target == nil || e.hp < target.hp {
			target = e
		}

		if _, nochange := minReadingOrder(util.Point{target.x, target.y}, util.Point{e.x, e.y}); e.hp == target.hp && !nochange {
			target = e
		}
	}

	if target == nil {
		return false
	}

	target.hp -= u.ap

	return true
}

func unitSort(lhs, rhs *unit) int {
	if lhs.y == rhs.y {
		return lhs.x - rhs.x
	}
	return lhs.y - rhs.y
}

func removeDead(grid util.Grid, units []*unit) []*unit {
	r := []*unit{}
	for _, u := range units {
		if u.hp > 0 {
			r = append(r, u)
			continue
		}
		grid.Set(u.x, u.y, '.')
	}

	return r
}

func isCombatOver(units []*unit) bool {
	r := units[0].t
	for _, rr := range units[1:] {
		if rr.t != r {
			return false
		}
	}
	return true
}

func p1() {
	grid := util.Grid(util.GetFileRuneGrid("2018/Day15/input"))

	var units []*unit
	grid.Each(func(x, y int, r rune) {
		if UnitType(r) == Elf || UnitType(r) == Goblin {
			units = append(units, &unit{UnitType(r), x, y, grid, 200, 3})
		}
	})

	slices.SortFunc(units, unitSort)
	round := 0

	for {
		round++
		slices.SortFunc(units, unitSort)
		actionPerformed := false
		for _, u := range units {
			if u.hp <= 0 {
				continue
			}
			actionPerformed = u.Move(grid, units) || actionPerformed
			actionPerformed = u.Attack(units) || actionPerformed
			units = removeDead(grid, units)
		}
		if isCombatOver(units) {
			break
		}
	}

	fullRounds := round - 1
	hitPoints := 0
	for _, u := range units {
		hitPoints += u.hp
	}

	fmt.Printf("Battle outcome after %d rounds and %d hitpoints left (part1): %d\n", fullRounds, hitPoints, fullRounds*hitPoints)
}
func isElfDead(units []*unit) bool {
	for _, u := range units {
		if u.t == Elf && u.hp <= 0 {
			return true
		}
	}
	return false
}

func p2() {

IncreaseAP:
	for ap := 4; ap < 1000; ap++ {
		grid := util.Grid(util.GetFileRuneGrid("2018/Day15/input"))

		var units []*unit
		grid.Each(func(x, y int, r rune) {
			if UnitType(r) == Elf || UnitType(r) == Goblin {
				iap := 3
				if UnitType(r) == Elf {
					iap = ap
				}
				units = append(units, &unit{UnitType(r), x, y, grid, 200, iap})
			}
		})

		slices.SortFunc(units, unitSort)
		round := 0

		for {
			round++
			slices.SortFunc(units, unitSort)
			actionPerformed := false
			for _, u := range units {
				if u.hp <= 0 {
					continue
				}
				actionPerformed = u.Move(grid, units) || actionPerformed
				actionPerformed = u.Attack(units) || actionPerformed
				if isElfDead(units) {
					continue IncreaseAP
				}
				units = removeDead(grid, units)
			}
			if isCombatOver(units) {
				break
			}
		}

		fullRounds := round - 1
		hitPoints := 0
		for _, u := range units {
			hitPoints += u.hp
		}

		fmt.Printf("Combat over, elves won with AP %d for outcome (part2): %d\n", ap, fullRounds*hitPoints)
		break
	}
}

func main() {
	p1()
	p2()
}
