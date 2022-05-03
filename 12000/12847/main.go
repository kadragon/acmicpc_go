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

	var n, m, max int64
	fmt.Fscan(rd, &n, &m)

	d := make([]int64, n+m)
	p := make([]int64, n+m)
	for i := m; i < n+m; i++ {
		fmt.Fscan(rd, &d[i])
		p[i] = p[i-1] + d[i] - d[i-m]

		if p[i] > max {
			max = p[i]
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
