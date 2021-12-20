package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type image struct {
	pixels [][]rune
	it     int
	algo   map[string]rune
}

func parseImage(rows []string, algo map[string]rune) image {
	im := image{
		pixels: make([][]rune, len(rows)),
		algo:   algo,
	}
	for y, row := range rows {
		im.pixels[y] = make([]rune, len(row))
		for x, r := range row {
			im.pixels[y][x] = r
		}
	}

	return im
}

func pixelStringToBinary(s string) string {
	var sb strings.Builder
	for _, ss := range s {
		switch ss {
		case '.':
			sb.WriteRune('0')
		case '#':
			sb.WriteRune('1')
		}
	}
	return sb.String()
}

func (i image) expand() image {
	im := image{
		pixels: make([][]rune, len(i.pixels)+2),
		it:     i.it,
		algo:   i.algo,
	}
	for y := range im.pixels {
		im.pixels[y] = make([]rune, len(i.pixels)+2)
	}
	for y := range im.pixels {
		for x := range im.pixels[y] {
			im.pixels[y][x] = i.pixelAt(x-1, y-1)
		}
	}

	return im
}

func (i image) enhance() image {
	im := image{
		pixels: make([][]rune, len(i.pixels)),
		it:     i.it + 1,
		algo:   i.algo,
	}
	for y := range i.pixels {
		im.pixels[y] = make([]rune, len(i.pixels[y]))
		for x := range i.pixels[y] {
			section := i.section(x, y)
			im.pixels[y][x] = i.algo[pixelStringToBinary(section)]
		}
	}

	return im
}

func (i image) litPixels() int {
	n := 0
	for y := range i.pixels {
		for x := range i.pixels[y] {
			if i.pixels[y][x] == '#' {
				n++
			}
		}
	}
	return n
}

func (i image) pixelAt(x, y int) rune {
	if y < 0 || y >= len(i.pixels) || x < 0 || x >= len(i.pixels[0]) {
		if i.algo["000000000"] == '#' {
			if i.it%2 == 0 {
				return i.algo["111111111"]
			} else {
				return i.algo["000000000"]
			}
		} else {
			return '.'
		}
	}
	return i.pixels[y][x]
}

func (i image) section(x, y int) string {
	return string([]rune{
		i.pixelAt(x-1, y-1),
		i.pixelAt(x, y-1),
		i.pixelAt(x+1, y-1),
		i.pixelAt(x-1, y),
		i.pixelAt(x, y),
		i.pixelAt(x+1, y),
		i.pixelAt(x-1, y+1),
		i.pixelAt(x, y+1),
		i.pixelAt(x+1, y+1),
	})
}

func (i image) String() string {
	var sb strings.Builder
	for y := range i.pixels {
		for x := range i.pixels[y] {
			sb.WriteRune(i.pixels[y][x])
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	return sb.String()
}

func parseAlgorithm(s string) map[string]rune {
	res := map[string]rune{}
	for i, r := range s {
		res[fmt.Sprintf("%09b", i)] = r
	}

	return res
}

func main() {
	input := "2021/Day20/input"
	lines := util.GetFileStrings(input)

	algo := parseAlgorithm(lines[0])
	im := parseImage(lines[2:], algo)

	for i := 0; i < 50; i++ {
		if i == 2 {
			log.Printf("Part 1: Lit pixels after 2 iterations: %d", im.litPixels())
		}
		im = im.expand()
		im = im.enhance()
	}
	log.Printf("Part 2: Lit pixels after 50 iterations: %d", im.litPixels())

}
