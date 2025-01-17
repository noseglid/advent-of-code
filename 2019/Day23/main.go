package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/noseglid/advent-of-code/util"
	"github.com/noseglid/advent-of-code/util/intcode"
)

type packet struct {
	src, dest int
	x, y      int
}

// Packets, receiver id as key, packet as value
var packets = map[int]chan packet{}

func bootComputer(id int, p *intcode.Program) {
	go p.Run()
	<-p.RequestInput
	p.Input <- id
}

func io(programs []*intcode.Program) (int, int) {
	p1 := math.MaxInt
	natYs := []int{}
	natPkg := packet{}
	idle := 0
	for {
		bidle := true
		for id, p := range programs {
			select {
			case dest := <-p.Output:
				bidle = false
				x, y := <-p.Output, <-p.Output
				packet := packet{src: id, dest: dest, x: x, y: y}

				if dest == 255 {
					if p1 == math.MaxInt {
						p1 = packet.y
					}
					natPkg = packet
				} else {
					packets[packet.dest] <- packet
				}
			case <-p.RequestInput:
				select {
				case packet := <-packets[id]:
					p.Input <- packet.x
					<-p.RequestInput
					p.Input <- packet.y
				default:
					p.Input <- -1
				}
			default:
			}
		}

		if bidle {
			idle++
		}
		if idle > 100 {
			idle = 0
			natPkg.dest = 0
			if slices.Contains(natYs, natPkg.y) {
				return p1, natPkg.y
			}
			natYs = append(natYs, natPkg.y)
			packets[natPkg.dest] <- natPkg
			natPkg = packet{}
		}
	}
}

func main() {
	src := util.GetCSVFileNumbers("2019/Day23/input")
	programs := make([]*intcode.Program, 50)

	for id := 0; id < 50; id++ {
		p := intcode.NewProgram(src, intcode.WithPingForInput(true))
		bootComputer(id, p)
		programs[id] = p
		packets[id] = make(chan packet, 10000)
		programs = append(programs, p)
	}

	p1, p2 := io(programs)
	fmt.Printf("Y of first packet to 255 (part1): %d\n", p1)
	fmt.Printf("Y of first repeated NAT package (part2): %d\n", p2)
}
