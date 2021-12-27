package main

import (
	"fmt"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

type bot struct {
	id     int
	chips  []int
	lo, hi recipient
}

func (b *bot) process(bots map[int]*bot, bins map[int][]int) {
	if b.chips[0] > b.chips[1] {
		b.chips[0], b.chips[1] = b.chips[1], b.chips[0]
	}

	switch b.lo.t {
	case BotRecipient:
		bots[b.lo.id].chips = append(bots[b.lo.id].chips, b.chips[0])
	case OutputRecipient:
		bins[b.lo.id] = append(bins[b.lo.id], b.chips[0])
	}

	switch b.hi.t {
	case BotRecipient:
		bots[b.hi.id].chips = append(bots[b.hi.id].chips, b.chips[1])
	case OutputRecipient:
		bins[b.hi.id] = append(bins[b.hi.id], b.chips[1])
	}

	b.chips = b.chips[:0]
}

func (b bot) HasChips(i, j int) bool {
	if len(b.chips) < 2 {
		return false
	}

	return (b.chips[0] == i && b.chips[1] == j) || (b.chips[0] == j && b.chips[1] == i)
}

type RecipientType string

var (
	BotRecipient    RecipientType = "bot"
	OutputRecipient RecipientType = "output"
	InputRecipient  RecipientType = "input"
)

type recipient struct {
	t  RecipientType
	id int
}

func parseBot(def string) (int, int) {
	var v, id int
	if _, err := fmt.Sscanf(def, "value %d goes to bot %d", &v, &id); err != nil {
		panic(err)
	}

	return id, v
}

func parseInstruction(def string) (int, recipient, recipient) {
	var (
		bot, lo, hi  int
		loRec, hiRec string
	)
	if _, err := fmt.Sscanf(def, "bot %d gives low to %s %d and high to %s %d", &bot, &loRec, &lo, &hiRec, &hi); err != nil {
		panic(err)
	}

	return bot, recipient{RecipientType(loRec), lo}, recipient{RecipientType(hiRec), hi}

}

func main() {
	input := "2016/Day10/input"
	lines := util.GetFileStrings(input)

	bots := map[int]*bot{}
	bins := map[int][]int{}

	for _, l := range lines {
		switch l[0] {
		case 'v':
			id, v := parseBot(l)
			if eb, ok := bots[id]; ok {
				eb.chips = append(eb.chips, v)
			} else {
				bots[id] = &bot{id: id, chips: []int{v}}
			}
		case 'b':
			id, lo, hi := parseInstruction(l)
			if eb, ok := bots[id]; ok {
				eb.lo, eb.hi = lo, hi
			} else {
				bots[id] = &bot{id: id, lo: lo, hi: hi}
			}
		}
	}

	for {
		didProcess := false
		for _, b := range bots {
			if b.HasChips(61, 17) {
				log.Printf("Part 1: Bot comparing 61 and 17 chips: %d", b.id)
			}
			if len(b.chips) >= 2 {
				b.process(bots, bins)
				didProcess = true
			}
		}

		if !didProcess {
			break
		}
	}
	log.Printf("Part 2: Multiplied value: %d", bins[0][0]*bins[1][0]*bins[2][0])

}
