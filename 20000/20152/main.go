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

	var h, n int
	fmt.Fscan(rd, &h, &n)

	if h > n {
		h, n = n, h
	}

	d := [31][31]int{}
	d[0][0] = 1
	for x := 1; x <= n-h; x++ {
		d[x][0] = 1
		for y := 1; y <= x; y++ {
			d[x][y] = d[x-1][y] + d[x][y-1]
		}
	}

	fmt.Fprintf(wr, "%d\n", d[n-h][n-h])
}
