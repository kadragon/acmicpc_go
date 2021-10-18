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

	var t int
	fmt.Fscan(rd, &t)

	for i := 1; i <= t; i++ {
		var m, n, x, y int
		ans := -1

		fmt.Fscan(rd, &m, &n, &x, &y)

		for k := 0; k <= n; k++ {
			if (k*m+x)%n == y%n {
				ans = k*m + x
				break

			}
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
