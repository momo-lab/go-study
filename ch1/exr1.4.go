package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "q1.4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filenames := range counts {
		n := 0
		// TODO: mapのvalueリストを取得する方法はある？
		var v []string
		for file, count := range filenames {
			n += count
			v = append(v, file)
		}
		if n > 1 {
			fmt.Printf("%d\t%s %v\n", n, line, v)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// TODO: mapのmapについて、要素が増えたときに初期化を自動でする方法はある？
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}
