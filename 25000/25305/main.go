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

	var n, k, t int
	fmt.Fscan(rd, &n, &k)

	d := []int{}

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		d = append(d, t)
	}

	sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })

	fmt.Fprintf(wr, "%d\n", d[k-1])
}
