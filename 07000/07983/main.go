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

type Data [][2]int

func (d Data) Len() int {
	return len(d)
}

func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Data) Less(i, j int) bool {
	return d[i][1] < d[j][1]
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d := make(Data, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i][0], &d[i][1])
	}

	sort.Sort(d)

	p := make([]int, 0)
	c := 0

	for i, v := range d {
		p = append(p, v[1]-c)
		c = v[1]
		p[i] -= v[0]
	}

	ans, k := p[0], 0
	for _, v := range p {
		k += v
		if ans > k {
			ans = k
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
