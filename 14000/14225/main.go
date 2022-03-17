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

	var n, t, max int
	fmt.Fscan(rd, &n)

	d := make([]bool, 2000001)
	d[0] = true

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		for i := max; i >= 0; i-- {
			if d[i] {
				d[i+t] = true
			}
		}

		max += t
	}
	ans := false
	for i := 1; i <= max; i++ {
		if !d[i] {
			fmt.Fprintf(wr, "%d\n", i)
			ans = true
			break
		}
	}

	if !ans {
		fmt.Fprintf(wr, "%d\n", max+1)
	}
}
