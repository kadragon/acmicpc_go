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

	var n, a, b, t int
	fmt.Fscan(rd, &n)

	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
		a += d[i]
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if t == 0 {
			b += d[i]
		}
	}

	fmt.Fprintf(wr, "%d\n%d\n", a, b)
}
