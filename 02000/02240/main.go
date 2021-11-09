package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dt [1001][31]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dt = [1001][31]int{}
}

func main() {
	defer wr.Flush()

	var t, w int
	fmt.Fscan(rd, &t, &w)

	d := make([]int, t+1)
	for i := 1; i <= t; i++ {
		fmt.Fscan(rd, &d[i])
	}
	if d[1] == 1 {
		dt[1][w] = 1
	} else {
		dt[1][w-1] = 1
	}

	for i := 2; i <= t; i++ {
		for j := w; j >= 0; j-- {
			if (w-j-1)%2+1 == d[i] {
				dt[i][j] = dt[i-1][j] + 1
				if j > 0 {
					dt[i][j-1] = dt[i-1][j]
				}
			} else {
				if j > 0 {
					dt[i][j-1] = dt[i-1][j] + 1
				}
				dt[i][j] = dt[i-1][j]
			}
		}
	}

	for i := 1; i <= t; i++ {
		fmt.Fprintf(wr, "%d %#v\n", i, dt[i][:w+1])
	}

	fmt.Fprintf(wr, "%s\n", "Hello World!")
}
