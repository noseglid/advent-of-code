package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Card struct {
	id      int
	winning []int
	numbers []int
}

func parseCard(line string) Card {
	card := Card{}
	s1 := strings.Split(line, ":")
	s2 := strings.Split(s1[1], "|")
	w := strings.Split(s2[0], " ")
	h := strings.Split(s2[1], " ")

	for _, wn := range w {
		if wn == "" {
			continue
		}
		num, _ := strconv.Atoi(wn)
		card.winning = append(card.winning, num)
	}

	for _, hn := range h {
		if hn == "" {
			continue
		}
		num, _ := strconv.Atoi(hn)
		card.numbers = append(card.numbers, num)
	}

	s3 := strings.Split(s1[0], " ")
	card.id, _ = strconv.Atoi(s3[len(s3)-1])

	return card
}

func cardScore(c Card) (int, int) {
	s := 0
	m := 0
	for _, h := range c.numbers {
		for _, w := range c.winning {
			if h == w {
				m++
				if s == 0 {
					s = 1
				} else {
					s = s * 2
				}
			}
		}
	}
	return s, m
}

func main() {
	lines := util.GetFileStrings("2023/Day4/input")

	s := 0
	cards := map[int][]Card{}
	for _, l := range lines {
		c := parseCard(l)
		cards[c.id] = append(cards[c.id], c)
		cs, _ := cardScore(c)
		s += cs
	}

	sc := 0
	for i := 1; ; i++ {
		if len(cards[i]) == 0 {
			break
		}
		sc += len(cards[i])

		_, m := cardScore(cards[i][0])
		for j := 1; j <= m; j++ {
			for f := 0; f < len(cards[i]); f++ {
				cards[i+j] = append(cards[i+j], cards[i+j][0])
			}
		}
	}

	fmt.Printf("Score (part1): %d\n", s)
	fmt.Printf("Scratchcards (part2): %d\n", sc)

}
