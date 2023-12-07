package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type HandStrength int

const (
	HighCard HandStrength = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardValue = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type entry struct {
	Hand       string
	SortedHand string
	Strength   HandStrength
	StrengthP2 HandStrength
	Bid        int
}

func jokerModified(cards []rune) []rune {
	most := map[rune]int{}
	for _, c := range cards {
		most[c] = most[c] + 1
	}

	var rr rune
	count := 0
	for r, c := range most {
		if c > count && r != 'J' {
			rr = r
			count = c
		}
	}

	var result []rune
	for _, c := range cards {
		if c == 'J' {
			result = append(result, rr)
		} else {
			result = append(result, c)
		}
	}

	return result
}

func isFiveOfAKind(cards []rune) (HandStrength, bool) {
	return FiveOfAKind, cards[0] == cards[4]
}
func isFourOfAKind(cards []rune) (HandStrength, bool) {
	return FourOfAKind, cards[0] == cards[3] || // XXXX_
		cards[1] == cards[4] // _XXXX
}
func isFullHouse(cards []rune) (HandStrength, bool) {
	return FullHouse, (cards[0] == cards[1] && cards[2] == cards[4]) || // XXYYY
		cards[0] == cards[2] && cards[3] == cards[4] // XXXYY
}
func isThreeOfAKind(cards []rune) (HandStrength, bool) {
	return ThreeOfAKind, cards[0] == cards[2] || // XXX__
		cards[1] == cards[3] || // _XXX_
		cards[2] == cards[4] // __XXX
}
func isTwoPair(cards []rune) (HandStrength, bool) {
	return TwoPair, (cards[0] == cards[1] && cards[2] == cards[3]) || // XXYY_
		(cards[0] == cards[1] && cards[3] == cards[4]) || // XX_YY
		(cards[1] == cards[2] && cards[3] == cards[4]) // _XXYY
}
func isOnePair(cards []rune) (HandStrength, bool) {
	return OnePair, cards[0] == cards[1] || // XX___
		cards[1] == cards[2] || // _XX__
		cards[2] == cards[3] || // __XX_
		cards[3] == cards[4] // ___XX
}
func isHighCard(cards []rune) (HandStrength, bool) {
	return HighCard, true
}

type Checker func(cards []rune) (HandStrength, bool)

var checkers = []Checker{
	isFiveOfAKind,
	isFourOfAKind,
	isFullHouse,
	isThreeOfAKind,
	isTwoPair,
	isOnePair,
	isHighCard,
}

func handStrength(hand string) (HandStrength, string) {
	cards := []rune(strings.Clone(hand))
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] < cards[j]
	})

	for _, c := range checkers {
		if strength, ok := c(cards); ok {
			return strength, string(cards)
		}
	}

	panic("impossible!")
}

func parse(input []string) []entry {
	var entries []entry
	for _, s := range input {
		if s == "" {
			continue
		}
		spl := strings.Split(s, " ")
		n, _ := strconv.Atoi(spl[1])
		entries = append(entries, entry{
			Hand: spl[0],
			Bid:  n,
		})
	}
	return entries
}

func entryless(rhs, lhs entry, strprovider func(entry) HandStrength, valprovider func(rune) int) bool {
	if strprovider(rhs) != strprovider(lhs) {
		return strprovider(rhs) < strprovider(lhs)
	}

	for i := range rhs.Hand {
		v1, v2 := valprovider(rune(rhs.Hand[i])), valprovider(rune(lhs.Hand[i]))
		if v1 == v2 {
			continue
		}
		return v1 < v2
	}

	panic("totally equal undefined order")
}

func main() {
	entries := parse(util.GetFileStrings("2023/Day7/input"))
	var entriesp2 []entry
	for i, e := range entries {
		entries[i].Strength, entries[i].SortedHand = handStrength(e.Hand)
		jm := jokerModified([]rune(e.Hand))
		entries[i].StrengthP2, _ = handStrength(string(jm))
		entriesp2 = append(entriesp2, entries[i])
	}

	sort.Slice(entries, func(i, j int) bool {
		strprovider := func(e entry) HandStrength { return e.Strength }
		valprovider := func(r rune) int { return cardValue[r] }
		return entryless(entries[i], entries[j], strprovider, valprovider)
	})
	sort.Slice(entriesp2, func(i, j int) bool {
		strprovider := func(e entry) HandStrength { return e.StrengthP2 }
		valprovider := func(r rune) int {
			if r == 'J' {
				return 1
			}
			return cardValue[r]
		}
		return entryless(entriesp2[i], entriesp2[j], strprovider, valprovider)
	})

	total := 0
	for i, e := range entries {
		total += (i + 1) * e.Bid
	}

	totalp2 := 0
	for i, e := range entriesp2 {
		totalp2 += (i + 1) * e.Bid
	}

	fmt.Printf("Sum of rank * bid (part1): %d\n", total)
	fmt.Printf("Sum of rank * bid (part2): %d\n", totalp2)

}
