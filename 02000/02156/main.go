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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	dt := make([][3]int, n+1)
	dt[1][1] = d[1]

	for i := 2; i <= n; i++ {
		dt[i][0] = dt[i-1][1] + d[i]
		dt[i][1] = dt[i-1][2] + d[i]
		dt[i][2] = max(max(dt[i-1][2], dt[i-1][1]), dt[i-1][0])
	}

	fmt.Fprintf(wr, "%d", max(max(dt[n][0], dt[n][1]), dt[n][2]))
}
