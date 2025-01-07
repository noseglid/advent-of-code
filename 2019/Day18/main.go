package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"

	"github.com/noseglid/advent-of-code/util"
)

type P = util.Point

func movablefn(obtainedKeys []rune) func(int, int, rune) bool {
	return func(i1, i2 int, r rune) bool {
		if r == '#' {
			return false
		}
		for _, k := range obtainedKeys {
			if r == unicode.ToUpper(k) {
				return true
			}
		}
		return !unicode.IsUpper(r)
	}
}
func findAllKeys(grid [][]rune) []rune {
	var k []rune
	for _, row := range grid {
		for _, c := range row {
			if c >= 'a' && c <= 'z' {
				k = append(k, c)
			}
		}
	}
	return k
}

var memo, memo2 = map[string]int{}, map[string]int{}

func cacheKey(pos []P, keys []rune) string {
	var sb strings.Builder
	for _, p := range pos {
		sb.WriteString(fmt.Sprintf("%d,%d:", p.X, p.Y))
	}
	sb.WriteString(string(keys))
	return sb.String()
}

func pathContainsKey(grid util.Grid, path []P, obtainedKeys []rune) bool {
	if len(path) <= 2 {
		return false
	}
	for _, p := range path[1 : len(path)-1] {
		if r := grid[p.Y][p.X]; r >= 'a' && r <= 'z' && !util.Contains(obtainedKeys, r) {
			return true
		}
	}
	return false
}

func findKeys(grid util.Grid, pos P, keys, obtainedKeys []rune) int {
	if len(keys) == 0 {
		return 0
	}

	ckey := cacheKey([]P{pos}, keys)
	if v, ok := memo[ckey]; ok {
		return v
	}

	minSteps := math.MaxInt
	for _, k := range keys {
		ex, ey := grid.Find(k)
		ep := P{ex, ey}

		steps, path, ok := grid.ShortestPath(pos, ep, movablefn(obtainedKeys))
		if !ok || pathContainsKey(grid, path, obtainedKeys) {
			// No path to the key
			continue
		}

		nextKeys, nextObtainedKeys := util.CopySlice(keys), append(util.CopySlice(obtainedKeys), k)
		nextKeys, _ = util.RemoveByValue(nextKeys, k)
		slices.Sort(nextKeys)

		if s := steps + findKeys(grid, ep, nextKeys, nextObtainedKeys); s < minSteps {
			minSteps = s
		}
	}

	memo[ckey] = minSteps
	return minSteps
}

func findKeysP2(grid util.Grid, pos []P, keys, obtainedKeys []rune) int {
	if len(keys) == 0 {
		return 0
	}
	ckey := cacheKey(pos, keys)
	if v, ok := memo2[ckey]; ok {
		return v
	}

	minSteps := math.MaxInt

	for _, p := range pos {
		for _, k := range keys {
			ex, ey := grid.Find(k)
			ep := P{ex, ey}

			steps, path, ok := grid.ShortestPath(p, ep, movablefn(obtainedKeys))
			if !ok || pathContainsKey(grid, path, obtainedKeys) {
				continue
			}

			nextKeys, nextObtainedKeys, nextPos := util.CopySlice(keys), append(util.CopySlice(obtainedKeys), k), util.CopySlice(pos)

			nextKeys, _ = util.RemoveByValue(nextKeys, k)
			slices.Sort(nextKeys)

			nextPos, _ = util.RemoveByValue(nextPos, p)
			nextPos = append(nextPos, ep)

			if s := steps + findKeysP2(grid, nextPos, nextKeys, nextObtainedKeys); s < minSteps {
				minSteps = s
			}
		}
	}

	memo2[ckey] = minSteps
	return minSteps
}

func main() {
	grid := util.Grid(util.GetFileRuneGrid("2019/Day18/input"))

	allKeys, obtainedKeys := findAllKeys(grid), []rune{}
	sx, sy := grid.Find('@')

	minSteps := findKeys(grid, P{sx, sy}, allKeys, obtainedKeys)
	fmt.Printf("minimum steps required to pick all keys (part1): %d\n", minSteps)

	grid[sy][sx] = '#'
	grid[sy][sx-1] = '#'
	grid[sy][sx+1] = '#'
	grid[sy-1][sx] = '#'
	grid[sy+1][sx] = '#'
	grid[sy-1][sx-1] = '@'
	grid[sy-1][sx+1] = '@'
	grid[sy+1][sx-1] = '@'
	grid[sy+1][sx+1] = '@'

	pos := []P{{sx - 1, sy - 1}, {sx + 1, sy - 1}, {sx - 1, sy + 1}, {sx + 1, sy + 1}}
	minStepsp2 := findKeysP2(grid, pos, allKeys, obtainedKeys)

	fmt.Printf("minimum steps in split maze (part2): %d\n", minStepsp2)

}
