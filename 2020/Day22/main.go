package main

import (
	"log"
	"strconv"
	"strings"
)

// var p1 = []int{
// 	9,
// 	2,
// 	6,
// 	3,
// 	1,
// }

// var p2 = []int{
// 	5,
// 	8,
// 	4,
// 	7,
// 	10,
// }

var p1 = []int{
	2,
	31,
	14,
	45,
	33,
	18,
	29,
	36,
	44,
	47,
	38,
	6,
	9,
	5,
	48,
	17,
	50,
	41,
	4,
	21,
	42,
	23,
	25,
	28,
	3,
}

var p2 = []int{
	26,
	16,
	27,
	12,
	49,
	32,
	19,
	46,
	37,
	15,
	10,
	30,
	11,
	24,
	1,
	40,
	7,
	8,
	43,
	34,
	20,
	35,
	22,
	39,
	13,
}

func deckScore(deck []int) int {
	m := 1
	s := 0
	for i := len(deck) - 1; i >= 0; i-- {
		s += m * deck[i]
		m++
	}

	return s
}

func playCombat(p1, p2 []int) {
	r := 0
	for {
		r++
		// log.Printf("round %d", r)
		if len(p1) == 0 {
			// log.Printf("p2 won")
			break
		}
		if len(p2) == 0 {
			// log.Printf("p1 won")
			break
		}

		c1, c2 := p1[0], p2[0]
		if c1 > c2 {
			p1 = append(p1[1:], []int{c1, c2}...)
			p2 = p2[1:]
		} else {
			p2 = append(p2[1:], []int{c2, c1}...)
			p1 = p1[1:]
		}
	}

	score := 0
	if len(p1) > 0 {
		score = deckScore(p1)
	} else {
		score = deckScore(p2)
	}

	log.Printf("winning deck score (part1): %d", score)
}

func cfgString(p1, p2 []int) string {
	var s strings.Builder
	s.WriteString(deckString(p1))
	s.WriteRune('|')
	s.WriteString(deckString(p2))
	return s.String()
}

func deckString(deck []int) string {
	var s strings.Builder
	sep := ""
	for _, c := range deck {
		s.WriteString(sep)
		s.WriteString(strconv.Itoa(c))
		sep = ","
	}
	return s.String()
}

func copyDeck(deck []int) []int {
	d := make([]int, len(deck))
	copy(d, deck)
	return d
}

func playRecursiveCombat(p1, p2 []int, depth int) ([]int, bool) {
	// log.Printf("playing recursive combat with decks %v and %v", p1, p2)
	p1 = copyDeck(p1)
	p2 = copyDeck(p2)

	knownConfigs := map[string]struct{}{}
	// log.Println()
	r := 0
	for {
		r++
		// log.Printf(" == Game %d, round %d == ", depth, r)
		cfg := cfgString(p1, p2)
		if _, ok := knownConfigs[cfg]; ok {
			// log.Printf("already seen config. terminating!")
			return p1, true
		}
		knownConfigs[cfg] = struct{}{}

		if len(p1) == 0 {
			return p2, false
		} else if len(p2) == 0 {
			return p1, true
		}

		c1, c2 := p1[0], p2[0]
		// log.Printf("player1 draw %d (deck: %s)", c1, deckString(p1[1:]))
		// log.Printf("player2 draw %d (deck: %s)", c2, deckString(p2[1:]))

		if len(p1)-1 >= c1 && len(p2)-1 >= c2 {
			if _, p1win := playRecursiveCombat(p1[1:c1+1], p2[1:c2+1], depth+1); p1win {
				p1 = append(p1[1:], []int{c1, c2}...)
				p2 = p2[1:]
				// log.Printf("Player1 wins!")
			} else {
				p2 = append(p2[1:], []int{c2, c1}...)
				p1 = p1[1:]
				// log.Printf("Player2 wins!")
			}
		} else {
			if c1 > c2 {
				p1 = append(p1[1:], []int{c1, c2}...)
				p2 = p2[1:]
				// log.Printf("Player1 wins!\n")
			} else {
				p2 = append(p2[1:], []int{c2, c1}...)
				p1 = p1[1:]
				// log.Printf("Player2 wins!\n")
			}
		}
		// log.Println()
	}
}

func main() {
	p1deck := copyDeck(p1)
	p2deck := copyDeck(p2)
	p1deck2 := copyDeck(p1)
	p2deck2 := copyDeck(p2)

	playCombat(p1deck, p2deck)

	deck, _ := playRecursiveCombat(p1deck2, p2deck2, 1)

	log.Printf("winning deck: %v", deck)
	log.Printf("score recursive combat winning deck (part2): %d", deckScore(deck))

}
