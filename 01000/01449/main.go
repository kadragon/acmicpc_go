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

	var n, l, ans, s int
	fmt.Fscan(rd, &n, &l)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	for _, v := range d {
		if v > s {
			ans++
			s = v + l - 1
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
