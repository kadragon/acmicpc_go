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

	for {
		var n int
		fmt.Fscan(rd, &n)

		if n == 0 {
			break
		}

		var max int64 = -10000
		d := make([]int64, n+1)

		for i := 1; i <= n; i++ {
			fmt.Fscan(rd, &d[i])
			if d[i]+d[i-1] > d[i] {
				d[i] = d[i] + d[i-1]
			}
			if d[i] > max {
				max = d[i]
			}
		}

		fmt.Fprintf(wr, "%d\n", max)
	}
}
