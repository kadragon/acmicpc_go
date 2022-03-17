package main

import (
	"bufio"
	"fmt"
	"os"
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

	var n, m, ans int
	var t string
	fmt.Fscan(rd, &n, &m)

	d := map[string]bool{}
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		d[t] = true
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &t)
		if _, ok := d[t]; ok {
			ans++
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
