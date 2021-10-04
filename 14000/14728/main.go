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

	var n, t, k, s int
	dt := make([]int, 10001)

	fmt.Fscan(rd, &n, &t)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &k, &s)

		for j := t; j >= k; j-- {
			if dt[j] < dt[j-k]+s {
				dt[j] = dt[j-k] + s
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[t])
}
