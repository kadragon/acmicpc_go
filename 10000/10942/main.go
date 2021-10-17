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

	var n, t, a, b int
	fmt.Fscan(rd, &n)

	d := make([]int, n+1)
	dt := [2001][2001]int{}

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
		dt[i][i] = 1
	}

	for r := 1; r < n; r++ {
		for s := 1; s <= n-r; s++ {
			e := s + r

			if dt[s+1][e-1] == 1 && d[s] == d[e] {
				dt[s][e] = 1
			} else if r == 1 && d[s] == d[e] {
				dt[s][e] = 1
			}
		}
	}

	fmt.Fscan(rd, &t)

	for i := 1; i <= t; i++ {
		fmt.Fscan(rd, &a, &b)
		fmt.Fprintf(wr, "%d\n", dt[a][b])
	}
}
