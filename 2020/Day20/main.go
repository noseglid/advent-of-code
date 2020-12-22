package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

const tileSize = 10

type Rotation int32
type Flip rune

const (
	Rot0   = Rotation(0)
	Rot90  = Rotation(90)
	Rot180 = Rotation(180)
	Rot270 = Rotation(270)
)

const (
	FlipNone = Flip('n')
	FlipH    = Flip('h')
	FlipV    = Flip('v')
)

var rots = []Rotation{Rot0, Rot90, Rot180, Rot270}
var flips = []Flip{FlipV, FlipNone}

type tile struct {
	id     int
	layout [][]int
}

type placedTile struct {
	tile     tile
	layout   [][]int
	rotation Rotation
	flip     Flip
}

func printLayout(layout [][]int) {
	for y := range layout {
		for x := range layout[y] {
			switch layout[y][x] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func flip(layout [][]int, f Flip) [][]int {
	if f == FlipNone {
		return layout
	}
	rsize := len(layout)
	transform := func(x, y int) (int, int) {
		switch f {
		case FlipH:
			return rsize - 1 - x, y
		case FlipV:
			return x, rsize - 1 - y
		}
		panic("bad flip")
	}

	flipped := make([][]int, rsize)
	for i := range layout {
		flipped[i] = make([]int, rsize)
	}

	for y := range layout {
		for x := range layout[y] {
			nx, ny := transform(x, y)
			flipped[ny][nx] = layout[y][x]
		}
	}

	return flipped
}

func rotate(layout [][]int, r Rotation) [][]int {
	if r == Rot0 {
		return layout
	}
	rsize := len(layout)
	transform := func(x, y int) (int, int) {
		switch r {
		case Rot90:
			return rsize - y - 1, x
		case Rot180:
			return rsize - x - 1, rsize - y - 1
		case Rot270:
			return y, rsize - x - 1
		}
		panic("bad rotation")
	}

	rotated := make([][]int, rsize)
	for i := range layout {
		rotated[i] = make([]int, rsize)
	}

	for y := range layout {
		for x := range layout[y] {
			nx, ny := transform(x, y)
			rotated[ny][nx] = layout[y][x]
		}
	}

	return rotated
}

var re = regexp.MustCompile(`^Tile (\d+):$`)

func parseTile(s string) tile {
	parts := strings.Split(s, "\n")
	m := re.FindStringSubmatch(parts[0])

	t := tile{
		id:     util.MustAtoi(m[1]),
		layout: make([][]int, 0, tileSize),
	}

	for i, row := range parts[1:] {
		if len(row) == 0 {
			continue
		}

		t.layout = append(t.layout, make([]int, 0, tileSize))
		for _, r := range row {
			entry := 0
			if r == '#' {
				entry = 1
			}
			t.layout[i] = append(t.layout[i], entry)
		}
	}

	return t
}

func nextCoord(x, y, max int) (int, int) {
	y++
	if y >= max {
		y = 0
		x++
	}
	return x, y
}

func lineUpBelow(topLayout, bottomLayout [][]int) bool {
	for i := 0; i < 10; i++ {
		if topLayout[9][i] != bottomLayout[0][i] {
			return false
		}
	}
	return true
}

func lineUpRightOf(leftLayout, rightLayout [][]int) bool {
	for i := 0; i < 10; i++ {
		if leftLayout[i][9] != rightLayout[i][0] {
			return false
		}
	}
	return true
}

func refineGrid(tiles [][]placedTile) [][]int {
	rsize := (tileSize - 2) * len(tiles)
	grid := make([][]int, rsize)

	for i := 0; i < (tileSize-2)*len(tiles); i++ {
		grid[i] = make([]int, rsize)
	}

	for ty := range tiles {
		for tx := range tiles[ty] {
			for y := 1; y <= tileSize-2; y++ {
				for x := 1; x <= tileSize-2; x++ {
					grid[ty*(tileSize-2)+y-1][tx*(tileSize-2)+x-1] = tiles[ty][tx].layout[y][x]
				}
			}
		}
	}
	return grid
}

func mapTiles(remainingTiles []tile, x, y int, grid [][]placedTile) ([][]placedTile, bool) {
	if len(remainingTiles) == 0 {
		return grid, true
	}
	for i, tt := range remainingTiles {
		for _, flp := range flips {
			for _, rot := range rots {
				// log.Printf("at %d,%d attempting to put %d with flip %c and rotation %v", x, y, tt.id, flp, rot)
				placed := placedTile{
					rotation: rot,
					flip:     flp,
					layout:   rotate(flip(tt.layout, flp), rot),
					tile:     tt,
				}

				canPlace := true
				if y > 0 && !lineUpBelow(grid[y-1][x].layout, placed.layout) {
					canPlace = false
				}
				if x > 0 && !lineUpRightOf(grid[y][x-1].layout, placed.layout) {
					canPlace = false
				}

				if !canPlace {
					continue
				}

				grid[y][x] = placed
				rem := make([]tile, len(remainingTiles))
				copy(rem, remainingTiles)
				nx, ny := nextCoord(x, y, len(grid))
				// log.Printf("successful placed %d at %d,%d!", tt.id, x, y)
				if fullGrid, ok := mapTiles(append(rem[:i], rem[i+1:]...), nx, ny, grid); ok {
					return fullGrid, true
				}
			}
		}
	}

	return nil, false
}

func monsterAt(x, y int, l [][]int) bool {
	if x+19 >= len(l) {
		return false
	}

	if y-1 < 0 || y+1 >= len(l) {
		return false
	}

	/*           1111111111
	   01234567890123456789
	-1                   #
	 0 #    ##    ##    ###
	 1  #  #  #  #  #  #
	*/
	check := make([]int, 15)
	check[0] = l[y][x]
	check[1] = l[y+1][x+1]
	check[2] = l[y+1][x+4]
	check[3] = l[y][x+5]
	check[4] = l[y][x+6]
	check[5] = l[y+1][x+7]
	check[6] = l[y+1][x+10]
	check[7] = l[y][x+11]
	check[8] = l[y][x+12]
	check[9] = l[y+1][x+13]
	check[10] = l[y+1][x+16]
	check[11] = l[y][x+17]
	check[12] = l[y-1][x+18]
	check[13] = l[y][x+18]
	check[14] = l[y][x+19]

	for _, c := range check {
		if c != 1 {
			return false
		}
	}

	return true
}

func countObjects(layout [][]int) int {
	n := 0
	for _, row := range layout {
		for _, e := range row {
			if e == 1 {
				n++
			}
		}
	}

	return n
}

type coord struct {
	x, y int
}

func main() {
	input := util.GetFile("2020/Day20/input")

	var tiles []tile
	for _, td := range strings.Split(input, "\n\n") {
		t := parseTile(td)
		tiles = append(tiles, t)
	}

	placedTiles := make([][]placedTile, int(math.Sqrt(float64(len(tiles)))))
	for i := range placedTiles {
		placedTiles[i] = make([]placedTile, len(placedTiles))
	}
	g, ok := mapTiles(tiles, 0, 0, placedTiles)
	if !ok {
		panic("no grid")
	}

	t1, t2, t3, t4 := g[0][0], g[0][len(g[0])-1], g[len(g[0])-1][0], g[len(g[0])-1][len(g[0])-1]
	log.Printf("product of corners %d*%d*%d*%d (part1): %d", t1.tile.id, t2.tile.id, t3.tile.id, t4.tile.id, t1.tile.id*t2.tile.id*t3.tile.id*t4.tile.id)

	refinedGrid := rotate(flip(refineGrid(placedTiles), FlipV), Rot90)
	printLayout(refinedGrid)

	for _, flp := range flips {
		for _, rot := range rots {
			useGrid := rotate(flip(refinedGrid, flp), rot)

			partOfMonster := map[coord]bool{}
			for y := range useGrid {
				for x := range useGrid[y] {
					if monsterAt(x, y, useGrid) {
						partOfMonster[coord{x, y}] = true
					}
				}
			}
			if len(partOfMonster) == 0 {
				continue
			}
			log.Printf("non-monster objects (part2): %d", countObjects(useGrid)-len(partOfMonster)*15)
			break
		}
	}
}
