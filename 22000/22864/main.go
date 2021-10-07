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

	var a, b, c, m, s, p int
	fmt.Fscan(rd, &a, &b, &c, &m)

	for i := 1; i <= 24; i++ {
		if p+a <= m {
			s += b
			p += a
		} else {
			p -= c
			if p < 0 {
				p = 0
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", s)
}
