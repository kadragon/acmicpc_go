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

	var n, m, d, p int
	fmt.Fscan(rd, &n, &m)

	dt := make([]int, n+1)

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &d, &p)
		for j := n; j >= d; j-- {
			if dt[j] < dt[j-d]+p {
				dt[j] = dt[j-d] + p
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[n])
}
