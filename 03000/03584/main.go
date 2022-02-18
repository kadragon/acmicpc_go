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

func solve() {
	var n, a, b int
	fmt.Fscan(rd, &n)

	p := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(rd, &a, &b)
		p[b] = a
	}

	fmt.Fscan(rd, &a, &b)
	c := make([]int, n+1)
	c[a] = 1

	for p[a] > 0 {
		c[p[a]] = 1
		a = p[a]
	}

	for {
		if c[b] == 1 {
			fmt.Fprintf(wr, "%d\n", b)
			return
		}
		b = p[b]
	}
}

func main() {
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &t)

	for i := 0; i < t; i++ {
		solve()
	}
}
