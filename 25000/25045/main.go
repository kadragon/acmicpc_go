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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	a := make([]int64, n)
	b := make([]int64, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &b[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	if n < m {
		n, m = m, n
	}

	var ans int64
	for i := 0; i < m; i++ {
		if a[i]-b[i] > 0 {
			ans += a[i] - b[i]
		} else {
			break
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
