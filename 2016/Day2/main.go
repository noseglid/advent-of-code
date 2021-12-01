package main

import (
	"bufio"
	"log"

	"github.com/noseglid/advent-of-code/util"
)

func main() {
	s := util.FileScanner("2016/Day2/input", bufio.ScanLines)

	keypadp2 := [][]rune{
		{'_', '_', '1', '_', '_'},
		{'_', '2', '3', '4', '_'},
		{'5', '6', '7', '8', '9'},
		{'_', 'A', 'B', 'C', '_'},
		{'_', '_', 'D', '_', '_'},
	}

	xp1, yp1 := 1, 1
	xp2, yp2 := 0, 2
	log.Printf("buttons: p1    p2")
	for s.Scan() {
		for _, d := range s.Text() {
			switch d {
			case 'U':
				if yp1-1 >= 0 {
					yp1--
				}

				if yp2-1 < 0 || keypadp2[yp2-1][xp2] == '_' {
					break
				}

				yp2--
			case 'R':
				if xp1+1 < 3 {
					xp1++
				}

				if xp2+1 > 4 || keypadp2[yp2][xp2+1] == '_' {
					break
				}
				xp2++
			case 'D':
				if yp1+1 < 3 {
					yp1++
				}

				if yp2+1 > 4 || keypadp2[yp2+1][xp2] == '_' {
					break
				}
				yp2++
			case 'L':
				if xp1-1 >= 0 {
					xp1--
				}
				if xp2-1 < 0 || keypadp2[yp2][xp2-1] == '_' {
					break
				}
				xp2--
			default:
				panic("bad direction")
			}
		}
		log.Printf("         %d     %s", yp1*3+xp1+1, string(keypadp2[yp2][xp2]))
	}
}
