package main

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

var groupRe = regexp.MustCompile(`(\d+) units each with (\d+) hit points(?: \((?:weak to ([^;)]+))?(?:; )?(?:immune to ([^)]+))?\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)`)

type System string

const (
	ImmuneSystem System = "immune"
	Infection    System = "infection"
)

type group struct {
	system     System
	id         int
	units      int
	hp         int
	weak       []string
	immune     []string
	ap         int
	apType     string
	initiative int
}

func (g group) EP() int {
	return g.units * g.ap
}

func (g group) String() string {
	return fmt.Sprintf("%s:%d", g.system, g.id)
}

// const immune = `
// 17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
// 989 units each with 1274 hit points (weak to bludgeoning, slashing; immune to fire) with an attack that does 25 slashing damage at initiative 3`

// const infection = `
// 801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
// 4485 units each with 2961 hit points (weak to fire, cold; immune to radiation) with an attack that does 12 slashing damage at initiative 4`

const immune = `
2208 units each with 6238 hit points (immune to slashing) with an attack that does 23 bludgeoning damage at initiative 20
7603 units each with 6395 hit points (weak to radiation) with an attack that does 6 cold damage at initiative 15
4859 units each with 5904 hit points (weak to fire) with an attack that does 12 cold damage at initiative 11
1608 units each with 7045 hit points (weak to fire, cold; immune to bludgeoning, radiation) with an attack that does 31 radiation damage at initiative 10
39 units each with 4208 hit points with an attack that does 903 radiation damage at initiative 7
6969 units each with 9562 hit points (immune to slashing, cold) with an attack that does 13 slashing damage at initiative 3
2483 units each with 6054 hit points (immune to fire) with an attack that does 20 cold damage at initiative 19
506 units each with 3336 hit points with an attack that does 64 radiation damage at initiative 6
2260 units each with 10174 hit points (weak to fire) with an attack that does 34 slashing damage at initiative 5
2817 units each with 9549 hit points (weak to bludgeoning; immune to cold, fire) with an attack that does 31 cold damage at initiative 2`

const infection = `
3650 units each with 25061 hit points (weak to fire, bludgeoning) with an attack that does 11 slashing damage at initiative 12
508 units each with 48731 hit points (weak to bludgeoning) with an attack that does 172 cold damage at initiative 13
724 units each with 27385 hit points with an attack that does 69 radiation damage at initiative 1
188 units each with 41786 hit points with an attack that does 416 bludgeoning damage at initiative 4
3045 units each with 36947 hit points (weak to slashing; immune to fire, bludgeoning) with an attack that does 24 slashing damage at initiative 9
7006 units each with 42545 hit points (immune to cold, slashing, fire) with an attack that does 9 fire damage at initiative 16
853 units each with 55723 hit points (weak to cold, fire) with an attack that does 114 bludgeoning damage at initiative 17
3268 units each with 43027 hit points (immune to slashing, fire) with an attack that does 25 slashing damage at initiative 8
1630 units each with 47273 hit points (weak to cold, bludgeoning) with an attack that does 57 slashing damage at initiative 14
3383 units each with 12238 hit points with an attack that does 7 radiation damage at initiative 18`

func buildGroup(s string, system System) []*group {
	var g []*group
	for i, l := range strings.Split(s, "\n")[1:] {
		m := groupRe.FindStringSubmatch(l)
		g = append(g, &group{
			system:     system,
			id:         i + 1,
			units:      util.MustAtoi(m[1]),
			hp:         util.MustAtoi(m[2]),
			weak:       strings.Split(m[3], ", "),
			immune:     strings.Split(m[4], ", "),
			ap:         util.MustAtoi(m[5]),
			apType:     m[6],
			initiative: util.MustAtoi(m[7]),
		})
	}
	return g
}

func damage(attacker, defender *group) int {
	if slices.Contains(defender.immune, attacker.apType) {
		return 0
	}
	ep := attacker.EP()
	if slices.Contains(defender.weak, attacker.apType) {
		ep *= 2
	}
	return ep
}

func SelectSort(a, b *group) int {
	if d := b.EP() - a.EP(); d == 0 {
		return b.initiative - a.initiative
	} else {
		return d
	}
}

func AttackSort(a, b *group) int {
	return b.initiative - a.initiative
}
func combinedEP(list []*group) int {
	ep := 0
	for _, g := range list {
		ep += g.EP()
	}
	return ep
}

func units(list []*group) int {
	c := 0
	for _, g := range list {
		c += g.units
	}
	return c
}

func simulate(immunes, infections []*group) (System, int) {
	for len(immunes) > 0 && len(infections) > 0 {
		selectedTargets := map[*group]*group{}
		all := append(append([]*group{}, infections...), immunes...)
		slices.SortFunc(all, SelectSort)

		for _, attacker := range all {
			targets := append([]*group{}, infections...)
			if attacker.system == Infection {
				targets = append([]*group{}, immunes...)
			}
			for t := range maps.Values(selectedTargets) {
				targets, _ = util.RemoveByValue(targets, t)
			}

			m, defgroup := 0, (*group)(nil)
			for _, defender := range targets {
				if d := damage(attacker, defender); d > m || defgroup == nil || (d == m && defgroup.EP() < defender.EP()) || (d == m && defgroup.EP() == defender.EP() && defgroup.initiative < defender.initiative) {
					m = d
					defgroup = defender
				}
			}

			if defgroup != nil && m > 0 {
				selectedTargets[attacker] = defgroup
			}
		}

		slices.SortFunc(all, AttackSort)
		didAttack := false
		for _, attacker := range all {
			defender, ok := selectedTargets[attacker]
			if !ok {
				continue
			}
			didAttack = true
			d := damage(attacker, defender)
			kills := util.Min(defender.units, d/defender.hp)
			defender.units -= kills
			if defender.units == 0 {
				immunes, _ = util.RemoveByValue(immunes, defender)
				infections, _ = util.RemoveByValue(infections, defender)
			}
		}
		if !didAttack {
			return Infection, -1
		}
	}

	if len(immunes) > 0 {
		return ImmuneSystem, units(immunes)
	} else {
		return Infection, units(infections)
	}
}

func boost(list []*group, n int) {
	for _, g := range list {
		g.ap += n
	}
}

func main() {
	immunes, infections := buildGroup(immune, ImmuneSystem), buildGroup(infection, Infection)

	_, units := simulate(immunes, infections)
	fmt.Printf("Units left for conquering army (part1): %d\n", units)

	for i := 1; i < 1000000; i++ {
		immunes, infections := buildGroup(immune, ImmuneSystem), buildGroup(infection, Infection)
		imm, inf := append([]*group{}, immunes...), append([]*group{}, infections...)
		boost(imm, i)
		winner, units := simulate(imm, inf)
		if winner == ImmuneSystem {
			fmt.Printf("Units lefter after boost, letting immune system win (part2): %d\n", units)
			break
		}
	}

}
