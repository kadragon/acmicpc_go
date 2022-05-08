package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	d = make([]int, 2)
}

func main() {
	defer wr.Flush()

	var t, n int
	fmt.Fscan(rd, &t)

	d[0] = 0
	d[1] = 1

	for i := 2; d[i-1]+d[i-2] <= 1000000000; i++ {
		d = append(d, d[i-1]+d[i-2])
	}

	for ; t > 0; t-- {
		fmt.Fscan(rd, &n)

		solve(n)
		fmt.Fprintln(wr)
	}
}

func solve(x int) {
	t := lower_bound(x)
	if t != x {
		solve(x - t)
	}
	fmt.Fprintf(wr, "%d ", t)
}

func lower_bound(x int) int {
	a, b := 0, len(d)
	for a < b {
		m := (a + b) / 2

		if d[m] <= x {
			a = m + 1
		} else {
			b = m
		}
	}

	return d[b-1]
}
