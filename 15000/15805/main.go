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

	var n, t, max int
	fmt.Fscan(rd, &n)

	d := []int{}
	checked := make([]bool, n)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)

		if t > max {
			max = t
		}

		if !checked[t] {
			d = append(d, t)
			checked[t] = true
		} else {
			l := len(d)
			f[d[l-1]] = d[l-2]
			d = d[:l-1]
		}
	}

	f[d[0]] = -1

	fmt.Fprintf(wr, "%d\n", max+1)
	for _, v := range f[:max+1] {
		fmt.Fprintf(wr, "%d ", v)
	}
}
