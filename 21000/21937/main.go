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

	var n, m, a, b, q int
	fmt.Fscan(rd, &n, &m)

	d := make([][]int, n+1)

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &a, &b)
		d[b] = append(d[b], a)
	}

	fmt.Fscan(rd, &q)

	checked := make([]bool, n+1)
	checked[q] = true

	p := make([]int, 0)
	p = append(p, q)
	c := 0

	for len(p) > 0 {
		now := p[0]

		for _, v := range d[now] {
			if !checked[v] {
				checked[v] = true
				c++
				p = append(p, v)
			}
		}

		p = p[1:]
	}

	fmt.Fprintf(wr, "%d\n", c)
}
