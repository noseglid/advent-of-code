package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

const totalTime = 30

type valve struct {
	name        string
	rate        int
	connections map[string]int
}

func (v valve) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Valve %s, r=%d: ", v.name, v.rate))
	for n, c := range v.connections {
		sb.WriteString(fmt.Sprintf("(%s:%d)", n, c))
	}

	return sb.String()
}

func serialize(timeLeft int, pos string, opened []string) string {
	var sb strings.Builder
	sb.WriteString(pos)
	sb.WriteString(strconv.Itoa(timeLeft))
	sort.Slice(opened, func(i, j int) bool { return opened[i] < opened[j] })
	for _, e := range opened {
		sb.WriteString(e)
	}

	return sb.String()
}
func serialize2(timeLeft int, pos1, pos2 string, opened []string) string {
	var sb strings.Builder
	sb.WriteString(pos1)
	sb.WriteString(pos2)
	sb.WriteString(strconv.Itoa(timeLeft))
	sort.Slice(opened, func(i, j int) bool { return opened[i] < opened[j] })
	for _, e := range opened {
		sb.WriteString(e)
	}

	return sb.String()
}

func calculateMinDistanceConnections(valve valve, allValves map[string]valve, connMap map[string][]string) map[string]int {
	distanceToNode := map[string]int{}
	unvisitedNodes := []string{}
	for _, n := range allValves {
		d := math.MaxInt
		if n.name == valve.name {
			d = 0
		}
		distanceToNode[n.name] = d
		unvisitedNodes = append(unvisitedNodes, n.name)
	}
	for len(unvisitedNodes) > 0 {
		sort.Slice(unvisitedNodes, func(i, j int) bool {
			return distanceToNode[unvisitedNodes[i]] < distanceToNode[unvisitedNodes[j]]
		})
		node := unvisitedNodes[0]
		unvisitedNodes = unvisitedNodes[1:]

		for _, c := range connMap[node] {
			if distanceToNode[c] > distanceToNode[node]+1 {
				distanceToNode[c] = distanceToNode[node] + 1
			}
		}
	}

	return distanceToNode
}

func openAndTravel(valves map[string]valve, curr string, timeLeft, ratePerMin int, opened []string) int {
	if timeLeft == 0 {
		return 0
	}
	if util.Contains(opened, curr) || valves[curr].rate == 0 {
		return travelOnly(valves, curr, timeLeft, ratePerMin, opened)
	}
	released := ratePerMin
	ratePerMin += valves[curr].rate
	opened = append([]string{}, opened...)
	opened = append(opened, curr)
	return released + travelOnly(valves, curr, timeLeft-1, ratePerMin, opened)
}

func travelOnly(valves map[string]valve, curr string, timeLeft, ratePerMin int, opened []string) int {
	var candidates []int

	var conns []struct {
		d  string
		tt int
	}
	for dest, tt := range valves[curr].connections {
		conns = append(conns, struct {
			d  string
			tt int
		}{dest, tt})
	}

	sort.Slice(conns, func(i, j int) bool { return conns[i].d < conns[j].d })

	for _, conn := range conns {
		if _, ok := valves[conn.d]; !ok {
			// Not part of graph (0 rate), skip it
			continue
		}
		if util.Contains(opened, conn.d) {
			// no need to travel if it's already opened
			continue
		}
		if conn.d == curr {
			// Not traveling to self
			continue
		}
		if timeLeft-conn.tt <= 0 {
			// No time to go here but still generates, so maybe idle...
			continue
		}

		candidates = append(candidates, ratePerMin*conn.tt+rescue(valves, conn.d, timeLeft-conn.tt, ratePerMin, append([]string{}, opened...)))
	}

	if len(candidates) == 0 {
		// no travel made sense, just idle
		return timeLeft * ratePerMin
	}

	return util.Max(candidates...)
}

var seen = map[string]int{}
var seen2 = map[string]int{}

func rescue(valves map[string]valve, curr string, timeLeft, ratePerMin int, opened []string) int {
	ser := serialize(timeLeft, curr, opened)
	if d, ok := seen[ser]; ok {
		return d
	}

	open := openAndTravel(valves, curr, timeLeft, ratePerMin, opened)
	travel := travelOnly(valves, curr, timeLeft, ratePerMin, opened)
	seen[ser] = util.Max(open, travel)

	return seen[ser]
}

func main() {
	lines := util.GetFileStrings("2022/Day16/input")

	allValves := map[string]valve{}
	connMap := map[string][]string{}
	for _, l := range lines {
		v := valve{}
		if _, err := fmt.Sscanf(l, "Valve %s has flow rate=%d;", &v.name, &v.rate); err != nil {
			panic(err)
		}
		n := strings.LastIndex(l, "leads to valves ") + len("leads to valves ")
		if n == -1 {
			panic("not part of string")
		}
		connMap[v.name] = strings.Split(l[n:], ", ")
		allValves[v.name] = v
	}

	for _, v := range allValves {
		v.connections = calculateMinDistanceConnections(v, allValves, connMap)
		allValves[v.name] = v
	}
	valves := map[string]valve{}
	for _, v := range allValves {
		if v.rate > 0 || v.name == "AA" {
			valves[v.name] = v
		}
	}

	fmt.Printf("Max pressure released (part1): %d\n", rescue(valves, "AA", totalTime, 0, []string{}))

}
