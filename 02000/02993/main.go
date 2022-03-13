package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var s string
	fmt.Fscan(rd, &s)

	d := s
	for i := 0; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			p := rev(s[0:i+1]) + rev(s[i+1:j+1]) + rev(s[j+1:])
			if p < d {
				d = p
			}
		}
	}

	fmt.Fprintf(wr, "%s", d)
}

func rev(s string) string {
	p := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		p = append(p, s[i])
	}

	return string(p)
}
