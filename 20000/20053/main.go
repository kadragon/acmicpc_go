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

	var t int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		var n, a int
		max, min := -1000001, 1000001
		fmt.Fscan(rd, &n)

		for ; n > 0; n-- {
			fmt.Fscan(rd, &a)
			if max < a {
				max = a
			}
			if min > a {
				min = a
			}
		}

		fmt.Fprintf(wr, "%d %d\n", min, max)
	}
}
