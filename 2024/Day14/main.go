package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Robot struct {
	x, y, vx, vy int
}

func safetyFactor(robots []Robot, w, h int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		if r.x < w/2 && r.y < h/2 {
			q1++
		} else if r.x > w/2 && r.y < h/2 {
			q2++
		} else if r.x < w/2 && r.y > h/2 {
			q3++
		} else if r.x > w/2 && r.y > h/2 {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func main() {
	lines := util.GetFileStrings("2024/Day14/input")

	robots := []Robot{}
	w, h, seconds := 101, 103, 6888

	for _, l := range lines {
		var r Robot
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy)
		robots = append(robots, r)
	}

	for i := 0; i < seconds; i++ {
		for r := range robots {
			robots[r].x = (robots[r].x + robots[r].vx) % w
			if robots[r].x < 0 {
				robots[r].x += w
			}
			robots[r].y = (robots[r].y + robots[r].vy) % h
			if robots[r].y < 0 {
				robots[r].y += h
			}
		}
		if i == 99 {
			fmt.Printf("Safety factor (part1): %d\n", safetyFactor(robots, w, h))
		}
	}
	fmt.Printf("Christmas tree after seconds (part2): %d\n", seconds)
	printRobots(robots)

}

var memo = map[string]bool{}

func printRobots(robots []Robot) bool {

	m := make([][]bool, 103)
	for i := range m {
		m[i] = make([]bool, 101)
	}

	for _, r := range robots {
		m[r.y][r.x] = true
	}

	var sb strings.Builder
	for _, row := range m {
		for _, cell := range row {
			if cell {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	if memo[sb.String()] {
		fmt.Printf("!!! Repeated !!!]\n")
		return true
	}
	memo[sb.String()] = true
	fmt.Print(sb.String())
	return false
}
