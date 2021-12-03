package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
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
