package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	results, _ := wordfreq(os.Stdin)
	fmt.Printf("type\tcount\n")
	for t, c := range results {
		fmt.Printf("%s\t%d\n", t, c)
	}
}

// TODO: 引数の r をリネームしたい
func wordfreq(reader io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	counter := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		counter[word]++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return counter, nil
}
