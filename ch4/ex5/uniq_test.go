package main

import (
	"fmt"
	"testing"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		s    []string
		want []string
	}{
		{
			[]string{"a", "b", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "b", "b", "c", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			[]string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		description := fmt.Sprintf("%v", test.s)
		result := uniq(test.s)
		if !equals(result, test.want) {
			t.Errorf("case %v: got %v, want %v", description, result, test.want)
		}
	}
}

func equals(a, b []string) bool {
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
