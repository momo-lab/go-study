package main

import "fmt"

func main() {
	output([]int{1, 2, 3, 4}, 3)
	output([]int{1, 2, 3, 4, 5}, 3)
	output([]int{1, 2, 3, 4, 5, 6}, 3)
	output([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	output([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)

	output([]int{1, 2, 3, 4}, 1)
	output([]int{1, 2, 3, 4, 5}, 1)
	output([]int{1, 2, 3, 4, 5, 6}, 1)
	output([]int{1, 2, 3, 4, 5, 6, 7}, 1)
	output([]int{1, 2, 3, 4, 5, 6, 7, 8}, 1)

	output([]int{1, 2, 3, 4}, 4)
	output([]int{1, 2, 3, 4, 5}, 4)
	output([]int{1, 2, 3, 4, 5, 6}, 4)
	output([]int{1, 2, 3, 4, 5, 6, 7}, 4)
	output([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4)
}

func output(s []int, n int) {
	fmt.Printf("%v rotate ", s)
	rotate(s, n)
	fmt.Printf("%v\n", s)
}

// 練習問題4.4
// TODO: 動いてはいるんだけど、自分で思いついたアルゴリズムではない。
func rotate(s []int, n int) {
	l := len(s)
	g := gcd(l, n)

	for i := 0; i < g; i++ {
		for j := (i + n) % l; j != i; j = (j + n) % l {
			s[i], s[j] = s[j], s[i]
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
