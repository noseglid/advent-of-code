package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type CacheEntry struct {
	p1wins int
	p2wins int
}

var cache map[Universe]CacheEntry
var rollOutcomes []int

func init() {
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				rollOutcomes = append(rollOutcomes, a+b+c)
			}
		}
	}
	cache = make(map[Universe]CacheEntry)
}

type DeterministicDice struct {
	next  int
	rolls int
}

func NewDeterministicDice() *DeterministicDice {
	return &DeterministicDice{
		next:  1,
		rolls: 0,
	}
}

func (d *DeterministicDice) Roll() int {
	r := ((d.next-1)%100 + 1) + ((d.next)%100 + 1) + ((d.next+1)%100 + 1)
	d.next = (d.next+2)%100 + 1
	d.rolls += 3
	return r
}

type Player struct {
	id    int
	pos   int
	score int
}

func NewPlayer(def string) Player {
	p := Player{}
	if _, err := fmt.Sscanf(def, "Player %d starting position: %d", &p.id, &p.pos); err != nil {
		panic(err)
	}
	return p
}

func (p *Player) Turn(dice *DeterministicDice) {
	p.Move(dice.Roll())
}

func (p *Player) Move(n int) {
	p.pos = (p.pos+n-1)%10 + 1
	p.score += p.pos
}

func part1(p1def, p2def string) {
	p1 := NewPlayer(p1def)
	p2 := NewPlayer(p2def)

	dice := NewDeterministicDice()
	var looser Player
	for {
		p1.Turn(dice)
		if p1.score >= 1000 {
			looser = p2
			break
		}

		p2.Turn(dice)
		if p2.score >= 1000 {
			looser = p1
			break
		}
	}

	log.Printf("Part 1: Loosing score times rolled: %d", dice.rolls*looser.score)
}

type Universe struct {
	p1, p2        Player
	currentPlayer int
}

func (u Universe) Step() []Universe {
	spawned := []Universe{}

	for _, r := range rollOutcomes {
		nu := Universe{u.p1, u.p2, (u.currentPlayer + 1) % 2}
		if u.currentPlayer == 0 {
			nu.p1.Move(r)
		} else {
			nu.p2.Move(r)
		}
		spawned = append(spawned, nu)
	}

	return spawned
}

func StepUniverse(u Universe) (int, int) {
	if c, ok := cache[u]; ok {
		return c.p1wins, c.p2wins
	}

	if u.p1.score >= 21 {
		return 1, 0
	} else if u.p2.score >= 21 {
		return 0, 1
	}

	p1wins, p2wins := 0, 0
	for _, univ := range u.Step() {
		w1, w2 := StepUniverse(univ)
		cache[univ] = CacheEntry{w1, w2}
		p1wins += w1
		p2wins += w2
	}
	cache[u] = CacheEntry{p1wins, p2wins}
	return p1wins, p2wins
}

func part2(p1def, p2def string) {
	p1, p2 := NewPlayer(p1def), NewPlayer(p2def)
	u := Universe{p1, p2, 0}
	log.Printf("Part 2: Most wins: %d", util.MaxInt(StepUniverse(u)))

}

func main() {
	input := "2021/Day21/input"
	lines := util.GetFileStrings(input)
	part1(lines[0], lines[1])
	part2(lines[0], lines[1])

}
