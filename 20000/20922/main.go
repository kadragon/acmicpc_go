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

	var n, k, max int
	fmt.Fscan(rd, &n, &k)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	a := make([]int, 100001)
	for s, e := 0, 0; e < n; e++ {
		a[d[e]]++
		for a[d[e]] > k {
			a[d[s]]--
			s++
		}
		if max < e-s+1 {
			max = e - s + 1
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
