package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	output("x", "X")
	output("X", "X")
}

func output(s1, s2 string) {
	fmt.Printf("%s == %s is %v\n", s1, s2, countDiffBitsForHash(s1, s2))
}

func countDiffBitsForHash(s1, s2 string) int {
	hash1 := sha256.Sum256([]byte(s1))
	hash2 := sha256.Sum256([]byte(s2))
	count := 0
	for i, c1 := range hash1 {
		c2 := hash2[i]
		count += countDiffBits(c1, c2)
	}
	return count
}

func countDiffBits(c1, c2 uint8) int {
	return popCount(c1 ^ c2)
}

func popCount(x uint8) int {
	count := 0
	for ; x > 0; x &= x - 1 {
		count++
	}
	return count
}
