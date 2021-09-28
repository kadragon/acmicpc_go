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

	var n, max int
	fmt.Fscan(rd, &n)

	d := make([]int, n+1)
	dt := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])

		for j := i - 1; j >= 0; j-- {
			if d[i] > d[j] && dt[i] < dt[j]+1 {
				dt[i] = dt[j] + 1
				if max < dt[i] {
					max = dt[i]
					break
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", n-max)
}
