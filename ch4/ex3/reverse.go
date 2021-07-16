package main

import "fmt"

func main() {
	output([]int{1, 2, 3, 4, 5})
	output2([...]int{1, 2, 3, 4, 5})
}

func output(s []int) {
	fmt.Printf("%v reverse to ", s)
	reverse(s)
	fmt.Printf("%v\n", s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func output2(s [5]int) {
	fmt.Printf("%v reverse to ", s)
	reverse2(&s)
	fmt.Printf("%v\n", s)
}

// 練習問題4.3
func reverse2(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
