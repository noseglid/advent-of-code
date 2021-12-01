package main

import (
	"log"
)

func findPrivateKey(sub, pub int) int {
	key := 1
	loopSize := 0
	for key != pub {
		key = (key * sub) % 20201227
		loopSize++
	}

	return loopSize
}

func transform(sub, loopSize int) int {
	v := 1
	for i := 0; i < loopSize; i++ {
		v = v * sub
		v = v % 20201227
	}

	return v
}

func main() {
	doorpub := 14082811
	cardpub := 5249543

	// doorpub := 17807724 // sample
	// cardpub := 5764801  // sample

	cardpriv := findPrivateKey(7, cardpub)
	log.Printf("private card: %d", cardpriv)
	encKey := transform(doorpub, cardpriv)

	log.Printf("encryption key (part1): %d", encKey)

}
