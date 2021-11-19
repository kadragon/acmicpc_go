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

	var n, t int
	var as, ts string
	m := 5

	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &ts, &t)
		if t < m {
			as = ts
			m = t
		}
	}

	fmt.Fprintf(wr, "%s\n", as)
}
