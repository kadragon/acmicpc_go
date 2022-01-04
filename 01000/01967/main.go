package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	next, cost int
}

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d [][]node
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bfs(s int) int {
	m := 0

	for i := 0; i < len(d[s]); i++ {
		m = max(m, d[s][i].cost+bfs(d[s][i].next))
	}

	return m
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d = make([][]node, n+1)

	var s, e, c int

	for i := 0; i < n-1; i++ {
		fmt.Fscan(rd, &s, &e, &c)
		d[s] = append(d[s], node{e, c})
	}

	ans := 0
	for i := 1; i <= n; i++ {
		t := []int{}
		for j := 0; j < len(d[i]); j++ {
			t = append(t, d[i][j].cost+bfs(d[i][j].next))
		}

		sort.Ints(t)

		tmp := 0
		if len(t) >= 2 {
			tmp = t[len(t)-1] + t[len(t)-2]
		} else if len(t) == 1 {
			tmp = t[0]
		}

		ans = max(ans, tmp)
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
