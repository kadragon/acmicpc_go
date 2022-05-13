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

	var n, m, ans, t int
	fmt.Fscan(rd, &n, &m)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		ans += t
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &t)
		ans -= t
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
