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

	d := make([]int, n*3)
	for i := 0; i < n*3; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	fmt.Fprintf(wr, "%d\n", d[n*2-1]-d[n])
}

// 1 2 2 2 3 3 4 4 4
