package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

func repeatedAny(s string, c int) (rune, bool) {
	n := 1
	cr := rune(s[0])
	for _, r := range s[1:] {
		if r == cr {
			n++
		} else {
			cr = r
			n = 1
		}
		if n == c {
			return cr, true
		}
	}

	return ' ', false
}

func isRepeated(s string, r rune, c int) bool {
	n := 0
	for _, rr := range s {
		if rr == r {
			n++
		} else {
			n = 0
		}
		if n == c {
			return true
		}
	}
	return false
}

func isKey(salt string, idx int, r rune, hashfn func(string, int) string) bool {
	for i := idx; i < idx+1000; i++ {
		h := hashfn(salt, i)
		if isRepeated(h, r, 5) {
			return true
		}
	}
	return false

}

var cache map[int]string
var cache2 map[int]string

func init() {
	cache = make(map[int]string)
	cache2 = make(map[int]string)
}

func hash(salt string, idx int) string {
	if h, ok := cache[idx]; ok {
		return h
	}
	s := fmt.Sprintf("%s%d", salt, idx)
	md5 := md5.New()
	md5.Write([]byte(s))
	h := hex.EncodeToString(md5.Sum(nil))
	cache[idx] = h
	return h
}

func hash2(salt string, idx int) string {
	if h, ok := cache2[idx]; ok {
		return h
	}

	s := fmt.Sprintf("%s%d", salt, idx)
	for i := 0; i <= 2016; i++ {
		md5 := md5.New()
		md5.Write([]byte(s))
		h := hex.EncodeToString(md5.Sum(nil))
		s = h
	}

	cache2[idx] = s

	return s
}

func main() {
	salt := "jlmsuwbz"
	i := 0
	keys := 0
	keys2 := 0
	index64thKey := -1
	index64thKey2 := -1
	for {
		if index64thKey == -1 {
			h := hash(salt, i)
			if r, ok := repeatedAny(h, 3); ok {
				if isKey(salt, i+1, r, hash) {
					keys++
					if keys == 64 {
						index64thKey = i
					}
				}
			}
		}

		if index64thKey2 == -1 {
			h2 := hash2(salt, i)
			if r, ok := repeatedAny(h2, 3); ok {
				if isKey(salt, i+1, r, hash2) {
					keys2++
					if keys2 == 64 {
						index64thKey2 = i
					}
				}
			}
		}
		if index64thKey != -1 && index64thKey2 != -1 {
			break
		}
		i++
	}

	log.Printf("Part 1: 64th key generated at index: %d", index64thKey)
	log.Printf("Part 2: 64th key generated at index: %d", index64thKey2)
}
