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

	for j := 1; j <= t; j++ {
		fmt.Fprintf(wr, "Case #%d:", j)

		var n, p int
		k := map[int]int{}
		fmt.Fscan(rd, &n)

		for i := 0; i < n*2; i++ {
			fmt.Fscan(rd, &p)
			q := p / 3 * 4

			if k[p] > 0 {
				k[p]--
			} else {
				k[q]++
				fmt.Fprintf(wr, " %d", p)
			}
		}

		fmt.Fprintln(wr)
	}
}
