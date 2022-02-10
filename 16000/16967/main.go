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

	var h, w, x, y int
	d := [1001][1001]int{}
	a := [1001][1001]int{}

	fmt.Fscan(rd, &h, &w, &x, &y)

	for i := 0; i < h+x; i++ {
		for j := 0; j < w+y; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i >= x && j >= y {
				a[i][j] = d[i][j] - a[i-x][j-y]
			} else {
				a[i][j] = d[i][j]
			}

			fmt.Fprintf(wr, "%d ", a[i][j])
		}
		fmt.Fprintf(wr, "\n")
	}
}
