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

	d := make([]int, n+1)
	dt := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 2; i <= n; i++ {
		min, max := 10001, 0

		for j := i; j > 0; j-- {
			if min > d[j] {
				min = d[j]
			}
			if max < d[j] {
				max = d[j]
			}
			if dt[i] < max-min+dt[j-1] {
				dt[i] = max - min + dt[j-1]
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[n])
}
