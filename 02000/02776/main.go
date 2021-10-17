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

	for i := 0; i < t; i++ {
		solve()
	}
}

func solve() {
	var n, m, p int
	fmt.Fscan(rd, &n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	fmt.Fscan(rd, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &p)
		fmt.Fprintf(wr, "%d\n", find(d, 0, n-1, p))
	}
}

func find(d []int, s, e, t int) int {
	if s > e {
		return 0
	}
	m := (s + e) / 2
	if d[m] == t {
		return 1
	} else if d[m] > t {
		return find(d, s, m-1, t)
	} else {
		return find(d, m+1, e, t)
	}
}
