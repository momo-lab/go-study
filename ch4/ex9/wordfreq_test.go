package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestWordfreq(t *testing.T) {
	tests := []struct {
		s    string
		want map[string]int
	}{
		{
			"abc abc def",
			map[string]int{"abc": 2, "def": 1},
		},
		{
			"abc abc\tabc",
			map[string]int{"abc": 3},
		},
	}

	for _, test := range tests {
		description := fmt.Sprintf("%s", test.s)
		reader := strings.NewReader(test.s)
		ret, _ := wordfreq(reader)
		if !equals(ret, test.want) {
			t.Errorf("case %s: got %v, want %v", description, ret, test.want)
		}
	}
}

func equals(a, b map[string]int) bool {
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
