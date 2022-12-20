package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type blueprint struct {
	id                     uint64
	oreRobotCostOre        uint64
	clayRobotCostOre       uint64
	obsidianRobotCostOre   uint64
	obsidianRobotCostClay  uint64
	geodeRobotCostOre      uint64
	geodeRobotCostObsidian uint64
}

func (bp blueprint) HighestOreCost() uint64 {
	return util.Max(bp.oreRobotCostOre, bp.clayRobotCostOre, bp.obsidianRobotCostOre, bp.geodeRobotCostOre)
}
func (bp blueprint) HighestClayCost() uint64 {
	return util.Max(bp.obsidianRobotCostClay)
}
func (bp blueprint) HighestObsidianCost() uint64 {
	return util.Max(bp.geodeRobotCostObsidian)
}

type state struct {
	bp                                                 blueprint
	oreRobots, clayRobots, obsidianRobots, geodeRobots uint64
	ore, clay, obsidian, geode                         uint64
	timeLeft                                           uint64
	didIdle                                            bool
}

func (s state) String() string {
	return fmt.Sprintf(
		"OreRobots(%2d), ClayRobots(%2d), ObsidianRobots(%2d), GeodeRobots(%2d)\n%-9s(%2d), %-10s(%2d), %-14s(%2d), %-11s(%2d)",
		s.oreRobots, s.clayRobots, s.obsidianRobots, s.geodeRobots, "Ore", s.ore, "Clay", s.clay, "Obsidian", s.obsidian, "Geode", s.geode)
}

func generate(s state) state {
	ns := s
	ns.ore += s.oreRobots
	ns.clay += s.clayRobots
	ns.obsidian += s.obsidianRobots
	ns.geode += s.geodeRobots
	ns.timeLeft -= 1
	return ns
}

func idle(s state) state {
	ns := generate(s)
	ns.didIdle = true
	return ns
}

func createOreRobot(s state) state {
	ns := generate(s)
	ns.ore -= ns.bp.oreRobotCostOre
	ns.oreRobots += 1
	ns.didIdle = false
	return ns
}

func createClayRobot(s state) state {
	ns := generate(s)
	ns.ore -= ns.bp.clayRobotCostOre
	ns.clayRobots += 1
	ns.didIdle = false
	return ns
}

func createObsidianRobot(s state) state {
	ns := generate(s)
	ns.ore -= ns.bp.obsidianRobotCostOre
	ns.clay -= ns.bp.obsidianRobotCostClay
	ns.obsidianRobots += 1
	ns.didIdle = false
	return ns
}

func createGeodeRobot(s state) state {
	ns := generate(s)
	ns.ore -= ns.bp.geodeRobotCostOre
	ns.obsidian -= ns.bp.geodeRobotCostObsidian
	ns.geodeRobots += 1
	ns.didIdle = false
	return ns
}

var seen = map[state]state{}

func pickMostGeodeState(states ...state) state {
	s := states[0]
	for _, ss := range states[1:] {
		if ss.geode > s.geode {
			s = ss
		}
	}
	return s
}

func CouldBuiltGeodeLastStep(s state) bool {
	return s.didIdle && s.ore >= (s.bp.geodeRobotCostOre+s.oreRobots) && s.obsidian >= (s.bp.geodeRobotCostObsidian+s.obsidianRobots)
}

func CouldBuiltObsidianLastStep(s state) bool {
	return s.didIdle && s.ore >= (s.bp.obsidianRobotCostOre+s.oreRobots) && s.clay >= (s.bp.obsidianRobotCostClay+s.clayRobots)
}

func CouldBuiltClayLastStep(s state) bool {
	return s.didIdle && s.ore >= (s.bp.clayRobotCostOre+s.oreRobots)
}

func CouldBuiltOreLastStep(s state) bool {
	return s.didIdle && s.ore >= (s.bp.oreRobotCostOre+s.oreRobots)
}

func step(s state) state {
	if v, ok := seen[s]; ok {
		return v
	}
	if s.timeLeft == 0 {
		return s
	}

	var candidates []state
	if !CouldBuiltGeodeLastStep(s) && s.ore >= s.bp.geodeRobotCostOre && s.obsidian >= s.bp.geodeRobotCostObsidian {
		candidates = append(candidates, step(createGeodeRobot(s)))
	}

	if !CouldBuiltObsidianLastStep(s) && s.obsidianRobots < s.bp.HighestObsidianCost() && s.ore >= s.bp.obsidianRobotCostOre && s.clay >= s.bp.obsidianRobotCostClay {
		candidates = append(candidates, step(createObsidianRobot(s)))
	}

	if !CouldBuiltClayLastStep(s) && s.clayRobots < s.bp.HighestClayCost() && s.ore >= s.bp.clayRobotCostOre {
		candidates = append(candidates, step(createClayRobot(s)))
	}

	if !CouldBuiltOreLastStep(s) && s.oreRobots < s.bp.HighestOreCost() && s.ore >= s.bp.oreRobotCostOre {
		candidates = append(candidates, step(createOreRobot(s)))
	}

	candidates = append(candidates, step(idle(s)))

	r := pickMostGeodeState(candidates...)
	seen[s] = r
	return r
}

func main() {
	input := util.GetFileStrings("2022/Day19/input")

	var blueprints []blueprint
	for _, l := range input {
		bp := blueprint{}
		_, err := fmt.Sscanf(
			l,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id, &bp.oreRobotCostOre, &bp.clayRobotCostOre, &bp.obsidianRobotCostOre, &bp.obsidianRobotCostClay, &bp.geodeRobotCostOre, &bp.geodeRobotCostObsidian,
		)
		if err != nil {
			panic(err)
		}
		blueprints = append(blueprints, bp)
	}

	qualitySum := 0
	for _, bp := range blueprints {
		fmt.Printf("Starting blueprint %d ... ", bp.id)
		seen = map[state]state{}
		state := state{
			bp,
			1, 0, 0, 0,
			0, 0, 0, 0,
			24,
			false,
		}
		best := step(state)
		q := int(state.bp.id) * int(best.geode)
		qualitySum += q
		fmt.Printf("done! geodes=%d, quality=%d\n", best.geode, q)
	}
	fmt.Printf("Quality level sum (part1) %d\n", qualitySum)

	qq := 1
	for _, bp := range blueprints[0:3] {
		fmt.Printf("Starting blueprint %d ... ", bp.id)
		seen = map[state]state{}
		state := state{
			bp,
			1, 0, 0, 0,
			0, 0, 0, 0,
			32,
			false,
		}
		best := step(state)
		qq *= int(best.geode)
		fmt.Printf("done! geodes %d\n", best.geode)
	}
	fmt.Printf("Multiplied # of goedes (part2): %d\n", qq)

}
