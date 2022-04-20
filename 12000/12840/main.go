package main

import (
	"bufio"
	"fmt"
	"os"
)

var MOD int = 86400

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

	var h, m, s, q, t, c int
	fmt.Fscan(rd, &h, &m, &s, &q)

	s += h*60*60 + m*60

	for ; q > 0; q-- {
		fmt.Fscan(rd, &t)

		if t == 3 {
			if s < 0 {
				s += MOD
			}
			fmt.Fprintf(wr, "%d %d %d\n", s/3600, s%3600/60, s%60)
		} else {
			fmt.Fscan(rd, &c)

			if t == 1 {
				s = (s + c) % MOD
			}
			if t == 2 {
				s = (s - c) % MOD
			}
		}
	}
}
