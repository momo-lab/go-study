package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{
			[]byte("abcde"),
			[]byte("edcba"),
		},
		{
			[]byte("こんにちは 世界"),
			[]byte("界世 はちにんこ"),
		},
	}

	for _, test := range tests {
		description := fmt.Sprintf("%s", test.s)
		reverse(test.s)
		if !equals(test.s, test.want) {
			t.Errorf("case %s: got %s, want %s", description, test.s, test.want)
		}
	}
}

func equals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
