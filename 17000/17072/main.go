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

	var n, m, a, b, c int
	fmt.Fscan(rd, &n, &m)

	d := [401][401]byte{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(rd, &a, &b, &c)
			t := a*2126 + b*7152 + c*722

			if t < 510000 {
				d[i][j] = '#'
			} else if t < 1020000 {
				d[i][j] = 'o'
			} else if t < 1530000 {
				d[i][j] = '+'
			} else if t < 2040000 {
				d[i][j] = '-'
			} else {
				d[i][j] = '.'
			}
		}

		fmt.Fprintf(wr, "%s\n", string(d[i][:m]))
	}
}
