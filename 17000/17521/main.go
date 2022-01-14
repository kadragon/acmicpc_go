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

	var n int
	var w int64
	fmt.Fscan(rd, &n, &w)

	d := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 1; i < n; i++ {
		if d[i-1] < d[i] {
			coin := w / d[i-1]
			w = w % d[i-1]
			w += coin * d[i]
		}
	}

	fmt.Fprintf(wr, "%d\n", w)
}
