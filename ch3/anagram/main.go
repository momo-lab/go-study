package main

import (
	"fmt"
	"sort"
	"strings"
)

// 練習問題 3.12
func main() {
	output("abc", "cba")
	output("abc", "cbd")
	output("しんぶんし", "ししぶんん")
	output("しんぶんし", "ししぶん")
}

// 出力
func output(s1, s2 string) {
	fmt.Printf("%s, %s is %v\n", s1, s2, isAnagram(s1, s2))
}

func isAnagram(s1, s2 string) bool {
	x, y := sortChars(s1), sortChars(s2)
	return x == y
}

func sortChars(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
