package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

const allbits = math.MaxInt64

func tryParseMask(s string) (string, bool) {
	if s[:4] != "mask" {
		return "", false
	}

	return s[7:43], true
}

var re = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

func parseSet(s string) (int, int) {
	m := re.FindStringSubmatch(s)
	return util.MustAtoi(m[1]), util.MustAtoi(m[2])
}

func printBits(n int) {
	for i := 35; i >= 0; i-- {
		fmt.Printf("%d", n>>i&1)
	}
	fmt.Println()
}

func applyMask(number int, mask string) int {
	for i, m := range mask {

		if m == 'X' {
			continue
		} else if m == '0' {
			number &= ^(1 << (36 - i - 1))
		} else if m == '1' {
			number |= 1 << (36 - i - 1)
		} else {
			panic("bad mask")
		}
	}

	return number
}

func expandAddressMask(s string) []string {
	if len(s) == 1 {
		if s == "X" {
			return []string{"0", "1"}
		} else {
			return []string{s}
		}
	}

	var result []string
	for _, mm := range expandAddressMask(s[1:]) {
		if s[0] == 'X' {
			result = append(result, fmt.Sprintf("0%s", mm))
			result = append(result, fmt.Sprintf("1%s", mm))
		} else {
			result = append(result, fmt.Sprintf("%s%s", string(s[0]), mm))
		}
	}
	return result
}

func applyAddressMask(number int, mask string) string {
	var addr strings.Builder
	for i, m := range mask {
		if m == '0' {
			addr.WriteString(strconv.Itoa(number >> (36 - i - 1) & 1))
		} else if m == '1' || m == 'X' {
			addr.WriteRune(m)
		} else {
			panic("bad mask")
		}
	}

	return addr.String()
}

func memSum(mem map[int]int) int {
	sum := 0
	for _, n := range mem {
		sum += n
	}

	return sum
}

func main() {
	s := util.FileScanner("2020/Day14/input", bufio.ScanLines)

	mem := map[int]int{}
	var mask string

	for s.Scan() {
		if m, ok := tryParseMask(s.Text()); ok {
			mask = m
		} else {
			addr, n := parseSet(s.Text())
			mem[addr] = applyMask(n, mask)
		}
	}
	log.Printf("sum of memory (part1): %d", memSum(mem))

	s2 := util.FileScanner("2020/Day14/input", bufio.ScanLines)

	mem2 := map[int]int{}
	var mask2 string

	for s2.Scan() {
		if m, ok := tryParseMask(s2.Text()); ok {
			mask2 = m
		} else {
			addr, n := parseSet(s2.Text())
			for _, exp := range expandAddressMask(applyAddressMask(addr, mask2)) {
				v, err := strconv.ParseInt(exp, 2, 64)
				if err != nil {
					log.Fatal(err)
				}
				mem2[int(v)] = n
			}
		}
	}

	log.Printf("sum of memory (part2): %d", memSum(mem2))
}
