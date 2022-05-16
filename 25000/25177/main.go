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

	var n, m, max int
	fmt.Fscan(rd, &n, &m)

	a := make([]int, 1001)
	b := make([]int, 1001)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &b[i])
		if b[i]-a[i] > max {
			max = b[i] - a[i]
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
