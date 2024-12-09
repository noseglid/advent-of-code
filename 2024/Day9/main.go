package main

import (
	"fmt"
	"slices"

	"github.com/noseglid/advent-of-code/util"
)

type btype int

const (
	File btype = 0
	Free btype = 1
)

type block struct {
	id   int
	t    btype
	size int
}

func dumpDisk(disk []block) {
	for _, b := range disk {
		for i := 0; i < b.size; i++ {
			if b.t == File {
				fmt.Printf("%d", b.id)
			} else {
				fmt.Printf(".")
			}
		}
	}
	fmt.Println()
}

func asArray(disk []block) []int {
	var arr []int

	for _, b := range disk {
		for i := 0; i < b.size; i++ {
			arr = append(arr, b.id)
		}
	}
	return arr
}

func IndexFromEnd[T any](s []T, cmp func(i T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if cmp(s[i]) {
			return i
		}
	}

	return -1
}

func defrag(arr []int) []int {
	for {
		i1, i2 := slices.Index(arr, -1), IndexFromEnd(arr, func(i int) bool { return i != -1 })
		if i1 == -1 || i2 == -1 || i2 <= i1 {
			break
		}
		arr[i1], arr[i2] = arr[i2], arr[i1]
	}
	return arr
}

func defragBlocks(disk []block) []block {
	for {
		did := false
	Inner:
		for i := len(disk) - 1; i >= 0; i-- {
			if disk[i].t == Free {
				continue
			}

			for j := 0; j < len(disk); j++ {
				if disk[j].t != Free {
					continue
				}
				if disk[j].size >= disk[i].size && j < i {
					did = true
					f := disk[j].size - disk[i].size
					disk[j], disk[i] = disk[i], disk[j]
					if f > 0 {
						disk[i].size -= f
						b := block{id: -1, size: f, t: Free}
						if j+1 == len(disk) {
							disk = append(disk, b)
						} else {
							disk = slices.Insert(disk, j+1, b)
						}
					}
					break Inner
				}
			}
		}
		if !did {
			break
		}
	}
	return disk
}

func checksum(arr []int) int {
	s := 0
	for p, v := range arr {
		if v == -1 {
			continue
		}
		s += p * v
	}
	return s
}

func main() {

	disk := []block{}
	nt := File
	id := 0
	for _, r := range util.GetFile("2024/Day9/input") {
		if r < '0' || r > '9' {
			continue
		}
		tid := -1
		if nt == File {
			tid = id
			id++
		}
		disk = append(disk, block{id: tid, t: nt, size: int(r - '0')})
		nt = (nt + 1) % 2
	}

	fmt.Printf("Checksum after defrag (part1): %d\n", checksum(defrag(asArray(disk))))
	fmt.Printf("Checksum after defrag (part2): %d\n", checksum(asArray(defragBlocks(disk))))
}
