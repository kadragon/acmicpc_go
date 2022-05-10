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

	var n, m, ans, k, t int
	fmt.Fscan(rd, &n, &m)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if k+t > m {
			ans++
			k = t
		} else {
			k += t
		}
	}

	if n == 0 {
		ans = -1
	}
	fmt.Fprintf(wr, "%d\n", ans+1)
}
