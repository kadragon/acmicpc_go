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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var n, k int
	fmt.Fscan(rd, &n, &k)

	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
		d[i] += d[i-1]
	}
	max := d[k]
	for i := k + 1; i <= n; i++ {
		if max < d[i]-d[i-k] {
			max = d[i] - d[i-k]
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
