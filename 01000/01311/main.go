package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dt []int
	d  [][]int
	n  int
)

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

func f(k, bit int) int {
	if k == n {
		return 0
	}
	if dt[bit] == 0 {
		dt[bit] = 987654321
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) == 0 {
				dt[bit] = min(dt[bit], f(k+1, bit+1<<i)+d[k][i])
			}
		}
	}

	return dt[bit]
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	d = make([][]int, n)
	dt = make([]int, 1<<n)

	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	fmt.Fprintf(wr, "%d\n", f(0, 0))
}
