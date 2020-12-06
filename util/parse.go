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
