package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type docs struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid *string
}

func (d docs) validnr(field *string, min, max int) bool {
	if field == nil {
		return false
	}
	nbyr, err := strconv.Atoi(*field)
	if err != nil {
		return false
	}

	return nbyr >= min && nbyr <= max
}

func (d docs) ValidByr() bool {
	return d.validnr(d.byr, 1920, 2002)
}

func (d docs) ValidIyr() bool {
	return d.validnr(d.iyr, 2010, 2020)
}

func (d docs) ValidEyr() bool {
	return d.validnr(d.eyr, 2020, 2030)
}

var hgtre = regexp.MustCompile(`(\d+)(cm|in)`)

func (d docs) ValidHgt() bool {
	if d.hgt == nil {
		return false
	}
	m := hgtre.FindStringSubmatch(*d.hgt)
	if len(m) != 3 {
		return false
	}
	v := util.MustAtoi(m[1])
	switch m[2] {
	case "cm":
		return v >= 150 && v <= 193
	case "in":
		return v >= 59 && v <= 75
	default:
		return false
	}
}

var hclre = regexp.MustCompile(`#([0-9a-f]+)`)

func (d docs) ValidHcl() bool {
	if d.hcl == nil {
		return false
	}

	m := hclre.FindStringSubmatch(*d.hcl)
	return len(m) == 2
}

func (d docs) ValidEcl() bool {
	if d.ecl == nil {
		return false
	}
	e := *d.ecl

	return e == "amb" || e == "blu" || e == "brn" || e == "gry" || e == "grn" || e == "hzl" || e == "oth"
}

var pidre = regexp.MustCompile(`\d{9}`)

func (d docs) ValidPid() bool {
	if d.pid == nil {
		return false
	}

	return pidre.MatchString(*d.pid)
}

func (d docs) Valid() bool {
	return d.ValidByr() && d.ValidIyr() && d.ValidEyr() && d.ValidHgt() && d.ValidHcl() && d.ValidEcl() && d.ValidPid()
}

func parseDocs(s string) docs {
	log.Printf("building from: %s", s)
	d := docs{}
	for _, f := range strings.Fields(s) {
		sp := strings.Split(f, ":")
		key, value := sp[0], sp[1]
		switch key {
		case "byr":
			d.byr = &value
		case "iyr":
			d.iyr = &value
		case "eyr":
			d.eyr = &value
		case "hgt":
			d.hgt = &value
		case "hcl":
			d.hcl = &value
		case "ecl":
			d.ecl = &value
		case "pid":
			d.pid = &value
		case "cid":
			d.cid = &value
		}
	}
	return d
}

func main() {
	f, err := os.Open("2020/Day4/input")
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	nvalid := 0
	for _, pp := range strings.Split(string(input), "\n\n") {
		if parseDocs(pp).Valid() {
			nvalid++
		}

	}

	// s := util.FileScanner("2020/Day4/input", bufio.ScanLines)
	// docs := []docs{}

	// var buf strings.Builder
	// for s.Scan() {
	// 	if len(s.Text()) == 0 {
	// 		docs = append(docs, parseDocs(buf.String()))
	// 		buf.Reset()
	// 	} else {
	// 		buf.WriteString(s.Text())
	// 		buf.WriteRune(' ')
	// 	}
	// }

	// nvalid := 0
	// for _, d := range docs {
	// 	if d.Valid() {
	// 		log.Printf("valid!")
	// 		nvalid++
	// 	}
	// }

	log.Printf("Valid documents (part1): %d", nvalid)
}
