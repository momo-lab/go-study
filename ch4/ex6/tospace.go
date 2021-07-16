package main

import (
	"unicode"
	"unicode/utf8"
)

func toSpace(s []byte) []byte {
	pos := 0
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if size == 0 {
			break
		}
		if unicode.IsSpace(r) {
			s[pos] = ' '
			pos++
		} else {
			utf8.EncodeRune(s[pos:], r)
			pos += size
		}
		i += size
	}
	return s[:pos]
}
