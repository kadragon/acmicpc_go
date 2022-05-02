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

	var n, m, t, ans int
	fmt.Fscan(rd, &n, &m)
	ans = n + m

	d := map[int]bool{}
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		d[t] = true
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &t)
		if d[t] {
			ans -= 2
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
