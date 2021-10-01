package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
	d  [41]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	d = [41]int{1, 1, 2}
}

func dp(k int) int {
	if d[k] == 0 {
		d[k] = dp(k-1) + dp(k-2)
	}
	return d[k]
}

func main() {
	defer wr.Flush()

	var n, k, s, t, a int
	fmt.Fscan(rd, &n, &k)

	s = 1
	a = 1

	for i := 0; i < k; i++ {
		fmt.Fscan(rd, &t)
		a *= dp(t - s)
		s = t + 1
	}

	fmt.Fprintf(wr, "%d", a*dp(n-s+1))
}
