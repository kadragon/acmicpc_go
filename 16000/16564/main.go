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

	var n, k int
	fmt.Fscan(rd, &n, &k)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	for i := 0; i < n; i++ {
		k += d[i]
		if i+1 < n && k/(i+1) < d[i+1] {
			fmt.Fprintf(wr, "%d\n", k/(i+1))
			return
		}
	}

	fmt.Fprintf(wr, "%d\n", k/n)
}
