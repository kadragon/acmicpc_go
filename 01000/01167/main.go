package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	next, cost int
}

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d       [100001][]node
	visited []bool

	maxCost, maxNode int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func dfs(k, c int) {
	if maxCost < c {
		maxCost = c
		maxNode = k
	}

	visited[k] = true
	for i := 0; i < len(d[k]); i++ {
		next, cost := d[k][i].next, d[k][i].cost
		if visited[next] {
			continue
		}

		dfs(next, c+cost)
	}
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d = [100001][]node{}

	var s, e, c int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &s)

		for {
			fmt.Fscan(rd, &e)
			if e == -1 {
				break
			}
			fmt.Fscan(rd, &c)

			d[s] = append(d[s], node{e, c})
			d[e] = append(d[e], node{s, c})
		}
	}

	visited = make([]bool, n+1)
	dfs(1, 0)

	visited = make([]bool, n+1)
	maxCost = 0
	dfs(maxNode, 0)

	fmt.Fprintf(wr, "%d\n", maxCost)
}
