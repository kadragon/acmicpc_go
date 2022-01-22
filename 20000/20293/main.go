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

const INF int = 6001

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer wr.Flush()

	var r, c, n int
	fmt.Fscan(rd, &r, &c, &n)

	d := [3002][3002]int{}

	var x, y, f int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &x, &y, &f)
		d[x][y] = f
	}

	for i := r; i > 0; i-- {
		d[i][c+1] = INF

		for j := c; j > 0; j-- {
			d[r+1][j] = INF

			if i == r && j == c {
				d[i][j] = 0
				continue
			}

			t := min(d[i+1][j], d[i][j+1]) + 1
			if d[i][j] > 0 {
				t -= d[i][j]
				if t < 0 {
					t = 0
				}
			}
			d[i][j] = t
		}
	}

	fmt.Fprintf(wr, "%d\n", d[1][1])
}
