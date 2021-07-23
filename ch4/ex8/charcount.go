package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

var classes = []struct {
	key string
	f   func(rune) bool
}{
	// {key: "Control", f: unicode.IsControl},
	// {key: "Digit", f: unicode.IsDigit},
	// {key: "Graphic", f: unicode.IsGraphic},
	// {key: "Letter", f: unicode.IsLetter},
	{key: "Lower", f: unicode.IsLower},
	// {key: "Mark", f: unicode.IsMark},
	{key: "Number", f: unicode.IsNumber},
	// {key: "Print", f: unicode.IsPrint},
	{key: "Punct", f: unicode.IsPunct},
	{key: "Space", f: unicode.IsSpace},
	{key: "Symbol", f: unicode.IsSymbol},
	// { key: "Title", f: unicode.IsTitle},
	{key: "Upper", f: unicode.IsUpper},
}

func main() {
	results := charcount(os.Stdin)
	fmt.Printf("type\tcount\n")
	for t, c := range results {
		fmt.Printf("%s\t%d\n", t, c)
	}
}

// TODO: 引数の r をリネームしたい
func charcount(r io.Reader) map[string]int {
	in := bufio.NewReader(r)
	invalid := 0
	counter := make(map[string]int)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		for _, c := range classes {
			if c.f(r) {
				counter[c.key]++
			}
		}
	}
	return counter
}
