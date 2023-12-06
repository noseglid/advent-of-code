package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func FileScanner(file string, split bufio.SplitFunc) *bufio.Scanner {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Split(split)

	return s
}

func GetFile(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(buf)
}

func GetFileStrings(file string) []string {
	s := FileScanner(file, bufio.ScanLines)
	var ss []string

	for s.Scan() {
		ss = append(ss, s.Text())
	}
	return ss
}

func GetFileNumbers(file string) []int {
	s := FileScanner(file, bufio.ScanLines)
	var n []int

	for s.Scan() {
		n = append(n, MustAtoi(s.Text()))
	}
	return n
}

func GetCSVFileNumbers(file string) []int {
	t := GetFile(file)
	ll := strings.Split(t, ",")

	var ii []int
	for _, l := range ll {
		ii = append(ii, MustAtoi(strings.TrimSpace(l)))
	}

	return ii
}

func GetFileSingleDigitGrid(file string) [][]int {
	grid := [][]int{}
	lines := GetFileStrings(file)
	for _, l := range lines {
		row := make([]int, len(l))
		for x, r := range l {
			row[x] = MustAtoi(string(r))
		}
		grid = append(grid, row)
	}
	return grid
}

func GetFileRuneGrid(file string) [][]rune {
	grid := [][]rune{}
	lines := GetFileStrings(file)
	for _, l := range lines {
		row := make([]rune, len(l))
		for x, r := range l {
			row[x] = r
		}
		grid = append(grid, row)
	}
	return grid
}

func NumberList(s string) []int {
	numbers := []int{}
	for _, s := range strings.Split(s, " ") {
		if s == "" {
			continue
		}
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}
