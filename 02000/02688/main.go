package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dt [1001][10]int64
)

func f(k, p int) int64 {
	if k == 1 {
		return 1
	}

	if dt[k][p] == 0 {
		for i := p; i <= 9; i++ {
			dt[k][p] += f(k-1, i)
		}
	}

	return dt[k][p]
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	dt = [1001][10]int64{}

	var t, n int
	fmt.Fscan(rd, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(rd, &n)
		var ans int64

		for j := 0; j <= 9; j++ {
			ans += f(n, j)
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
