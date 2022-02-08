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
	var n int
	fmt.Fscan(rd, &n)

	d := make([]int, n+1)
	dt := make([]int, n+1)

	var max int = -1000001

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
		dt[i] = d[i]
		if dt[i-1]+d[i] > dt[i] {
			dt[i] = dt[i-1] + d[i]
		}
		if max < dt[i] {
			max = dt[i]
		}
	}
	fmt.Fprintf(wr, "%d\n", max)
}

func main() {
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		solve()
	}
}
