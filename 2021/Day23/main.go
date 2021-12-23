package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func roomHallwayX(roomId int) int {
	switch roomId {
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 6
	case 4:
		return 8
	}
	panic("bad room id")
}

func hallwayCanMove(a, b int, hallway [11]amphipod) (int, bool) {
	// fmt.Printf("testing moving from %d to %d\n", a, b)
	dx := 1
	if a > b {
		dx = -1
	}
	for tx := a + dx; tx != b+dx; tx += dx {
		if hallway[tx].ok {
			// occupied, can't move to this room
			return 0, false
		}
	}
	return util.Absolute(b - a), true

}

var amphiEmpty = amphipod{false, 0, 0}
var amphiA = amphipod{true, 1, 1}
var amphiB = amphipod{true, 10, 2}
var amphiC = amphipod{true, 100, 3}
var amphiD = amphipod{true, 1000, 4}

type amphipod struct {
	ok         bool
	stepCost   int
	targetRoom int
}

func (a amphipod) String() string {
	if !a.ok {
		return "."
	}
	switch a.targetRoom {
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	}
	panic("bad target room")
}

type room struct {
	id        int
	amphipods []amphipod
}

func (r room) canAccept(a amphipod) bool {
	if !a.ok {
		return false
	}
	if a.targetRoom != r.id {
		return false
	}
	for i := range r.amphipods {
		if !r.amphipods[i].ok {
			continue
		}
		if r.amphipods[i].targetRoom == r.id {
			continue
		}
		return false
	}

	return true
}

func (r *room) add(a amphipod, hallwayX int) int {
	xdist := util.Absolute(roomHallwayX(r.id) - hallwayX)
	for i := len(r.amphipods) - 1; i >= 0; i-- {
		if !r.amphipods[i].ok {
			r.amphipods[i] = a
			return (xdist + i + 1) * a.stepCost
		}
	}

	panic("cant add to room")
}

func (r room) String() string {
	var sb strings.Builder
	var ss []string
	for _, a := range r.amphipods {
		ss = append(ss, a.String())
	}
	sb.WriteString(strings.Join(ss, ""))
	return sb.String()
}

func (r room) done() bool {
	for _, a := range r.amphipods {
		if !a.ok {
			return false
		}
		if a.targetRoom != r.id {
			return false
		}
	}

	return true
}

func (r room) existingIsCorrect() bool {
	for _, a := range r.amphipods {
		if !a.ok {
			continue
		}
		if a.targetRoom != r.id {
			return false
		}
	}

	return true
}

func (r room) clone() room {
	return room{
		id:        r.id,
		amphipods: append([]amphipod{}, r.amphipods...),
	}
}

func (r room) amphiToHallway() (int, int, bool) {
	if r.existingIsCorrect() {
		return 0, 0, false
	}

	for i, a := range r.amphipods {
		if !a.ok {
			continue
		}

		return i, (i + 1) * a.stepCost, true
	}

	return 0, 0, false
}

type burrow struct {
	hallway [11]amphipod
	rooms   [4]room
	cost    int
}

type cacheEntry struct {
	cost    uint32
	hallway uint32
	rooms   uint32
}

func (b burrow) cacheEntry() cacheEntry {
	var ce cacheEntry

	ce.cost = uint32(b.cost)
	ce.hallway |= uint32(b.hallway[0].targetRoom << 0)
	ce.hallway |= uint32(b.hallway[1].targetRoom << 2)
	ce.hallway |= uint32(b.hallway[2].targetRoom << 4)
	ce.hallway |= uint32(b.hallway[3].targetRoom << 6)
	ce.hallway |= uint32(b.hallway[4].targetRoom << 8)
	ce.hallway |= uint32(b.hallway[5].targetRoom << 10)
	ce.hallway |= uint32(b.hallway[6].targetRoom << 12)
	ce.hallway |= uint32(b.hallway[7].targetRoom << 14)
	ce.hallway |= uint32(b.hallway[8].targetRoom << 16)
	ce.hallway |= uint32(b.hallway[9].targetRoom << 18)
	ce.hallway |= uint32(b.hallway[10].targetRoom << 20)

	for i := 0; i < len(b.rooms[0].amphipods); i++ {
		ce.rooms |= uint32((b.rooms[0].amphipods[i].targetRoom)) << (30 - i*2)
		ce.rooms |= uint32((b.rooms[0].amphipods[i].targetRoom)) << (22 - i*2)
		ce.rooms |= uint32((b.rooms[0].amphipods[i].targetRoom)) << (14 - i*2)
		ce.rooms |= uint32((b.rooms[0].amphipods[i].targetRoom)) << (6 - i*2)
	}

	return ce
}

var cache map[cacheEntry]int

