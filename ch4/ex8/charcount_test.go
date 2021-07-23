package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		s    string
		want map[string]int
	}{
		{
			"abcABC",
			map[string]int{"Lower": 3, "Upper": 3},
		},
		{
			" \t\n\r",
			map[string]int{"Space": 4},
		},
		{
			"123.45",
			map[string]int{"Number": 5, "Punct": 1},
		},
		{
			"$`=~^|<>+",
			map[string]int{"Symbol": 9},
		},
		{
			"!?\"#%&'@()[]{}_-\\,./*;:",
			map[string]int{"Punct": 23},
		},
	}

	for _, test := range tests {
		description := fmt.Sprintf("%s", test.s)
		reader := strings.NewReader(test.s)
		ret := charcount(reader)
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
