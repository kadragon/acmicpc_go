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
	d := make([]int, 1001)
	dt := make([]int, 1001)

	fmt.Fscan(rd, &n)
	for i := 1; i <= n; i++ {
		dt[i] = 1

		fmt.Fscan(rd, &d[i])
		for j := i - 1; j > 0; j-- {
			if d[j] < d[i] && dt[j]+1 > dt[i] {
				dt[i] = dt[j] + 1
			}

			if max < dt[i] {
				max = dt[i]
				break
			}
		}
	}

	fmt.Fprintf(wr, "%d", max)
}