func (b burrow) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Burrow (cost=%d)\n", b.cost))
	sb.WriteString("#############\n")
	sb.WriteString("#")
	sb.WriteString(b.hallway[0].String())
	sb.WriteString(b.hallway[1].String())
	sb.WriteString(b.hallway[2].String())
	sb.WriteString(b.hallway[3].String())
	sb.WriteString(b.hallway[4].String())
	sb.WriteString(b.hallway[5].String())
	sb.WriteString(b.hallway[6].String())
	sb.WriteString(b.hallway[7].String())
	sb.WriteString(b.hallway[8].String())
	sb.WriteString(b.hallway[9].String())
	sb.WriteString(b.hallway[10].String())
	sb.WriteString("#\n")
	sb.WriteString("##")
	for i := 0; i < len(b.rooms[0].amphipods); i++ {
		sb.WriteString("#")
		for _, r := range b.rooms {
			sb.WriteString(fmt.Sprintf("%s#", r.amphipods[i]))
		}
		if i == 0 {
			sb.WriteString("##")
		}
		sb.WriteString("\n  ")
	}
	sb.WriteString("#########\n")
	return sb.String()
}

func (b burrow) done() bool {
	for _, r := range b.rooms {
		if !r.done() {
			return false
		}
	}
	return true
}

func (b burrow) clone() burrow {
	var hh [11]amphipod
	for i, a := range b.hallway {
		hh[i] = a
	}

	var rr [4]room
	for i, r := range b.rooms {
		rr[i] = r.clone()
	}

	return burrow{
		hallway: hh,
		rooms:   rr,
	}
}

var globalMin int = math.MaxInt

func minCost(b burrow) (int, bool) {
	if b.done() {
		return b.cost, true
	}
	if cost, ok := cache[b.cacheEntry()]; ok {
		return cost, true
	}

	candidates := []burrow{}

	for hallwayX, ha := range b.hallway {
		if !ha.ok {
			continue
		}
		for j, r := range b.rooms {
			// fmt.Printf("testing %s at %d for %s\n", ha, hallwayX, r)
			if !r.canAccept(ha) {
				// Room cannot accept amphipod
				continue
			}

			if _, ok := hallwayCanMove(hallwayX, roomHallwayX(r.id), b.hallway); !ok {
				// Cannot move to this room
				continue
			}

			nb := b.clone()
			nb.cost += b.cost + nb.rooms[j].add(ha, hallwayX)
			nb.hallway[hallwayX] = amphiEmpty
			candidates = append(candidates, nb)
		}
	}

	for ri, r := range b.rooms {
		if amphiIndex, cost, ok := r.amphiToHallway(); ok {
			hallwayX := roomHallwayX(r.id)
			for i := range b.hallway {
				if i == 2 || i == 4 || i == 6 || i == 8 {
					continue
				}
				steps, ok := hallwayCanMove(hallwayX, i, b.hallway)
				if !ok {
					continue
				}

				hallwayCost := steps * r.amphipods[amphiIndex].stepCost

				nb := b.clone()
				nb.cost += b.cost + cost + hallwayCost
				nb.hallway[i] = nb.rooms[ri].amphipods[amphiIndex]
				nb.rooms[ri].amphipods[amphiIndex] = amphiEmpty
				candidates = append(candidates, nb)
			}

		}
	}

	if len(candidates) == 0 {
		// Unsolvable
		return 0, false
	}

	min := math.MaxInt
	for _, bb := range candidates {
		if bb.cost >= globalMin {
			continue
		}
		if cost, solved := minCost(bb); solved {
			if cost < min {
				min = cost
			}
			if cost < globalMin {
				globalMin = cost
			}
		}
	}

	cache[b.cacheEntry()] = min
	return min, true
}

func parseAmphipod(char rune) amphipod {
	switch char {
	case 'A':
		return amphiA
	case 'B':
		return amphiB
	case 'C':
		return amphiC
	case 'D':
		return amphiD
	}
	panic("bad amphipod")
}

func parseBurrow(rooms []string) burrow {
	b := burrow{
		rooms: [4]room{
			{id: 1},
			{id: 2},
			{id: 3},
			{id: 4},
		},
	}
	for _, r := range rooms {
		if strings.HasPrefix(r, "  #########") {
			return b
		}

		b.rooms[0].amphipods = append(b.rooms[0].amphipods, parseAmphipod(rune(r[3])))
		b.rooms[1].amphipods = append(b.rooms[1].amphipods, parseAmphipod(rune(r[5])))
		b.rooms[2].amphipods = append(b.rooms[2].amphipods, parseAmphipod(rune(r[7])))
		b.rooms[3].amphipods = append(b.rooms[3].amphipods, parseAmphipod(rune(r[9])))
	}

	return b
}

func main() {
	base := "input"

	input1 := fmt.Sprintf("2021/Day23/%s", base)
	burrow1 := parseBurrow(util.GetFileStrings(input1)[2:])
	cache = map[cacheEntry]int{}
	globalMin = math.MaxInt
	min1, _ := minCost(burrow1)
	log.Printf("Part 1: minimum cost: %d", min1)

	input2 := fmt.Sprintf("2021/Day23/%s2", base)
	burrow2 := parseBurrow(util.GetFileStrings(input2)[2:])
	cache = map[cacheEntry]int{}
	globalMin = math.MaxInt
	min2, _ := minCost(burrow2)
	log.Printf("Part 2: minimum cost: %d", min2)
}
