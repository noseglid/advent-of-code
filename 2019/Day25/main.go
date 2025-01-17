package main

import (
	"fmt"
	"sync"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

const code = "north\nnorth\ntake monolith\nnorth\ntake hypercube\nsouth\nsouth\neast\neast\ntake easter egg\neast\nsouth\ntake ornament\nwest\nsouth\nwest\n"

func main() {
	src := util.GetCSVFileNumbers("2019/Day25/input")

	program := intcode.NewProgram(src)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		program.Run()
	}()

	go func() {
		for c := range program.Output {
			fmt.Printf("%c", c)
		}
	}()

	// Interactive
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	str, _, _ := reader.ReadLine()
	//	for _, c := range str {
	//		program.Input <- int(c)
	//	}
	//	program.Input <- '\n'
	//}

	// Automatic
	for _, c := range code {
		program.Input <- int(c)
	}

	wg.Wait()

}
