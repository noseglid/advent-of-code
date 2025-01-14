package main

import (
	"fmt"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

func write(program *intcode.Program, cmd []string) {
	for _, cc := range cmd {
		for _, r := range cc {
			program.Input <- int(r)
		}
		program.Input <- '\n'
	}
}

/***********
    ABCD
(1)    #
(2)   ##
(3) #  #
(4)  ###
(5) # ##
(6) ## #
(7) ####

***********/

func main() {
	src := util.GetCSVFileNumbers("2019/Day21/input")
	program := intcode.NewProgram(src)

	go program.Run()

	write(program, []string{
		"NOT C J",
		"AND D J",
		"NOT A T",
		"OR T J",
		"WALK",
	})

	damage := 0
	for c := range program.Output {
		if c > 256 {
			damage = c
			break
		}

		fmt.Printf("%c", c)
	}
	fmt.Printf("Total hull damage (part1): %d\n", damage)

	program.Reset()

	go program.Run()
	write(program, []string{
		"NOT C T",
		"AND D T",
		"AND H T",
		"OR T J",

		"NOT B T",
		"AND D T",
		"OR T J",

		"NOT A T",
		"OR T J",

		"RUN",
	})

	damage = 0
	for c := range program.Output {
		if c > 256 {
			damage = c
			break
		}

		fmt.Printf("%c", c)
	}
	fmt.Printf("Total hull damage with extended senors (part2): %d\n", damage)

}
