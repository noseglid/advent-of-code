package util

import (
	"bufio"
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
