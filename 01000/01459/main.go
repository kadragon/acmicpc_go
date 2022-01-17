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

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer wr.Flush()

	var x, y, w, s int64
	fmt.Fscan(rd, &x, &y, &w, &s)

	var ans int64

	if w*2 > s {
		m := min(x, y)
		x -= m
		y -= m
		ans += m * s
	}

	t := x + y
	ans += (t - t%2) * min(w, s)
	ans += t % 2 * w

	fmt.Fprintf(wr, "%d\n", ans)
}
