package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func pwComplete(rr []rune) bool {
	for _, r := range rr {
		if int(r) == '_' {
			return false
		}
	}

	return true
}

func main() {

	input := "ojvtpuvg"
	// input := "abc"
	i := 0
	var pwp1 strings.Builder
	var pwp2 = []rune{'_', '_', '_', '_', '_', '_', '_', '_'}
	for len(pwp1.String()) < 8 || !pwComplete(pwp2) {
		hh := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		if len(pwp1.String()) < 8 && hh[0] == 0 && hh[1] == 0 && hh[2] < 16 {
			pwp1.WriteString(strconv.FormatInt(int64(hh[2]), 16))
			log.Printf("matched at %d, password=%s", i, pwp1.String())
		}

		if hh[0] == 0 && hh[1] == 0 && hh[2] < 16 {
			c := strconv.FormatInt(int64(hh[3]>>4), 16)[0]
			pos := int64(hh[2])
			if pos < 8 && pwp2[pos] == '_' {
				pwp2[pos] = rune(c)
			}
			log.Printf("pw2 matched %c at %d, password=%s", c, pos, string(pwp2))
		}

		i++
	}
	log.Printf("Part 1: Password is %s", pwp1.String())
	log.Printf("Part 2: Password is %s", string(pwp2))

}
