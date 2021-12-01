package main

import (
	"fmt"
	"log"
	"math"

	"github.com/noseglid/advent-of-code/util"
)

const width, height = 25, 6

func printImage(image [][]int) {
	for _, rows := range image {
		for _, v := range rows {
			if v == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func countPixels(layer [][]int, v int) int {
	n := 0
	for _, row := range layer {
		for _, vv := range row {
			if v == vv {
				n++
			}
		}
	}

	return n
}

func findMinPixels(layers [][][]int, v int) int {
	min := math.MaxInt64
	i := -1
	for index, l := range layers {
		n := countPixels(l, v)
		if n < min {
			min = n
			i = index
		}
	}

	return i
}

func decodeImage(layers [][][]int) [][]int {
	image := make([][]int, height)
	for i := range image {
		image[i] = make([]int, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for _, l := range layers {
				if l[y][x] == 2 {
					continue
				}
				image[y][x] = l[y][x]
				break
			}
		}
	}

	return image

}

func main() {
	input := util.GetFile("2019/Day8/input")
	layers := make([][][]int, 0, 8)

	layer, x, y := 0, 0, 0
	for _, r := range input {
		if len(layers) == layer {
			layers = append(layers, [][]int{})
		}
		if len(layers[layer]) == y {
			layers[layer] = append(layers[layer], []int{})
		}

		v := util.MustAtoi(string(r))
		layers[layer][y] = append(layers[layer][y], v)

		x++
		if x == width {
			x = 0
			y++
			if y == height {
				y = 0
				layer++
			}
		}
	}

	l := findMinPixels(layers, 0)
	ones := countPixels(layers[l], 1)
	twos := countPixels(layers[l], 2)

	log.Printf("ones and twos in min 0 layer (part1): %d", ones*twos)

	log.Printf("decodec image (part2):")
	img := decodeImage(layers)
	printImage(img)

}
