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

	var n int
	for {
		fmt.Fscan(rd, &n)
		if n == 0 {
			break
		}

		var p, t int
		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &t)
			if p != t {
				fmt.Fprintf(wr, "%d ", t)
				p = t
			}
		}
		fmt.Fprintf(wr, "$\n")
	}
}
