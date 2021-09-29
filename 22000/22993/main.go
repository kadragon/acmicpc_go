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
	var s int64
	fmt.Fscan(rd, &n, &s)

	d := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)
	for i := 0; i < n-1; i++ {
		if s <= int64(d[i]) {
			fmt.Fprintf(wr, "No")
			return
		}
		s += int64(d[i])
	}

	fmt.Fprintf(wr, "Yes")
}
