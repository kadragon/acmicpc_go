package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d [10][2]int

	ans, n int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	d = [10][2]int{}
}

func gap(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func f(k, a, b int) {
	if k == n {
		if b > 0 {
			t := gap(a, b)
			if ans > t {
				ans = t
			}
		}
		return
	}

	f(k+1, a, b)
	f(k+1, a*d[k][0], b+d[k][1])
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	ans = 1000000001

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i][0], &d[i][1])
	}

	f(0, 1, 0)

	fmt.Fprintf(wr, "%d\n", ans)
}
