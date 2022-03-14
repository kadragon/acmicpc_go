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
	var j, n, r, c, k int

	fmt.Fscan(rd, &j, &n)

	d := []int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &r, &c)
		d = append(d, r*c)
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	for i := 0; i < n; i++ {
		j -= d[i]
		k++
		if j <= 0 {
			break
		}
	}

	fmt.Fprintf(wr, "%d\n", k)
}
