package main

import (
	"log"
	"strings"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

func main() {
	s := util.GetFile("2019/Day9/input")

	var memory []int

	for _, op := range strings.Split(s, ",") {
		memory = append(memory, util.MustAtoi(strings.TrimSpace(op)))
	}

	program := intcode.NewProgram(memory)
	program.Input <- 1
	program.Run()

	log.Printf("BOOST keycode in test mode (part1): %d", <-program.Output)

	program.Reset()
	program.Input <- 2
	program.Run()
	log.Printf("Distress signal coordinates (part2): %d", <-program.Output)
}
