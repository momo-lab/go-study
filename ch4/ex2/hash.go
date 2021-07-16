package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

// TODO: ハッシュ値を算出するときにデータをすべて読み込んでいるのがイヤ
func main() {
	var (
		s2 = flag.Bool("sha256", false, "use SHA256. default")
		s3 = flag.Bool("sha384", false, "use SHA384")
		s5 = flag.Bool("sha512", false, "use SHA512")
	)
	flag.Parse()

	var h hash.Hash
	switch {
	case *s2:
		h = sha256.New()
	case *s3:
		h = sha512.New384()
	case *s5:
		h = sha512.New()
	default:
		h = sha256.New()
	}

	if _, err := io.Copy(h, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "hash err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
