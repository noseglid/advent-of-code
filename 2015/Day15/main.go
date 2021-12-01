package main

import (
	"bufio"
	"log"
	"regexp"

	"github.com/noseglid/advent-of-code/util"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

//Sprinkles: capacity 2, durability 0, flavor -2, texture 0, calories 3
// var re = regexp.MustCompile(`([[:alpha:]]+): capacity (\d+), durability (\d+), flavor (\d+), texture (\d+), calories (\d+)`)
var re = regexp.MustCompile(`([[:alpha:]]+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func parseIngredient(s string) ingredient {
	m := re.FindStringSubmatch(s)
	if len(m) != 7 {
		panic("not 7 matches")
	}

	return ingredient{
		name:       m[1],
		capacity:   util.MustAtoi(m[2]),
		durability: util.MustAtoi(m[3]),
		flavor:     util.MustAtoi(m[4]),
		texture:    util.MustAtoi(m[5]),
		calories:   util.MustAtoi(m[6]),
	}
}

func ingredientPropertyScore(ingredient ingredient, teaspoons int) (int, int, int, int) {
	return teaspoons * ingredient.capacity,
		teaspoons * ingredient.durability,
		teaspoons * ingredient.flavor,
		teaspoons * ingredient.texture
}

func main() {

	s := util.FileScanner("2015/Day15/input", bufio.ScanLines)

	var ingredients []ingredient

	for s.Scan() {
		ingredients = append(ingredients, parseIngredient(s.Text()))
	}

	highestScore := 0
	for a := 100; a >= 0; a-- {
		for b := 100 - a; b >= 0; b-- {
			for c := 100 - a - b; c >= 0; c-- {
				d := 100 - a - b - c

				c1, d1, f1, t1 := ingredientPropertyScore(ingredients[0], a)
				c2, d2, f2, t2 := ingredientPropertyScore(ingredients[1], b)
				c3, d3, f3, t3 := ingredientPropertyScore(ingredients[2], c)
				c4, d4, f4, t4 := ingredientPropertyScore(ingredients[3], d)

				if ingredients[0].calories*a+ingredients[1].calories*b+ingredients[2].calories*c+ingredients[3].calories*d != 500 {
					continue
				}
				cscore := c1 + c2 + c3 + c4
				dscore := d1 + d2 + d3 + d4
				fscore := f1 + f2 + f3 + f4
				tscore := t1 + t2 + t3 + t4
				if cscore < 0 || dscore < 0 || fscore < 0 || tscore < 0 {
					continue
				}
				total := cscore * dscore * fscore * tscore
				if total > highestScore {
					log.Printf("new high at %d,%d,%d,%d", a, b, c, d)
					highestScore = total
				}
			}
		}
	}

	log.Printf("highest score (part1): %d", highestScore)
}
