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

	var n, m, t int
	fmt.Fscan(rd, &n, &m)

	for ; m > 0; m-- {
		fmt.Fscan(rd, &t)

		d := make([]int, t+1)
		d[0] = 200001

		for i := 1; i <= t; i++ {
			fmt.Fscan(rd, &d[i])
			if d[i-1] < d[i] {
				fmt.Fprintf(wr, "No\n")
				return
			}
		}
	}

	fmt.Fprintln(wr, "Yes")
}
