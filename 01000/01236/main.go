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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	x := make([]bool, n)
	y := make([]bool, m)

	cx, cy := n, m

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if d[i][j] == 'X' {
				if !x[i] {
					x[i] = true
					cx--
				}
				if !y[j] {
					y[j] = true
					cy--
				}
			}
		}
	}

	var ans int
	if cx < cy {
		ans = cy
	} else {
		ans = cx
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
