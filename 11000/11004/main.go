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

	fmt.Fprintf(wr, "%d\n", d[k-1])
}
