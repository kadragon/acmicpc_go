package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n, a, b, c int
	fmt.Fscan(rd, &n, &a, &b, &c)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	k, cost := c, a
	for i := 0; i < n; i++ {
		if k/cost <= (k+d[i])/(cost+b) {
			k += d[i]
			cost += b
		}
	}

	fmt.Fprintf(wr, "%d\n", k/cost)
}
