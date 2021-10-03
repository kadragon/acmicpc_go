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

	var n int
	fmt.Fscan(rd, &n)

	dt := make([]int, 101)
	c := make([]int, n+1)
	v := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &c[i])
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &v[i])

		for k := 99; k >= c[i]; k-- {
			if dt[k-c[i]]+v[i] > dt[k] {
				dt[k] = dt[k-c[i]] + v[i]
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[99])
}
