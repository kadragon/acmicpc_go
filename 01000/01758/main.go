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

	var n int
	fmt.Fscan(rd, &n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	ans := 0
	for i := n - 1; i >= 0; i-- {
		if d[i] > n-(i+1) {
			ans += d[i] - (n - (i + 1))
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
