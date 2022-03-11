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

	var r, c, k int
	fmt.Fscan(rd, &r, &c, &k)

	d := [1001][1001]int{}
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			fmt.Fscan(rd, &d[i][j])
			d[i][j] += d[i-1][j] + d[i][j-1] - d[i-1][j-1]
		}
	}

	var r1, c1, r2, c2 int

	for i := 0; i < k; i++ {
		fmt.Fscan(rd, &r1, &c1, &r2, &c2)
		t := d[r2][c2] - d[r1-1][c2] - d[r2][c1-1] + d[r1-1][c1-1]
		fmt.Fprintf(wr, "%d\n", t/((r2-r1+1)*(c2-c1+1)))
	}
}
