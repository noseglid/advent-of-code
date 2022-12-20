package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
)

type entry struct {
	v int
}

func stepsAndDirection(v int, listLen int) (int, int) {
	removeSize := listLen - 1 // full lap
	nToRemove := util.Absolute(v) / removeSize
	steps := util.Absolute(v) - nToRemove*removeSize
	if v < 0 {
		return steps, -1
	} else {
		return steps, 1
	}
}

func indexCorrect(list []*entry, index int) int {
	for index < 0 {
		index += len(list)
	}
	for index >= len(list) {
		index -= len(list)
	}
	return index
}

func mixList(entries, decrypted []*entry) {
Outer:
	for _, current := range entries {
		for index, s := range decrypted {
			if s == current {
				steps, direction := stepsAndDirection(s.v, len(entries))
				moveFrom := indexCorrect(decrypted, index)
				moveTo := indexCorrect(decrypted, index+direction)
				for i := 0; i < steps; i++ {
					decrypted[moveFrom], decrypted[moveTo] = decrypted[moveTo], decrypted[moveFrom]
					moveFrom = indexCorrect(decrypted, moveFrom+direction)
					moveTo = indexCorrect(decrypted, moveTo+direction)
				}

				continue Outer
			}
		}
	}
}

func findNumbers(decrypted []*entry) (int, int, int) {
	indexOfZero := -1
	for i, c := range decrypted {
		if c.v == 0 {
			indexOfZero = i
		}
	}
	if indexOfZero == -1 {
		panic("no zero found")
	}
	i1000 := indexCorrect(decrypted, indexOfZero+1000)
	i2000 := indexCorrect(decrypted, indexOfZero+2000)
	i3000 := indexCorrect(decrypted, indexOfZero+3000)
	return decrypted[i1000].v, decrypted[i2000].v, decrypted[i3000].v
}

func main() {
	input := util.GetFileNumbers("2022/Day20/input")
	var entries, decrypted, entriesp2, decryptedp2 []*entry

	for _, n := range input {
		e := &entry{n}
		ep2 := &entry{n * 811589153}
		entries = append(entries, e)
		entriesp2 = append(entriesp2, ep2)
		decrypted = append(decrypted, e)
		decryptedp2 = append(decryptedp2, ep2)
	}

	mixList(entries, decrypted)
	v1, v2, v3 := findNumbers(decrypted)
	fmt.Printf("Grove coordinates summed (part1): %d+%d+%d =  %d\n", v1, v2, v3, v1+v2+v3)

	for i := 1; i <= 10; i++ {
		mixList(entriesp2, decryptedp2)
	}
	v11, v22, v33 := findNumbers(decryptedp2)
	fmt.Printf("Grove coordinates summed with encryption key (part2): %d\n", v11+v22+v33)

}
