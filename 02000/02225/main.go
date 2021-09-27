package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
	dt [201][201]int64
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dt = [201][201]int64{}
}

func f(n, k int) int64 {
	if n > 0 && k == 0 {
		return 0
	}
	if n == 0 || k == 1 {
		return 1
	}
	if dt[n][k] == 0 {
		var s int64
		for i := 0; i <= n && n-i >= 0; i++ {
			s += f(n-i, k-1)
		}
		dt[n][k] = s % 1000000000
	}

	return dt[n][k]
}

func main() {
	defer wr.Flush()

	var n, k int
	fmt.Fscan(rd, &n, &k)

	fmt.Fprintf(wr, "%d\n", f(n, k))
}
