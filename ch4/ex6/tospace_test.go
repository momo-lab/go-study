package main

import (
	"fmt"
	"testing"
)

func TestToSpace(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{
			[]byte("こんにちは\t世界"),
			[]byte("こんにちは 世界"),
		},
		{
			[]byte("hello\nworld"),
			[]byte("hello world"),
		},
		{
			[]byte("こんにちは　世界　　hoge"),
			[]byte("こんにちは 世界  hoge"),
		},
	}

	for _, test := range tests {
		description := fmt.Sprintf("%s", test.s)
		result := toSpace(test.s)
		if !equals(result, test.want) {
			t.Errorf("case %s: got %s, want %s", description, result, test.want)
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
