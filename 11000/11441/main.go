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

	var n, t, a, b int
	fmt.Fscan(rd, &n)

	d := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
		d[i] += d[i-1]
	}

	fmt.Fscan(rd, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(rd, &a, &b)
		fmt.Fprintf(wr, "%d\n", d[b]-d[a-1])
	}
}
