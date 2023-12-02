package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

var mostCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines := util.GetFileStrings("2023/Day2/input")

	sum := 0
	minSum := 0
	for _, game := range lines {
		s1 := strings.Split(game, ":")
		gameID, _ := strconv.Atoi(strings.TrimPrefix(s1[0], "Game "))
		gameOK := true

		minResult := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, pp := range strings.Split(s1[1], ";") {
			result := map[string]int{}
			for _, g := range strings.Split(pp, ",") {
				var count int
				var tt string
				if _, err := fmt.Sscanf(g, "%d %s", &count, &tt); err != nil {
					panic(err)
				}
				result[tt] += count
			}

			for k, v := range result {
				if v > mostCubes[k] {
					gameOK = false
				}
				if v > minResult[k] {
					minResult[k] = v
				}
			}
		}

		if gameOK {
			sum += gameID
		}

		product := 1
		for _, v := range minResult {
			product *= v
		}
		minSum += product
	}

	fmt.Printf("Sum possible games (part1): %d\n", sum)
	fmt.Printf("Sum of min games power (part2): %d\n", minSum)

}
