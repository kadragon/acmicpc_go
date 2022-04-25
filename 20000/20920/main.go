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

type p struct {
	s    string
	c, l int
}

func main() {
	defer wr.Flush()

	var n, m int
	fmt.Fscan(rd, &n, &m)

	d := map[string]int{}

	var s string
	for ; n > 0; n-- {
		fmt.Fscan(rd, &s)
		if len(s) >= m {
			d[s]++
		}
	}

	q := []p{}
	for i, v := range d {
		q = append(q, p{i, v, len(i)})
	}

	sort.Slice(q, func(i, j int) bool {
		if q[i].c == q[j].c {
			if q[i].l == q[j].l {
				return q[i].s < q[j].s
			}
			return q[i].l > q[j].l
		}
		return q[i].c > q[j].c
	})

	for _, v := range q {
		fmt.Fprintln(wr, v.s)
	}
}
