package main

func reverse(s []byte) {
	// 単純に入れ替える
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	// UTF-8の多バイト文字を再度入れ替える
	for i := 0; i < len(s); i++ {
		if s[i]>>5 == 0b110 {
			// 2バイト文字
			s[i], s[i-1] = s[i-1], s[i]
		} else if s[i]>>4 == 0b1110 {
			// 3バイト文字
			s[i], s[i-2] = s[i-2], s[i]
		} else if s[i]>>3 == 0b11110 {
			// 4バイト文字
			s[i], s[i-3] = s[i-3], s[i]
			s[i-1], s[i-2] = s[i-2], s[i-1]
		}
	}
}
