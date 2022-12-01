package main

import (
	"fmt"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

func listString(l []int) string {
	var sb strings.Builder
	sep := ""
	for _, v := range l {
		sb.WriteString(fmt.Sprintf("%s%d", sep, v))
		sep = " -> "
	}
	return sb.String()
}

func reverse(l []int) {
	for i := 0; i < len(l)/2; i++ {
		l[i], l[len(l)-i-1] = l[len(l)-i-1], l[i]
	}
}

func main() {
	// input := "206,63,255,131,65,80,238,157,254,24,133,2,16,0,1,3"
	// size := 256
	input := "3, 4, 1, 5"
	size := 5

	var lengths []int
	for _, s := range strings.Split(input, ",") {
		lengths = append(lengths, util.MustAtoi(strings.TrimSpace(s)))
	}
	currentPosition := 0
	skipSize := 0
	_ = skipSize

	list := make([]int, size)
	for i := range list {
		list[i] = i
	}

	for _, ll := range lengths {
		fmt.Println("=== ROUND START ===")
		fmt.Printf("currentPosition=%d, skipSize=%d, length=%d\n", currentPosition, skipSize, ll)
		fmt.Printf(">> %s\n", listString(list))
		newList := append([]int{}, list[currentPosition:]...)
		newList = append(newList, list[:currentPosition]...)
		fmt.Printf("rotated list >> %s\n", listString(newList))
		reverse(newList[:ll])
		fmt.Printf("reversed and rotated list >> %s\n", listString(newList))
		list = append([]int{}, newList[currentPosition+1:]...)
		list = append(list, newList[:currentPosition+1]...)
		fmt.Printf("list >> %s\n", listString(list))

		currentPosition = (currentPosition + ll + skipSize) % len(list)
		skipSize++

		fmt.Printf("=== ROUND END ===\n\n")
	}

	fmt.Printf("Final >> %s\n", listString(list))

	// list := util.NewCircularLinkedList[int]()
	// for i := 0; i < size; i++ {
	// 	list.Add(i)
	// }

	// // // cc := 3
	// for _, ll := range lengths {
	// 	fmt.Print("=== ROUND START ===\n")
	// 	fmt.Printf("list=%s\n", list.String())
	// 	fmt.Printf("length=%d, currentPosition=%d, skipSize=%d\n", ll, currentPosition, skipSize)
	// 	from := currentPosition
	// 	to := (currentPosition + ll - 1) % list.Len()
	// 	fmt.Printf("from=%d, to=%d\n", from, to)

	// 	list.Reverse(from, to)
	// 	fmt.Printf("reversed %s\n", list.String())

	// 	currentPosition = (currentPosition + ll + skipSize) % list.Len()
	// 	skipSize++
	// 	fmt.Printf("=== ROUND END ===\n\n")

	// 	// 	fmt.Printf("reversed\n")
	// 	// 	fmt.Printf("using length: %d\n", ll)
	// 	// 	ptr.reverse(ll - 1)
	// 	// 	head.print()
	// 	// 	fmt.Printf("---\n")

	// 	// 	currentPosition += lengths[0] + skipSize
	// 	// 	lengths = lengths[1:]
	// 	// 	skipSize++
	// 	// 	fmt.Printf("currentPosition=%d, skipSize=%d\n", currentPosition, skipSize)

	// 	// 	// cc--
	// 	// 	// if cc == 0 {
	// 	// 	// 	break
	// 	// 	// }
	// }

	fmt.Printf("Part 1: Multiplied numbers: %d", list[0]*list[1])

}
