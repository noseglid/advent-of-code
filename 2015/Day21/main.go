package main

import (
	"log"
	"math"
	"sort"
)

type char struct {
	hp     int
	damage int
	armor  int
}

type item struct {
	cost   int
	damage int
	armor  int
}

func itemEffects(items []item) (int, int) {
	dmg, arm := 0, 0
	for _, item := range items {
		dmg += item.damage
		arm += item.armor
	}
	return dmg, arm
}

func itemCosts(items []item) int {
	c := 0
	for _, item := range items {
		c += item.cost
	}
	return c
}

func battle(boss, player char, items []item) bool {
	playerPlusDamage, playerPlusArmor := itemEffects(items)
	bossDeal := boss.damage - (player.armor + playerPlusArmor)
	if bossDeal <= 0 {
		bossDeal = 1
	}
	playerDeal := player.damage + playerPlusDamage - boss.armor
	if playerDeal <= 0 {
		playerDeal = 1
	}
	return math.Ceil(float64(boss.hp)/float64(playerDeal)) <= math.Ceil(float64(player.hp)/float64(bossDeal))
}

func main() {

	boss := char{104, 8, 1}
	player := char{100, 0, 0}

	weapons := []*item{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
	armors := []*item{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}

	rings := []*item{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	var winCosts []int
	var looseCosts []int

	for _, w := range weapons {
		for _, a := range armors {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1 == r2 {
						// Cannot use the same ring
						continue
					}

					items := []item{*w, *a, *r1, *r2}
					cost := itemCosts(items)
					if battle(boss, player, items) {
						winCosts = append(winCosts, cost)
					} else {
						looseCosts = append(looseCosts, cost)
					}
				}
			}
		}
	}

	sort.Ints(winCosts)
	sort.Sort(sort.Reverse(sort.IntSlice(looseCosts)))

	log.Printf("least gold (part1): %d", winCosts[0])
	log.Printf("most gold (part2): %d", looseCosts[0])
}
