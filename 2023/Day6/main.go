package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func dorace(t int, h int) int {
	return h * (t - h)
}

func simulate(times, distance []int) int {
	product := 1
	for race, t := range times {
		distanceToBeat := distance[race]
		winWays := 0
		for i := 0; i < t; i++ {
			distance := dorace(t, i)
			if distance > distanceToBeat {
				winWays++
			}
		}

		product *= winWays
	}
	return product
}

func main() {
	file := util.GetFileStrings("2023/Day6/input")
	times := util.NumberList(strings.TrimPrefix(file[0], "Time:"))
	distance := util.NumberList(strings.TrimPrefix(file[1], "Distance:"))

	filep2 := util.GetFileStrings("2023/Day6/inputp2")
	timesp2 := util.NumberList(strings.TrimPrefix(filep2[0], "Time:"))
	distancep2 := util.NumberList(strings.TrimPrefix(filep2[1], "Distance:"))

	fmt.Printf("Ways to win (part1): %d\n", simulate(times, distance))
	fmt.Printf("Ways to win (part2): %d\n", simulate(timesp2, distancep2))

}
