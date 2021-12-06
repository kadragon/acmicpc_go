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

	p []int
)

func find(a, b, t int) int {
	if a > b {
		return -1
	}
	mid := (a + b) / 2

	if p[mid] == t {
		return mid
	}
	if p[mid] > t {
		return find(a, mid-1, t)
	}
	return find(mid+1, b, t)
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var n, k, ans int
	fmt.Fscan(rd, &n)

	p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &p[i])
	}

	fmt.Fscan(rd, &k)

	sort.Ints(p)
	for i := 0; i < n; i++ {
		if k-p[i] > p[i] {
			if find(i, n-1, k-p[i]) != -1 {
				ans++
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
