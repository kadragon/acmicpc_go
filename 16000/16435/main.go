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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var n, l int
	fmt.Fscan(rd, &n, &l)

	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	for i := 0; i < n; i++ {
		if l >= d[i] {
			l++
		} else {
			break
		}
	}

	fmt.Fprintf(wr, "%d\n", l)
}
