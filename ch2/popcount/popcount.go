package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 練習問題2.3
func PopCount2(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// 練習問題2.4
func PopCount3(x uint64) int {
	count := 0
	for ; x > 0; x >>= 1 {
		if x & 1 == 1 {
			count++
		}
	}
	return count
}

// 練習問題2.5
func PopCount4(x uint64) int {
	count := 0
	for ; x > 0; x &= x - 1 {
		count++
	}
	return count
}

func benchmark(x uint64, f func(uint64) int) {
	const times = 1000000
	start := time.Now()
	var count int
	for i := 0; i < times; i++ {
		count = f(x)
	}
	secs := time.Since(start).Milliseconds()
	fmt.Printf("%d(%b) = count %d   %dms\n", x, x, count, secs)
}

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			continue
		}
		benchmark(x, PopCount)
		benchmark(x, PopCount2)
		benchmark(x, PopCount3)
		benchmark(x, PopCount4)
	}
}
