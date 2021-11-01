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

	var n, t int
	fmt.Fscan(rd, &n)

	dt := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dt[i] = 10000
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &t)

		for j := 1; j <= t && i+j <= n; j++ {
			if dt[i+j] > dt[i]+1 {
				dt[i+j] = dt[i] + 1
			}
		}
	}

	if dt[n] == 10000 {
		fmt.Fprintf(wr, "-1")
	} else {
		fmt.Fprintf(wr, "%d", dt[n])
	}
}
