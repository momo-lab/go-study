package main

import (
	"fmt"
//	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	benchmark(exr1_1, "stringsを使う    ")
	benchmark(exr1_2, "stringsを使わない")
}

func benchmark(f func() string, msg string) {
	const times = 100000
	start := time.Now()
	for i := 0; i < times; i++ {
		f()
		//fmt.Fprintf(ioutil.Discard, "%v\n", f())
	}
	score := time.Since(start).Milliseconds()
	fmt.Printf("%s: %.2dms\n", msg, score)
	fmt.Println(f())
}

func exr1_1() string {
	return strings.Join(os.Args[1:], " ")
}

func exr1_2() string {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}
