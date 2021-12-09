package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n      int
	m, ans int64

	d []int
)

func f(mul, t, r, depth int) {
	if r == 0 {
		if t%2 == 1 {
			ans += m / int64(mul)
		} else {
			ans -= m / int64(mul)
		}
		return
	}

	if depth == n {
		return
	}

	f(mul*d[depth], t, r-1, depth+1)
	f(mul, t, r, depth+1)
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	fmt.Fscan(rd, &n, &m)

	d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 1; i <= n; i++ {
		f(1, i, i, 0)
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
