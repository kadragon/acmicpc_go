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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	d := [301][301]int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(rd, &d[i][j])
			d[i][j] += max(d[i-1][j], d[i][j-1])
		}
	}

	fmt.Fprintf(wr, "%d\n", d[n][m])
}
