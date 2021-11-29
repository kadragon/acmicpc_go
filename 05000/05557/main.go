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

	var n int
	d := [101]int{}
	dt := [101][21]int64{}

	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	dt[0][d[0]] = 1

	for i := 1; i < n-1; i++ {
		for j := 20; j >= 0; j-- {
			if dt[i-1][j] > 0 {
				if j+d[i] <= 20 {
					dt[i][j+d[i]] += dt[i-1][j]
				}

				if j >= d[i] {
					dt[i][j-d[i]] += dt[i-1][j]
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[n-2][d[n-1]])
}
