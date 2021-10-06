package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dt, t, p []int
	n        int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func f(k int) int {
	if k > n {
		return 0
	}

	if dt[k] > 0 {
		return dt[k]
	}

	if k+t[k] <= n+1 && f(k+t[k])+p[k] > f(k+1) {
		dt[k] = f(k+t[k]) + p[k]
	} else {
		dt[k] = f(k + 1)
	}

	return dt[k]
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	dt = make([]int, n+1)
	t = make([]int, n+1)
	p = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &t[i], &p[i])

	}

	fmt.Fprintf(wr, "%d\n", f(1))
}
