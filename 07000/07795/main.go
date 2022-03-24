package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	for ; t > 0; t-- {
		solve()
	}
}

func solve() {
	var n, m, ans, bp int
	fmt.Fscan(rd, &n, &m)

	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	for _, v := range a {
		for ; bp < m && b[bp] < v; bp++ {
		}
		ans += bp
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
