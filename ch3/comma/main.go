package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	output("")
	output("1")
	output("12")
	output("123")
	output("1234")
	output("12345")
	output("123456")
	output("1234567")
	output("12345678")
	output("123456789")
	output("1234567890")

	outputf("")
	outputf("0")
	outputf("-1000")
	outputf("-999")
	outputf("0.1")
	outputf("-0.1")
	outputf("0.000001")
	outputf("-0.000001")
	outputf("+1234567.89012345")
	outputf("1234567.89012345")
	outputf("-1234567.89012345")
}

func output(s string) {
	fmt.Printf("[%s] = [%s] = [%s]\n", s, comma(s), comma2(s))
}

func outputf(s string) {
	fmt.Printf("[%s] = [%s]\n", s, commaf(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 練習問題 3.10
func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	if n > 0 {
		buf.WriteString(s[:n])
	}
	for ss := s[n:]; len(ss) > 0; ss = ss[3:] {
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(ss[:3])
	}
	return buf.String()
}

// 練習問題 3.11
func commaf(s string) string {
	ss := s[:]

	// 符号記号
	sign := ""
	if len(ss) > 0 && (ss[0] == '+' || ss[0] == '-') {
		sign = ss[0:1]
		ss = ss[1:]
	}

	// 小数点以下
	scale := ""
	pos := strings.Index(ss, ".")
	if pos >= 0 {
		scale = ss[pos:]
		ss = ss[:pos]
	}

	return sign + comma2(ss) + scale
}
